package graphql

import (
	"context"

	gqbzmodels "github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/internal/repository/rtbsource"
	"github.com/sspserver/api/internal/repository/statistic"
	"github.com/sspserver/api/internal/server/graphql/models"
)

// QueryResolver is the resolver for the statistics.
type QueryResolver struct {
	uc        statistic.Usecase
	rtbsource rtbsource.Usecase
}

// NewQueryResolver is the constructor for the QueryResolver.
func NewQueryResolver(uc statistic.Usecase, rtbSource rtbsource.Usecase) *QueryResolver {
	return &QueryResolver{
		uc:        uc,
		rtbsource: rtbSource,
	}
}

// Statistic is the resolver for the statistic.
func (r *QueryResolver) Statistic(ctx context.Context, group []models.StatisticKey, order []*models.StatisticAdKeyOrder, filter *models.StatisticAdListFilter, page *models.Page) (*models.StatisticAdItemConnection, error) {
	resp, err := r.uc.Statistic(ctx, group, order, filter, page)
	if err != nil {
		return nil, err
	}

	// Convert response to the connection
	stats := &models.StatisticAdItemConnection{Edges: nil, PageInfo: nil}

	// Iterate over items and map each
	for _, i := range resp.Items {
		// Convert pb item
		item, err := itemFromPb(i)
		if err != nil {
			return nil, err
		}
		stats.List = append(stats.List, item)
		stats.Edges = append(stats.Edges, &models.StatisticAdItemEdge{
			Node: item,
		})
	}
	stats.TotalCount = int(resp.TotalCount)

	// Set page info
	stats.PageInfo = &gqbzmodels.PageInfo{
		HasNextPage: false,
		Total:       int(resp.TotalCount),
		Page:        0,
		Count:       0,
	}

	// Iterate over items and map each
	for sid, stat := range stats.List {
		for kid, key := range stat.Keys {
			if key.Key == models.StatisticKeySourceID {
				src, err := r.rtbsource.Get(ctx, key.Value.(uint64))
				if err != nil {
					return nil, err
				}
				stats.List[sid].Keys[kid].Text = src.Title
				stats.Edges[sid].Node.Keys[kid].Text = src.Title
			}
		}
	}

	return stats, nil
}
