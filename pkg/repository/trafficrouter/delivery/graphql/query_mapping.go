package graphql

import (
	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/adcorelib/admodels/types"

	"github.com/sspserver/api/pkg/repository/trafficrouter"
	"github.com/sspserver/api/pkg/repository/trafficrouter/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func TrafficRouterFilterFromGraphQL(fl *gqlmodels.TrafficRouterListFilter) *trafficrouter.Filter {
	if fl == nil {
		return nil
	}
	return &trafficrouter.Filter{
		ID:              fl.ID,
		AccountID:       gocast.PtrAsValue(fl.AccountID, 0),
		Active:          activeStatusPtrFromGraphQL(fl.Active),
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

func activeStatusPtrFromGraphQL(status *gqlmodels.ActiveStatus) *types.ActiveStatus {
	if status == nil {
		return nil
	}
	st := activeStatusFromGraphQL(*status)
	return &st
}

func activeStatusFromGraphQL(status gqlmodels.ActiveStatus) types.ActiveStatus {
	if status == gqlmodels.ActiveStatusActive {
		return types.StatusActive
	}
	return types.StatusPause
}

func FillTrafficRouterListOrderFromGraphQL(src *gqlmodels.TrafficRouterListOrder, dst *trafficrouter.ListOrder) {
	if src == nil || dst == nil {
		return
	}
	if src.ID != nil {
		dst.ID = src.ID.AsOrder()
	}
	if src.Title != nil {
		dst.Title = src.Title.AsOrder()
	}
	if src.Active != nil {
		dst.Active = src.Active.AsOrder()
	}
	if src.Percent != nil {
		dst.Percent = src.Percent.AsOrder()
	}
	if src.CreatedAt != nil {
		dst.CreatedAt = src.CreatedAt.AsOrder()
	}
	if src.UpdatedAt != nil {
		dst.UpdatedAt = src.UpdatedAt.AsOrder()
	}
}

func TrafficRouterListOrderFromGraphQL(order []*gqlmodels.TrafficRouterListOrder) trafficrouter.ListOrder {
	var out trafficrouter.ListOrder
	for _, item := range order {
		FillTrafficRouterListOrderFromGraphQL(item, &out)
	}
	return out
}

func FillTrafficRouterCreateInputModel(inp gqlmodels.TrafficRouterCreateInput, obj *models.TrafficRouter) {
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

func FillTrafficRouterUpdateInputModel(inp gqlmodels.TrafficRouterUpdateInput, obj *models.TrafficRouter) {
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
