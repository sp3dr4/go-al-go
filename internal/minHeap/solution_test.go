package minheap

import "testing"

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap()

	if heap.Length != 0 {
		t.Errorf("expected 0, got %v", heap.Length)
	}

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	if heap.Length != 8 {
		t.Errorf("expected 8, got %v", heap.Length)
	}
	expectDelete(t, heap, 1)
	expectDelete(t, heap, 3)
	expectDelete(t, heap, 4)
	expectDelete(t, heap, 5)
	if heap.Length != 4 {
		t.Errorf("expected 4, got %v", heap.Length)
	}
	expectDelete(t, heap, 7)
	expectDelete(t, heap, 8)
	expectDelete(t, heap, 69)
	expectDelete(t, heap, 420)
	if heap.Length != 0 {
		t.Errorf("expected 0, got %v", heap.Length)
	}
}

func expectDelete(t *testing.T, heap *MinHeap, expected int) {
	res, err := heap.Delete()
	if err != nil {
		t.Errorf("expected %v, got err %v", expected, err)
	}
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
