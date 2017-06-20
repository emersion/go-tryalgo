package graph_test

import (
	"github.com/emersion/go-tryalgo/graph"
)

// Same as https://en.wikipedia.org/wiki/File:Graph.traversal.example.svg
var testGraph = graph.NewIntList([][]int{
	{1, 2, 3},
	{0, 4, 5},
	{0, 6},
	{0, 5},
	{1},
	{1, 3},
	{2},
})

var n = -1

var testValuedGraph = graph.NewIntMatrix([][]int{
	//0  1  2  3  4  5  6  7  8
	{ n, 1, n, 0, n, n, n, n, n}, // 0
	{ 1, 1, 4, n, n, n, n, n, n}, // 1
	{ n, 4, n, n, n, 3, n, n, 2}, // 2
	{ 0, n, n, n, 1, n, 3, n, n}, // 3
	{ n, n, n, 1, n, 5, n, 1, n}, // 4
	{ n, n, 3, n, 5, n, n, n, 2}, // 5
	{ n, n, n, 3, n, n, n, 6, n}, // 6
	{ n, n, n, n, 1, n, 6, n, 2}, // 7
	{ n, n, 2, n, n, 2, n, 2, n}, // 8
})
