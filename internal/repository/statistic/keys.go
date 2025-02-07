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
