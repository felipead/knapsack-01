package model

import (
	"github.com/shopspring/decimal"
)

type Knapsack struct {
	originalCapacity  ValueType
	allocatedCapacity ValueType
	storedItems       []Item
}

func NewKnapsack(capacity decimal.Decimal) *Knapsack {
	return &Knapsack{
		originalCapacity:  capacity,
		allocatedCapacity: decimal.Zero,
		storedItems:       make([]Item, 0),
	}
}

func (k *Knapsack) Store(item Item) {
	k.allocatedCapacity = k.allocatedCapacity.Add(item.Cost)
	k.storedItems = append(k.storedItems, item)
}

func (k *Knapsack) CanHold(item Item) bool {
	return item.Cost.LessThanOrEqual(k.RemainingCapacity())
}

func (k *Knapsack) RemainingCapacity() ValueType {
	return k.originalCapacity.Sub(k.allocatedCapacity)
}

func (k *Knapsack) GetStoredItems() []Item {
	return k.storedItems
}

func (k *Knapsack) Print() {
	// TODO
}
