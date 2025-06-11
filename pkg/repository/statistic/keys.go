package statistic

type Key string

const (
	KeyUndefined    Key = "undefined"
	KeyDatemark     Key = "datemark"
	KeyTimemark     Key = "timemark"
	KeySourceID     Key = "source_id"
	KeyPlatformType Key = "platform_type"
	KeyDomain       Key = "domain"
	KeyAppID        Key = "app_id"
	KeyZoneID       Key = "zone_id"
	KeyFormatID     Key = "format_id"
	KeyFormatCode   Key = "format_code"
	KeyCarrierID    Key = "carrier_id"
	KeyCountry      Key = "country"
	KeyLanguage     Key = "language"
	KeyIP           Key = "ip"
	KeyDeviceID     Key = "device_id"
	KeyDeviceType   Key = "device_type"
	KeyOsID         Key = "os_id"
	KeyBrowserID    Key = "browser_id"
)

func (key Key) String() string {
	return string(key)
}

type OrderingKey string

const (
	OrderingUndefined       OrderingKey = "undefined"
	OrderingKeyDatemark     OrderingKey = "datemark"
	OrderingKeyTimemark     OrderingKey = "timemark"
	OrderingKeySourceID     OrderingKey = "source_id"
	OrderingKeyPlatformType OrderingKey = "platform_type"
	OrderingKeyDomain       OrderingKey = "domain"
	OrderingKeyAppID        OrderingKey = "app_id"
	OrderingKeyZoneID       OrderingKey = "zone_id"
	OrderingKeyFormatID     OrderingKey = "format_id"
	OrderingKeyFormatCode   OrderingKey = "format_code"
	OrderingKeyCarrierID    OrderingKey = "carrier_id"
	OrderingKeyCountry      OrderingKey = "country"
	OrderingKeyLanguage     OrderingKey = "language"
	OrderingKeyIP           OrderingKey = "ip"
	OrderingKeyDeviceID     OrderingKey = "device_id"
	OrderingKeyDeviceType   OrderingKey = "device_type"
	OrderingKeyOsID         OrderingKey = "os_id"
	OrderingKeyBrowserID    OrderingKey = "browser_id"
	// Counters
	OrderingKeyRequests OrderingKey = "imps+views+directs+clicks"
	OrderingKeyImps     OrderingKey = "imps"
	OrderingKeyViews    OrderingKey = "views"
	OrderingKeyDirects  OrderingKey = "directs"
	OrderingKeyClicks   OrderingKey = "clicks"
	OrderingKeyBids     OrderingKey = "bid_requests"
	OrderingKeyWins     OrderingKey = "bid_wins"
	OrderingKeySkips    OrderingKey = "bid_skips"
	OrderingKeyNobids   OrderingKey = "bid_nobids"
	OrderingKeyErrors   OrderingKey = "bid_errors"
	OrderingKeyRevenue  OrderingKey = "revenue"
	OrderingKeyBidPrice OrderingKey = "potential_revenue"
	OrderingKeyCTR      OrderingKey = "IF(imps > 0, (clicks / imps) * 100, 0)"
	OrderingKeyECPM     OrderingKey = "IF(imps > 0, (revenue / imps) * 1000, 0)"
	OrderingKeyECPC     OrderingKey = "IF(clicks > 0, (revenue / clicks), 0)"
)

func (key OrderingKey) String() string {
	return string(key)
}
