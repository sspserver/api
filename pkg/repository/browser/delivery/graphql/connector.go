package graphql

import (
	"context"

	"github.com/geniusrabbit/blaze-api/repository"
	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	"github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/pkg/repository/browser"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// BrowserConnection implements collection accessor interface with pagination.
type BrowserConnection = connectors.CollectionConnection[*gqlmodels.Browser]

// NewBrowserConnection based on query object
func NewBrowserConnection(
	ctx context.Context,
	browserAccessor browser.Usecase,
	filter *gqlmodels.BrowserListFilter,
	order []*gqlmodels.BrowserListOrder,
	page *models.Page,
) *BrowserConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.Browser]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Browser, error) {
			newOrder := BrowserListOrderFromGraphQL(order)
			newFilter := BrowserFilterFromGraphQL(filter)
			preloadOptions := (*repository.PreloadOption)(nil)
			if newFilter.IsChildrensPreload() {
				preloadOptions = &repository.PreloadOption{Fields: []string{`Versions`}}
			}
			list, err := browserAccessor.FetchList(ctx, preloadOptions, newFilter, &newOrder, page.Pagination())
			return FromBrowserModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return browserAccessor.Count(ctx, BrowserFilterFromGraphQL(filter))
		},
	}, page)
}
