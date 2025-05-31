package models

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"

	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/trafficrouter"
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
		FormatCodes:     m.Formats,
		DeviceTypeIDs:   m.DeviceTypes,
		DeviceIDs:       m.Devices,
		OSIDs:           m.OS,
		BrowserIDs:      m.Browsers,
		CarrierIDs:      m.Carriers,
		CategoryIDs:     m.Categories,
		CountryCodes:    m.Countries,
		LanguageCodes:   m.Languages,
		Domains:         m.Domains,
		ApplicationIDs:  m.Applications,
		ZoneIDs:         m.Zones,
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
		Formats:         fl.FormatCodes,
		DeviceTypes:     fl.DeviceTypeIDs,
		Devices:         fl.DeviceIDs,
		OS:              fl.OSIDs,
		Browsers:        fl.BrowserIDs,
		Carriers:        fl.CarrierIDs,
		Categories:      fl.CategoryIDs,
		Countries:       fl.CountryCodes,
		Languages:       fl.LanguageCodes,
		Domains:         fl.Domains,
		Applications:    fl.ApplicationIDs,
		Zones:           fl.ZoneIDs,
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
	if ord.Title != nil {
		order.Title = ord.Title.AsOrder()
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
	obj.Title = inp.Title
	obj.Description = gocast.PtrAsValue(inp.Description, obj.Description)
	obj.RTBSourceIDs = inp.RTBSourceIDs
	obj.Formats = inp.FormatCodes
	obj.DeviceTypes = inp.DeviceTypeIDs
	obj.Devices = inp.DeviceIDs
	obj.OS = inp.OSIDs
	obj.Browsers = inp.BrowserIDs
	obj.Carriers = inp.CarrierIDs
	obj.Categories = inp.CategoryIDs
	obj.Countries = inp.CountryCodes
	obj.Languages = inp.LanguageCodes
	obj.Domains = inp.Domains
	obj.Applications = inp.ApplicationIDs
	obj.Zones = inp.ZoneIDs
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
	obj.Title = gocast.PtrAsValue(inp.Title, obj.Title)
	obj.Description = gocast.PtrAsValue(inp.Description, obj.Description)
	obj.RTBSourceIDs = gocast.IfThen(inp.RTBSourceIDs != nil, inp.RTBSourceIDs, obj.RTBSourceIDs)
	obj.Formats = gocast.IfThen(inp.FormatCodes != nil, inp.FormatCodes, obj.Formats)
	obj.DeviceTypes = gocast.IfThen(inp.DeviceTypeIDs != nil, inp.DeviceTypeIDs, obj.DeviceTypes)
	obj.Devices = gocast.IfThen(inp.DeviceIDs != nil, inp.DeviceIDs, obj.Devices)
	obj.OS = gocast.IfThen(inp.OSIDs != nil, inp.OSIDs, obj.OS)
	obj.Browsers = gocast.IfThen(inp.BrowserIDs != nil, inp.BrowserIDs, obj.Browsers)
	obj.Carriers = gocast.IfThen(inp.CarrierIDs != nil, inp.CarrierIDs, obj.Carriers)
	obj.Categories = gocast.IfThen(inp.CategoryIDs != nil, inp.CategoryIDs, obj.Categories)
	obj.Countries = gocast.IfThen(inp.CountryCodes != nil, inp.CountryCodes, obj.Countries)
	obj.Languages = gocast.IfThen(inp.LanguageCodes != nil, inp.LanguageCodes, obj.Languages)
	obj.Domains = gocast.IfThen(inp.Domains != nil, inp.Domains, obj.Domains)
	obj.Applications = gocast.IfThen(inp.ApplicationIDs != nil, inp.ApplicationIDs, obj.Applications)
	obj.Zones = gocast.IfThen(inp.ZoneIDs != nil, inp.ZoneIDs, obj.Zones)
	obj.Secure = gocast.PtrAsValue(inp.Secure.IntPtr(), obj.Secure)
	obj.AdBlock = gocast.PtrAsValue(inp.AdBlock.IntPtr(), obj.AdBlock)
	obj.PrivateBrowsing = gocast.PtrAsValue(inp.PrivateBrowsing.IntPtr(), obj.PrivateBrowsing)
	obj.IP = gocast.PtrAsValue(inp.IP.IntPtr(), obj.IP)
}
