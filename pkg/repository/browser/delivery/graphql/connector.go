package graphql

import (
	"context"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/repository"
	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	"github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/pkg/repository/browser"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// BrowserConnection implements collection accessor interface with pagination.
type BrowserConnection = connectors.CollectionConnection[gqlmodels.Browser, gqlmodels.BrowserEdge]

// NewBrowserConnection based on query object
func NewBrowserConnection(
	ctx context.Context,
	browserAccessor browser.Usecase,
	filter *gqlmodels.BrowserListFilter,
	order []*gqlmodels.BrowserListOrder,
	page *models.Page,
) *BrowserConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[gqlmodels.Browser, gqlmodels.BrowserEdge]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Browser, error) {
			newOrder := xtypes.SliceReduce(order,
				func(val *gqlmodels.BrowserListOrder, res *browser.ListOrder) { val.Fill(res) })
			newFilter := filter.Filter()
			preloadOptions := (*repository.PreloadOption)(nil)
			if newFilter.IsChildrensPreload() {
				preloadOptions = &repository.PreloadOption{Fields: []string{`Versions`}}
			}
			list, err := browserAccessor.FetchList(ctx, preloadOptions, filter.Filter(), &newOrder, page.Pagination())
			return FromBrowserModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return browserAccessor.Count(ctx, filter.Filter())
		},
		ConvertToEdgeFunc: func(obj *gqlmodels.Browser) *gqlmodels.BrowserEdge {
			return &gqlmodels.BrowserEdge{
				Cursor: gocast.Str(obj.ID),
				Node:   obj,
			}
		},
	}, page)
}
