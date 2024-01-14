package model

import "github.com/shopspring/decimal"

type Knapsack struct {
	Capacity decimal.Decimal
}

func NewKnapsack(capacity decimal.Decimal) *Knapsack {
	return &Knapsack{
		Capacity: capacity,
	}
}

func (k *Knapsack) Solve(strategy Strategy, items []Item) error {
	return strategy(k, items)
}

func (k *Knapsack) Print() {
	// TODO
}
