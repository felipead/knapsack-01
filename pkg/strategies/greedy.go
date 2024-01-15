package strategies

import (
	"github.com/felipead/knapsack-01/pkg/model"
	"github.com/felipead/knapsack-01/pkg/sorting"
)

func CompareByValueCostRatio(a, b model.Item) int {
	return a.ValueCostRatio().Cmp(
		b.ValueCostRatio(),
	)
}

func Greedy(knapsack *model.Knapsack, items []model.Item) error {
	sorting.Quicksort(items, CompareByValueCostRatio)

	for i := len(items) - 1; i >= 0; i-- {
		item := items[i]
		if knapsack.CanHold(item) {
			knapsack.Store(item)
		}
	}

	return nil
}
