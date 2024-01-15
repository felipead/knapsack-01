package model

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItem_ValueCostRatio(t *testing.T) {
	assert.Equal(t,
		"0.3125",
		Item{
			Value: decimal.NewFromInt(5),
			Cost:  decimal.NewFromInt(16),
		}.ValueCostRatio().String(),
	)

	assert.Equal(t,
		"16.32",
		Item{
			Value: decimal.RequireFromString("16.32"),
			Cost:  decimal.NewFromInt(1),
		}.ValueCostRatio().String(),
	)

	assert.Equal(t,
		"0.0001545595054096",
		Item{
			Value: decimal.RequireFromString("0.0005"),
			Cost:  decimal.RequireFromString("3.235"),
		}.ValueCostRatio().String(),
	)

	assert.Equal(t,
		"1",
		Item{
			Value: decimal.NewFromInt(1),
			Cost:  decimal.NewFromInt(1),
		}.ValueCostRatio().String(),
	)
}
