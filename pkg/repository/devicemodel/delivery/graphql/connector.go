package graphql

import (
	"context"

	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	"github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/pkg/repository/devicemodel"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// DeviceModelConnection implements collection accessor interface with pagination.
type DeviceModelConnection = connectors.CollectionConnection[*gqlmodels.DeviceModel]

// NewDeviceModelConnection based on query object
func NewDeviceModelConnection(
	ctx context.Context,
	deviceModelAccessor devicemodel.Usecase,
	filter *gqlmodels.DeviceModelListFilter,
	order []*gqlmodels.DeviceModelListOrder,
	page *models.Page,
) *DeviceModelConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.DeviceModel]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.DeviceModel, error) {
			newOrder := DeviceModelListOrderFromGraphQL(order)
			list, err := deviceModelAccessor.FetchList(ctx, DeviceModelFilterFromGraphQL(filter), &newOrder, page.Pagination())
			return FromDeviceModelModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return deviceModelAccessor.Count(ctx, DeviceModelFilterFromGraphQL(filter))
		},
	}, page)
}
