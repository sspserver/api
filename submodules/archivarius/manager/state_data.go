package manager

type StateData struct {
	Balance     int64
	Profit      int64
	TotalProfit int64
	Spend       int64
	TotalSpend  int64
	TestSpend   int64
	Views       uint64
	Clicks      uint64
	Leads       uint64
	LastSync    uint64
}
