package adformat

import (
	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/blaze-api/repository"
	"gorm.io/gorm"

	"github.com/sspserver/api/models"
)

// Filter of the objects list
type Filter struct {
	ID           []uint64
	Codename     []string
	CodenameLike string
	Type         []string
	Active       *types.ActiveStatus
}

func (fl *Filter) PrepareQuery(query *gorm.DB) *gorm.DB {
	if fl == nil {
		return query
	}
	if len(fl.ID) > 0 {
		query = query.Where(`id IN (?)`, fl.ID)
	}
	if len(fl.Codename) > 0 {
		query = query.Where(`codename IN (?)`, fl.Codename)
	}
	if fl.CodenameLike != `` {
		query = query.Where(`codename LIKE ?`, fl.CodenameLike)
	}
	if len(fl.Type) > 0 {
		query = query.Where(`type IN (?)`, fl.Type)
	}
	if fl.Active != nil {
		query = query.Where(`active = ?`, fl.Active.Name())
	}
	return query
}

// ListOrder of the objects list
type ListOrder struct {
	Title     models.Order
	Codename  models.Order
	Type      models.Order
	Active    models.Order
	CreatedAt models.Order
	UpdatedAt models.Order
}

func (ol *ListOrder) PrepareQuery(query *gorm.DB) *gorm.DB {
	if ol == nil {
		return query
	}
	query = ol.Title.PrepareQuery(query, `title`)
	query = ol.Codename.PrepareQuery(query, `codename`)
	query = ol.Type.PrepareQuery(query, `type`)
	query = ol.Active.PrepareQuery(query, `active`)
	query = ol.CreatedAt.PrepareQuery(query, `created_at`)
	query = ol.UpdatedAt.PrepareQuery(query, `updated_at`)
	return query
}

type (
	Option  = repository.QOption
	Options = repository.ListOptions
)
