package queue

import "testing"

func TestQueue(t *testing.T) {
	v := NewMyQueue[int]()

	v.Enqueue(5)
	v.Enqueue(7)
	v.Enqueue(9)
	expectLen(t, v, 3)

	expectDeque[int](t, v, 5)
	expectLen(t, v, 2)

	v.Enqueue(11)
	expectDeque[int](t, v, 7)
	expectDeque[int](t, v, 9)
	expectPeek[int](t, v, 11)
	expectDeque[int](t, v, 11)
	if _, err := v.Deque(); err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectLen(t, v, 0)

	v.Enqueue(69)
	expectPeek[int](t, v, 69)
	expectLen(t, v, 1)
}

func expectLen[T any](t *testing.T, q Queue[T], expected int) {
	o := q.Len()
	if o != expected {
		t.Errorf("Expected %v, got %v", expected, o)
	}
}

func expectDeque[T comparable](t *testing.T, q Queue[T], expected T) {
	o, err := q.Deque()
	if err != nil {
		t.Errorf("Expected %v, got %v", expected, err)
	}
	if o != expected {
		t.Errorf("Expected %v, got %v", expected, o)
	}
}

func expectPeek[T comparable](t *testing.T, q Queue[T], expected T) {
	o, err := q.Peek()
	if err != nil {
		t.Errorf("Expected %v, got %v", expected, err)
	}
	if o != expected {
		t.Errorf("Expected %v, got %v", expected, o)
	}
}
