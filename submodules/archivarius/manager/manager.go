package manager

import (
	"context"
	"sync"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/adcorelib/billing"
	"github.com/geniusrabbit/adcorelib/fasttime"
)

type ServerCounter interface {
	Value() uint64
}

// Action represents different types of actions that can be performed on a state.
type Action int

const (
	// ActionView represents a "view" action on an advertisement or campaign.
	ActionView Action = iota + 1
	// ActionClick represents a "click" action on an advertisement or campaign.
	ActionClick
	// ActionLead represents a "lead" action, indicating a conversion or user interest.
	ActionLead
)

// Manager manages different types of advertisement-related states, providing
// safe concurrent access to state items for campaigns, ads, access points, etc.
type Manager struct {
	serviceCount         ServerCounter // Tracks the total number of services managed
	lastSync             uint64        // Timestamp of the last synchronization
	syncInterval         uint64        // Interval between synchronization attempts
	lastSyncUpdatedCount uint64        //
	lastSyncTotalCount   uint64        //

	loader StateLoader // StateLoader interface for fetching state updates

	// Separate sync.Maps to cache different types of StateItems by ID.
	// Each sync.Map provides concurrent access and thread-safe storage.
	campCache sync.Map // Caches campaign states
	adCache   sync.Map // Caches ad states
	apCache   sync.Map // Caches access point states
	asCache   sync.Map // Caches ad source states
	acCache   sync.Map // Caches account states
}

// NewManager creates a new Manager instance with initialized caches.
func NewManager(loader StateLoader, serverCounter ServerCounter, syncInterval uint64) *Manager {
	return &Manager{
		serviceCount: serverCounter,
		lastSync:     0,
		syncInterval: syncInterval,
		loader:       loader,
	}
}

// Do performs a specified action on the given state, updating its metrics based on
// whether it's a test action or not. Uses a specific amount `vol` as part of the update.
func (m *Manager) Do(state State, vol billing.Money, action Action, test bool) {
	state.(*StateItem).Do(test, vol, action)
}

// CampaignState retrieves the campaign state by ID, creating a new state if it doesn't exist.
func (m *Manager) CampaignState(id uint64) *StateItem {
	return m.getState(&m.campCache, id)
}

// AdState retrieves the ad state by ID, creating a new state if it doesn't exist.
func (m *Manager) AdState(id, campID uint64) *StateItem {
	campState := m.CampaignState(campID)
	state := m.getState(&m.adCache, id)
	state.parent = campState
	return state
}

// AccessPointState retrieves the access point state by ID, creating a new state if it doesn't exist.
func (m *Manager) AccessPointState(id uint64) *StateItem {
	return m.getState(&m.apCache, id)
}

// AdSourceState retrieves the ad source state by ID, creating a new state if it doesn't exist.
func (m *Manager) AdSourceState(id uint64) *StateItem {
	return m.getState(&m.asCache, id)
}

// AccountState retrieves the account state by ID, creating a new state if it doesn't exist.
func (m *Manager) AccountState(id uint64) *BalanceStateItem {
	// Attempt to retrieve the state from the cache
	if state, ok := m.acCache.Load(id); ok {
		return state.(*BalanceStateItem)
	}
	// If not found, create a new BalanceStateItem and attempt to store it in the cache
	newState := &BalanceStateItem{}
	actual, loaded := m.acCache.LoadOrStore(id, newState)
	if loaded {
		return actual.(*BalanceStateItem)
	}
	return newState
}

// Update updates the state of all managed entities by fetching the latest data from the data source.
func (m *Manager) Update(ctx context.Context) (_, _ uint64, _ error) {
	m.lastSyncUpdatedCount = 0
	m.lastSyncTotalCount = 0
	tm := fasttime.UnixTimestampNano()
	err := m.loader.StreamUpdate(ctx,
		gocast.IfThen(m.lastSync+m.syncInterval > tm, m.lastSync, 0),
		m.updateItemIter)
	if err == nil {
		m.lastSync = tm
	}
	return m.lastSyncTotalCount, m.lastSyncUpdatedCount, err
}

func (m *Manager) updateItemIter(itemType ItemType, id uint64, state *StateData) error {
	updated := false
	m.lastSyncTotalCount++
	switch itemType {
	case ItemCampaign:
		updated = m.updateItem(&m.campCache, id, state)
	case ItemAd:
		updated = m.updateItem(&m.adCache, id, state)
	case ItemAccessPoint:
		updated = m.updateItem(&m.apCache, id, state)
	case ItemAdSource:
		updated = m.updateItem(&m.asCache, id, state)
	case ItemAccount:
		accState := m.AccountState(id)
		updated = accState.update(state, m.serviceCount.Value())
	}
	if updated {
		m.lastSyncUpdatedCount++
	}
	return nil
}

// updateItem updates the existing state in the specified cache with a new StateItem's values.
// If the state doesn't exist, it will be created by calling `getState`.
func (m *Manager) updateItem(cache *sync.Map, id uint64, newState *StateData) bool {
	oldState := m.getState(cache, id)
	return oldState.update(newState, m.serviceCount.Value())
}

// getState retrieves the state from the specified cache by ID.
// If the state does not exist, a new StateItem is created, stored in the cache, and returned.
func (m *Manager) getState(cache *sync.Map, id uint64) *StateItem {
	// Attempt to retrieve the state from the cache
	if state, ok := cache.Load(id); ok {
		return state.(*StateItem)
	}
	// If not found, create a new StateItem and attempt to store it in the cache
	newState := &StateItem{}
	actual, loaded := cache.LoadOrStore(id, newState)
	if loaded {
		return actual.(*StateItem) // If another goroutine stored it first, use the existing one
	}
	return newState // Otherwise, return the newly created StateItem
}
