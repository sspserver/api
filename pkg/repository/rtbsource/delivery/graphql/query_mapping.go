package graphql

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/admodels/types"
	bzgqlmodels "github.com/geniusrabbit/blaze-api/server/graphql/models"
	"github.com/geniusrabbit/gosql/v2"

	"github.com/sspserver/api/pkg/repository/rtbsource"
	rtbmodels "github.com/sspserver/api/pkg/repository/rtbsource/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func RTBSourceFilterFromGraphQL(fl *gqlmodels.RTBSourceListFilter) *rtbsource.Filter {
	if fl == nil {
		return nil
	}
	return &rtbsource.Filter{
		ID:        fl.ID,
		AccountID: gocast.PtrAsValue(fl.AccountID, 0),
		Protocol:  fl.Protocol,
		Status:    approveStatusPtrFromGraphQL(fl.Status),
		Active:    activeStatusPtrFromGraphQL(fl.Active),
		Method:    fl.Method,
		RequestType: xtypes.SliceApply(fl.RequestType,
			func(t gqlmodels.RTBRequestFormatType) types.RTBRequestType { return rtbRequestTypeFromGraphQL(t) }),
		AuctionType: xtypes.SliceApply(fl.AuctionType,
			func(t gqlmodels.AuctionType) types.AuctionType { return auctionTypeFromGraphQL(t) }),
	}
}

func rtbRequestTypeFromGraphQL(t gqlmodels.RTBRequestFormatType) types.RTBRequestType {
	switch t {
	case gqlmodels.RTBRequestFormatTypeJSON:
		return types.RTBRequestTypeJSON
	case gqlmodels.RTBRequestFormatTypeXML:
		return types.RTBRequestTypeXML
	}
	return types.RTBRequestTypeUndefined
}

func auctionTypeFromGraphQL(t gqlmodels.AuctionType) types.AuctionType {
	switch t {
	case gqlmodels.AuctionTypeFirstPrice:
		return types.FirstPriceAuctionType
	case gqlmodels.AuctionTypeSecondPrice:
		return types.SecondPriceAuctionType
	}
	return types.UndefinedAuctionType
}

func approveStatusPtrFromGraphQL(status *bzgqlmodels.ApproveStatus) *types.ApproveStatus {
	if status == nil {
		return nil
	}
	st := approveStatusFromGraphQL(*status)
	return &st
}

func approveStatusFromGraphQL(status bzgqlmodels.ApproveStatus) types.ApproveStatus {
	if status == bzgqlmodels.ApproveStatusApproved {
		return types.StatusApproved
	}
	if status == bzgqlmodels.ApproveStatusRejected {
		return types.StatusRejected
	}
	return types.StatusPending
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

func FillRTBSourceListOrderFromGraphQL(src *gqlmodels.RTBSourceListOrder, dst *rtbsource.ListOrder) {
	if src == nil || dst == nil {
		return
	}
	if src.ID != nil {
		dst.ID = src.ID.AsOrder()
	}
	if src.Title != nil {
		dst.Title = src.Title.AsOrder()
	}
	if src.AccountID != nil {
		dst.AccountID = src.AccountID.AsOrder()
	}
	if src.Protocol != nil {
		dst.Protocol = src.Protocol.AsOrder()
	}
	if src.Status != nil {
		dst.Status = src.Status.AsOrder()
	}
	if src.Active != nil {
		dst.Active = src.Active.AsOrder()
	}
	if src.Method != nil {
		dst.Method = src.Method.AsOrder()
	}
	if src.RequestType != nil {
		dst.RequestType = src.RequestType.AsOrder()
	}
	if src.AuctionType != nil {
		dst.AuctionType = src.AuctionType.AsOrder()
	}
	if src.CreatedAt != nil {
		dst.CreatedAt = src.CreatedAt.AsOrder()
	}
	if src.UpdatedAt != nil {
		dst.UpdatedAt = src.UpdatedAt.AsOrder()
	}
}

func RTBSourceListOrderFromGraphQL(order []*gqlmodels.RTBSourceListOrder) rtbsource.ListOrder {
	var out rtbsource.ListOrder
	for _, item := range order {
		FillRTBSourceListOrderFromGraphQL(item, &out)
	}
	return out
}

func FillRTBSourceCreateInputModel(inp gqlmodels.RTBSourceCreateInput, m *rtbmodels.RTBSource) {
	if m == nil {
		return
	}
	*m = rtbmodels.RTBSource{
		AccountID:   gocast.PtrAsValue(inp.AccountID, 0),
		Title:       inp.Title,
		Description: gocast.PtrAsValue(inp.Description, ""),

		Flags: gocast.IfThenExec(inp.Flags != nil,
			func() gosql.NullableJSON[rtbmodels.RTBSourceFlags] {
				return *gosql.MustNullableJSON[rtbmodels.RTBSourceFlags](inp.Flags.DataOr(nil))
			},
			func() gosql.NullableJSON[rtbmodels.RTBSourceFlags] {
				return gosql.NullableJSON[rtbmodels.RTBSourceFlags]{}
			}),
		Protocol:      inp.Protocol,
		MinimalWeight: gocast.PtrAsValue(inp.MinimalWeight, 0),

		URL:    inp.URL,
		Method: inp.Method,
		RequestType: gocast.IfThenExec(rtbRequestTypeFromGraphQL(inp.RequestType) == types.RTBRequestTypeUndefined,
			func() types.RTBRequestType { return types.RTBRequestTypeJSON },
			func() types.RTBRequestType { return rtbRequestTypeFromGraphQL(inp.RequestType) }),
		Headers: gocast.IfThenExec(inp.Headers != nil,
			func() gosql.NullableJSON[map[string]string] {
				return *gosql.MustNullableJSON[map[string]string](inp.Headers.DataOr(nil))
			},
			func() gosql.NullableJSON[map[string]string] { return gosql.NullableJSON[map[string]string]{} }),
		RPS:     inp.Rps,
		Timeout: inp.Timeout,

		Accuracy:              inp.Accuracy,
		PriceCorrectionReduce: inp.PriceCorrectionReduce,
		AuctionType: gocast.IfThenExec(auctionTypeFromGraphQL(inp.AuctionType) == types.UndefinedAuctionType,
			func() types.AuctionType { return types.FirstPriceAuctionType },
			func() types.AuctionType { return auctionTypeFromGraphQL(inp.AuctionType) }),
		MinBid: inp.MinBid,
		MaxBid: inp.MaxBid,

		Formats:         inp.FormatCodes,
		DeviceTypes:     inp.DeviceTypeIDs,
		Devices:         inp.DeviceIDs,
		OS:              inp.OSIDs,
		Browsers:        inp.BrowserIDs,
		Carriers:        inp.CarrierIDs,
		Categories:      inp.CategoryIDs,
		Countries:       inp.CountryCodes,
		Languages:       inp.LanguageCodes,
		Applications:    inp.ApplicationIDs,
		Domains:         inp.Domains,
		Zones:           inp.ZoneIDs,
		Secure:          gocast.PtrAsValue(inp.Secure.IntPtr(), 0),
		AdBlock:         gocast.PtrAsValue(inp.AdBlock.IntPtr(), 0),
		PrivateBrowsing: gocast.PtrAsValue(inp.PrivateBrowsing.IntPtr(), 0),
		IP:              gocast.PtrAsValue(inp.IP.IntPtr(), 0),

		Config: gocast.IfThenExec(inp.Config != nil,
			func() gosql.NullableJSON[any] { return *(*gosql.NullableJSON[any])(inp.Config) },
			func() gosql.NullableJSON[any] { return gosql.NullableJSON[any]{} }),
	}
}

func FillRTBSourceUpdateInputModel(inp gqlmodels.RTBSourceUpdateInput, m *rtbmodels.RTBSource) {
	if m == nil {
		return
	}
	m.AccountID = gocast.PtrAsValue(inp.AccountID, m.AccountID)
	m.Title = gocast.PtrAsValue(inp.Title, m.Title)
	m.Description = gocast.PtrAsValue(inp.Description, m.Description)

	m.Flags = gocast.IfThenExec(inp.Flags != nil,
		func() gosql.NullableJSON[rtbmodels.RTBSourceFlags] {
			return *gosql.MustNullableJSON[rtbmodels.RTBSourceFlags](inp.Flags.DataOr(nil))
		},
		func() gosql.NullableJSON[rtbmodels.RTBSourceFlags] { return m.Flags })
	m.Protocol = gocast.PtrAsValue(inp.Protocol, m.Protocol)
	m.MinimalWeight = gocast.PtrAsValue(inp.MinimalWeight, m.MinimalWeight)

	m.URL = gocast.PtrAsValue(inp.URL, m.URL)
	m.Method = gocast.PtrAsValue(inp.Method, m.Method)
	if inp.RequestType != nil {
		m.RequestType = rtbRequestTypeFromGraphQL(*inp.RequestType)
	}
	m.Headers = gocast.IfThenExec(inp.Headers != nil,
		func() gosql.NullableJSON[map[string]string] {
			return *gosql.MustNullableJSON[map[string]string](inp.Headers.DataOr(nil))
		},
		func() gosql.NullableJSON[map[string]string] { return m.Headers })
	m.RPS = gocast.PtrAsValue(inp.Rps, m.RPS)
	m.Timeout = gocast.PtrAsValue(inp.Timeout, m.Timeout)

	m.Accuracy = gocast.PtrAsValue(inp.Accuracy, m.Accuracy)
	m.PriceCorrectionReduce = gocast.PtrAsValue(inp.PriceCorrectionReduce, m.PriceCorrectionReduce)
	if inp.AuctionType != nil {
		m.AuctionType = auctionTypeFromGraphQL(*inp.AuctionType)
	}
	m.MinBid = gocast.PtrAsValue(inp.MinBid, m.MinBid)
	m.MaxBid = gocast.PtrAsValue(inp.MaxBid, m.MaxBid)

	m.Formats = gocast.IfThen(inp.FormatCodes != nil, inp.FormatCodes, []string(m.Formats))
	m.DeviceTypes = gocast.IfThen(inp.DeviceTypeIDs != nil, inp.DeviceTypeIDs, []uint64(m.DeviceTypes))
	m.Devices = gocast.IfThen(inp.DeviceIDs != nil, inp.DeviceIDs, []uint64(m.Devices))
	m.OS = gocast.IfThen(inp.OSIDs != nil, inp.OSIDs, []uint64(m.OS))
	m.Browsers = gocast.IfThen(inp.BrowserIDs != nil, inp.BrowserIDs, []uint64(m.Browsers))
	m.Carriers = gocast.IfThen(inp.CarrierIDs != nil, inp.CarrierIDs, []uint64(m.Carriers))
	m.Categories = gocast.IfThen(inp.CategoryIDs != nil, inp.CategoryIDs, []uint64(m.Categories))
	m.Countries = gocast.IfThen(inp.CountryCodes != nil, inp.CountryCodes, []string(m.Countries))
	m.Languages = gocast.IfThen(inp.LanguageCodes != nil, inp.LanguageCodes, []string(m.Languages))
	m.Applications = gocast.IfThen(inp.ApplicationIDs != nil, inp.ApplicationIDs, []uint64(m.Applications))
	m.Domains = gocast.IfThen(inp.Domains != nil, inp.Domains, []string(m.Domains))
	m.Zones = gocast.IfThen(inp.ZoneIDs != nil, inp.ZoneIDs, []uint64(m.Zones))
	m.Secure = gocast.PtrAsValue(inp.Secure.IntPtr(), m.Secure)
	m.AdBlock = gocast.PtrAsValue(inp.AdBlock.IntPtr(), m.AdBlock)
	m.PrivateBrowsing = gocast.PtrAsValue(inp.PrivateBrowsing.IntPtr(), m.PrivateBrowsing)
	m.IP = gocast.PtrAsValue(inp.IP.IntPtr(), m.IP)

	m.Config = gocast.IfThenExec(inp.Config != nil,
		func() gosql.NullableJSON[any] { return *(*gosql.NullableJSON[any])(inp.Config) },
		func() gosql.NullableJSON[any] { return m.Config })
}
