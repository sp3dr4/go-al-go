package quicksort

import (
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}

	QuickSort(arr)

	expected := []int{3, 4, 7, 9, 42, 69, 420}
	comparison := slices.Compare(arr, expected)
	if comparison != 0 {
		t.Errorf("Expected compare 0, got %v", comparison)
	}
}
