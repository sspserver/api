package graphql

import (
	"context"

	"github.com/sspserver/api/internal/repository/statistic"
	"github.com/sspserver/api/internal/server/graphql/connectors"
	qmodels "github.com/sspserver/api/internal/server/graphql/models"
)

type QueryResolver struct {
	uc statistic.Usecase
}

func NewQueryResolver(uc statistic.Usecase) *QueryResolver {
	return &QueryResolver{uc: uc}
}

func (r *QueryResolver) Statistic(ctx context.Context, filter *qmodels.StatisticAdListFilter, group []qmodels.StatisticKey, order []*qmodels.StatisticAdKeyOrder, page *qmodels.Page) (*connectors.StatisticAdItemConnection, error) {
	return connectors.NewStatisticAdItemConnection(ctx, r.uc, filter, group, order, page), nil
}
