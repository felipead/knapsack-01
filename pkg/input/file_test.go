package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldLoadSmallProblemInputFromFile(t *testing.T) {
	problem, err := LoadProblemFromFile("./testdata/problem_small.knapsack")
	assert.NoError(t, err)
	assert.Equal(t, len(problem.Items), 20)
	assert.Equal(t, problem.Capacity.String(), "2.5")

	assert.Equal(t, problem.Items[0].Value.String(), "0.703562")
	assert.Equal(t, problem.Items[0].Cost.String(), "0.751231")

	assert.Equal(t, problem.Items[1].Value.String(), "0.658012")
	assert.Equal(t, problem.Items[1].Cost.String(), "0.0562173")

	assert.Equal(t, problem.Items[12].Value.String(), "0.530008")
	assert.Equal(t, problem.Items[12].Cost.String(), "0.556193")

	assert.Equal(t, problem.Items[13].Value.String(), "0.180224")
	assert.Equal(t, problem.Items[13].Cost.String(), "0.507427")

	assert.Equal(t, problem.Items[18].Value.String(), "0.0189656")
	assert.Equal(t, problem.Items[18].Cost.String(), "0.5895")

	assert.Equal(t, problem.Items[19].Value.String(), "0.725904")
	assert.Equal(t, problem.Items[19].Cost.String(), "0.320655")
}

func TestShouldLoadLargeProblemInputFromFile(t *testing.T) {
	problem, err := LoadProblemFromFile("./testdata/problem_large.knapsack")
	assert.NoError(t, err)
	assert.Equal(t, len(problem.Items), 10000)
	assert.Equal(t, problem.Capacity.String(), "1250")

	assert.Equal(t, problem.Items[0].Value.String(), "0.690495")
	assert.Equal(t, problem.Items[0].Cost.String(), "0.0187375")

	assert.Equal(t, problem.Items[9999].Value.String(), "0.0952599")
	assert.Equal(t, problem.Items[9999].Cost.String(), "0.521027")
}

func TestShouldFailToLoadProblemInputFromInvalidFiles(t *testing.T) {
	var err error

	_, err = LoadProblemFromFile("./testdata/problem_empty.knapsack")
	assert.ErrorContains(t, err, "unable to read first line as Number of Items: file stream ended with EOF")

	_, err = LoadProblemFromFile("./testdata/problem_invalid_number_of_items.knapsack")
	assert.ErrorContains(t, err, "unable to read first line as Number of Items")

	_, err = LoadProblemFromFile("./testdata/problem_invalid_capacity.knapsack")
	assert.ErrorContains(t, err, "unable to read second line as Capacity")

	_, err = LoadProblemFromFile("./testdata/problem_invalid_item_value.knapsack")
	assert.ErrorContains(t, err, "unable to read Item #6: failed to parse Value field")

	_, err = LoadProblemFromFile("./testdata/problem_invalid_item_cost.knapsack")
	assert.ErrorContains(t, err, "unable to read Item #11: failed to parse Cost field")

	_, err = LoadProblemFromFile("./testdata/problem_invalid_item_line.knapsack")
	assert.ErrorContains(t, err, "unable to read Item #11: line was expected to have 2 fields, got 4")

	_, err = LoadProblemFromFile("./testdata/problem_fewer_items_than_declared.knapsack")
	assert.ErrorContains(t, err, "unable to read Item #21: file stream ended with EOF")
}
