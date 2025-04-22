package rtbsource

import (
	"github.com/geniusrabbit/blaze-api/repository"
	"gorm.io/gorm"

	"github.com/sspserver/api/pkg/models"
)

// Filter of the objects list
type Filter struct {
	ID        []uint64
	AccountID uint64
	Protocol  []string
	Status    *models.ApproveStatus
	Active    *models.ActiveStatus
	// Request
	Method      []string
	RequestType []models.RTBRequestType
	// Auction
	AuctionType []models.AuctionType
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
	if len(fl.Protocol) > 0 {
		query = query.Where(`protocol IN (?)`, fl.Protocol)
	}
	if fl.Status != nil {
		query = query.Where(`status = ?`, *fl.Status)
	}
	if fl.Active != nil {
		query = query.Where(`active = ?`, *fl.Active)
	}
	if len(fl.Method) > 0 {
		query = query.Where(`method IN (?)`, fl.Method)
	}
	if len(fl.RequestType) > 0 {
		query = query.Where(`request_type IN (?)`, fl.RequestType)
	}
	return query
}

// ListOrder of the objects list
type ListOrder struct {
	ID        models.Order
	Title     models.Order
	AccountID models.Order
	Protocol  models.Order
	Status    models.Order
	Active    models.Order

	// Request
	Method      models.Order
	RequestType models.Order

	// Auction
	AuctionType models.Order

	CreatedAt models.Order
	UpdatedAt models.Order
	DeletedAt models.Order
}

func (ol *ListOrder) PrepareQuery(query *gorm.DB) *gorm.DB {
	if ol == nil {
		return query
	}
	query = ol.ID.PrepareQuery(query, `id`)
	query = ol.Title.PrepareQuery(query, `title`)
	query = ol.AccountID.PrepareQuery(query, `account_id`)
	query = ol.Protocol.PrepareQuery(query, `protocol`)
	query = ol.Status.PrepareQuery(query, `status`)
	query = ol.Active.PrepareQuery(query, `active`)

	query = ol.Method.PrepareQuery(query, `method`)
	query = ol.RequestType.PrepareQuery(query, `request_type`)

	query = ol.AuctionType.PrepareQuery(query, `auction_type`)

	query = ol.CreatedAt.PrepareQuery(query, `created_at`)
	query = ol.UpdatedAt.PrepareQuery(query, `updated_at`)
	query = ol.DeletedAt.PrepareQuery(query, `deleted_at`)
	return query
}

// List select options
type (
	Option  = repository.QOption
	Options = repository.ListOptions
)
