package graphql

import (
	"slices"
	"strings"

	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/repository"

	"github.com/sspserver/api/pkg/repository/statistic"
	"github.com/sspserver/api/pkg/repository/statistic/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FromStatisticAdItemKeyModel(st *models.StatisticAdItemKey) *gqlmodels.StatisticItemKey {
	if st == nil {
		return nil
	}
	return &gqlmodels.StatisticItemKey{
		Key: gqlmodels.FromRepoStatisticKey(
			statistic.Key(strings.ToLower(st.Key)),
		),
		Value: st.Value,
		Text:  st.Text,
	}
}

func FromStatisticAdItemModel(st *models.StatisticAdItem) *gqlmodels.StatisticAdItem {
	if st == nil {
		return nil
	}
	return &gqlmodels.StatisticAdItem{
		Keys: xtypes.SliceApply(st.Keys, func(k models.StatisticAdItemKey) *gqlmodels.StatisticItemKey {
			return FromStatisticAdItemKeyModel(&k)
		}),
		// Money counters
		Revenue:  st.Revenue,
		BidPrice: st.BidPrice,
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

func FromStatisticAdItemModelList(st []*models.StatisticAdItem) []*gqlmodels.StatisticAdItem {
	return xtypes.SliceApply(st, FromStatisticAdItemModel)
}

func StatisticGroup(group []gqlmodels.StatisticKey) *repository.GroupOption {
	if len(group) == 0 {
		return nil
	}
	return statistic.WithGroup(
		xtypes.SliceApply(group, func(k gqlmodels.StatisticKey) statistic.Key {
			return k.AsQueryKey()
		})...,
	)
}

func StatisticAdListOrder(ord []*gqlmodels.StatisticAdKeyOrder, group []gqlmodels.StatisticKey) *statistic.ListOrder {
	if len(ord) == 0 {
		return nil
	}
	nord := &statistic.ListOrder{}
	groupKeys := xtypes.SliceApply(group, func(k gqlmodels.StatisticKey) string {
		return k.AsQueryKey().String()
	})
	for _, o := range ord {
		oKey := o.Key.AsQueryOrderKey()
		if !statistic.IsAggregationKey(oKey.String()) ||
			slices.Contains(groupKeys, oKey.String()) {
			nord.SetOrder(oKey, o.Order.AsOrder())
		}
	}
	return nord
}
