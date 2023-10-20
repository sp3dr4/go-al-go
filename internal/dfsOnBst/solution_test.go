package dfsonbst

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

func TestDfs(t *testing.T) {
	if !Dfs(tree, 45) {
		t.Errorf("Expected true, got false")
	}
	if !Dfs(tree, 7) {
		t.Errorf("Expected true, got false")
	}
	if Dfs(tree, 69) {
		t.Errorf("Expected false, got true")
	}
}
