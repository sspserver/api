package graphql

import (
	"context"

	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	"github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/pkg/repository/zone"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// ZoneConnection implements collection accessor interface with pagination.
type ZoneConnection = connectors.CollectionConnection[*gqlmodels.Zone]

// NewZoneConnection based on query object
func NewZoneConnection(
	ctx context.Context,
	zoneAccessor zone.Usecase,
	filter *gqlmodels.ZoneListFilter,
	order *gqlmodels.ZoneListOrder,
	page *models.Page,
) *ZoneConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.Zone]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Zone, error) {
			list, err := zoneAccessor.FetchList(ctx, ZoneFilterFromGraphQL(filter), ZoneOrderFromGraphQL(order), page.Pagination())
			return FromZoneModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return zoneAccessor.Count(ctx, ZoneFilterFromGraphQL(filter))
		},
	}, page)
}
