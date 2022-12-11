package aoc

type Heap[T any] struct {
	less func(v1, v2 T) bool
	data []T
}

func NewHeap[T any](less func(v1, v2 T) bool) Heap[T] {
	return Heap[T]{less: less}
}
func (h *Heap[T]) Push(values ...T) {
	for _, v := range values {
		h.data = append(h.data, v)
		h.heapifyUp()
	}
}

func (h *Heap[T]) Pop() T {
	v := h.data[0]
	h.swap(0, len(h.data)-1)
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown()
	return v
}

func (h *Heap[T]) Poll() T {
	return h.data[0]
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) heapifyDown() {
	i := 0
	for {
		li, ri := h.leftChild(i), h.rightChild(i)
		if li < 0 && ri < 0 {
			return
		}
		if ri < 0 || (h.less(h.data[li], h.data[ri]) && h.less(h.data[li], h.data[i])) {
			h.swap(li, i)
			i = li
			continue
		}
		if h.less(h.data[ri], h.data[i]) {
			h.swap(ri, i)
			i = ri
			continue
		}
		return
	}
}

func (h *Heap[T]) heapifyUp() {
	i := len(h.data) - 1
	for i > 0 {
		pi := h.parent(i)
		if h.less(h.data[pi], h.data[i]) {
			return
		}
		h.swap(pi, i)
		i = pi
	}
}

func (h *Heap[T]) swap(i1, i2 int) {
	h.data[i1], h.data[i2] = h.data[i2], h.data[i1]
}

func (h *Heap[T]) parent(index int) int {
	return (index - 1) / 2
}

func (h *Heap[T]) leftChild(index int) int {
	i := 2*index + 1
	if i >= len(h.data) {
		i = -1
	}
	return i
}

func (h *Heap[T]) rightChild(index int) int {
	i := 2*index + 2
	if i >= len(h.data) {
		i = -1
	}
	return i
}
