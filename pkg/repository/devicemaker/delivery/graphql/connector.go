package graphql

import (
	"context"

	"github.com/geniusrabbit/blaze-api/repository"
	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	"github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/pkg/repository/devicemaker"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// DeviceMakerConnection implements collection accessor interface with pagination.
type DeviceMakerConnection = connectors.CollectionConnection[*gqlmodels.DeviceMaker]

// NewDeviceMakerConnection based on query object
func NewDeviceMakerConnection(
	ctx context.Context,
	deviceMakerAccessor devicemaker.Usecase,
	filter *gqlmodels.DeviceMakerListFilter,
	order []*gqlmodels.DeviceMakerListOrder,
	page *models.Page,
) *DeviceMakerConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.DeviceMaker]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.DeviceMaker, error) {
			newOrder := DeviceMakerListOrderFromGraphQL(order)
			list, err := deviceMakerAccessor.FetchList(ctx,
				&repository.PreloadOption{Fields: []string{`Models`}},
				DeviceMakerFilterFromGraphQL(filter), &newOrder, page.Pagination())
			return FromDeviceMakerModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return deviceMakerAccessor.Count(ctx, DeviceMakerFilterFromGraphQL(filter))
		},
	}, page)
}
