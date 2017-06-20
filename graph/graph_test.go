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
