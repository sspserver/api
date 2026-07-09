package graphql

import (
	"context"

	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	"github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/pkg/repository/application"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// ApplicationConnection implements collection accessor interface with pagination.
type ApplicationConnection = connectors.CollectionConnection[*gqlmodels.Application]

// NewApplicationConnection based on query object
func NewApplicationConnection(
	ctx context.Context,
	applicationAccessor application.Usecase,
	filter *gqlmodels.ApplicationListFilter,
	order *gqlmodels.ApplicationListOrder,
	page *models.Page,
) *ApplicationConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.Application]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Application, error) {
			list, err := applicationAccessor.FetchList(ctx, ApplicationFilterFromGraphQL(filter), ApplicationOrderFromGraphQL(order), page.Pagination())
			return FromApplicationModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return applicationAccessor.Count(ctx, ApplicationFilterFromGraphQL(filter))
		},
	}, page)
}
