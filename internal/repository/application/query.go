package application

import (
	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/blaze-api/repository"
	"gorm.io/gorm"

	"github.com/sspserver/api/models"
)

// Filter of the objects list
type Filter struct {
	ID       []uint64
	Title    string
	URI      string
	Type     []models.ApplicationType
	Platform []models.PlatformType
	Permium  *bool
	Status   *types.ApproveStatus
	Active   *types.ActiveStatus
}

func (fl *Filter) PrepareQuery(query *gorm.DB) *gorm.DB {
	if fl == nil {
		return query
	}
	if len(fl.ID) > 0 {
		query = query.Where(`id IN (?)`, fl.ID)
	}
	if len(fl.Title) > 0 {
		query = query.Where(`title ILIKE ?`, fl.Title)
	}
	if len(fl.URI) > 0 {
		query = query.Where(`uri ILIKE ?`, fl.URI)
	}
	if len(fl.Type) > 0 {
		query = query.Where(`type IN (?)`, fl.Type)
	}
	if len(fl.Platform) > 0 {
		query = query.Where(`platform IN (?)`, fl.Platform)
	}
	if fl.Permium != nil {
		query = query.Where(`premium = ?`, *fl.Permium)
	}
	if fl.Status != nil {
		query = query.Where(`status = ?`, *fl.Status)
	}
	if fl.Active != nil {
		query = query.Where(`active = ?`, *fl.Active)
	}
	return query
}

// ListOrder of the objects list
type ListOrder struct {
	ID models.Order

	Title models.Order
	URI   models.Order

	Type     models.Order
	Platform models.Order
	Premium  models.Order
	Status   models.Order
	Active   models.Order

	CreatedAt models.Order
	UpdatedAt models.Order
}

func (ol *ListOrder) PrepareQuery(query *gorm.DB) *gorm.DB {
	if ol == nil {
		return query
	}
	query = ol.ID.PrepareQuery(query, `id`)
	query = ol.Type.PrepareQuery(query, `type`)
	query = ol.Platform.PrepareQuery(query, `platform`)
	query = ol.Premium.PrepareQuery(query, `premium`)
	query = ol.Status.PrepareQuery(query, `status`)
	query = ol.Active.PrepareQuery(query, `active`)
	query = ol.CreatedAt.PrepareQuery(query, `created_at`)
	query = ol.UpdatedAt.PrepareQuery(query, `updated_at`)
	return query
}

type (
	Option  = repository.QOption
	Options = repository.ListOptions
)
