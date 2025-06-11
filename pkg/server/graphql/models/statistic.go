package models

import (
	"strings"
	"time"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/repository"

	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/statistic"
)

func FromStatisticAdItemKeyModel(st *models.StatisticAdItemKey) *StatisticItemKey {
	if st == nil {
		return nil
	}
	return &StatisticItemKey{
		Key: FromRepoStatisticKey(
			statistic.Key(strings.ToLower(st.Key)),
		),
		Value: st.Value,
		Text:  st.Text,
	}
}

func FromStatisticAdItemModel(st *models.StatisticAdItem) *StatisticAdItem {
	if st == nil {
		return nil
	}
	return &StatisticAdItem{
		Keys: xtypes.SliceApply(st.Keys, func(k models.StatisticAdItemKey) *StatisticItemKey {
			return FromStatisticAdItemKeyModel(&k)
		}),
		// Money counters
		Revenue:  st.Revenue,
		BidPrice: st.BidPrice, // Sum of all bids prices
		// Counters
		Requests:    st.Requests,
		Impressions: st.Impressions,
		Views:       st.Views,
		Directs:     st.Directs,
		Clicks:      st.Clicks,
		Bids:        st.Bids,
		Wins:        st.Wins,
		Skips:       st.Skips,
		Nobids:      st.Nobids,
		Errors:      st.Errors,
		// Calculated fields
		Ctr:  st.CTR(),
		ECpm: st.ECPM(),
		ECpc: st.ECPC(),
	}
}

func FromStatisticAdItemModelList(st []*models.StatisticAdItem) []*StatisticAdItem {
	return xtypes.SliceApply(st, FromStatisticAdItemModel)
}

func (fl *StatisticAdListFilter) Filter() *statistic.Filter {
	if fl == nil {
		return nil
	}
	return &statistic.Filter{
		Conditions: xtypes.SliceApply(fl.Conditions, func(c *StatisticAdKeyCondition) *statistic.Condition {
			return &statistic.Condition{
				Key:   c.Key.AsQueryKey(),
				Op:    c.Op.AsQueryOp(),
				Value: c.Value,
			}
		}),
		StartDate: gocast.PtrAsValue((*time.Time)(fl.StartDate), time.Time{}),
		EndDate:   gocast.PtrAsValue((*time.Time)(fl.EndDate), time.Time{}),
	}
}

func StatisticGroup(group []StatisticKey) *repository.GroupOption {
	if len(group) == 0 {
		return nil
	}
	return statistic.WithGroup(
		xtypes.SliceApply(group, func(k StatisticKey) statistic.Key {
			return k.AsQueryKey()
		})...,
	)
}

func StatisticAdListOrder(ord []*StatisticAdKeyOrder) *statistic.ListOrder {
	if len(ord) == 0 {
		return nil
	}
	nord := &statistic.ListOrder{}
	for _, o := range ord {
		nord.SetOrder(o.Key.AsQueryOrderKey(), o.Order.AsOrder())
	}
	return nord
}

func FromRepoStatisticKey(key statistic.Key) StatisticKey {
	switch key {
	case statistic.KeyDatemark:
		return StatisticKeyDatemark
	case statistic.KeyTimemark:
		return StatisticKeyTimemark
	case statistic.KeySourceID:
		return StatisticKeySourceID
	case statistic.KeyPlatformType:
		return StatisticKeyPlatformType
	case statistic.KeyDomain:
		return StatisticKeyDomain
	case statistic.KeyAppID:
		return StatisticKeyAppID
	case statistic.KeyZoneID:
		return StatisticKeyZoneID
	case statistic.KeyFormatID:
		return StatisticKeyFormatID
	case statistic.KeyFormatCode:
		return StatisticKeyFormatCode
	case statistic.KeyCarrierID:
		return StatisticKeyCarrierID
	case statistic.KeyCountry:
		return StatisticKeyCountry
	case statistic.KeyLanguage:
		return StatisticKeyLanguage
	case statistic.KeyIP:
		return StatisticKeyIP
	case statistic.KeyDeviceID:
		return StatisticKeyDeviceID
	case statistic.KeyDeviceType:
		return StatisticKeyDeviceType
	case statistic.KeyOsID:
		return StatisticKeyOsID
	case statistic.KeyBrowserID:
		return StatisticKeyBrowserID
	}
	return StatisticKeyUndefined
}

func (key StatisticKey) AsQueryKey() statistic.Key {
	switch key {
	case StatisticKeyDatemark:
		return statistic.KeyDatemark
	case StatisticKeyTimemark:
		return statistic.KeyTimemark
	case StatisticKeySourceID:
		return statistic.KeySourceID
	case StatisticKeyPlatformType:
		return statistic.KeyPlatformType
	case StatisticKeyDomain:
		return statistic.KeyDomain
	case StatisticKeyAppID:
		return statistic.KeyAppID
	case StatisticKeyZoneID:
		return statistic.KeyZoneID
	case StatisticKeyFormatID:
		return statistic.KeyFormatID
	case StatisticKeyFormatCode:
		return statistic.KeyFormatCode
	case StatisticKeyCarrierID:
		return statistic.KeyCarrierID
	case StatisticKeyCountry:
		return statistic.KeyCountry
	case StatisticKeyLanguage:
		return statistic.KeyLanguage
	case StatisticKeyIP:
		return statistic.KeyIP
	case StatisticKeyDeviceID:
		return statistic.KeyDeviceID
	case StatisticKeyDeviceType:
		return statistic.KeyDeviceType
	case StatisticKeyOsID:
		return statistic.KeyOsID
	case StatisticKeyBrowserID:
		return statistic.KeyBrowserID
	}
	return statistic.KeyUndefined
}

func (op StatisticCondition) AsQueryOp() statistic.Operation {
	switch op {
	case StatisticConditionEq:
		return statistic.ConditionEq
	case StatisticConditionNotEq:
		return statistic.ConditionNotEq
	case StatisticConditionGt:
		return statistic.ConditionGt
	case StatisticConditionGtEq:
		return statistic.ConditionGtEq
	case StatisticConditionLt:
		return statistic.ConditionLt
	case StatisticConditionLtEq:
		return statistic.ConditionLtEq
	case StatisticConditionIn:
		return statistic.ConditionIn
	case StatisticConditionNotIn:
		return statistic.ConditionNotIn
	case StatisticConditionBetween:
		return statistic.ConditionBetween
	case StatisticConditionNotBetween:
		return statistic.ConditionNotBetween
	case StatisticConditionLike:
		return statistic.ConditionLike
	case StatisticConditionNotLike:
		return statistic.ConditionNotLike
	case StatisticConditionIsNull:
		return statistic.ConditionIsNull
	case StatisticConditionIsNotNull:
		return statistic.ConditionIsNotNull
	}
	return statistic.ConditionUndefined
}

func (ord StatisticOrderingKey) AsQueryOrderKey() statistic.OrderingKey {
	switch ord {
	case StatisticOrderingKeyDatemark:
		return statistic.OrderingKeyDatemark
	case StatisticOrderingKeyTimemark:
		return statistic.OrderingKeyTimemark
	case StatisticOrderingKeySourceID:
		return statistic.OrderingKeySourceID
	case StatisticOrderingKeyPlatformType:
		return statistic.OrderingKeyPlatformType
	case StatisticOrderingKeyDomain:
		return statistic.OrderingKeyDomain
	case StatisticOrderingKeyAppID:
		return statistic.OrderingKeyAppID
	case StatisticOrderingKeyZoneID:
		return statistic.OrderingKeyZoneID
	case StatisticOrderingKeyFormatID:
		return statistic.OrderingKeyFormatID
	case StatisticOrderingKeyFormatCode:
		return statistic.OrderingKeyFormatCode
	case StatisticOrderingKeyCarrierID:
		return statistic.OrderingKeyCarrierID
	case StatisticOrderingKeyCountry:
		return statistic.OrderingKeyCountry
	case StatisticOrderingKeyLanguage:
		return statistic.OrderingKeyLanguage
	case StatisticOrderingKeyIP:
		return statistic.OrderingKeyIP
	case StatisticOrderingKeyDeviceID:
		return statistic.OrderingKeyDeviceID
	case StatisticOrderingKeyDeviceType:
		return statistic.OrderingKeyDeviceType
	case StatisticOrderingKeyOsID:
		return statistic.OrderingKeyOsID
	case StatisticOrderingKeyBrowserID:
		return statistic.OrderingKeyBrowserID
		// Counters
	case StatisticOrderingKeyRequests:
		return statistic.OrderingKeyRequests
	case StatisticOrderingKeyImpressions:
		return statistic.OrderingKeyImps
	case StatisticOrderingKeyViews:
		return statistic.OrderingKeyViews
	case StatisticOrderingKeyDirects:
		return statistic.OrderingKeyDirects
	case StatisticOrderingKeyClicks:
		return statistic.OrderingKeyClicks
	case StatisticOrderingKeyBids:
		return statistic.OrderingKeyBids
	case StatisticOrderingKeyWins:
		return statistic.OrderingKeyWins
	case StatisticOrderingKeySkips:
		return statistic.OrderingKeySkips
	case StatisticOrderingKeyNobids:
		return statistic.OrderingKeyNobids
	case StatisticOrderingKeyErrors:
		return statistic.OrderingKeyErrors
	case StatisticOrderingKeyRevenue:
		return statistic.OrderingKeyRevenue
	case StatisticOrderingKeyBidPrice:
		return statistic.OrderingKeyBidPrice
	case StatisticOrderingKeyCtr:
		return statistic.OrderingKeyCTR
	case StatisticOrderingKeyEcpm:
		return statistic.OrderingKeyECPM
	case StatisticOrderingKeyEcpc:
		return statistic.OrderingKeyECPC
	}
	return statistic.OrderingUndefined
}
