package graphql

import (
	"context"

	"github.com/sspserver/api/pkg/repository/statistic"
	qmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type QueryResolver struct {
	uc statistic.Usecase
}

func NewQueryResolver(uc statistic.Usecase) *QueryResolver {
	return &QueryResolver{uc: uc}
}

func (r *QueryResolver) Statistic(ctx context.Context, filter *qmodels.StatisticAdListFilter, group []qmodels.StatisticKey, order []*qmodels.StatisticAdKeyOrder, page *qmodels.Page) (*StatisticAdItemConnection, error) {
	return NewStatisticAdItemConnection(ctx, r.uc, filter, group, order, page), nil
}
