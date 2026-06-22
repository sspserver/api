package graphql

import (
	"context"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/repository"
	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	"github.com/geniusrabbit/blaze-api/server/graphql/models"

	osrepo "github.com/sspserver/api/pkg/repository/os"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// OSConnection implements collection accessor interface with pagination.
type OSConnection = connectors.CollectionConnection[gqlmodels.Os, gqlmodels.OSEdge]

// NewOSConnection based on query object
func NewOSConnection(
	ctx context.Context,
	osAccessor osrepo.Usecase,
	filter *gqlmodels.OSListFilter,
	order []*gqlmodels.OSListOrder,
	page *models.Page,
) *OSConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[gqlmodels.Os, gqlmodels.OSEdge]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Os, error) {
			newOrder := xtypes.SliceReduce(order,
				func(val *gqlmodels.OSListOrder, res *osrepo.ListOrder) { val.Fill(res) })
			newFilter := filter.Filter()
			preloadOptions := (*repository.PreloadOption)(nil)
			if newFilter.IsChildrensPreload() {
				preloadOptions = &repository.PreloadOption{Fields: []string{`Versions`}}
			}
			list, err := osAccessor.FetchList(ctx, preloadOptions, newFilter, &newOrder, page.Pagination())
			return FromOSModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return osAccessor.Count(ctx, filter.Filter())
		},
		ConvertToEdgeFunc: func(obj *gqlmodels.Os) *gqlmodels.OSEdge {
			return &gqlmodels.OSEdge{
				Cursor: gocast.Str(obj.ID),
				Node:   obj,
			}
		},
	}, page)
}
