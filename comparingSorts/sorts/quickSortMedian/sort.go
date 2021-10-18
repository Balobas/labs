package quickSortMedian

import (
	"MagistraturaLabsASD/labs/comparingSorts/compareFuncs"
)

type QuickSortMedian struct {}

func(qsm QuickSortMedian) Sort(array []int, compareFunc compareFuncs.CompareFunc) []int {
	arr := array
	reqSort(arr, compareFunc)
	return arr
}

func reqSort(array []int, compareFunc compareFuncs.CompareFunc) {
	if len(array) > 1 {
		m := Partition(array, compareFunc)

		reqSort(array[:m], compareFunc)
		reqSort(array[m:], compareFunc)
	}
}

func Partition(arr []int, compareFunc compareFuncs.CompareFunc) int {

	pivotI := len(arr) / 2

	SortThree(&arr[0], &arr[pivotI], &arr[len(arr) - 1], compareFunc)

	pivot := arr[pivotI]
	l, r := 0, len(arr) - 1

	for {

		for compareFunc(arr[l], pivot) {
			l++
		}

		for compareFunc(pivot, arr[r]) {
			r--
		}

		if l >= r {
			return r
		}

		swap(&arr[l], &arr[r])
	}
}

func SortThree(a, b, c *int, compareFunc compareFuncs.CompareFunc) {
	if !compareFunc(*a, *c) {
		swap(a, c)
	}

	if !compareFunc(*a, *b) {
		swap(a, b)
	} else if !compareFunc(*b, *c) {
		swap(b, c)
	}
}

func swap(a, b *int) {
	t := *b
	*b = *a
	*a = t
}
