# Manager Module Documentation

## Table of Contents

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Overview](#overview)
    - [Types](#types)
    - [Constants](#constants)
4. [Manager](#manager)
    - [Fields](#fields)
    - [Constructor](#constructor)
    - [Methods](#methods)
5. [StateItem](#stateitem)
    - [Fields description](#fields-description)
    - [Methods description](#methods-description)
6. [Usage Example](#usage-example)
7. [Interfaces](#interfaces)
    - [StateLoader](#stateloader)
    - [State](#state)
    - [ItemType](#itemtype)
8. [Concurrency Considerations](#concurrency-considerations)

---

## Introduction

The **Manager** module is designed to handle and track various advertisement-related states in a concurrent environment. It efficiently manages different types of states such as campaigns, ads, access points, ad sources, and accounts using thread-safe data structures and atomic operations. This ensures high performance and reliability in multi-threaded applications.

---

## Installation

To use the Manager module in your Go project, ensure you have Go installed and properly set up. Then, you can import the module as follows:

```go
import "github.com/geniusrabbit/adcorelib/manager"
```

Ensure you also have the necessary dependencies installed:

```bash
go get github.com/geniusrabbit/adcorelib/billing
go get github.com/geniusrabbit/adcorelib/fasttime
go get github.com/demdxx/gocast/v2
```

---

## Overview

The Manager module consists of several key components:

- **Types**: Defines the primary structures and interfaces used.
- **Constants**: Enumerates the different action types.
- **Manager**: Central structure managing various state caches.
- **StateItem**: Represents individual state entities with atomic counters.

### Types

- `Action`: Enum representing different types of actions (`View`, `Click`, `Lead`).
- `Manager`: Manages different advertisement-related states.
- `StateItem`: Tracks performance metrics for campaigns, ads, etc.
- `StateLoader`: Interface for fetching state updates (assumed to be defined elsewhere).

### Constants

- `ActionView`, `ActionClick`, `ActionLead`: Enumerated actions that can be performed on states.

---

## Manager

The `Manager` struct is responsible for managing different types of advertisement-related states. It ensures thread-safe access and updates to these states using `sync.Map` and atomic operations.

### Fields

```go
type Manager struct {
    serviceCount uint64       // Tracks the total number of services managed
    lastSync     uint64       // Timestamp of the last synchronization
    syncInterval uint64       // Interval between synchronization attempts

    loader       StateLoader  // StateLoader interface for fetching state updates

    campCache    sync.Map     // Caches campaign states
    adCache      sync.Map     // Caches ad states
    apCache      sync.Map     // Caches access point states
    asCache      sync.Map     // Caches ad source states
    acCache      sync.Map     // Caches account states
}
```

### Constructor

Creates a new instance of `Manager` with initialized caches.

```go
func NewManager(loader StateLoader, serverCount, syncInterval uint64) *Manager
```

**Parameters:**

- `loader` (`StateLoader`): Interface for fetching state updates.
- `serverCount` (`uint64`): Number of services managed.
- `syncInterval` (`uint64`): Interval between synchronization attempts.

**Returns:**

- `*Manager`: A pointer to the newly created `Manager` instance.

### Methods

#### `Do`

Performs a specified action on the given state, updating its metrics based on whether it's a test action or not. Uses a specific amount `vol` as part of the update.

```go
func (m *Manager) Do(state State, vol billing.Money, action Action, test bool)
```

**Parameters:**

- `state` (`State`): The state on which to perform the action.
- `vol` (`billing.Money`): The monetary volume associated with the action.
- `action` (`Action`): The type of action to perform (`View`, `Click`, `Lead`).
- `test` (`bool`): Indicates whether the action is a test action.

#### `CampaignState`

Retrieves the campaign state by ID, creating a new state if it doesn't exist.

```go
func (m *Manager) CampaignState(id uint64) *StateItem
```

**Parameters:**

- `id` (`uint64`): The unique identifier for the campaign.

**Returns:**

- `*StateItem`: Pointer to the campaign's `StateItem`.

#### `AdState`

Retrieves the ad state by ID, creating a new state if it doesn't exist.

```go
func (m *Manager) AdState(id uint64) *StateItem
```

**Parameters:**

- `id` (`uint64`): The unique identifier for the ad.

**Returns:**

- `*StateItem`: Pointer to the ad's `StateItem`.

#### `AccessPointState`

Retrieves the access point state by ID, creating a new state if it doesn't exist.

```go
func (m *Manager) AccessPointState(id uint64) *StateItem
```

**Parameters:**

- `id` (`uint64`): The unique identifier for the access point.

**Returns:**

- `*StateItem`: Pointer to the access point's `StateItem`.

#### `AdSourceState`

Retrieves the ad source state by ID, creating a new state if it doesn't exist.

```go
func (m *Manager) AdSourceState(id uint64) *StateItem
```

**Parameters:**

- `id` (`uint64`): The unique identifier for the ad source.

**Returns:**

- `*StateItem`: Pointer to the ad source's `StateItem`.

#### `AccountState`

Retrieves the account state by ID, creating a new state if it doesn't exist.

```go
func (m *Manager) AccountState(id uint64) *StateItem
```

**Parameters:**

- `id` (`uint64`): The unique identifier for the account.

**Returns:**

- `*StateItem`: Pointer to the account's `StateItem`.

#### `Update`

Updates the state of all managed entities by fetching the latest data from the data source.

```go
func (m *Manager) Update(ctx context.Context) error
```

**Parameters:**

- `ctx` (`context.Context`): The context for managing cancellation and timeouts.

**Returns:**

- `error`: An error if the update fails, otherwise `nil`.

#### `updateItemIter`

Internal method to iterate and update items based on their type.

```go
func (m *Manager) updateItemIter(itemType ItemType, id uint64, state *StateItem) error
```

**Parameters:**

- `itemType` (`ItemType`): The type of the item being updated.
- `id` (`uint64`): The unique identifier for the item.
- `state` (`*StateItem`): The new state data to update.

**Returns:**

- `error`: An error if the update fails, otherwise `nil`.

#### `updateItem`

Updates the existing state in the specified cache with a new `StateItem`'s values. If the state doesn't exist, it will be created by calling `getState`.

```go
func (m *Manager) updateItem(cache *sync.Map, id uint64, newState *StateItem)
```

**Parameters:**

- `cache` (`*sync.Map`): The cache map where the state is stored.
- `id` (`uint64`): The unique identifier for the state.
- `newState` (`*StateItem`): The new state data to update.

#### `getState`

Retrieves the state from the specified cache by ID. If the state does not exist, a new `StateItem` is created, stored in the cache, and returned.

```go
func (m *Manager) getState(cache *sync.Map, id uint64) *StateItem
```

**Parameters:**

- `cache` (`*sync.Map`): The cache map where the state is stored.
- `id` (`uint64`): The unique identifier for the state.

**Returns:**

- `*StateItem`: Pointer to the retrieved or newly created `StateItem`.

---

## StateItem

The `StateItem` struct represents an individual state entity, tracking various performance metrics such as profit, spend, views, clicks, and leads. It uses atomic operations to ensure thread-safe updates in a concurrent environment.

### Fields description

```go
type StateItem struct {
    parent    *StateItem    // Parent StateItem for hierarchical updates
    profit    atomic.Int64  // Total profit, atomically managed
    spend     atomic.Int64  // Total spend, atomically managed
    testSpend atomic.Int64  // Total test spend, atomically managed
    views     atomic.Uint64 // Number of views, atomically managed
    clicks    atomic.Uint64 // Number of clicks, atomically managed
    leads     atomic.Uint64 // Number of leads (conversions), atomically managed
    lastSync  uint64        // Timestamp of the last synchronization
}
```

### Methods description

#### `Profit`

Retrieves the current profit amount.

```go
func (state *StateItem) Profit() billing.Money
```

**Returns:**

- `billing.Money`: The current profit.

#### `Spend`

Retrieves the current spend amount.

```go
func (state *StateItem) Spend() billing.Money
```

**Returns:**

- `billing.Money`: The current spend.

#### `TestSpend`

Retrieves the current test spend amount.

```go
func (state *StateItem) TestSpend() billing.Money
```

**Returns:**

- `billing.Money`: The current test spend.

#### `Views`

Retrieves the current view count.

```go
func (state *StateItem) Views() uint64
```

**Returns:**

- `uint64`: The number of views.

#### `Clicks`

Retrieves the current click count.

```go
func (state *StateItem) Clicks() uint64
```

**Returns:**

- `uint64`: The number of clicks.

#### `Leads`

Retrieves the current lead (conversion) count.

```go
func (state *StateItem) Leads() uint64
```

**Returns:**

- `uint64`: The number of leads.

#### `Do`

Performs an action on the state item (e.g., incrementing views, clicks, or leads) and adjusts the spend. If a test action is indicated, the test spend is also adjusted. Updates can propagate to the parent if defined.

```go
func (state *StateItem) Do(test bool, vol billing.Money, action Action)
```

**Parameters:**

- `test` (`bool`): Indicates if the action is a test action.
- `vol` (`billing.Money`): The monetary volume associated with the action.
- `action` (`Action`): The type of action to perform (`View`, `Click`, `Lead`).

**Behavior:**

1. Increments the corresponding metric (`views`, `clicks`, `leads`) based on the action.
2. Adds the volume `vol` to `spend`.
3. If `test` is `true`, adds the volume `vol` to `testSpend`.
4. Propagates the update to the `parent` if it exists.

#### `update`

Updates the `StateItem` based on a new `StateItem`'s values, averaging the values by `serviceCount`. The update occurs only if the new state is more recent and has a higher spend amount.

```go
func (state *StateItem) update(newState *StateItem, serviceCount uint64) bool
```

**Parameters:**

- `newState` (`*StateItem`): The new state data to update from.
- `serviceCount` (`uint64`): The total number of services, used for averaging.

**Returns:**

- `bool`: `true` if the update was successful, `false` otherwise.

**Behavior:**

1. Checks if the `newState` has a more recent `lastSync` timestamp and a higher `spend` value.
2. If conditions are met, updates the current state by averaging the new state's values with `serviceCount`.
3. Updates the `lastSync` timestamp.
4. Returns `true` if the update was performed, `false` otherwise.

---

## Usage Example

Below is an example demonstrating how to use the `Manager` and `StateItem` to manage advertisement states.

```go
package main

import (
    "context"
    "fmt"

    "github.com/geniusrabbit/adcorelib/billing"
    "github.com/geniusrabbit/adcorelib/manager"
)

// MockStateLoader is a mock implementation of the StateLoader interface for demonstration.
type MockStateLoader struct{}

// StreamUpdate is a mock implementation that does nothing.
func (loader *MockStateLoader) StreamUpdate(ctx context.Context, since uint64, iterFunc func(itemType manager.ItemType, id uint64, state *manager.StateItem) error) error {
    // Mock implementation: no updates
    return nil
}

func main() {
    // Initialize a new Manager with a mock StateLoader, service count, and sync interval
    loader := &MockStateLoader{}
    serviceCount := uint64(10)
    syncInterval := uint64(60) // e.g., 60 seconds

    mgr := manager.NewManager(loader, serviceCount, syncInterval)

    // Example ID for a campaign
    campaignID := uint64(12345)

    // Retrieve or create the campaign state
    campaignState := mgr.CampaignState(campaignID)

    // Perform a "view" action
    mgr.Do(campaignState, billing.Money(100), manager.ActionView, false)

    // Perform a "click" action as a test
    mgr.Do(campaignState, billing.Money(50), manager.ActionClick, true)

    // Retrieve and print the updated state
    fmt.Printf("Campaign ID: %d\n", campaignID)
    fmt.Printf("Views: %d\n", campaignState.Views())
    fmt.Printf("Clicks: %d\n", campaignState.Clicks())
    fmt.Printf("Spend: %d\n", campaignState.Spend())
    fmt.Printf("Test Spend: %d\n", campaignState.TestSpend())

    // Update all states (mocked)
    ctx := context.Background()
    err := mgr.Update(ctx)
    if err != nil {
        fmt.Printf("Update failed: %v\n", err)
    } else {
        fmt.Println("Update successful.")
    }

    // Delete the campaign state when it's no longer needed
    mgr.DeleteCampaignState(campaignID)
}
```

**Output:**

```ini
Campaign ID: 12345
Views: 1
Clicks: 1
Spend: 150
Test Spend: 50
Update successful.
```

---

## Interfaces

### StateLoader

The `StateLoader` interface is responsible for fetching state updates from a data source. It is used by the `Manager` to keep the states synchronized.

```go
type StateLoader interface {
    // StreamUpdate streams updates starting from a given timestamp.
    // It takes a context for cancellation, a timestamp `since`, and an iterator function `iterFunc` 
    // to process each update.
    StreamUpdate(ctx context.Context, since uint64, iterFunc func(itemType ItemType, id uint64, state *StateItem) error) error
}
```

**Methods:**

- `StreamUpdate(ctx context.Context, since uint64, iterFunc func(itemType ItemType, id uint64, state *StateItem) error) error`: Streams updates from a data source starting from the specified timestamp and applies them using the provided iterator function.

### State

The `State` interface defines the methods that any state entity must implement. This allows the `Manager` to interact with different state types uniformly.

```go
type State interface {
    Profit() billing.Money
    Spend() billing.Money
    TestSpend() billing.Money
    Views() uint64
    Clicks() uint64
    Leads() uint64
    Do(test bool, vol billing.Money, action Action)
}
```

**Methods:**

- `Profit() billing.Money`: Retrieves the current profit.
- `Spend() billing.Money`: Retrieves the current spend.
- `TestSpend() billing.Money`: Retrieves the current test spend.
- `Views() uint64`: Retrieves the number of views.
- `Clicks() uint64`: Retrieves the number of clicks.
- `Leads() uint64`: Retrieves the number of leads.
- `Do(test bool, vol billing.Money, action Action)`: Performs an action on the state.

### ItemType

The `ItemType` enum represents the different types of items managed by the `Manager`. It is used to categorize state updates.

```go
type ItemType int

const (
    ItemCampaign ItemType = iota + 1
    ItemAd
    ItemAccessPoint
    ItemAdSource
    ItemAccount
)
```

---

## Concurrency Considerations

The Manager module is designed for concurrent environments where multiple goroutines may access and modify states simultaneously. Key concurrency features include:

- **`sync.Map`**: Utilized for caching different state types, providing thread-safe access without manual locking.
- **Atomic Operations**: The `StateItem` struct uses atomic types (`atomic.Int64`, `atomic.Uint64`) to ensure thread-safe read and write operations on metrics.
- **Parent State Propagation**: When updating a state, changes can propagate to parent states safely using atomic operations.

These features ensure that the Manager module can handle high-throughput and concurrent access patterns efficiently.

---
