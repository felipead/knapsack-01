package cmd

import (
	"github.com/felipead/knapsack-01/pkg/strategies"
	"log"

	"github.com/felipead/knapsack-01/pkg/input"
	"github.com/felipead/knapsack-01/pkg/model"
)

func Main() {
	// TODO: read problem file and strategy from command-line, instead of hardcoding

	problem, err := input.LoadProblemFromFile("examples/dataset1.knapsack")
	if err != nil {
		log.Fatal(err)
	}

	knapsack := model.NewKnapsack(problem.Capacity)
	err = knapsack.Solve(strategies.Greedy, problem.Items)
	if err != nil {
		log.Fatal(err)
	}
	knapsack.Print()
}
