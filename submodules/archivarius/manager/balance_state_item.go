package manager

import (
	"sync/atomic"

	"github.com/geniusrabbit/adcorelib/billing"
)

type BalanceStateItem struct {
	balance  atomic.Int64
	spend    atomic.Int64
	lastSync uint64 // Timestamp of the last synchronization
}

func (state *BalanceStateItem) Balance() billing.Money {
	return billing.Money(state.balance.Load())
}

func (state *BalanceStateItem) Spend() billing.Money {
	return billing.Money(state.spend.Load())
}

func (state *BalanceStateItem) Charge(amount billing.Money) {
	state.spend.Add(int64(amount))
}

func (state *BalanceStateItem) update(newState *StateData, serviceCount uint64) bool {
	// Only proceed if the new state is more recent and has higher spend
	if state.lastSync >= newState.LastSync || state.Spend() >= billing.Money(newState.Spend) {
		return false
	}
	state.balance.Store(newState.Balance / int64(serviceCount))
	state.spend.Store(newState.Spend / int64(serviceCount))
	state.lastSync = newState.LastSync
	return true
}
