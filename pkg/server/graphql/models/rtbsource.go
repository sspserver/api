package models

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/server/graphql/types"
	"github.com/geniusrabbit/gosql/v2"

	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/rtbsource"
)

func FromRTBSourceModel(src *models.RTBSource) *RTBSource {
	return &RTBSource{
		ID:          src.ID,
		AccountID:   src.AccountID,
		Title:       src.Title,
		Description: src.Description,

		Status: FromApproveStatus(src.Status),
		Active: FromActiveStatus(src.Active),

		Flags:         *types.MustNullableJSONFrom(src.Flags.Data),
		Protocol:      src.Protocol,
		MinimalWeight: src.MinimalWeight,

		URL:         src.URL,
		Method:      src.Method,
		RequestType: FromRTBRequestType(src.RequestType),
		Headers:     *types.MustNullableJSONFrom(src.Headers.DataOr(nil)),
		Rps:         src.RPS,
		Timeout:     src.Timeout,

		Accuracy:              src.Accuracy,
		PriceCorrectionReduce: src.PriceCorrectionReduce,
		AuctionType:           FromAuctionType(src.AuctionType),
		MinBid:                src.MinBid,
		MaxBid:                src.MaxBid,

		FormatCodes:     src.Formats,
		DeviceTypeIDs:   src.DeviceTypes,
		DeviceIDs:       src.Devices,
		OSIDs:           src.OS,
		BrowserIDs:      src.Browsers,
		CarrierIDs:      src.Carriers,
		CategoryIDs:     src.Categories,
		CountryCodes:    src.Countries,
		LanguageCodes:   src.Languages,
		ApplicationIDs:  src.Applications,
		Domains:         src.Domains,
		ZoneIDs:         src.Zones,
		Secure:          FromAnyOnlyExclude(src.Secure),
		AdBlock:         FromAnyOnlyExclude(src.AdBlock),
		PrivateBrowsing: FromAnyOnlyExclude(src.PrivateBrowsing),
		IP:              FromAnyIPv4IPv6(src.IP),

		Config: *types.MustNullableJSONFrom(src.Config.DataOr(nil)),

		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
		DeletedAt: DeletedAt(src.DeletedAt),
	}
}

func FromRTBSourceModelList(src []*models.RTBSource) []*RTBSource {
	return xtypes.SliceApply(src, FromRTBSourceModel)
}

func (fl *RTBSourceListFilter) Filter() *rtbsource.Filter {
	if fl == nil {
		return nil
	}
	return &rtbsource.Filter{
		ID:        fl.ID,
		AccountID: gocast.PtrAsValue(fl.AccountID, 0),
		Protocol:  fl.Protocol,
		Status:    ApproveStatusPtr(fl.Status),
		Active:    ActiveStatusPtr(fl.Active),
		Method:    fl.Method,
		RequestType: xtypes.SliceApply(fl.RequestType,
			func(t RTBRequestFormatType) models.RTBRequestType { return t.RequestType() }),
		AuctionType: xtypes.SliceApply(fl.AuctionType,
			func(t AuctionType) models.AuctionType { return t.AuctionType() }),
	}
}

func (ord *RTBSourceListOrder) Order() *rtbsource.ListOrder {
	if ord == nil {
		return nil
	}
	return &rtbsource.ListOrder{
		ID:          ord.ID.AsOrder(),
		Title:       ord.Title.AsOrder(),
		AccountID:   ord.AccountID.AsOrder(),
		Protocol:    ord.Protocol.AsOrder(),
		Status:      ord.Status.AsOrder(),
		Active:      ord.Active.AsOrder(),
		Method:      ord.Method.AsOrder(),
		RequestType: ord.RequestType.AsOrder(),
		AuctionType: ord.AuctionType.AsOrder(),
		CreatedAt:   ord.CreatedAt.AsOrder(),
		UpdatedAt:   ord.UpdatedAt.AsOrder(),
	}
}

func (ord *RTBSourceListOrder) Fill(order *rtbsource.ListOrder) {
	if ord == nil || order == nil {
		return
	}
	if ord.ID != nil {
		order.ID = ord.ID.AsOrder()
	}
	if ord.Title != nil {
		order.Title = ord.Title.AsOrder()
	}
	if ord.AccountID != nil {
		order.AccountID = ord.AccountID.AsOrder()
	}
	if ord.Protocol != nil {
		order.Protocol = ord.Protocol.AsOrder()
	}
	if ord.Status != nil {
		order.Status = ord.Status.AsOrder()
	}
	if ord.Active != nil {
		order.Active = ord.Active.AsOrder()
	}
	if ord.Method != nil {
		order.Method = ord.Method.AsOrder()
	}
	if ord.RequestType != nil {
		order.RequestType = ord.RequestType.AsOrder()
	}
	if ord.AuctionType != nil {
		order.AuctionType = ord.AuctionType.AsOrder()
	}
	if ord.CreatedAt != nil {
		order.CreatedAt = ord.CreatedAt.AsOrder()
	}
	if ord.UpdatedAt != nil {
		order.UpdatedAt = ord.UpdatedAt.AsOrder()
	}
}

func (inp *RTBSourceCreateInput) FillModel(m *models.RTBSource) {
	*m = models.RTBSource{
		AccountID:   gocast.PtrAsValue(inp.AccountID, 0),
		Title:       inp.Title,
		Description: gocast.PtrAsValue(inp.Description, ""),

		Flags: gocast.IfThenExec(inp.Flags != nil,
			func() gosql.NullableJSON[models.RTBSourceFlags] {
				return *gosql.MustNullableJSON[models.RTBSourceFlags](
					inp.Flags.DataOr(nil),
				)
			},
			func() gosql.NullableJSON[models.RTBSourceFlags] {
				return gosql.NullableJSON[models.RTBSourceFlags]{}
			}),
		Protocol:      inp.Protocol,
		MinimalWeight: gocast.PtrAsValue(inp.MinimalWeight, 0),

		URL:         inp.URL,
		Method:      inp.Method,
		RequestType: gocast.PtrAsValue(inp.RequestType.RequestTypePtr(), models.RTBRequestTypeJSON),
		Headers: gocast.IfThenExec(inp.Headers != nil,
			func() gosql.NullableJSON[map[string]string] {
				return *gosql.MustNullableJSON[map[string]string](
					inp.Headers.DataOr(nil),
				)
			},
			func() gosql.NullableJSON[map[string]string] {
				return gosql.NullableJSON[map[string]string]{}
			}),
		RPS:     inp.Rps,
		Timeout: inp.Timeout,

		Accuracy:              inp.Accuracy,
		PriceCorrectionReduce: inp.PriceCorrectionReduce,
		AuctionType:           gocast.PtrAsValue(inp.AuctionType.AuctionTypePtr(), models.AuctionTypeFirstPrice),

		MinBid: inp.MinBid,
		MaxBid: inp.MaxBid,

		// Targeting filters
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
			func() gosql.NullableJSON[any] {
				return *(*gosql.NullableJSON[any])(inp.Config)
			},
			func() gosql.NullableJSON[any] {
				return gosql.NullableJSON[any]{}
			}),
	}
}

func (inp *RTBSourceUpdateInput) FillModel(m *models.RTBSource) {
	m.AccountID = gocast.PtrAsValue(inp.AccountID, m.AccountID)
	m.Title = gocast.PtrAsValue(inp.Title, m.Title)
	m.Description = gocast.PtrAsValue(inp.Description, m.Description)

	m.Flags = gocast.IfThenExec(inp.Flags != nil,
		func() gosql.NullableJSON[models.RTBSourceFlags] {
			return *gosql.MustNullableJSON[models.RTBSourceFlags](
				inp.Flags.DataOr(nil),
			)
		},
		func() gosql.NullableJSON[models.RTBSourceFlags] { return m.Flags })
	m.Protocol = gocast.PtrAsValue(inp.Protocol, m.Protocol)
	m.MinimalWeight = gocast.PtrAsValue(inp.MinimalWeight, m.MinimalWeight)

	m.URL = gocast.PtrAsValue(inp.URL, m.URL)
	m.Method = gocast.PtrAsValue(inp.Method, m.Method)
	m.RequestType = gocast.PtrAsValue(inp.RequestType.RequestTypePtr(), m.RequestType)
	m.Headers = gocast.IfThenExec(inp.Headers != nil,
		func() gosql.NullableJSON[map[string]string] {
			return *gosql.MustNullableJSON[map[string]string](
				inp.Headers.DataOr(nil),
			)
		},
		func() gosql.NullableJSON[map[string]string] { return m.Headers })
	m.RPS = gocast.PtrAsValue(inp.Rps, m.RPS)
	m.Timeout = gocast.PtrAsValue(inp.Timeout, m.Timeout)

	m.Accuracy = gocast.PtrAsValue(inp.Accuracy, m.Accuracy)
	m.PriceCorrectionReduce = gocast.PtrAsValue(inp.PriceCorrectionReduce, m.PriceCorrectionReduce)
	m.AuctionType = gocast.PtrAsValue(inp.AuctionType.AuctionTypePtr(), m.AuctionType)
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
