package models

import (
	"github.com/geniusrabbit/archivarius/client"
)

func (op StatisticCondition) ClientCondition() client.Condition {
	switch op {
	case StatisticConditionEq:
		return client.Condition_EQ
	case StatisticConditionNe:
		return client.Condition_NE
	case StatisticConditionGt:
		return client.Condition_GT
	case StatisticConditionGe:
		return client.Condition_GE
	case StatisticConditionLt:
		return client.Condition_LT
	case StatisticConditionLe:
		return client.Condition_LE
	case StatisticConditionIn:
		return client.Condition_IN
	case StatisticConditionNi:
		return client.Condition_NI
	case StatisticConditionBt:
		return client.Condition_BT
	case StatisticConditionNb:
		return client.Condition_NB
	case StatisticConditionLi:
		return client.Condition_LI
	case StatisticConditionNl:
		return client.Condition_NL
	}
	return client.Condition_UNKNOWN
}

func (key StatisticKey) ClientKey() client.Key {
	switch key {
	case StatisticKeyDatemark:
		return client.Key_KEY_DATEMARK
	case StatisticKeyTimemark:
		return client.Key_KEY_TIMEMARK
	case StatisticKeyCluster:
		return client.Key_KEY_CLUSTER
	case StatisticKeyProjectID:
		return client.Key_KEY_PROJECT_ID
	case StatisticKeyAccountID:
		return client.Key_KEY_ACCOUNT_ID
	case StatisticKeyPubAccountID:
		return client.Key_KEY_PUB_ACCOUNT_ID
	case StatisticKeyAdvAccountID:
		return client.Key_KEY_ADV_ACCOUNT_ID
	case StatisticKeySourceID:
		return client.Key_KEY_SOURCE_ID
	case StatisticKeyAccessPointID:
		return client.Key_KEY_ACCESS_POINT_ID
	case StatisticKeyPlatform:
		return client.Key_KEY_PLATFORM
	case StatisticKeyDomain:
		return client.Key_KEY_DOMAIN
	case StatisticKeyAppID:
		return client.Key_KEY_APP_ID
	case StatisticKeyZoneID:
		return client.Key_KEY_ZONE_ID
	case StatisticKeyCampaignID:
		return client.Key_KEY_CAMPAIGN_ID
	case StatisticKeyAdID:
		return client.Key_KEY_AD_ID
	case StatisticKeyFormatID:
		return client.Key_KEY_FORMAT_ID
	case StatisticKeyJumperID:
		return client.Key_KEY_JUMPER_ID
	case StatisticKeyCarrierID:
		return client.Key_KEY_CARRIER_ID
	case StatisticKeyCountry:
		return client.Key_KEY_COUNTRY
	case StatisticKeyCity:
		return client.Key_KEY_CITY
	case StatisticKeyLanguage:
		return client.Key_KEY_LANGUAGE
	case StatisticKeyIP:
		return client.Key_KEY_IP
	case StatisticKeyDeviceType:
		return client.Key_KEY_DEVICE_TYPE
	case StatisticKeyDeviceID:
		return client.Key_KEY_DEVICE_ID
	case StatisticKeyOsID:
		return client.Key_KEY_OS_ID
	case StatisticKeyBrowserID:
		return client.Key_KEY_BROWSER_ID
	}
	return client.Key_UNKNOWN_KEY
}

func (ordKey StatisticOrderingKey) ClientOrderingKey() client.OrderingKey {
	switch ordKey {
	case StatisticOrderingKeyDatemark:
		return client.OrderingKey_ORDERING_KEY_DATEMARK
	case StatisticOrderingKeyTimemark:
		return client.OrderingKey_ORDERING_KEY_TIMEMARK
	case StatisticOrderingKeyCluster:
		return client.OrderingKey_ORDERING_KEY_CLUSTER
	case StatisticOrderingKeyProjectID:
		return client.OrderingKey_ORDERING_KEY_PROJECT_ID
	case StatisticOrderingKeyPubAccountID:
		return client.OrderingKey_ORDERING_KEY_PUB_ACCOUNT_ID
	case StatisticOrderingKeyAdvAccountID:
		return client.OrderingKey_ORDERING_KEY_ADV_ACCOUNT_ID
	case StatisticOrderingKeySourceID:
		return client.OrderingKey_ORDERING_KEY_SOURCE_ID
	case StatisticOrderingKeyAccessPointID:
		return client.OrderingKey_ORDERING_KEY_ACCESS_POINT_ID
	case StatisticOrderingKeyPlatform:
		return client.OrderingKey_ORDERING_KEY_PLATFORM
	case StatisticOrderingKeyDomain:
		return client.OrderingKey_ORDERING_KEY_DOMAIN
	case StatisticOrderingKeyAppID:
		return client.OrderingKey_ORDERING_KEY_APP_ID
	case StatisticOrderingKeyZoneID:
		return client.OrderingKey_ORDERING_KEY_ZONE_ID
	case StatisticOrderingKeyCampaignID:
		return client.OrderingKey_ORDERING_KEY_CAMPAIGN_ID
	case StatisticOrderingKeyAdID:
		return client.OrderingKey_ORDERING_KEY_AD_ID
	case StatisticOrderingKeyFormatID:
		return client.OrderingKey_ORDERING_KEY_FORMAT_ID
	case StatisticOrderingKeyJumperID:
		return client.OrderingKey_ORDERING_KEY_JUMPER_ID
	case StatisticOrderingKeyCarrierID:
		return client.OrderingKey_ORDERING_KEY_CARRIER_ID
	case StatisticOrderingKeyCountry:
		return client.OrderingKey_ORDERING_KEY_COUNTRY
	case StatisticOrderingKeyCity:
		return client.OrderingKey_ORDERING_KEY_CITY
	case StatisticOrderingKeyLanguage:
		return client.OrderingKey_ORDERING_KEY_LANGUAGE
	case StatisticOrderingKeyIP:
		return client.OrderingKey_ORDERING_KEY_IP
	case StatisticOrderingKeyDeviceType:
		return client.OrderingKey_ORDERING_KEY_DEVICE_TYPE
	case StatisticOrderingKeyDeviceID:
		return client.OrderingKey_ORDERING_KEY_DEVICE_ID
	case StatisticOrderingKeyOsID:
		return client.OrderingKey_ORDERING_KEY_OS_ID
	case StatisticOrderingKeyBrowserID:
		return client.OrderingKey_ORDERING_KEY_BROWSER_ID
	case StatisticOrderingKeySpent:
		return client.OrderingKey_ORDERING_KEY_SPENT
	case StatisticOrderingKeyProfit:
		return client.OrderingKey_ORDERING_KEY_PROFIT
	case StatisticOrderingKeyBidPrice:
		return client.OrderingKey_ORDERING_KEY_BID_PRICE
	case StatisticOrderingKeyRequests:
		return client.OrderingKey_ORDERING_KEY_REQUESTS
	case StatisticOrderingKeyImpressions:
		return client.OrderingKey_ORDERING_KEY_IMPRESSIONS
	case StatisticOrderingKeyViews:
		return client.OrderingKey_ORDERING_KEY_VIEWS
	case StatisticOrderingKeyDirects:
		return client.OrderingKey_ORDERING_KEY_DIRECTS
	case StatisticOrderingKeyClicks:
		return client.OrderingKey_ORDERING_KEY_CLICKS
	case StatisticOrderingKeyLeads:
		return client.OrderingKey_ORDERING_KEY_LEADS
	case StatisticOrderingKeyBids:
		return client.OrderingKey_ORDERING_KEY_BIDS
	case StatisticOrderingKeyWins:
		return client.OrderingKey_ORDERING_KEY_WINS
	case StatisticOrderingKeySkips:
		return client.OrderingKey_ORDERING_KEY_SKIPS
	case StatisticOrderingKeyNobids:
		return client.OrderingKey_ORDERING_KEY_NOBIDS
	case StatisticOrderingKeyErrors:
		return client.OrderingKey_ORDERING_KEY_ERRORS
	case StatisticOrderingKeyCtr:
		return client.OrderingKey_ORDERING_KEY_CTR
	case StatisticOrderingKeyEcpm:
		return client.OrderingKey_ORDERING_KEY_ECPM
	case StatisticOrderingKeyEcpc:
		return client.OrderingKey_ORDERING_KEY_ECPC
	case StatisticOrderingKeyEcpa:
		return client.OrderingKey_ORDERING_KEY_ECPA
	}
	return client.OrderingKey_UNKNOWN_ORDERING_KEY
}
