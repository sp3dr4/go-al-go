package lru

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type LRU[K comparable, V any] struct {
	length        int
	capacity      int
	head          *Node[V]
	tail          *Node[V]
	lookup        *map[K]*Node[V]
	reverseLookup *map[*Node[V]]K
}

func NewLRU[K comparable, V any](capacity int) *LRU[K, V] {
	return &LRU[K, V]{
		length:        0,
		capacity:      capacity,
		head:          nil,
		tail:          nil,
		lookup:        &map[K]*Node[V]{},
		reverseLookup: &map[*Node[V]]K{},
	}
}

func (l *LRU[K, V]) Update(key K, value V) {
	node, ok := (*l.lookup)[key]
	if !ok {
		node = &Node[V]{value: value}
		l.length += 1
		l.prepend(node)
		l.trimCache()

		(*l.lookup)[key] = node
		(*l.reverseLookup)[node] = key
	} else {
		l.detach(node)
		l.prepend(node)
		node.value = value
	}
}

func (l *LRU[K, V]) Get(key K) (V, error) {
	node, ok := (*l.lookup)[key]
	if !ok {
		var empty V
		return empty, errors.New("no item found")
	}

	l.detach(node)
	l.prepend(node)

	return node.value, nil
}

func (l *LRU[K, V]) detach(node *Node[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if l.head == node {
		l.head = l.head.next
	}
	if l.tail == node {
		l.tail = l.tail.prev
	}

	node.prev = nil
	node.next = nil
}

func (l *LRU[K, V]) prepend(node *Node[V]) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}
	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *LRU[K, V]) trimCache() {
	if l.length <= l.capacity {
		return
	}

	tail := l.tail
	l.detach(l.tail)

	key, ok := (*l.reverseLookup)[tail]
	if !ok {
		panic(fmt.Sprintf("key %v should have been found", key))
	}
	delete(*l.lookup, key)
	delete(*l.reverseLookup, tail)
	l.length -= 1
}
