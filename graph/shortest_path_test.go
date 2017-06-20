package graph_test

import (
	"reflect"
	"testing"

	"github.com/emersion/go-tryalgo/graph"
)

func TestDijkstra(t *testing.T) {
	l := graph.Dijkstra(testValuedGraph, 0, 8)
	want := []int{0, 1, 2, 8}
	if !reflect.DeepEqual(l, want) {
		t.Errorf("Dijkstra(0, 8) = %v, want %v", l, want)
	}
}
