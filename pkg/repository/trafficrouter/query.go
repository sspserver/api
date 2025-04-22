package trafficrouter

import (
	"github.com/geniusrabbit/blaze-api/repository"
	"github.com/sspserver/api/pkg/models"
	"gorm.io/gorm"
)

// Filter of the objects list
type Filter struct {
	ID              []uint64
	AccountID       uint64
	Active          *models.ActiveStatus
	RTBSourceIDs    []uint64
	Formats         []string
	DeviceTypes     []uint64
	Devices         []uint64
	OS              []uint64
	Browsers        []uint64
	Carriers        []uint64
	Categories      []uint64
	Countries       []string
	Languages       []string
	Domains         []string
	Applications    []uint64
	Zones           []uint64
	Secure          int
	AdBlock         int
	PrivateBrowsing int
	IP              int
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
	if fl.Active != nil {
		query = query.Where(`active = ?`, *fl.Active)
	}
	if len(fl.RTBSourceIDs) > 0 {
		query = query.Where(`rtb_source_ids IN (?)`, fl.RTBSourceIDs)
	}
	if len(fl.Formats) > 0 {
		query = query.Where(`formats IN (?)`, fl.Formats)
	}
	if len(fl.DeviceTypes) > 0 {
		query = query.Where(`device_types IN (?)`, fl.DeviceTypes)
	}
	if len(fl.Devices) > 0 {
		query = query.Where(`devices IN (?)`, fl.Devices)
	}
	if len(fl.OS) > 0 {
		query = query.Where(`os IN (?)`, fl.OS)
	}
	if len(fl.Browsers) > 0 {
		query = query.Where(`browsers IN (?)`, fl.Browsers)
	}
	if len(fl.Carriers) > 0 {
		query = query.Where(`carriers IN (?)`, fl.Carriers)
	}
	if len(fl.Categories) > 0 {
		query = query.Where(`categories IN (?)`, fl.Categories)
	}
	if len(fl.Countries) > 0 {
		query = query.Where(`countries IN (?)`, fl.Countries)
	}
	if len(fl.Languages) > 0 {
		query = query.Where(`languages IN (?)`, fl.Languages)
	}
	if len(fl.Domains) > 0 {
		query = query.Where(`domains IN (?)`, fl.Domains)
	}
	if len(fl.Applications) > 0 {
		query = query.Where(`apps IN (?)`, fl.Applications)
	}
	if len(fl.Zones) > 0 {
		query = query.Where(`zones IN (?)`, fl.Zones)
	}
	if fl.Secure > 0 {
		query = query.Where(`secure = ?`, fl.Secure)
	}
	if fl.AdBlock > 0 {
		query = query.Where(`adblock = ?`, fl.AdBlock)
	}
	if fl.PrivateBrowsing > 0 {
		query = query.Where(`private_browsing = ?`, fl.PrivateBrowsing)
	}
	if fl.IP > 0 {
		query = query.Where(`ip = ?`, fl.IP)
	}
	return query
}

// ListOrder of the objects list
type ListOrder struct {
	ID        models.Order
	Title     models.Order
	Active    models.Order
	Percent   models.Order
	CreatedAt models.Order
	UpdatedAt models.Order
}

func (ol *ListOrder) PrepareQuery(query *gorm.DB) *gorm.DB {
	if ol == nil {
		return query
	}
	query = ol.ID.PrepareQuery(query, `id`)
	query = ol.Title.PrepareQuery(query, `title`)
	query = ol.Active.PrepareQuery(query, `active`)
	query = ol.Percent.PrepareQuery(query, `percent`)
	query = ol.CreatedAt.PrepareQuery(query, `created_at`)
	query = ol.UpdatedAt.PrepareQuery(query, `updated_at`)
	return query
}

// List select options
type (
	Option  = repository.QOption
	Options = repository.ListOptions
)
