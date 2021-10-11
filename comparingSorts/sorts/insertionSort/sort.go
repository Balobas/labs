package insertionSort

import "MagistraturaLabsASD/labs/comparingSorts/compareFuncs"

type InsertionSort struct {}

func(i InsertionSort) Sort(array []int, compareFunc compareFuncs.CompareFunc) []int {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && !compareFunc(array[j-1], array[j]); j-- {
			tmp := array[j]
			array[j] = array[j-1]
			array[j-1] = tmp
		}
	}
	return array
}
