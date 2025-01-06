package protobuf

import (
	"github.com/geniusrabbit/archivarius/internal/archivarius"
	"github.com/geniusrabbit/archivarius/internal/server/grpc"
)

// keyFromPb converts grpc.Key to archivarius.Key.
func keyFromPb(key grpc.Key) archivarius.Key {
	switch key {
	case grpc.Key_KEY_DATEMARK:
		return archivarius.KeyDatemark
	case grpc.Key_KEY_TIMEMARK:
		return archivarius.KeyTimemark
	case grpc.Key_KEY_CLUSTER:
		return archivarius.KeyCluster
	case grpc.Key_KEY_PROJECT_ID:
		return archivarius.KeyProjectID
	case grpc.Key_KEY_ACCOUNT_ID:
		return archivarius.KeyAccountID
	case grpc.Key_KEY_PUB_ACCOUNT_ID:
		return archivarius.KeyPubAccountID
	case grpc.Key_KEY_ADV_ACCOUNT_ID:
		return archivarius.KeyAdvAccountID
	case grpc.Key_KEY_SOURCE_ID:
		return archivarius.KeySourceID
	case grpc.Key_KEY_ACCESS_POINT_ID:
		return archivarius.KeyAccessPointID
	case grpc.Key_KEY_PLATFORM:
		return archivarius.KeyPlatform
	case grpc.Key_KEY_DOMAIN:
		return archivarius.KeyDomain
	case grpc.Key_KEY_APP_ID:
		return archivarius.KeyAppID
	case grpc.Key_KEY_ZONE_ID:
		return archivarius.KeyZoneID
	case grpc.Key_KEY_CAMPAIGN_ID:
		return archivarius.KeyCampaignID
	case grpc.Key_KEY_AD_ID:
		return archivarius.KeyAdID
	case grpc.Key_KEY_FORMAT_ID:
		return archivarius.KeyFormatID
	case grpc.Key_KEY_JUMPER_ID:
		return archivarius.KeyJumperID
	case grpc.Key_KEY_CARRIER_ID:
		return archivarius.KeyCarrierID
	case grpc.Key_KEY_COUNTRY:
		return archivarius.KeyCountry
	case grpc.Key_KEY_CITY:
		return archivarius.KeyCity
	case grpc.Key_KEY_LANGUAGE:
		return archivarius.KeyLanguage
	case grpc.Key_KEY_IP:
		return archivarius.KeyIP
	case grpc.Key_KEY_DEVICE_TYPE:
		return archivarius.KeyDeviceType
	case grpc.Key_KEY_DEVICE_ID:
		return archivarius.KeyDeviceID
	case grpc.Key_KEY_OS_ID:
		return archivarius.KeyOSID
	case grpc.Key_KEY_BROWSER_ID:
		return archivarius.KeyBrowserID
	}

	panic("unknown key")
}

// keyToPb converts archivarius.Key to grpc.Key.
func keyToPb(key archivarius.Key) grpc.Key {
	switch key {
	case archivarius.KeyDatemark:
		return grpc.Key_KEY_DATEMARK
	case archivarius.KeyTimemark:
		return grpc.Key_KEY_TIMEMARK
	case archivarius.KeyCluster:
		return grpc.Key_KEY_CLUSTER
	case archivarius.KeyProjectID:
		return grpc.Key_KEY_PROJECT_ID
	case archivarius.KeyAccountID:
		return grpc.Key_KEY_ACCOUNT_ID
	case archivarius.KeyPubAccountID:
		return grpc.Key_KEY_PUB_ACCOUNT_ID
	case archivarius.KeyAdvAccountID:
		return grpc.Key_KEY_ADV_ACCOUNT_ID
	case archivarius.KeySourceID:
		return grpc.Key_KEY_SOURCE_ID
	case archivarius.KeyAccessPointID:
		return grpc.Key_KEY_ACCESS_POINT_ID
	case archivarius.KeyPlatform:
		return grpc.Key_KEY_PLATFORM
	case archivarius.KeyDomain:
		return grpc.Key_KEY_DOMAIN
	case archivarius.KeyAppID:
		return grpc.Key_KEY_APP_ID
	case archivarius.KeyZoneID:
		return grpc.Key_KEY_ZONE_ID
	case archivarius.KeyCampaignID:
		return grpc.Key_KEY_CAMPAIGN_ID
	case archivarius.KeyAdID:
		return grpc.Key_KEY_AD_ID
	case archivarius.KeyFormatID:
		return grpc.Key_KEY_FORMAT_ID
	case archivarius.KeyJumperID:
		return grpc.Key_KEY_JUMPER_ID
	case archivarius.KeyCarrierID:
		return grpc.Key_KEY_CARRIER_ID
	case archivarius.KeyCountry:
		return grpc.Key_KEY_COUNTRY
	case archivarius.KeyCity:
		return grpc.Key_KEY_CITY
	case archivarius.KeyLanguage:
		return grpc.Key_KEY_LANGUAGE
	case archivarius.KeyIP:
		return grpc.Key_KEY_IP
	case archivarius.KeyDeviceType:
		return grpc.Key_KEY_DEVICE_TYPE
	case archivarius.KeyDeviceID:
		return grpc.Key_KEY_DEVICE_ID
	case archivarius.KeyOSID:
		return grpc.Key_KEY_OS_ID
	case archivarius.KeyBrowserID:
		return grpc.Key_KEY_BROWSER_ID
	}

	panic("unknown key")
}

// orderingKeyToPb converts archivarius.OrderingKey to grpc.OrderingKey.
func orderingKeyFromPb(key grpc.OrderingKey) archivarius.OrderingKey {
	switch key {
	case grpc.OrderingKey_ORDERING_KEY_DATEMARK:
		return archivarius.OrderingKeyDatemark
	case grpc.OrderingKey_ORDERING_KEY_TIMEMARK:
		return archivarius.OrderingKeyTimemark
	case grpc.OrderingKey_ORDERING_KEY_CLUSTER:
		return archivarius.OrderingKeyCluster
	case grpc.OrderingKey_ORDERING_KEY_PROJECT_ID:
		return archivarius.OrderingKeyProjectID
	case grpc.OrderingKey_ORDERING_KEY_PUB_ACCOUNT_ID:
		return archivarius.OrderingKeyPubAccountID
	case grpc.OrderingKey_ORDERING_KEY_ADV_ACCOUNT_ID:
		return archivarius.OrderingKeyAdvAccountID
	case grpc.OrderingKey_ORDERING_KEY_SOURCE_ID:
		return archivarius.OrderingKeySourceID
	case grpc.OrderingKey_ORDERING_KEY_ACCESS_POINT_ID:
		return archivarius.OrderingKeyAccessPointID
	case grpc.OrderingKey_ORDERING_KEY_PLATFORM:
		return archivarius.OrderingKeyPlatform
	case grpc.OrderingKey_ORDERING_KEY_DOMAIN:
		return archivarius.OrderingKeyDomain
	case grpc.OrderingKey_ORDERING_KEY_APP_ID:
		return archivarius.OrderingKeyAppID
	case grpc.OrderingKey_ORDERING_KEY_ZONE_ID:
		return archivarius.OrderingKeyZoneID
	case grpc.OrderingKey_ORDERING_KEY_CAMPAIGN_ID:
		return archivarius.OrderingKeyCampaignID
	case grpc.OrderingKey_ORDERING_KEY_AD_ID:
		return archivarius.OrderingKeyAdID
	case grpc.OrderingKey_ORDERING_KEY_FORMAT_ID:
		return archivarius.OrderingKeyFormatID
	case grpc.OrderingKey_ORDERING_KEY_JUMPER_ID:
		return archivarius.OrderingKeyJumperID
	case grpc.OrderingKey_ORDERING_KEY_CARRIER_ID:
		return archivarius.OrderingKeyCarrierID
	case grpc.OrderingKey_ORDERING_KEY_COUNTRY:
		return archivarius.OrderingKeyCountry
	case grpc.OrderingKey_ORDERING_KEY_CITY:
		return archivarius.OrderingKeyCity
	case grpc.OrderingKey_ORDERING_KEY_LANGUAGE:
		return archivarius.OrderingKeyLanguage
	case grpc.OrderingKey_ORDERING_KEY_IP:
		return archivarius.OrderingKeyIP
	case grpc.OrderingKey_ORDERING_KEY_DEVICE_TYPE:
		return archivarius.OrderingKeyDeviceType
	case grpc.OrderingKey_ORDERING_KEY_DEVICE_ID:
		return archivarius.OrderingKeyDeviceID
	case grpc.OrderingKey_ORDERING_KEY_OS_ID:
		return archivarius.OrderingKeyOSID
	case grpc.OrderingKey_ORDERING_KEY_BROWSER_ID:
		return archivarius.OrderingKeyBrowserID
	case grpc.OrderingKey_ORDERING_KEY_SPENT:
		return archivarius.OrderingKeySpent
	case grpc.OrderingKey_ORDERING_KEY_PROFIT:
		return archivarius.OrderingKeyProfit
	case grpc.OrderingKey_ORDERING_KEY_BID_PRICE:
		return archivarius.OrderingKeyBidPrice
	case grpc.OrderingKey_ORDERING_KEY_REQUESTS:
		return archivarius.OrderingKeyRequests
	case grpc.OrderingKey_ORDERING_KEY_IMPRESSIONS:
		return archivarius.OrderingKeyImpressions
	case grpc.OrderingKey_ORDERING_KEY_VIEWS:
		return archivarius.OrderingKeyViews
	case grpc.OrderingKey_ORDERING_KEY_DIRECTS:
		return archivarius.OrderingKeyDirects
	case grpc.OrderingKey_ORDERING_KEY_CLICKS:
		return archivarius.OrderingKeyClicks
	case grpc.OrderingKey_ORDERING_KEY_LEADS:
		return archivarius.OrderingKeyLeads
	case grpc.OrderingKey_ORDERING_KEY_BIDS:
		return archivarius.OrderingKeyBids
	case grpc.OrderingKey_ORDERING_KEY_WINS:
		return archivarius.OrderingKeyWins
	case grpc.OrderingKey_ORDERING_KEY_SKIPS:
		return archivarius.OrderingKeySkips
	case grpc.OrderingKey_ORDERING_KEY_NOBIDS:
		return archivarius.OrderingKeyNobids
	case grpc.OrderingKey_ORDERING_KEY_ERRORS:
		return archivarius.OrderingKeyErrors
	case grpc.OrderingKey_ORDERING_KEY_CTR:
		return archivarius.OrderingKeyCTR
	case grpc.OrderingKey_ORDERING_KEY_ECPM:
		return archivarius.OrderingKeyECPM
	case grpc.OrderingKey_ORDERING_KEY_ECPC:
		return archivarius.OrderingKeyECPC
	case grpc.OrderingKey_ORDERING_KEY_ECPA:
		return archivarius.OrderingKeyECPA
	}

	panic("unknown key")
}
