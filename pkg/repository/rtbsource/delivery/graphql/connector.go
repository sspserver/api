package graphql

import (
	"context"

	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	"github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/pkg/repository/rtbsource"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// RTBSourceConnection implements collection accessor interface with pagination.
type RTBSourceConnection = connectors.CollectionConnection[*gqlmodels.RTBSource]

// NewRTBSourceConnection based on query object
func NewRTBSourceConnection(
	ctx context.Context,
	rtbSourceAccessor rtbsource.Usecase,
	filter *gqlmodels.RTBSourceListFilter,
	order []*gqlmodels.RTBSourceListOrder,
	page *models.Page,
) *RTBSourceConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.RTBSource]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.RTBSource, error) {
			newOrder := xtypes.SliceReduce(order,
				func(val *gqlmodels.RTBSourceListOrder, res *rtbsource.ListOrder) { val.Fill(res) })
			list, err := rtbSourceAccessor.FetchList(ctx, filter.Filter(), &newOrder, page.Pagination())
			return FromRTBSourceModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return rtbSourceAccessor.Count(ctx, filter.Filter())
		},
	}, page)
}
