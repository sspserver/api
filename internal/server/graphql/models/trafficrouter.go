package models

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"

	"github.com/sspserver/api/internal/repository/trafficrouter"
	"github.com/sspserver/api/models"
)

func FromTrafficRouterModel(m *models.TrafficRouter) *TrafficRouter {
	if m == nil {
		return nil
	}
	return &TrafficRouter{
		ID:          m.ID,
		AccountID:   m.AccountID,
		Title:       m.Title,
		Percent:     m.Percent,
		Description: m.Description,

		Active: FromActiveStatus(m.Active),

		RTBSourceIDs: m.RTBSourceIDs,

		// Targeting filters
		Formats:         m.Formats,
		DeviceTypes:     m.DeviceTypes,
		Devices:         m.Devices,
		Os:              m.OS,
		Browsers:        m.Browsers,
		Carriers:        m.Carriers,
		Categories:      m.Categories,
		Countries:       m.Countries,
		Languages:       m.Languages,
		Domains:         m.Domains,
		Applications:    m.Applications,
		Zones:           m.Zones,
		Secure:          FromAnyOnlyExclude(m.Secure),
		AdBlock:         FromAnyOnlyExclude(m.AdBlock),
		PrivateBrowsing: FromAnyOnlyExclude(m.PrivateBrowsing),
		IP:              FromAnyIPv4IPv6(m.IP),
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

func FromTrafficRouterModelList(m []*models.TrafficRouter) []*TrafficRouter {
	return xtypes.SliceApply(m, FromTrafficRouterModel)
}

func (fl *TrafficRouterListFilter) Filter() *trafficrouter.Filter {
	if fl == nil {
		return nil
	}
	return &trafficrouter.Filter{
		ID:              fl.ID,
		AccountID:       gocast.PtrAsValue(fl.AccountID, 0),
		Active:          ActiveStatusPtr(fl.Active),
		RTBSourceIDs:    fl.RTBSourceIDs,
		Formats:         fl.Formats,
		DeviceTypes:     fl.DeviceTypes,
		Devices:         fl.Devices,
		OS:              fl.Os,
		Browsers:        fl.Browsers,
		Carriers:        fl.Carriers,
		Categories:      fl.Categories,
		Countries:       fl.Countries,
		Languages:       fl.Languages,
		Domains:         fl.Domains,
		Applications:    fl.Applications,
		Zones:           fl.Zones,
		Secure:          gocast.PtrAsValue(fl.Secure.IntPtr(), 0),
		AdBlock:         gocast.PtrAsValue(fl.AdBlock.IntPtr(), 0),
		PrivateBrowsing: gocast.PtrAsValue(fl.PrivateBrowsing.IntPtr(), 0),
		IP:              gocast.PtrAsValue(fl.IP.IntPtr(), 0),
	}
}

func (ord *TrafficRouterListOrder) Order() *trafficrouter.ListOrder {
	if ord == nil {
		return nil
	}
	return &trafficrouter.ListOrder{
		ID:        ord.ID.AsOrder(),
		Active:    ord.Active.AsOrder(),
		Percent:   ord.Percent.AsOrder(),
		CreatedAt: ord.CreatedAt.AsOrder(),
		UpdatedAt: ord.UpdatedAt.AsOrder(),
	}
}

func (ord *TrafficRouterListOrder) Fill(order *trafficrouter.ListOrder) {
	if ord == nil || order == nil {
		return
	}
	if ord.ID != nil {
		order.ID = ord.ID.AsOrder()
	}
	if ord.Active != nil {
		order.Active = ord.Active.AsOrder()
	}
	if ord.Percent != nil {
		order.Percent = ord.Percent.AsOrder()
	}
	if ord.CreatedAt != nil {
		order.CreatedAt = ord.CreatedAt.AsOrder()
	}
	if ord.UpdatedAt != nil {
		order.UpdatedAt = ord.UpdatedAt.AsOrder()
	}
}

func (inp *TrafficRouterCreateInput) FillModel(obj *models.TrafficRouter) {
	if obj == nil {
		return
	}
	obj.AccountID = gocast.PtrAsValue(inp.AccountID, obj.AccountID)
	obj.Percent = inp.Percent
	obj.Description = gocast.PtrAsValue(inp.Description, obj.Description)
	obj.Active = ActiveStatusFrom(inp.Active)
	obj.RTBSourceIDs = inp.RTBSourceIDs
	obj.Formats = inp.Formats
	obj.DeviceTypes = inp.DeviceTypes
	obj.Devices = inp.Devices
	obj.OS = inp.Os
	obj.Browsers = inp.Browsers
	obj.Carriers = inp.Carriers
	obj.Categories = inp.Categories
	obj.Countries = inp.Countries
	obj.Languages = inp.Languages
	obj.Domains = inp.Domains
	obj.Applications = inp.Applications
	obj.Zones = inp.Zones
	obj.Secure = inp.Secure.Int()
	obj.AdBlock = inp.AdBlock.Int()
	obj.PrivateBrowsing = inp.PrivateBrowsing.Int()
	obj.IP = inp.IP.Int()
}

func (inp *TrafficRouterUpdateInput) FillModel(obj *models.TrafficRouter) {
	if obj == nil {
		return
	}
	obj.Percent = gocast.PtrAsValue(inp.Percent, obj.Percent)
	obj.Description = gocast.PtrAsValue(inp.Description, obj.Description)
	obj.Active = gocast.PtrAsValue(ActiveStatusPtr(inp.Active), obj.Active)
	obj.RTBSourceIDs = gocast.IfThen(inp.RTBSourceIDs != nil, inp.RTBSourceIDs, obj.RTBSourceIDs)
	obj.Formats = gocast.IfThen(inp.Formats != nil, inp.Formats, obj.Formats)
	obj.DeviceTypes = gocast.IfThen(inp.DeviceTypes != nil, inp.DeviceTypes, obj.DeviceTypes)
	obj.Devices = gocast.IfThen(inp.Devices != nil, inp.Devices, obj.Devices)
	obj.OS = gocast.IfThen(inp.Os != nil, inp.Os, obj.OS)
	obj.Browsers = gocast.IfThen(inp.Browsers != nil, inp.Browsers, obj.Browsers)
	obj.Carriers = gocast.IfThen(inp.Carriers != nil, inp.Carriers, obj.Carriers)
	obj.Categories = gocast.IfThen(inp.Categories != nil, inp.Categories, obj.Categories)
	obj.Countries = gocast.IfThen(inp.Countries != nil, inp.Countries, obj.Countries)
	obj.Languages = gocast.IfThen(inp.Languages != nil, inp.Languages, obj.Languages)
	obj.Domains = gocast.IfThen(inp.Domains != nil, inp.Domains, obj.Domains)
	obj.Applications = gocast.IfThen(inp.Applications != nil, inp.Applications, obj.Applications)
	obj.Zones = gocast.IfThen(inp.Zones != nil, inp.Zones, obj.Zones)
	obj.Secure = gocast.PtrAsValue(inp.Secure.IntPtr(), obj.Secure)
	obj.AdBlock = gocast.PtrAsValue(inp.AdBlock.IntPtr(), obj.AdBlock)
	obj.PrivateBrowsing = gocast.PtrAsValue(inp.PrivateBrowsing.IntPtr(), obj.PrivateBrowsing)
	obj.IP = gocast.PtrAsValue(inp.IP.IntPtr(), obj.IP)
}
