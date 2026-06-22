package graphql

import (
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/server/graphql/types"

	"github.com/sspserver/api/pkg/repository/rtbsource/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FromRTBSourceModel(src *models.RTBSource) *gqlmodels.RTBSource {
	return &gqlmodels.RTBSource{
		ID:          src.ID,
		AccountID:   src.AccountID,
		Title:       src.Title,
		Description: src.Description,

		Status: gqlmodels.FromApproveStatus(src.Status),
		Active: gqlmodels.FromActiveStatus(src.Active),

		Flags:         *types.MustNullableJSONFrom(src.Flags.Data),
		Protocol:      src.Protocol,
		MinimalWeight: src.MinimalWeight,

		URL:         src.URL,
		Method:      src.Method,
		RequestType: gqlmodels.FromRTBRequestType(src.RequestType),
		Headers:     *types.MustNullableJSONFrom(src.Headers.DataOr(nil)),
		Rps:         src.RPS,
		Timeout:     src.Timeout,

		Accuracy:              src.Accuracy,
		PriceCorrectionReduce: src.PriceCorrectionReduce,
		AuctionType:           gqlmodels.FromAuctionType(src.AuctionType),
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
		Secure:          gqlmodels.FromAnyOnlyExclude(src.Secure),
		AdBlock:         gqlmodels.FromAnyOnlyExclude(src.AdBlock),
		PrivateBrowsing: gqlmodels.FromAnyOnlyExclude(src.PrivateBrowsing),
		IP:              gqlmodels.FromAnyIPv4IPv6(src.IP),

		Config: *types.MustNullableJSONFrom(src.Config.DataOr(nil)),

		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
		DeletedAt: gqlmodels.DeletedAt(src.DeletedAt),
	}
}

func FromRTBSourceModelList(src []*models.RTBSource) []*gqlmodels.RTBSource {
	return xtypes.SliceApply(src, FromRTBSourceModel)
}
