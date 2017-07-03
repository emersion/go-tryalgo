package knapsack_test

import (
	"testing"

	"github.com/emersion/go-tryalgo/knapsack"
)

var knapsackTests = []struct {
	weights, values []int
	cap             int
	opt             int
}{
	{
		weights: []int{2, 3, 4, 5},
		values:  []int{3, 4, 5, 6},
		cap:     5,
		opt:     7,
	},
	{
		weights: []int{580, 1616, 1906, 1942, 50, 294},
		values:  []int{874, 620, 345, 269, 360, 470},
		cap:     2000,
		opt:     1704,
	},
	{
		weights: []int{5, 4, 3, 2, 1},
		values:  []int{30, 19, 20, 10, 20},
		cap:     10,
		opt:     70,
	},
}

func TestKnapsack(t *testing.T) {
	for _, test := range knapsackTests {
		got := knapsack.Knapsack(test.weights, test.values, test.cap)
		if got != test.opt {
			t.Errorf("Knapsack(%v, %v, %v) = %v, want %v", test.weights, test.values, test.cap, got, test.opt)
		}
	}
}
