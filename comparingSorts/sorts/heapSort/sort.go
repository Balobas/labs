package heapSort

import (
	"MagistraturaLabsASD/labs/comparingSorts/compareFuncs"
)

type HeapSort struct {}

func heapify(array []int, compareFunc compareFuncs.CompareFunc) {
	largest := 0

	l, r := 1, 2

	if l < len(array) && compareFunc(array[largest], array[l]) {
		largest = l
	}

	if r < len(array) && compareFunc(array[largest], array[r]) {
		largest = r
	}

	if largest == 0 {
		return
	}

	swap(&array[0], &array[largest])
	heapify(array[largest:], compareFunc)
}

func(hs HeapSort) Sort(array []int, compareFunc compareFuncs.CompareFunc) []int {
	for i := len(array) / 2 ; i >= 0; i-- {
		heapify(array[i:], compareFunc)
	}

	for i := len(array) - 1; i > 0; i-- {
		swap(&array[i], &array[0])
		heapify(array[:i], compareFunc)
	}

	return array
}

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
