package input

import (
	"github.com/shopspring/decimal"

	"github.com/felipead/knapsack-01/pkg/model"
)

type Problem struct {
	Capacity decimal.Decimal
	Items    []model.Item
}
