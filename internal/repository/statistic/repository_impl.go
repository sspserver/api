package statistic

import (
	"context"

	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/archivarius/client"
	gqbzmodels "github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/internal/server/graphql/models"
)

// NewRepository is the constructor for the statistic repository.
func NewRepository(cl *client.APIClient) Repository {
	return &RepositoryImpl{cl: cl}
}

// RepositoryImpl is the implementation of the statistic repository
type RepositoryImpl struct {
	cl *client.APIClient
}

// Statistic returns a list of items by the given request.
func (r *RepositoryImpl) Statistic(
	ctx context.Context,
	group []models.StatisticKey,
	order []*models.StatisticAdKeyOrder,
	filter *models.StatisticAdListFilter,
	page *models.Page,
) (*client.StatisticResponse, error) {
	req := &client.StatisticRequest{}

	// Process filter
	if filter != nil {
		cond := make([]*client.FilterCondition, 0, len(filter.Conditions))
		// Iterate over conditions and map each
		for _, c := range filter.Conditions {
			var val []*client.Value
			if len(c.Value) > 0 {
				val = make([]*client.Value, 0, len(c.Value))
				// map value
				for _, v := range c.Value {
					nval, err := client.AnyExtValue(v)
					if err != nil {
						return nil, err
					}
					val = append(val, &client.Value{Value: nval})
				}
			}
			cond = append(cond, &client.FilterCondition{
				Key:   c.Key.ClientKey(),
				Op:    c.Op.ClientCondition(),
				Value: val,
			})
		}

		newFilter := &client.Filter{
			Conditions: cond,
		}

		if filter.StartDate != nil {
			newFilter.StartDate = client.TimeNew(filter.StartDate.GetTime())
		}
		if filter.EndDate != nil {
			newFilter.EndDate = client.TimeNew(filter.EndDate.GetTime())
		}

		req.Filter = newFilter
	}

	// Process order
	if len(order) > 0 {
		ord := make([]*client.Order, 0, len(order))
		// Iterate over order keys and map each
		for _, o := range order {
			ord = append(ord, &client.Order{
				Key: o.Key.ClientOrderingKey(),
				Asc: o.Order == gqbzmodels.OrderingAsc,
			})
		}

		req.Order = ord
	}

	// Process pagination
	if page != nil {
		req.PageOffset = uint64(page.Pagination().Offset)
		req.PageLimit = uint64(page.Pagination().Size)
	}

	// Process group
	req.Group = xtypes.SliceApply(group, func(k models.StatisticKey) client.Key {
		return k.ClientKey()
	})

	// Call the service
	return r.cl.Statistic(ctx, req)
}
