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

// при двух динаковых элементах в массиве может выбрать их значение за pivot тогда зависнет
func partitionBad(arr []int, compareFunc compareFuncs.CompareFunc) int {

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

	SortThree(&arr[0], &arr[p], &arr[len(arr) - 1], compareFunc)

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
