package btpreorder

import (
	"slices"
	"testing"
)

var tree = &BinaryNode[int]{
	value: 20,
	right: &BinaryNode[int]{
		value: 50,
		right: &BinaryNode[int]{
			value: 100,
			right: nil,
			left:  nil,
		},
		left: &BinaryNode[int]{
			value: 30,
			right: &BinaryNode[int]{
				value: 45,
				right: nil,
				left:  nil,
			},
			left: &BinaryNode[int]{
				value: 29,
				right: nil,
				left:  nil,
			},
		},
	},
	left: &BinaryNode[int]{
		value: 10,
		right: &BinaryNode[int]{
			value: 15,
			right: nil,
			left:  nil,
		},
		left: &BinaryNode[int]{
			value: 5,
			right: &BinaryNode[int]{
				value: 7,
				right: nil,
				left:  nil,
			},
			left: nil,
		},
	},
}

func TestPreOrderSearch(t *testing.T) {
	expected := []int{
		20,
		10,
		5,
		7,
		15,
		50,
		30,
		29,
		45,
		100,
	}
	result := PreOrderSearch(tree)
	if slices.Compare(*result, expected) != 0 {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
