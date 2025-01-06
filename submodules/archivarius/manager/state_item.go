package manager

import (
	"sync/atomic"

	"github.com/geniusrabbit/adcorelib/billing"
)

// StateItem represents an item that tracks the performance metrics of an ad campaign or other entities.
// It includes counters for views, clicks, leads, profit, and spend amounts.
type StateItem struct {
	parent      *StateItem    // Parent StateItem, used to aggregate metrics up the hierarchy
	profit      atomic.Int64  // Tracks total profit (atomic for thread-safe operations)
	totalProfit atomic.Int64  // Tracks total profit (atomic for thread-safe operations)
	spend       atomic.Int64  // Tracks total spend (atomic for thread-safe operations)
	totalSpend  atomic.Int64  // Tracks total spend (atomic for thread-safe operations)
	testSpend   atomic.Int64  // Tracks test spend (separate tracking for testing)
	views       atomic.Uint64 // Tracks the number of views
	clicks      atomic.Uint64 // Tracks the number of clicks
	leads       atomic.Uint64 // Tracks the number of leads (conversions)
	lastSync    uint64        // Timestamp of the last synchronization
}

// Profit retrieves the current profit amount as billing.Money.
func (state *StateItem) Profit() billing.Money {
	return billing.Money(state.profit.Load())
}

// TotalProfit retrieves the total profit amount as billing.Money.
func (state *StateItem) TotalProfit() billing.Money {
	return billing.Money(state.totalProfit.Load())
}

// Spend retrieves the current spend amount as billing.Money.
func (state *StateItem) Spend() billing.Money {
	return billing.Money(state.spend.Load())
}

// TotalSpend retrieves the total spend amount as billing.Money.
func (state *StateItem) TotalSpend() billing.Money {
	return billing.Money(state.totalSpend.Load())
}

// TestSpend retrieves the test spend amount as billing.Money.
func (state *StateItem) TestSpend() billing.Money {
	return billing.Money(state.testSpend.Load())
}

// Views retrieves the current view count.
func (state *StateItem) Views() uint64 {
	return state.views.Load()
}

// Clicks retrieves the current click count.
func (state *StateItem) Clicks() uint64 {
	return state.clicks.Load()
}

// Leads retrieves the current lead (conversion) count.
func (state *StateItem) Leads() uint64 {
	return state.leads.Load()
}

// Do performs an action on the state item (e.g., incrementing views, clicks, or leads) and adjusts the spend.
// If a test action is indicated, the test spend is also adjusted. Updates can propagate to the parent if defined.
func (state *StateItem) Do(test bool, vol billing.Money, action Action) {
	// Update metrics based on the specified action type
	switch action {
	case ActionView:
		state.views.Add(1)
	case ActionClick:
		state.clicks.Add(1)
	case ActionLead:
		state.leads.Add(1)
	}
	// Add volume to spend
	state.spend.Add(vol.Int64())
	// If this is a test, update testSpend as well
	if test {
		state.testSpend.Add(vol.Int64())
	}
	// If there is a parent state, propagate the update to the parent
	if state.parent != nil {
		state.parent.Do(test, vol, action)
	}
}

// update updates the StateItem based on a new StateItem's values, averaging the values by service count.
// The update occurs only if the new state is more recent and has a higher spend amount.
func (state *StateItem) update(newState *StateData, serviceCount uint64) bool {
	// Only proceed if the new state is more recent and has higher spend
	if state.lastSync >= newState.LastSync || state.Spend() >= billing.Money(newState.Spend) {
		return false
	}
	// Update state values, adjusting for the total service count (averaging across services)
	state.profit.Store(newState.Profit / int64(serviceCount))
	state.totalProfit.Store(newState.TotalProfit)
	state.spend.Store(newState.Spend / int64(serviceCount))
	state.totalSpend.Store(newState.TotalSpend)
	state.testSpend.Store(newState.TestSpend / int64(serviceCount))
	state.views.Store(newState.Views / serviceCount)
	state.clicks.Store(newState.Clicks / serviceCount)
	state.leads.Store(newState.Leads / serviceCount)
	state.lastSync = newState.LastSync // Update the last synchronization time
	return true
}
