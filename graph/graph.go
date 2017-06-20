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

type boolMatrix [][]bool

func (m boolMatrix) Neighbours(i int) []int {
	var l []int
	for j, v := range m[i] {
		if v {
			l = append(l, j)
		}
	}
	return l
}

// NewBoolMatrix creates a new integer graph from an adjacency matrix.
func NewBoolMatrix(m [][]bool) Graph {
	return boolMatrix(m)
}

// ValuedGraph is a graph with values on edges.
type ValuedGraph interface {
	Graph

	// Distance returns the distance from i to its neighbour j.
	Distance(i, j int) int
}

type intMatrix [][]int

func (m intMatrix) Neighbours(i int) []int {
	var l []int
	for j, v := range m[i] {
		if v >= 0 {
			// TODO: support negative values
			l = append(l, j)
		}
	}
	return l
}

func (m intMatrix) Distance(i, j int) int {
	return m[i][j]
}

// NewIntMatrix creates a new integer valued graph from an adjacency matrix.
func NewIntMatrix(m [][]int) ValuedGraph {
	return intMatrix(m)
}
