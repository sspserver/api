package repository

import (
	"bytes"
	"strings"
	"time"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/billing"

	"github.com/sspserver/api/models"
)

type AggregatedCountersLocal struct {
	Datemark     time.Time `db:"datemark" json:"datemark"`
	Timemark     time.Time `db:"timemark" json:"timemark"`
	SourceID     uint64    `db:"source_id" json:"source_id"`
	PlatformType uint8     `db:"platform_type" json:"platform_type"`
	Domain       string    `db:"domain" json:"domain"`
	AppID        uint64    `db:"app_id" json:"app_id"`
	ZoneID       uint64    `db:"zone_id" json:"zone_id"`
	FormatID     uint32    `db:"format_id" json:"format_id"`
	// Targeting
	CarrierID  uint64 `db:"carrier_id" json:"carrier_id"`
	Country    string `db:"country" json:"country"`
	Language   string `db:"language" json:"language"`
	DeviceID   uint32 `db:"device_id" json:"device_id"`
	DeviceType uint32 `db:"device_type" json:"device_type"`
	OSID       uint32 `db:"os_id" json:"os_id"`
	BrowserID  uint32 `db:"browser_id" json:"browser_id"`
	// Money
	PricingModel       uint8         `db:"pricing_model" json:"pricing_model"`
	PotentialRevenue   billing.Money `db:"potential_revenue" json:"potential_revenue"`
	FailedRevenue      billing.Money `db:"failed_revenue" json:"failed_revenue"`
	CompromisedRevenue billing.Money `db:"compromised_revenue" json:"compromised_revenue"`
	Revenue            billing.Money `db:"revenue" json:"revenue"`
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

	SrcBidRequests uint64 `db:"src_bid_requests" json:"src_bid_requests"`
	SrcBidWins     uint64 `db:"src_bid_wins" json:"src_bid_wins"`
	SrcBidSkips    uint64 `db:"src_bid_skips" json:"src_bid_skips"`
	SrcBidNobids   uint64 `db:"src_bid_nobids" json:"src_bid_nobids"`
	SrcBidErrors   uint64 `db:"src_bid_errors" json:"src_bid_errors"`

	Adblocks uint64 `db:"adblocks" json:"adblocks"`
	Privates uint64 `db:"privates" json:"privates"`
	Robots   uint64 `db:"robots" json:"robots"`
	Backups  uint64 `db:"backups" json:"backups"`
}

func (m *AggregatedCountersLocal) TableName() string {
	return "stats.aggregated_counters_local"
}

func (m *AggregatedCountersLocal) AsStatisticItem() *models.StatisticAdItem {
	allKeys := []models.StatisticAdItemKey{
		{Key: "datemark", Value: m.Datemark},
		{Key: "timemark", Value: m.Timemark},
		{Key: "source_id", Value: m.SourceID},
		{Key: "platform_type", Value: m.PlatformType},
		{Key: "domain", Value: m.Domain},
		{Key: "app_id", Value: m.AppID},
		{Key: "zone_id", Value: m.ZoneID},
		{Key: "format_id", Value: m.FormatID},
		{Key: "carrier_id", Value: m.CarrierID},
		{Key: "country", Value: m.Country},
		{Key: "language", Value: m.Language},
		{Key: "device_type", Value: m.DeviceType},
		{Key: "device_id", Value: m.DeviceID},
		{Key: "os_id", Value: m.OSID},
		{Key: "browser_id", Value: m.BrowserID},
	}
	return &models.StatisticAdItem{
		Keys: xtypes.Slice[models.StatisticAdItemKey](allKeys).
			Apply(func(val models.StatisticAdItemKey) models.StatisticAdItemKey {
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
			Filter(func(g models.StatisticAdItemKey) bool {
				return !gocast.IsEmpty(g.Value)
			}),
		Profit:      m.Revenue.Float64(),
		BidPrice:    m.PotentialRevenue.Float64(),
		Requests:    m.Imps + m.Views + m.Directs + m.Clicks,
		Impressions: m.Imps,
		Views:       m.Views,
		Directs:     m.Directs,
		Clicks:      m.Clicks,
		Wins:        m.SrcBidWins,
		Bids:        m.SrcBidRequests,
		Skips:       m.SrcBidSkips,
		Nobids:      m.SrcBidNobids,
		Errors:      m.SrcBidErrors,
	}
}
