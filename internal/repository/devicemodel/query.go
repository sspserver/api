package devicemodel

import (
	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/blaze-api/repository"
	"gorm.io/gorm"

	"github.com/sspserver/api/models"
)

// Filter of the objects list
type Filter struct {
	ID            []uint64
	Codename      []string
	Name          []string
	ParentID      []uint64
	Active        *types.ActiveStatus
	MakerCodename []string
	TypeCodename  []string
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
	if len(fl.Name) > 0 {
		query = query.Where(`name IN (?)`, fl.Name)
	}
	if len(fl.ParentID) > 0 {
		if len(fl.ParentID) == 1 && fl.ParentID[0] == 0 {
			query = query.Where(`parent_id IS NULL`)
		} else {
			query = query.Where(`parent_id IN (?)`, fl.ParentID)
		}
	}
	if fl.Active != nil {
		query = query.Where(`active = ?`, fl.Active.Name())
	}
	if len(fl.MakerCodename) > 0 {
		query = query.Where(`maker_codename IN (?)`, fl.MakerCodename)
	}
	if len(fl.TypeCodename) > 0 {
		query = query.Where(`type_codename IN (?)`, fl.TypeCodename)
	}
	return query
}

// ListOrder of the objects list
type ListOrder struct {
	ID            models.Order
	Codename      models.Order
	Name          models.Order
	TypeCodename  models.Order
	MakerCodename models.Order
	Active        models.Order
	CreatedAt     models.Order
	UpdatedAt     models.Order
	YearRelease   models.Order
}

func (ol *ListOrder) PrepareQuery(query *gorm.DB) *gorm.DB {
	if ol == nil {
		return query
	}
	query = ol.ID.PrepareQuery(query, `id`)
	query = ol.Codename.PrepareQuery(query, `codename`)
	query = ol.Name.PrepareQuery(query, `name`)
	query = ol.TypeCodename.PrepareQuery(query, `type_codename`)
	query = ol.MakerCodename.PrepareQuery(query, `maker_codename`)
	query = ol.Active.PrepareQuery(query, `active`)
	query = ol.CreatedAt.PrepareQuery(query, `created_at`)
	query = ol.UpdatedAt.PrepareQuery(query, `updated_at`)
	query = ol.YearRelease.PrepareQuery(query, `year_release`)
	return query
}

// List select options
type (
	Option  = repository.QOption
	Options = repository.ListOptions
)
