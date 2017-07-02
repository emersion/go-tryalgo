package knapsack

// Knapsack solves the knapsack problem: given some objects with weights and
// values, maximize the value without exceeding the capacity.
//
// Complexity: O(n * cap), with n number of objects
func Knapsack(weights, values []int, cap int) int {
	// TODO: optimize memory usage

	if len(values) != len(weights) {
		panic("values and weights don't have the same length")
	}

	n := len(weights)
	if n == 0 {
		return 0
	}

	// opt[i] is the set of values that can be reach with objects from 0 to i (inclusive)
	opt := make([][]int, n)

	opt[0] = make([]int, cap+1)
	for c := weights[0]; c <= cap; c++ {
		opt[0][c] = values[0]
	}

	for i := 1; i < n; i++ {
		opt[i] = make([]int, cap+1)

		for c := 0; c <= cap; c++ {
			without := opt[i-1][c]

			if c >= weights[i] {
				with := values[i] + opt[i-1][c - weights[i]]
				if with > without {
					// Take this object
					opt[i][c] = with
					continue
				}
			}

			// Don't take this object
			opt[i][c] = without
		}
	}

	return opt[n-1][cap]
}
