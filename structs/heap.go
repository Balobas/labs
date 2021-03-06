package structs

import "MagistraturaLabsASD/labs/comparingSorts/compareFuncs"

type Heap struct {
	data []int
	size int
	n int
	compareFunc compareFuncs.CompareFunc
}

func NewHeap(n int, compareFunc compareFuncs.CompareFunc) *Heap {
	return &Heap{make([]int, n*2), 0, n, compareFunc}
}

func(h *Heap) Push(elem int) {
	if h.size == 0 {
		h.data[0] = elem
		h.size++
		return
	}

	h.data[h.size] = elem
	h.size++
}

func(h *Heap) Pop() int {
	if h.size == 1 {
		res := h.data[0]
		h.data = make([]int, h.n * 2)
		h.size--
		return res
	}
	for i := h.size / 2 + 1; i >= 0; i-- {
		h.heapify(i)
	}
	res := h.data[0]
	h.data = h.data[1:]
	h.size--
	return res
}

func(h *Heap) heapify(i int) {
	maxOrMin := i

	l, r := 2*i + 1, 2*i + 2

	if l < h.size && h.compareFunc(h.data[maxOrMin], h.data[l]) {
		maxOrMin = l
	}

	if r < h.size && h.compareFunc(h.data[maxOrMin], h.data[r]) {
		maxOrMin = r
	}

	if maxOrMin == i {
		return
	}

	swap(&h.data[i], &h.data[maxOrMin])
	h.heapify(maxOrMin)
}

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
