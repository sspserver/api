package graphql

import (
	"context"

	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	"github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/pkg/repository/trafficrouter"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// TrafficRouterConnection implements collection accessor interface with pagination.
type TrafficRouterConnection = connectors.CollectionConnection[*gqlmodels.TrafficRouter]

// NewTrafficRouterConnection based on query object
func NewTrafficRouterConnection(
	ctx context.Context,
	trafficRouterAccessor trafficrouter.Usecase,
	filter *gqlmodels.TrafficRouterListFilter,
	order []*gqlmodels.TrafficRouterListOrder,
	page *models.Page,
) *TrafficRouterConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.TrafficRouter]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.TrafficRouter, error) {
			newOrder := TrafficRouterListOrderFromGraphQL(order)
			list, err := trafficRouterAccessor.FetchList(ctx, TrafficRouterFilterFromGraphQL(filter), &newOrder, page.Pagination())
			return FromTrafficRouterModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return trafficRouterAccessor.Count(ctx, TrafficRouterFilterFromGraphQL(filter))
		},
	}, page)
}
