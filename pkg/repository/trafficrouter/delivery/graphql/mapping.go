package graphql

import (
	"github.com/demdxx/xtypes"

	"github.com/sspserver/api/pkg/repository/trafficrouter/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FromTrafficRouterModel(m *models.TrafficRouter) *gqlmodels.TrafficRouter {
	if m == nil {
		return nil
	}
	return &gqlmodels.TrafficRouter{
		ID:          m.ID,
		AccountID:   m.AccountID,
		Title:       m.Title,
		Percent:     m.Percent,
		Description: m.Description,

		Active: gqlmodels.FromActiveStatus(m.Active),

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
		Secure:          gqlmodels.FromAnyOnlyExclude(m.Secure),
		AdBlock:         gqlmodels.FromAnyOnlyExclude(m.AdBlock),
		PrivateBrowsing: gqlmodels.FromAnyOnlyExclude(m.PrivateBrowsing),
		IP:              gqlmodels.FromAnyIPv4IPv6(m.IP),
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

func FromTrafficRouterModelList(m []*models.TrafficRouter) []*gqlmodels.TrafficRouter {
	return xtypes.SliceApply(m, FromTrafficRouterModel)
}
