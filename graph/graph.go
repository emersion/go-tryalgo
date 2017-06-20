// Package graph provides functions to handle graphs.
package graph

// A Graph contains vertices linked by edges.
type Graph interface {
	Neighbours(int) []int
}

type intList [][]int

func (l intList) Neighbours(i int) []int {
	return l[i]
}

// NewIntList creates a new integer graph from an adjacency list.
func NewIntList(l [][]int) Graph {
	return intList(l)
}

type intMatrix [][]bool

func (m intMatrix) Neighbours(i int) []int {
	var l []int
	for j, v := range m[i] {
		if v {
			l = append(l, j)
		}
	}
	return l
}

// NewIntMatrix creates a new integer graph from an adjacency matrix.
func NewIntMatrix(m [][]bool) Graph {
	return intMatrix(m)
}
