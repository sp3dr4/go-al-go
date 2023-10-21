package dfsgraphlist

import (
	"slices"
	"testing"
)

var list WeightedAdjacencyList = [][]GraphEdge{
	{{to: 1}, {to: 2}},
	{{to: 4}},
	{{to: 3}},
	{},
	{{to: 1}, {to: 3}, {to: 5}},
	{{to: 2}, {to: 6}},
	{{to: 3}},
}

func TestDfs(t *testing.T) {
	res := Dfs(&list, 0, 6)
	expected := []int{0, 1, 4, 5, 6}
	if slices.Compare(res, expected) != 0 {
		t.Errorf("expected %v, got %v", expected, res)
	}

	res = Dfs(&list, 6, 0)
	expected = []int{}
	if slices.Compare(res, expected) != 0 {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
