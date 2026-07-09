package graphql

import (
	"time"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"

	"github.com/sspserver/api/pkg/repository/statistic"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func StatisticFilterFromGraphQL(fl *gqlmodels.StatisticAdListFilter) *statistic.Filter {
	if fl == nil {
		return nil
	}
	return &statistic.Filter{
		Conditions: xtypes.SliceApply(fl.Conditions, func(c *gqlmodels.StatisticAdKeyCondition) *statistic.Condition {
			if c == nil {
				return nil
			}
			return &statistic.Condition{
				Key:   statisticKeyFromGraphQL(c.Key),
				Op:    statisticConditionFromGraphQL(c.Op),
				Value: c.Value,
			}
		}),
		StartDate: gocast.PtrAsValue((*time.Time)(fl.StartDate), time.Time{}),
		EndDate:   gocast.PtrAsValue((*time.Time)(fl.EndDate), time.Time{}),
	}
}

func FromRepoStatisticKey(key statistic.Key) gqlmodels.StatisticKey {
	switch key {
	case statistic.KeyDatemark:
		return gqlmodels.StatisticKeyDatemark
	case statistic.KeyTimemark:
		return gqlmodels.StatisticKeyTimemark
	case statistic.KeySourceID:
		return gqlmodels.StatisticKeySourceID
	case statistic.KeyPlatformType:
		return gqlmodels.StatisticKeyPlatformType
	case statistic.KeyDomain:
		return gqlmodels.StatisticKeyDomain
	case statistic.KeyAppID:
		return gqlmodels.StatisticKeyAppID
	case statistic.KeyZoneID:
		return gqlmodels.StatisticKeyZoneID
	case statistic.KeyFormatID:
		return gqlmodels.StatisticKeyFormatID
	case statistic.KeyFormatCode:
		return gqlmodels.StatisticKeyFormatCode
	case statistic.KeyCarrierID:
		return gqlmodels.StatisticKeyCarrierID
	case statistic.KeyCountry:
		return gqlmodels.StatisticKeyCountry
	case statistic.KeyLanguage:
		return gqlmodels.StatisticKeyLanguage
	case statistic.KeyIP:
		return gqlmodels.StatisticKeyIP
	case statistic.KeyDeviceID:
		return gqlmodels.StatisticKeyDeviceID
	case statistic.KeyDeviceType:
		return gqlmodels.StatisticKeyDeviceType
	case statistic.KeyOsID:
		return gqlmodels.StatisticKeyOsID
	case statistic.KeyBrowserID:
		return gqlmodels.StatisticKeyBrowserID
	}
	return gqlmodels.StatisticKeyUndefined
}

func statisticKeyFromGraphQL(key gqlmodels.StatisticKey) statistic.Key {
	switch key {
	case gqlmodels.StatisticKeyDatemark:
		return statistic.KeyDatemark
	case gqlmodels.StatisticKeyTimemark:
		return statistic.KeyTimemark
	case gqlmodels.StatisticKeySourceID:
		return statistic.KeySourceID
	case gqlmodels.StatisticKeyPlatformType:
		return statistic.KeyPlatformType
	case gqlmodels.StatisticKeyDomain:
		return statistic.KeyDomain
	case gqlmodels.StatisticKeyAppID:
		return statistic.KeyAppID
	case gqlmodels.StatisticKeyZoneID:
		return statistic.KeyZoneID
	case gqlmodels.StatisticKeyFormatID:
		return statistic.KeyFormatID
	case gqlmodels.StatisticKeyFormatCode:
		return statistic.KeyFormatCode
	case gqlmodels.StatisticKeyCarrierID:
		return statistic.KeyCarrierID
	case gqlmodels.StatisticKeyCountry:
		return statistic.KeyCountry
	case gqlmodels.StatisticKeyLanguage:
		return statistic.KeyLanguage
	case gqlmodels.StatisticKeyIP:
		return statistic.KeyIP
	case gqlmodels.StatisticKeyDeviceID:
		return statistic.KeyDeviceID
	case gqlmodels.StatisticKeyDeviceType:
		return statistic.KeyDeviceType
	case gqlmodels.StatisticKeyOsID:
		return statistic.KeyOsID
	case gqlmodels.StatisticKeyBrowserID:
		return statistic.KeyBrowserID
	}
	return statistic.KeyUndefined
}

func statisticConditionFromGraphQL(op gqlmodels.StatisticCondition) statistic.Operation {
	switch op {
	case gqlmodels.StatisticConditionEq:
		return statistic.ConditionEq
	case gqlmodels.StatisticConditionNotEq:
		return statistic.ConditionNotEq
	case gqlmodels.StatisticConditionGt:
		return statistic.ConditionGt
	case gqlmodels.StatisticConditionGtEq:
		return statistic.ConditionGtEq
	case gqlmodels.StatisticConditionLt:
		return statistic.ConditionLt
	case gqlmodels.StatisticConditionLtEq:
		return statistic.ConditionLtEq
	case gqlmodels.StatisticConditionIn:
		return statistic.ConditionIn
	case gqlmodels.StatisticConditionNotIn:
		return statistic.ConditionNotIn
	case gqlmodels.StatisticConditionBetween:
		return statistic.ConditionBetween
	case gqlmodels.StatisticConditionNotBetween:
		return statistic.ConditionNotBetween
	case gqlmodels.StatisticConditionLike:
		return statistic.ConditionLike
	case gqlmodels.StatisticConditionNotLike:
		return statistic.ConditionNotLike
	case gqlmodels.StatisticConditionIsNull:
		return statistic.ConditionIsNull
	case gqlmodels.StatisticConditionIsNotNull:
		return statistic.ConditionIsNotNull
	}
	return statistic.ConditionUndefined
}

func statisticOrderKeyFromGraphQL(ord gqlmodels.StatisticOrderingKey) statistic.OrderingKey {
	switch ord {
	case gqlmodels.StatisticOrderingKeyDatemark:
		return statistic.OrderingKeyDatemark
	case gqlmodels.StatisticOrderingKeyTimemark:
		return statistic.OrderingKeyTimemark
	case gqlmodels.StatisticOrderingKeySourceID:
		return statistic.OrderingKeySourceID
	case gqlmodels.StatisticOrderingKeyPlatformType:
		return statistic.OrderingKeyPlatformType
	case gqlmodels.StatisticOrderingKeyDomain:
		return statistic.OrderingKeyDomain
	case gqlmodels.StatisticOrderingKeyAppID:
		return statistic.OrderingKeyAppID
	case gqlmodels.StatisticOrderingKeyZoneID:
		return statistic.OrderingKeyZoneID
	case gqlmodels.StatisticOrderingKeyFormatID:
		return statistic.OrderingKeyFormatID
	case gqlmodels.StatisticOrderingKeyFormatCode:
		return statistic.OrderingKeyFormatCode
	case gqlmodels.StatisticOrderingKeyCarrierID:
		return statistic.OrderingKeyCarrierID
	case gqlmodels.StatisticOrderingKeyCountry:
		return statistic.OrderingKeyCountry
	case gqlmodels.StatisticOrderingKeyLanguage:
		return statistic.OrderingKeyLanguage
	case gqlmodels.StatisticOrderingKeyIP:
		return statistic.OrderingKeyIP
	case gqlmodels.StatisticOrderingKeyDeviceID:
		return statistic.OrderingKeyDeviceID
	case gqlmodels.StatisticOrderingKeyDeviceType:
		return statistic.OrderingKeyDeviceType
	case gqlmodels.StatisticOrderingKeyOsID:
		return statistic.OrderingKeyOsID
	case gqlmodels.StatisticOrderingKeyBrowserID:
		return statistic.OrderingKeyBrowserID
		// Counters
	case gqlmodels.StatisticOrderingKeyRequests:
		return statistic.OrderingKeyRequests
	case gqlmodels.StatisticOrderingKeyImpressions:
		return statistic.OrderingKeyImps
	case gqlmodels.StatisticOrderingKeyViews:
		return statistic.OrderingKeyViews
	case gqlmodels.StatisticOrderingKeyDirects:
		return statistic.OrderingKeyDirects
	case gqlmodels.StatisticOrderingKeyClicks:
		return statistic.OrderingKeyClicks
	case gqlmodels.StatisticOrderingKeyBids:
		return statistic.OrderingKeyBids
	case gqlmodels.StatisticOrderingKeyWins:
		return statistic.OrderingKeyWins
	case gqlmodels.StatisticOrderingKeySkips:
		return statistic.OrderingKeySkips
	case gqlmodels.StatisticOrderingKeyNobids:
		return statistic.OrderingKeyNobids
	case gqlmodels.StatisticOrderingKeyErrors:
		return statistic.OrderingKeyErrors
	case gqlmodels.StatisticOrderingKeyRevenue:
		return statistic.OrderingKeyRevenue
	case gqlmodels.StatisticOrderingKeyBidPrice:
		return statistic.OrderingKeyBidPrice
	case gqlmodels.StatisticOrderingKeyCtr:
		return statistic.OrderingKeyCTR
	case gqlmodels.StatisticOrderingKeyEcpm:
		return statistic.OrderingKeyECPM
	case gqlmodels.StatisticOrderingKeyEcpc:
		return statistic.OrderingKeyECPC
	}
	return statistic.OrderingUndefined
}
