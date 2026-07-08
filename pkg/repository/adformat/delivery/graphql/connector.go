package graphql

import (
	"context"

	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	"github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/pkg/repository/adformat"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// AdFormatConnection implements collection accessor interface with pagination.
type AdFormatConnection = connectors.CollectionConnection[*gqlmodels.AdFormat]

// NewAdFormatConnection based on query object
func NewAdFormatConnection(
	ctx context.Context,
	adFormatAccessor adformat.Usecase,
	filter *gqlmodels.AdFormatListFilter,
	order *gqlmodels.AdFormatListOrder,
	page *models.Page,
) *AdFormatConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.AdFormat]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.AdFormat, error) {
			list, err := adFormatAccessor.FetchList(ctx,
				filter.Filter(), order.Order(), page.Pagination())
			return FromAdFormatModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return adFormatAccessor.Count(ctx, filter.Filter())
		},
	}, page)
}
