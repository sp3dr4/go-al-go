package stack

import "errors"

type Stack[T any] interface {
	Push(T)
	Pop() (T, error)
	Peek() (T, error)
	Len() int
}

type MyNode[T any] struct {
	value T
	prev  *MyNode[T]
}

type MyStack[T any] struct {
	length int
	head   *MyNode[T]
}

func NewMyStack[T any]() *MyStack[T] {
	return &MyStack[T]{
		head:   nil,
		length: 0,
	}
}

func (s *MyStack[T]) Push(item T) {
	node := &MyNode[T]{value: item}
	if s.length == 0 {
		s.head = node
	} else {
		node.prev = s.head
		s.head = node
	}
	s.length += 1
}

func (s *MyStack[T]) Pop() (T, error) {
	if s.head == nil {
		var empty T
		return empty, errors.New("queue empty")
	}
	s.length -= 1
	head := s.head
	s.head = s.head.prev
	head.prev = nil
	if s.length == 0 {
		s.head = nil
	}
	return head.value, nil
}

func (s *MyStack[T]) Peek() (T, error) {
	if s.head == nil {
		var empty T
		return empty, errors.New("queue empty")
	}
	return s.head.value, nil
}

func (s *MyStack[T]) Len() int {
	return s.length
}
