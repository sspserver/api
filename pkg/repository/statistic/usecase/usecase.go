package usecase

import (
	"context"

	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/statistic"
)

type Usecase struct {
	repo statistic.Repository
}

func NewUsecase(repo statistic.Repository) *Usecase {
	return &Usecase{repo: repo}
}

func (uc *Usecase) Statistic(ctx context.Context, opts ...statistic.Option) ([]*models.StatisticAdItem, error) {
	return uc.repo.Statistic(ctx, opts...)
}

func (uc *Usecase) Count(ctx context.Context, opts ...statistic.Option) (int64, error) {
	return uc.repo.Count(ctx, opts...)
}
