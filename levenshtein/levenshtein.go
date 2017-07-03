package levenshtein

func min(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}

func letterDist(a, b byte) int {
	if a == b {
		return 0
	}
	return 1
}

// Levenshtein computes the Levenshtein distance between w1 and w2.
func Levenshtein(w1, w2 []byte) int {
	// TODO: reduce space usage

	n, m := len(w1), len(w2)
	dist := make([][]int, n+1)
	for i, c1 := range w1 {
		if i == 0 {
			dist[i] = make([]int, m+1)
		}
		dist[i+1] = make([]int, m+1)

		for j, c2 := range w2 {
			if i == 0 || j == 0 {
				dist[i][j] = i + j
			}

			dist[i+1][j+1] = min(
				dist[i][j] + letterDist(c1, c2), // Change
				dist[i+1][j] + 1, // Delete
				dist[i][j+1] + 1, // Insert
			)
		}
	}

	return dist[n][m]
}
