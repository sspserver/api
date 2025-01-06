package repository

import (
	"bytes"
	"net"
	"strings"
	"time"

	"github.com/geniusrabbit/archivarius/internal/archivarius"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/billing"
)

// AggregatedCountersLocal is a model for aggregated counters
type AggregatedCountersLocal struct {
	Datemark      time.Time `db:"datemark" json:"datemark"`
	Timemark      time.Time `db:"timemark" json:"timemark"`
	Cluster       string    `db:"cluster" json:"cluster"`
	ProjectID     uint64    `db:"project_id" json:"project_id"`
	PubAccountID  uint64    `db:"pub_account_id" json:"pub_account_id"`
	AdvAccountID  uint64    `db:"adv_account_id" json:"adv_account_id"`
	SourceID      uint64    `db:"source_id" json:"source_id"`
	AccessPointID uint64    `db:"access_point_id" json:"access_point_id"`
	Platform      uint8     `db:"platform" json:"platform"`
	Domain        string    `db:"domain" json:"domain"`
	AppID         uint64    `db:"app_id" json:"app_id"`
	ZoneID        uint64    `db:"zone_id" json:"zone_id"`
	CampaignID    uint64    `db:"campaign_id" json:"campaign_id"`
	AdID          uint64    `db:"ad_id" json:"ad_id"`
	FormatID      uint32    `db:"format_id" json:"format_id"`
	JumperID      uint64    `db:"jumper_id" json:"jumper_id"`
	// Targeting
	CarrierID  uint64 `db:"carrier_id" json:"carrier_id"`
	Country    string `db:"country" json:"country"`
	City       string `db:"city" json:"city"`
	Language   string `db:"language" json:"language"`
	IP         net.IP `db:"ip" json:"ip"` // Assuming IP could be NULL
	DeviceType uint32 `db:"device_type" json:"device_type"`
	DeviceID   uint32 `db:"device_id" json:"device_id"`
	OSID       uint32 `db:"os_id" json:"os_id"`
	BrowserID  uint32 `db:"browser_id" json:"browser_id"`
	// Money
	PricingModel        uint8         `db:"pricing_model" json:"pricing_model"`
	AdvSpend            billing.Money `db:"adv_spend" json:"adv_spend"`
	AdvPotentialSpend   billing.Money `db:"adv_potential_spend" json:"adv_potential_spend"`
	AdvFailedSpend      billing.Money `db:"adv_failed_spend" json:"adv_failed_spend"`
	AdvCompromisedSpend billing.Money `db:"adv_compromised_spend" json:"adv_compromised_spend"`
	PubRevenue          billing.Money `db:"pub_revenue" json:"pub_revenue"`
	SalesBudget         billing.Money `db:"sales_budget" json:"sales_budget"`
	BuyingBudget        billing.Money `db:"buying_budget" json:"buying_budget"`
	BidPrice            billing.Money `db:"bid_price" json:"bid_price"`
	NetworkRevenue      billing.Money `db:"network_revenue" json:"network_revenue"`
	// Counters
	Imps               uint64 `db:"imps" json:"imps"`
	SuccessImps        uint64 `db:"success_imps" json:"success_imps"`
	FailedImps         uint64 `db:"failed_imps" json:"failed_imps"`
	CompromisedImps    uint64 `db:"compromised_imps" json:"compromised_imps"`
	CustomImps         uint64 `db:"custom_imps" json:"custom_imps"`
	BackupImps         uint64 `db:"backup_imps" json:"backup_imps"`
	Views              uint64 `db:"views" json:"views"`
	FailedViews        uint64 `db:"failed_views" json:"failed_views"`
	CompromisedViews   uint64 `db:"compromised_views" json:"compromised_views"`
	CustomViews        uint64 `db:"custom_views" json:"custom_views"`
	BackupViews        uint64 `db:"backup_views" json:"backup_views"`
	Directs            uint64 `db:"directs" json:"directs"`
	SuccessDirects     uint64 `db:"success_directs" json:"success_directs"`
	FailedDirects      uint64 `db:"failed_directs" json:"failed_directs"`
	CompromisedDirects uint64 `db:"compromised_directs" json:"compromised_directs"`
	CustomDirects      uint64 `db:"custom_directs" json:"custom_directs"`
	BackupDirects      uint64 `db:"backup_directs" json:"backup_directs"`
	Clicks             uint64 `db:"clicks" json:"clicks"`
	FailedClicks       uint64 `db:"failed_clicks" json:"failed_clicks"`
	CompromisedClicks  uint64 `db:"compromised_clicks" json:"compromised_clicks"`
	CustomClicks       uint64 `db:"custom_clicks" json:"custom_clicks"`
	BackupClicks       uint64 `db:"backup_clicks" json:"backup_clicks"`
	Leads              uint64 `db:"leads" json:"leads"`
	SuccessLeads       uint64 `db:"success_leads" json:"success_leads"`
	FailedLeads        uint64 `db:"failed_leads" json:"failed_leads"`
	CompromisedLeads   uint64 `db:"compromised_leads" json:"compromised_leads"`

	SrcBidRequests uint64  `db:"src_bid_requests" json:"src_bid_requests"`
	SrcBidWins     uint64  `db:"src_bid_wins" json:"src_bid_wins"`
	SrcBidSkips    uint64  `db:"src_bid_skips" json:"src_bid_skips"`
	SrcBidNobids   uint64  `db:"src_bid_nobids" json:"src_bid_nobids"`
	SrcBidErrors   uint64  `db:"src_bid_errors" json:"src_bid_errors"`
	ApBidRequests  uint64  `db:"ap_bid_requests" json:"ap_bid_requests"`
	ApBidWins      uint64  `db:"ap_bid_wins" json:"ap_bid_wins"`
	ApBidSkips     uint64  `db:"ap_bid_skips" json:"ap_bid_skips"`
	ApBidNobids    uint64  `db:"ap_bid_nobids" json:"ap_bid_nobids"`
	ApBidErrors    uint64  `db:"ap_bid_errors" json:"ap_bid_errors"`
	Wins           uint64  `db:"wins" json:"wins"`
	Bids           uint64  `db:"bids" json:"bids"`
	Skips          uint64  `db:"skips" json:"skips"`
	Nobids         uint64  `db:"nobids" json:"nobids"`
	Errors         uint64  `db:"errors" json:"errors"`
	CTR            float64 `db:"ctr" json:"ctr"`
	ECPM           float64 `db:"ecpm" json:"ecpm"`
	ECPC           float64 `db:"ecpc" json:"ecpc"`
	ECPA           float64 `db:"ecpa" json:"ecpa"`
	Adblocks       uint64  `db:"adblocks" json:"adblocks"`
	Privates       uint64  `db:"privates" json:"privates"`
	Robots         uint64  `db:"robots" json:"robots"`
	Backups        uint64  `db:"backups" json:"backups"`
}

// TableName returns the table name for the model
func (m *AggregatedCountersLocal) TableName() string {
	return "stats.aggregated_counters_local"
}

// AsItem converts aggregated counter record to archivarius.AdItem
func (m *AggregatedCountersLocal) AsItem() *archivarius.AdItem {
	allKeys := []archivarius.AdItemKey{
		{Key: archivarius.KeyDatemark, Value: m.Datemark},
		{Key: archivarius.KeyTimemark, Value: m.Timemark},
		{Key: archivarius.KeyCluster, Value: m.Cluster},
		{Key: archivarius.KeyProjectID, Value: m.ProjectID},
		{Key: archivarius.KeyPubAccountID, Value: m.PubAccountID},
		{Key: archivarius.KeyAdvAccountID, Value: m.AdvAccountID},
		{Key: archivarius.KeySourceID, Value: m.SourceID},
		{Key: archivarius.KeyAccessPointID, Value: m.AccessPointID},
		{Key: archivarius.KeyPlatform, Value: m.Platform},
		{Key: archivarius.KeyDomain, Value: m.Domain},
		{Key: archivarius.KeyAppID, Value: m.AppID},
		{Key: archivarius.KeyZoneID, Value: m.ZoneID},
		{Key: archivarius.KeyCampaignID, Value: m.CampaignID},
		{Key: archivarius.KeyAdID, Value: m.AdID},
		{Key: archivarius.KeyFormatID, Value: m.FormatID},
		{Key: archivarius.KeyJumperID, Value: m.JumperID},
		{Key: archivarius.KeyCarrierID, Value: m.CarrierID},
		{Key: archivarius.KeyCountry, Value: m.Country},
		{Key: archivarius.KeyCity, Value: m.City},
		{Key: archivarius.KeyLanguage, Value: m.Language},
		{Key: archivarius.KeyIP, Value: m.IP},
		{Key: archivarius.KeyDeviceType, Value: m.DeviceType},
		{Key: archivarius.KeyDeviceID, Value: m.DeviceID},
		{Key: archivarius.KeyOSID, Value: m.OSID},
		{Key: archivarius.KeyBrowserID, Value: m.BrowserID},
	}

	return &archivarius.AdItem{
		Keys: xtypes.Slice[archivarius.AdItemKey](allKeys).
			Apply(func(val archivarius.AdItemKey) archivarius.AdItemKey {
				switch v := val.Value.(type) {
				case []byte:
					val.Value = bytes.Trim(v, "\u0000 \n\t")
				case string:
					val.Value = strings.Trim(v, "\u0000 \n\t")
				case time.Time:
					if v.IsZero() {
						val.Value = nil
					}
				}
				return val
			}).
			// Filter only filled keys
			Filter(func(g archivarius.AdItemKey) bool {
				return !gocast.IsEmpty(g.Value)
			}),
		Spent:       m.AdvSpend.Float64(),
		Profit:      m.PubRevenue.Float64(),
		BidPrice:    m.BidPrice.Float64(),
		Impressions: m.Imps,
		Views:       m.Views,
		Directs:     m.Directs,
		Clicks:      m.Clicks,
		Leads:       m.Leads,
		Wins:        m.Wins,
		Bids:        m.Bids,
		Skips:       m.Skips,
		Nobids:      m.Nobids,
		Errors:      m.Errors,
		CTR:         m.CTR,
		ECPM:        m.ECPM,
		ECPC:        m.ECPC,
		ECPA:        m.ECPA,
	}
}
