package input

import (
	"github.com/felipead/knapsack-01/pkg/model"
	"github.com/shopspring/decimal"
)

type Problem struct {
	Capacity decimal.Decimal
	Items    []model.Item
}
