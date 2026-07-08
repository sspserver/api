package graphql

import (
	"context"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/repository"
	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	"github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/pkg/repository/category"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// CategoryConnection implements collection accessor interface with pagination.
type CategoryConnection = connectors.CollectionConnection[*gqlmodels.Category]

// NewCategoryConnection based on query object
func NewCategoryConnection(
	ctx context.Context,
	categoryAccessor category.Usecase,
	filter *gqlmodels.CategoryListFilter,
	order *gqlmodels.CategoryListOrder,
	page *models.Page,
) *CategoryConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.Category]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Category, error) {
			newFilter := filter.Filter()
			list, err := categoryAccessor.FetchList(ctx,
				&repository.PreloadOption{
					Fields: gocast.IfThen(newFilter.IsChildrensPreload(),
						[]string{`Parent`, `Childrens`}, []string{`Parent`}),
				},
				newFilter, order.Order(), page.Pagination())
			return FromCategoryModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return categoryAccessor.Count(ctx, filter.Filter())
		},
	}, page)
}
