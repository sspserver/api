package archivarius

import (
	"context"
)

// OrderingKey is a key for ordering
type OrderingKey string

// List of ordering keys
const (
	OrderingKeyDatemark      OrderingKey = "datemark"
	OrderingKeyTimemark      OrderingKey = "timemark"
	OrderingKeyCluster       OrderingKey = "cluster"
	OrderingKeyProjectID     OrderingKey = "project_id"
	OrderingKeyPubAccountID  OrderingKey = "pub_account_id"
	OrderingKeyAdvAccountID  OrderingKey = "adv_account_id"
	OrderingKeySourceID      OrderingKey = "source_id"
	OrderingKeyAccessPointID OrderingKey = "access_point_id"
	OrderingKeyPlatform      OrderingKey = "platform"
	OrderingKeyDomain        OrderingKey = "domain"
	OrderingKeyAppID         OrderingKey = "app"
	OrderingKeyZoneID        OrderingKey = "zone_id"
	OrderingKeyCampaignID    OrderingKey = "campaign_id"
	OrderingKeyAdID          OrderingKey = "ad_id"
	OrderingKeyFormatID      OrderingKey = "format_id"
	OrderingKeyJumperID      OrderingKey = "jumper_id"
	OrderingKeyCarrierID     OrderingKey = "carrier_id"
	OrderingKeyCountry       OrderingKey = "country"
	OrderingKeyCity          OrderingKey = "city"
	OrderingKeyLanguage      OrderingKey = "language"
	OrderingKeyIP            OrderingKey = "ip"
	OrderingKeyDeviceType    OrderingKey = "device_type"
	OrderingKeyDeviceID      OrderingKey = "device_id"
	OrderingKeyOSID          OrderingKey = "os_id"
	OrderingKeyBrowserID     OrderingKey = "browser_id"
	OrderingKeySpent         OrderingKey = "spent"
	OrderingKeyProfit        OrderingKey = "profit"
	OrderingKeyBidPrice      OrderingKey = "bid_price"
	OrderingKeyRequests      OrderingKey = "requests"
	OrderingKeyImpressions   OrderingKey = "impressions"
	OrderingKeyViews         OrderingKey = "views"
	OrderingKeyDirects       OrderingKey = "directs"
	OrderingKeyClicks        OrderingKey = "clicks"
	OrderingKeyLeads         OrderingKey = "leads"
	OrderingKeyBids          OrderingKey = "bids"
	OrderingKeyWins          OrderingKey = "wins"
	OrderingKeySkips         OrderingKey = "skips"
	OrderingKeyNobids        OrderingKey = "nobids"
	OrderingKeyErrors        OrderingKey = "errors"
	OrderingKeyCTR           OrderingKey = "ctr"
	OrderingKeyECPM          OrderingKey = "ecpm"
	OrderingKeyECPC          OrderingKey = "ecpc"
	OrderingKeyECPA          OrderingKey = "ecpa"
)

// IsGroup checks if key is grouping key
func (k OrderingKey) IsGroup() bool {
	switch k {
	case OrderingKeySpent,
		OrderingKeyProfit,
		OrderingKeyBidPrice,
		OrderingKeyRequests,
		OrderingKeyImpressions,
		OrderingKeyViews,
		OrderingKeyDirects,
		OrderingKeyClicks,
		OrderingKeyLeads,
		OrderingKeyBids,
		OrderingKeyWins,
		OrderingKeySkips,
		OrderingKeyNobids,
		OrderingKeyErrors,
		OrderingKeyCTR,
		OrderingKeyECPM,
		OrderingKeyECPC,
		OrderingKeyECPA:
		return false
	}

	return true
}

// Key is a key for groups and ordering
type Key string

const (
	KeyDatemark      Key = "datemark"
	KeyTimemark      Key = "timemark"
	KeyCluster       Key = "cluster"
	KeyProjectID     Key = "project_id"
	KeyAccountID     Key = "account_id" // SPECIFIC CASE
	KeyPubAccountID  Key = "pub_account_id"
	KeyAdvAccountID  Key = "adv_account_id"
	KeySourceID      Key = "source_id"
	KeyAccessPointID Key = "access_point_id"
	KeyPlatform      Key = "platform"
	KeyDomain        Key = "domain"
	KeyAppID         Key = "app"
	KeyZoneID        Key = "zone_id"
	KeyCampaignID    Key = "campaign_id"
	KeyAdID          Key = "ad_id"
	KeyFormatID      Key = "format_id"
	KeyJumperID      Key = "jumper_id"
	KeyCarrierID     Key = "carrier_id"
	KeyCountry       Key = "country"
	KeyCity          Key = "city"
	KeyLanguage      Key = "language"
	KeyIP            Key = "ip"
	KeyDeviceType    Key = "device_type"
	KeyDeviceID      Key = "device_id"
	KeyOSID          Key = "os_id"
	KeyBrowserID     Key = "browser_id"
)

// AdItemKey is a grouping key
type AdItemKey struct {
	Key   Key `json:"key,omitempty"`
	Value any `json:"value,omitempty"`
}

// AdItem is a statistic item
type AdItem struct {
	// Grouping keys
	Keys []AdItemKey `json:"keys"`

	// Money counters
	Spent    float64 `json:"spent"`
	Profit   float64 `json:"profit"`
	BidPrice float64 `json:"bid_price"`

	// Counters
	Requests    uint64 `json:"requests"`
	Impressions uint64 `json:"impressions"`
	Views       uint64 `json:"views"`
	Directs     uint64 `json:"directs"`
	Clicks      uint64 `json:"clicks"`
	Leads       uint64 `json:"leads"`

	// Calculated fields
	Wins   uint64  `json:"wins"`
	Bids   uint64  `json:"bids"`
	Skips  uint64  `json:"skips"`
	Nobids uint64  `json:"nobids"`
	Errors uint64  `json:"errors"`
	CTR    float64 `json:"ctr"`
	ECPM   float64 `json:"ecpm"`
	ECPC   float64 `json:"ecpc"`
	ECPA   float64 `json:"ecpa"`
}

// Repository of access to the statistic
type Repository interface {
	Statistic(ctx context.Context, opts ...Option) ([]*AdItem, error)
	Count(ctx context.Context, opts ...Option) (int64, error)
}
