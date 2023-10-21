package dijkstralist

import (
	"slices"
	"testing"
)

var list WeightedAdjacencyList = [][]GraphEdge{
	{{to: 1, weight: 3}, {to: 2, weight: 1}},
	{{to: 0, weight: 3}, {to: 2, weight: 4}, {to: 4, weight: 1}},
	{{to: 1, weight: 4}, {to: 3, weight: 7}, {to: 0, weight: 1}},
	{{to: 2, weight: 7}, {to: 4, weight: 5}, {to: 6, weight: 1}},
	{{to: 1, weight: 1}, {to: 3, weight: 5}, {to: 5, weight: 2}},
	{{to: 6, weight: 1}, {to: 4, weight: 2}, {to: 2, weight: 18}},
	{{to: 3, weight: 1}, {to: 5, weight: 1}},
}

func TestDijkstraList(t *testing.T) {
	res := DijkstraList(&list, 0, 6)
	expected := []int{0, 1, 4, 5, 6}
	if slices.Compare(res, expected) != 0 {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
