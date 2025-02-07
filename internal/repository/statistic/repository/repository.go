package repository

import (
	"context"

	"github.com/demdxx/xtypes"
	"gorm.io/gorm"

	"github.com/sspserver/api/internal/repository"
	"github.com/sspserver/api/internal/repository/statistic"
	"github.com/sspserver/api/models"
)

type Repository struct {
	repository.Repository
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Statistic(ctx context.Context, opts ...statistic.Option) ([]*models.StatisticAdItem, error) {
	var items []*AggregatedCountersLocal
	query := r.conn(ctx)
	query = statistic.Options(opts).PrepareQuery(query)
	if err := query.Find(&items).Error; err != nil {
		return nil, err
	}

	return xtypes.SliceApply(items, func(it *AggregatedCountersLocal) *models.StatisticAdItem {
		return it.AsStatisticItem()
	}), nil
}

func (r *Repository) Count(ctx context.Context, opts ...statistic.Option) (int64, error) {
	var count int64
	query := r.conn(ctx).Model(&AggregatedCountersLocal{})
	query = statistic.Options(opts).PrepareQuery(query)
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) conn(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx)
}
