package manager

import "context"

type ItemType int

const (
	ItemCampaign ItemType = iota + 1
	ItemAd
	ItemAccessPoint
	ItemAdSource
	ItemAccount
)

type ItemIterator func(itemType ItemType, id uint64, state *StateData) error

type StateLoader interface {
	StreamUpdate(ctx context.Context, lastSync uint64, iter ItemIterator) error
}
