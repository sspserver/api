package graphql

import (
	"context"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/repository"
	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	"github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/pkg/repository/devicemaker"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// DeviceMakerConnection implements collection accessor interface with pagination.
type DeviceMakerConnection = connectors.CollectionConnection[gqlmodels.DeviceMaker, gqlmodels.DeviceMakerEdge]

// NewDeviceMakerConnection based on query object
func NewDeviceMakerConnection(
	ctx context.Context,
	deviceMakerAccessor devicemaker.Usecase,
	filter *gqlmodels.DeviceMakerListFilter,
	order []*gqlmodels.DeviceMakerListOrder,
	page *models.Page,
) *DeviceMakerConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[gqlmodels.DeviceMaker, gqlmodels.DeviceMakerEdge]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.DeviceMaker, error) {
			newOrder := xtypes.SliceReduce(order,
				func(val *gqlmodels.DeviceMakerListOrder, res *devicemaker.ListOrder) { val.Fill(res) })
			list, err := deviceMakerAccessor.FetchList(ctx,
				&repository.PreloadOption{Fields: []string{`Models`}},
				filter.Filter(), &newOrder, page.Pagination())
			return FromDeviceMakerModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return deviceMakerAccessor.Count(ctx, filter.Filter())
		},
		ConvertToEdgeFunc: func(obj *gqlmodels.DeviceMaker) *gqlmodels.DeviceMakerEdge {
			return &gqlmodels.DeviceMakerEdge{
				Cursor: gocast.Str(obj.ID),
				Node:   obj,
			}
		},
	}, page)
}
