package doublylinkedlist

import "testing"

func TestLinkedList(t *testing.T) {
	l := NewDoublyLinkedList[int]()

	l.Append(5)
	l.Append(7)
	l.Append(9)
	expectLen(t, l, 3)

	expectGet(t, l, 2, 9)
	expectRemoveAt(t, l, 1, 7)
	expectLen(t, l, 2)

	l.Append(11)
	expectRemoveAt(t, l, 1, 9)
	if _, err := l.Remove(9); err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectRemoveAt(t, l, 0, 5)
	expectRemoveAt(t, l, 0, 11)
	expectLen(t, l, 0)

	l.Prepend(5)
	l.Prepend(7)
	l.Prepend(9)

	expectGet(t, l, 2, 5)
	expectGet(t, l, 0, 9)
	expectRemove(t, l, 9)
	expectLen(t, l, 2)
	expectGet(t, l, 0, 7)
}

func expectLen[T any](t *testing.T, l LinkedList[T], expected int) {
	o := l.Len()
	if o != expected {
		t.Errorf("Expected %v, got %v", expected, o)
	}
}

func expectGet[T comparable](t *testing.T, l LinkedList[T], idx int, expected T) {
	v, err := l.Get(idx)
	if err != nil {
		t.Errorf("Expected %v, got %v", expected, err)
	}
	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func expectRemoveAt[T comparable](t *testing.T, l LinkedList[T], idx int, expected T) {
	v, err := l.RemoveAt(idx)
	if err != nil {
		t.Errorf("Expected %v, got %v", expected, err)
	}
	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func expectRemove[T comparable](t *testing.T, l LinkedList[T], v T) {
	o, err := l.Remove(v)
	if err != nil {
		t.Errorf("Expected %v, got %v", v, err)
	}
	if o != v {
		t.Errorf("Expected %v, got %v", v, o)
	}
}
