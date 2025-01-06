package manager

import "github.com/geniusrabbit/adcorelib/billing"

type State interface {
	Profit() billing.Money
	TotalProfit() billing.Money
	Spend() billing.Money
	TotalSpend() billing.Money
	TestSpend() billing.Money
	Views() uint64
	Clicks() uint64
	Leads() uint64
}
