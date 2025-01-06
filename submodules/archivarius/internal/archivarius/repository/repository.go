package repository

import (
	"context"

	"github.com/geniusrabbit/archivarius/internal/archivarius"

	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/repository"
	"gorm.io/gorm"
)

// Repository is a storage for archivarius
type Repository struct {
	repository.Repository
	db *gorm.DB
}

// NewRepository creates a new Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Statistic returns a list of items
func (r *Repository) Statistic(ctx context.Context, opts ...archivarius.Option) ([]*archivarius.AdItem, error) {
	var items []*AggregatedCountersLocal
	query := r.conn(ctx)
	query = archivarius.Options(opts).PrepareQuery(query)
	if err := query.Find(&items).Error; err != nil {
		return nil, err
	}

	return xtypes.SliceApply(items, func(it *AggregatedCountersLocal) *archivarius.AdItem {
		return it.AsItem()
	}), nil
}

// Count returns a count of items
func (r *Repository) Count(ctx context.Context, opts ...archivarius.Option) (int64, error) {
	var count int64
	query := r.conn(ctx).Model(&AggregatedCountersLocal{})
	query = archivarius.Options(opts).PrepareQuery(query)
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// conn is a helper to get a connection with context
func (r *Repository) conn(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx)
}
