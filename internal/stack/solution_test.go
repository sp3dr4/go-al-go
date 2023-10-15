package stack

import "testing"

func TestStack(t *testing.T) {
	v := NewMyStack[int]()

	v.Push(5)
	v.Push(7)
	v.Push(9)
	expectLen(t, v, 3)

	expectPop(t, v, 9)
	expectLen(t, v, 2)

	v.Push(11)
	expectPop(t, v, 11)
	expectPop(t, v, 7)
	expectPeek(t, v, 5)
	expectPop(t, v, 5)
	if _, err := v.Pop(); err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectLen(t, v, 0)

	v.Push(69)
	expectPeek(t, v, 69)
	expectLen(t, v, 1)
}

func expectLen[T any](t *testing.T, s Stack[T], expected int) {
	o := s.Len()
	if o != expected {
		t.Errorf("Expected %v, got %v", expected, o)
	}
}

func expectPop[T comparable](t *testing.T, s Stack[T], expected T) {
	o, err := s.Pop()
	if err != nil {
		t.Errorf("Expected %v, got %v", expected, err)
	}
	if o != expected {
		t.Errorf("Expected %v, got %v", expected, o)
	}
}

func expectPeek[T comparable](t *testing.T, s Stack[T], expected T) {
	o, err := s.Peek()
	if err != nil {
		t.Errorf("Expected %v, got %v", expected, err)
	}
	if o != expected {
		t.Errorf("Expected %v, got %v", expected, o)
	}
}
