package btbfs

import "testing"

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

func TestBtBfs(t *testing.T) {
	if !Bfs(tree, 45) {
		t.Errorf("(%v) Expected true, got false", 45)
	}
	if !Bfs(tree, 7) {
		t.Errorf("(%v) Expected true, got false", 7)
	}
	if Bfs(tree, 69) {
		t.Errorf("(%v) Expected false, got true", 69)
	}
}
