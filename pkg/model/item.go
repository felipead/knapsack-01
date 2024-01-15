package model

import "github.com/shopspring/decimal"

type Item struct {
	Index int
	Value decimal.Decimal
	Cost  decimal.Decimal
}
