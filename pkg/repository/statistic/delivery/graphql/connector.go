package graphql

import (
	"context"

	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	"github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/pkg/repository/statistic"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// StatisticAdItemConnection implements collection accessor interface with pagination.
type StatisticAdItemConnection = connectors.CollectionConnection[*gqlmodels.StatisticAdItem]

// NewStatisticAdItemConnection based on query object
func NewStatisticAdItemConnection(
	ctx context.Context,
	statisticAccessor statistic.Usecase,
	filter *gqlmodels.StatisticAdListFilter,
	group []gqlmodels.StatisticKey,
	order []*gqlmodels.StatisticAdKeyOrder,
	page *models.Page,
) *StatisticAdItemConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.StatisticAdItem]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.StatisticAdItem, error) {
			list, err := statisticAccessor.Statistic(ctx,
				StatisticFilterFromGraphQL(filter), StatisticGroup(group),
				StatisticAdListOrder(order, group), page.Pagination())
			return FromStatisticAdItemModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return statisticAccessor.Count(ctx, StatisticFilterFromGraphQL(filter), StatisticGroup(group))
		},
	}, page)
}
