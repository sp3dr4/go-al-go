package binarysearchlist

import "testing"

var haystack = []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

func TestLinearSearchList(t *testing.T) {
	tests := []struct {
		list     []int
		val      int
		expected bool
	}{
		{haystack, 69, true},
		{haystack, 1336, false},
		{haystack, 69420, true},
		{haystack, 69421, false},
		{haystack, 1, true},
		{haystack, 0, false},
	}

	for _, test := range tests {
		result := BinarySearchList(haystack, test.val)
		if result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}
