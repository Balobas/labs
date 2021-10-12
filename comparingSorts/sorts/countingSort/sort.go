package countingSort

import "MagistraturaLabsASD/labs/comparingSorts/compareFuncs"

type CountingSort struct {}


// для множества натуральных чисел
func(cs CountingSort) Sort(array []int, compareFunc compareFuncs.CompareFunc) []int {
	max := 0
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}

	count := make([]int, max + 1)

	for i := 0; i < len(array); i++ {
		count[array[i]] ++
	}

	result := make([]int, len(array))
	i := 0

	if compareFunc(1, 2) {
		for j := 0; j < len(count); j++ {
			if count[j] != 0 {
				for k := 0; k < count[j]; k++ {
					result[i] = j
					i++
				}
			}
		}
	} else {
		for j := len(count) - 1; j >= 0; j-- {
			if count[j] != 0 {
				for k := 0; k < count[j]; k++ {
					result[i] = j
					i++
				}
			}
		}
	}

	return result
}
