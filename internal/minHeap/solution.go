package minheap

import "errors"

type MinHeap struct {
	Length int
	data   []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		Length: 0,
		data:   []int{},
	}
}

func (h *MinHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(h.Length)
	h.Length += 1
}

func (h *MinHeap) Delete() (int, error) {
	if h.Length == 0 {
		return -1, errors.New("heap is empty")
	}
	out := h.data[0]
	h.Length -= 1
	if h.Length == 0 {
		h.data = []int{}
		return out, nil
	}
	h.data[0] = h.data[h.Length]
	h.heapifyDown(0)
	return out, nil
}

func (h *MinHeap) heapifyDown(idx int) {
	if idx >= h.Length {
		return
	}
	lIdx := h.leftChild(idx)
	rIdx := h.rightChild(idx)
	if lIdx >= h.Length {
		return
	}

	lV := h.data[lIdx]
	rV := h.data[rIdx]
	v := h.data[idx]

	if lV > rV && v > rV {
		h.data[idx] = rV
		h.data[rIdx] = v
		h.heapifyDown(rIdx)
	} else if rV > lV && v > lV {
		h.data[idx] = lV
		h.data[lIdx] = v
		h.heapifyDown(lIdx)
	}
}

func (h *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}
	p := h.parent(idx)
	parentV := h.data[p]
	v := h.data[idx]

	if parentV > v {
		h.data[idx] = parentV
		h.data[p] = v
		h.heapifyUp(p)
	}
}

func (h *MinHeap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h *MinHeap) leftChild(idx int) int {
	return idx*2 + 1
}

func (h *MinHeap) rightChild(idx int) int {
	return idx*2 + 2
}
