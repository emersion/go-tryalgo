package sort_test

var testArray = []int{7, 2, 10, 5, 3, 1, 6, 4, 9, 8}
var testSorted = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func newTestArray() []int {
	a := make([]int, len(testArray))
	for i, v := range testArray {
		a[i] = v
	}
	return a
}
