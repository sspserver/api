package models

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/server/graphql/types"
	"github.com/geniusrabbit/gosql/v2"

	"github.com/sspserver/api/internal/repository/rtbsource"
	"github.com/sspserver/api/models"
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

		Formats:         src.Formats,
		DeviceTypes:     src.DeviceTypes,
		Devices:         src.Devices,
		Os:              src.OS,
		Browsers:        src.Browsers,
		Carriers:        src.Carriers,
		Categories:      src.Categories,
		Countries:       src.Countries,
		Languages:       src.Languages,
		Applications:    src.Applications,
		Domains:         src.Domains,
		Zones:           src.Zones,
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
	}
}

func (ord *RTBSourceListOrder) Order() *rtbsource.ListOrder {
	if ord == nil {
		return nil
	}
	return &rtbsource.ListOrder{
		Title:     ord.Title.AsOrder(),
		AccountID: ord.AccountID.AsOrder(),
		CreatedAt: ord.CreatedAt.AsOrder(),
		UpdatedAt: ord.UpdatedAt.AsOrder(),
	}
}

func (inp *RTBSourceInput) FillModel(m *models.RTBSource) {
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

	m.Formats = gocast.IfThen(inp.Formats != nil, inp.Formats, []string(m.Formats))
	m.DeviceTypes = gocast.IfThen(inp.DeviceTypes != nil, inp.DeviceTypes, []int64(m.DeviceTypes))
	m.Devices = gocast.IfThen(inp.Devices != nil, inp.Devices, []int64(m.Devices))
	m.OS = gocast.IfThen(inp.Os != nil, inp.Os, []int64(m.OS))
	m.Browsers = gocast.IfThen(inp.Browsers != nil, inp.Browsers, []int64(m.Browsers))
	m.Carriers = gocast.IfThen(inp.Carriers != nil, inp.Carriers, []int64(m.Carriers))
	m.Categories = gocast.IfThen(inp.Categories != nil, inp.Categories, []int64(m.Categories))
	m.Countries = gocast.IfThen(inp.Countries != nil, inp.Countries, []string(m.Countries))
	m.Languages = gocast.IfThen(inp.Languages != nil, inp.Languages, []string(m.Languages))
	m.Applications = gocast.IfThen(inp.Applications != nil, inp.Applications, []int64(m.Applications))
	m.Domains = gocast.IfThen(inp.Domains != nil, inp.Domains, []string(m.Domains))
	m.Zones = gocast.IfThen(inp.Zones != nil, inp.Zones, []int64(m.Zones))
	m.Secure = gocast.PtrAsValue(inp.Secure.IntPtr(), m.Secure)
	m.AdBlock = gocast.PtrAsValue(inp.AdBlock.IntPtr(), m.AdBlock)
	m.PrivateBrowsing = gocast.PtrAsValue(inp.PrivateBrowsing.IntPtr(), m.PrivateBrowsing)
	m.IP = gocast.PtrAsValue(inp.IP.IntPtr(), m.IP)

	m.Config = gocast.IfThenExec(inp.Config != nil,
		func() gosql.NullableJSON[any] { return *(*gosql.NullableJSON[any])(inp.Config) },
		func() gosql.NullableJSON[any] { return m.Config })
}
