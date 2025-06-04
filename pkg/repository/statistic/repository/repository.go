package repository

import (
	"context"
	"slices"

	"github.com/demdxx/xtypes"
	"gorm.io/gorm"

	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository"
	adfrepo "github.com/sspserver/api/pkg/repository/adformat/repository"
	"github.com/sspserver/api/pkg/repository/statistic"
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

	hasFormatCode := false
	hasFormatID := false
	for _, opt := range opts {
		if gopt, ok := opt.(*repository.GroupOption); ok {
			if keys, _ := gopt.Ext.([]statistic.Key); len(keys) > 0 {
				if slices.Contains(keys, statistic.KeyFormatCode) {
					hasFormatCode = true
				}
				if slices.Contains(keys, statistic.KeyFormatID) {
					hasFormatID = true
				}
				if hasFormatCode && hasFormatID {
					break
				}
			}
		}
	}

	formats := ([]*models.Format)(nil)
	if hasFormatCode {
		fmts := adfrepo.New()
		formats, _ = fmts.FetchList(ctx)
	}

	return xtypes.SliceApply(items, func(it *AggregatedCountersLocal) *models.StatisticAdItem {
		return it.AsStatisticItem(ctx, hasFormatID, formats)
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
