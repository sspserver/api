package rtbsource

import (
	"github.com/geniusrabbit/blaze-api/repository"
	"gorm.io/gorm"

	"github.com/sspserver/api/models"
)

// Filter of the objects list
type Filter struct {
	ID        []uint64
	AccountID uint64
}

func (fl *Filter) PrepareQuery(query *gorm.DB) *gorm.DB {
	if fl == nil {
		return query
	}
	if len(fl.ID) > 0 {
		query = query.Where(`id IN (?)`, fl.ID)
	}
	if fl.AccountID > 0 {
		query = query.Where(`account_id = ?`, fl.AccountID)
	}
	return query
}

// ListOrder of the objects list
type ListOrder struct {
	Title     models.Order
	AccountID models.Order
	CreatedAt models.Order
	UpdatedAt models.Order
}

func (ol *ListOrder) PrepareQuery(query *gorm.DB) *gorm.DB {
	if ol == nil {
		return query
	}
	query = ol.Title.PrepareQuery(query, `title`)
	query = ol.AccountID.PrepareQuery(query, `account_id`)
	query = ol.CreatedAt.PrepareQuery(query, `created_at`)
	query = ol.UpdatedAt.PrepareQuery(query, `updated_at`)
	return query
}

// List select options
type (
	Option  = repository.QOption
	Options = repository.ListOptions
)
