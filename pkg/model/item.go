package model

import "github.com/shopspring/decimal"

type Item struct {
	Value decimal.Decimal
	Cost  decimal.Decimal
}
