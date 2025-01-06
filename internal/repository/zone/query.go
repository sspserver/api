package zone

import (
	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/blaze-api/repository"
	"gorm.io/gorm"

	"github.com/sspserver/api/models"
)

// Filter of the objects list
type Filter struct {
	ID        []uint64
	Codename  []string
	AccountID []uint64

	Type   *types.ZoneType
	Status *types.ApproveStatus
	Active *types.ActiveStatus

	MinECPM *float64
	MaxECPM *float64
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
	if len(fl.AccountID) > 0 {
		query = query.Where(`account_id IN (?)`, fl.AccountID)
	}
	if fl.Type != nil {
		query = query.Where(`type = ?`, *fl.Type)
	}
	if fl.Status != nil {
		query = query.Where(`status = ?`, *fl.Status)
	}
	if fl.Active != nil {
		query = query.Where(`active = ?`, *fl.Active)
	}
	if fl.MinECPM != nil && fl.MaxECPM != nil {
		query = query.Where(`min_ecpm BETWEEN ? AND ?`, *fl.MinECPM, *fl.MaxECPM)
	} else if fl.MinECPM != nil {
		query = query.Where(`min_ecpm >= ?`, *fl.MinECPM)
	} else if fl.MaxECPM != nil {
		query = query.Where(`min_ecpm <= ?`, *fl.MaxECPM)
	}
	return query
}

// ListOrder of the objects list
type ListOrder struct {
	ID       models.Order
	Codename models.Order

	Title     models.Order
	AccountID models.Order

	Type   models.Order
	Status models.Order
	Active models.Order

	MinECPM models.Order

	CreatedAt models.Order
	UpdatedAt models.Order
}

func (ol *ListOrder) PrepareQuery(query *gorm.DB) *gorm.DB {
	if ol == nil {
		return query
	}
	query = ol.ID.PrepareQuery(query, `id`)
	query = ol.Codename.PrepareQuery(query, `codename`)

	query = ol.Title.PrepareQuery(query, `title`)
	query = ol.AccountID.PrepareQuery(query, `account_id`)

	query = ol.Type.PrepareQuery(query, `type`)
	query = ol.Status.PrepareQuery(query, `status`)
	query = ol.Active.PrepareQuery(query, `active`)

	query = ol.MinECPM.PrepareQuery(query, `min_ecpm`)

	query = ol.CreatedAt.PrepareQuery(query, `created_at`)
	query = ol.UpdatedAt.PrepareQuery(query, `updated_at`)
	return query
}

type (
	Option  = repository.QOption
	Options = repository.ListOptions
)
