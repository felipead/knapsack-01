package main

import (
	"log"

	"github.com/felipead/knapsack-01/pkg/input"
	"github.com/felipead/knapsack-01/pkg/strategies"
)

func main() {
	// TODO: read problem file and strategy from command-line, instead of hardcoding

	problem, err := input.LoadProblemFromFile("examples/dataset1.knapsack")
	if err != nil {
		log.Fatal(err)
	}

	knapsack, err := problem.Solve(strategies.Greedy)
	if err != nil {
		log.Fatal(err)
	}
	knapsack.Print()
}
