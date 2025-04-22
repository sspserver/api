package models

type StatisticAdItemKey struct {
	Key   string `json:"key,omitempty"`
	Value any    `json:"value,omitempty"`
	Text  string `json:"text,omitempty"`
}

type StatisticAdItem struct {
	Keys []StatisticAdItemKey `json:"keys"`

	// Money counters
	Profit   float64 `json:"profit"`
	BidPrice float64 `json:"bid_price"`

	// Counters
	Requests    uint64 `json:"requests"`
	Impressions uint64 `json:"impressions"`
	Views       uint64 `json:"views"`
	Directs     uint64 `json:"directs"`
	Clicks      uint64 `json:"clicks"`
	Wins        uint64 `json:"wins"`
	Bids        uint64 `json:"bids"`
	Skips       uint64 `json:"skips"`
	Nobids      uint64 `json:"nobids"`
	Errors      uint64 `json:"errors"`
}

func (it *StatisticAdItem) CTR() float64 {
	if it.Views == 0 {
		return 0
	}
	return float64(it.Clicks) / float64(it.Views) * 100
}

func (it *StatisticAdItem) ECPM() float64 {
	if it.Impressions == 0 {
		return 0
	}
	return it.Profit / float64(it.Impressions) * 1000
}

func (it *StatisticAdItem) ECPC() float64 {
	if it.Clicks == 0 {
		return 0
	}
	return it.Profit / float64(it.Clicks)
}
