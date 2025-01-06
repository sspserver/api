package category

import (
	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/blaze-api/repository"
	"gorm.io/gorm"

	"github.com/sspserver/api/models"
)

// Filter of the objects list
type Filter struct {
	ID       []uint64
	Name     []string
	ParentID []uint64
	IABCode  []string
	Active   *types.ActiveStatus
}

func (fl *Filter) PrepareQuery(query *gorm.DB) *gorm.DB {
	if fl == nil {
		return query
	}
	if len(fl.ID) > 0 {
		query = query.Where(`id IN (?)`, fl.ID)
	}
	if len(fl.Name) > 0 {
		query = query.Where(`name IN (?)`, fl.Name)
	}
	if len(fl.ParentID) > 0 {
		query = query.Where(`parent_id IN (?)`, fl.ParentID)
	}
	if len(fl.IABCode) > 0 {
		query = query.Where(`iab_code IN (?)`, fl.IABCode)
	}
	if fl.Active != nil {
		query = query.Where(`active = ?`, fl.Active.Name())
	}
	return query
}

// ListOrder of the objects list
type ListOrder struct {
	ID        models.Order
	Name      models.Order
	IABCode   models.Order
	ParentID  models.Order
	Position  models.Order
	Active    models.Order
	CreatedAt models.Order
	UpdatedAt models.Order
}

func (ol *ListOrder) PrepareQuery(query *gorm.DB) *gorm.DB {
	if ol == nil {
		return query
	}
	query = ol.ID.PrepareQuery(query, `id`)
	query = ol.Name.PrepareQuery(query, `name`)
	query = ol.IABCode.PrepareQuery(query, `iab_code`)
	query = ol.ParentID.PrepareQuery(query, `parent_id`)
	query = ol.Position.PrepareQuery(query, `position`)
	query = ol.Active.PrepareQuery(query, `active`)
	query = ol.CreatedAt.PrepareQuery(query, `created_at`)
	query = ol.UpdatedAt.PrepareQuery(query, `updated_at`)
	return query
}

// List select options
type (
	Option  = repository.QOption
	Options = repository.ListOptions
)
