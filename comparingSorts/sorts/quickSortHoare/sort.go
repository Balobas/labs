package quickSortHoare

import (
	"MagistraturaLabsASD/labs/comparingSorts/compareFuncs"
)

type QuickSortHoare struct {}

func(qsh QuickSortHoare) Sort(array []int, compareFunc compareFuncs.CompareFunc) []int {
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
	if len(arr) == 2 {
		if compareFunc(arr[0], arr[1]) {
			return 1
		} else {
			swap(&arr[0], &arr[1])
			return 1
		}
	}

	var l, r []int
	p := len(arr) / 2
	pE := arr[p]

	for i := 0; i < len(arr); i++ {
		if i == p {
			continue
		}
		if compareFunc(arr[i], arr[p]) {
			l = append(l, arr[i])
		} else {
			r = append(r, arr[i])
		}
	}

	i := 0
	for j := 0; j < len(l); j++ {
		arr[i] = l[j]
		i++
	}

	arr[i] = pE
	p = i
	i++

	for k := 0; k < len(r); k++ {
		arr[i] = r[k]
		i++
	}

	return p
}

func swap(a, b *int) {
	t := *b
	*b = *a
	*a = t
}
