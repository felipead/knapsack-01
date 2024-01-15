package model

type Problem struct {
	Capacity ValueType
	Items    []Item
}

func (p *Problem) Solve(strategy Strategy) (*Knapsack, error) {
	k := NewKnapsack(p.Capacity)
	err := strategy(k, p.Items)
	return k, err
}
