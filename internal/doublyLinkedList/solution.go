package doublylinkedlist

import "errors"

type LinkedList[T any] interface {
	Len() int
	Append(T)
	Prepend(T)
	Get(int) (T, error)
	InsertAt(T, int) error
	Remove(T) (T, error)
	RemoveAt(int) (T, error)
}

type MyNode[T any] struct {
	value T
	prev  *MyNode[T]
	next  *MyNode[T]
}

type DoublyLinkedList[T comparable] struct {
	length int
	head   *MyNode[T]
	tail   *MyNode[T]
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (l *DoublyLinkedList[T]) Len() int {
	return l.length
}

func (l *DoublyLinkedList[T]) Append(item T) {
	node := &MyNode[T]{value: item}
	if l.length == 0 {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}
	l.length += 1
}

func (l *DoublyLinkedList[T]) Prepend(item T) {
	node := &MyNode[T]{value: item}
	if l.length == 0 {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.length += 1
}

func (l *DoublyLinkedList[T]) Get(idx int) (T, error) {
	node, err := l.getAt(idx)
	if err != nil {
		var empty T
		return empty, err
	}
	return node.value, nil
}

func (l *DoublyLinkedList[T]) InsertAt(item T, idx int) error {
	if idx > l.length {
		return errors.New("out of bounds index")
	}
	if idx == l.length {
		l.Append(item)
		return nil
	}
	if idx == 0 {
		l.Prepend(item)
		return nil
	}

	curr, err := l.getAt(idx)
	if err != nil {
		return err
	}

	node := &MyNode[T]{
		value: item,
		next:  curr,
		prev:  curr.prev,
	}
	curr.prev.next = node
	curr.prev = node
	l.length += 1
	return nil
}

func (l *DoublyLinkedList[T]) Remove(item T) (T, error) {
	var found *MyNode[T]
	curr := l.head
	for i := 0; i < l.length; i++ {
		if curr.value == item {
			found = curr
			break
		}
		curr = curr.next
	}
	if found == nil {
		var empty T
		return empty, errors.New("item not found")
	}
	return l.removeNode(found)
}

func (l *DoublyLinkedList[T]) RemoveAt(idx int) (T, error) {
	curr, err := l.getAt(idx)
	if err != nil {
		var empty T
		return empty, err
	}
	return l.removeNode(curr)
}

func (l *DoublyLinkedList[T]) removeNode(node *MyNode[T]) (T, error) {
	if node == l.head && node == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		if node.prev != nil {
			node.prev.next = node.next
		}
		if node.next != nil {
			node.next.prev = node.prev
		}
		if node == l.head {
			l.head = node.next
		}
		if node == l.tail {
			l.tail = node.prev
		}

		node.prev = nil
		node.next = nil
	}

	l.length -= 1
	return node.value, nil
}

func (l *DoublyLinkedList[T]) getAt(idx int) (*MyNode[T], error) {
	if idx < 0 || idx >= l.length {
		return nil, errors.New("invalid index")
	}
	curr := l.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}
	return curr, nil
}
