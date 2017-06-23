package sort

// Insertion is the insertion sort. Complexity: O(n^2)
func Insertion(a []int) {
	for i := 1; i < len(a); i++ {
		v := a[i]
		j := i
		for j > 0 {
			if v >= a[j-1] {
				break
			}
			a[j] = a[j-1]
			j--
		}
		a[j] = v
	}
}
