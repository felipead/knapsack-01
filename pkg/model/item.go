package model

type Item struct {
	Index int
	Value ValueType
	Cost  ValueType
}

func (i Item) ValueCostRatio() ValueType {
	// We don't support items with zero cost. That would create a division by zero.
	return i.Value.Div(i.Cost)
}
