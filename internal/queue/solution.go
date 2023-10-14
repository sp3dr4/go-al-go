package queue

import "errors"

type Queue[T any] interface {
	Enqueue(T)
	Deque() (T, error)
	Peek() (T, error)
	Len() int
}

type MyNode[T any] struct {
	value T
	next  *MyNode[T]
}

type MyQueue[T any] struct {
	length int
	head   *MyNode[T]
	tail   *MyNode[T]
}

func NewMyQueue[T any]() *MyQueue[T] {
	return &MyQueue[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (q *MyQueue[T]) Enqueue(item T) {
	node := &MyNode[T]{value: item}
	if q.length == 0 {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.length += 1
}

func (q *MyQueue[T]) Deque() (T, error) {
	if q.head == nil {
		var empty T
		return empty, errors.New("queue empty")
	}
	q.length -= 1
	head := q.head
	q.head = q.head.next
	head.next = nil
	if q.length == 0 {
		q.tail = nil
	}
	return head.value, nil
}

func (q *MyQueue[T]) Peek() (T, error) {
	if q.head == nil {
		var empty T
		return empty, errors.New("queue empty")
	}
	return q.head.value, nil
}

func (q *MyQueue[T]) Len() int {
	return q.length
}
