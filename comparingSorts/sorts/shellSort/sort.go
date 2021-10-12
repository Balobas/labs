package shellSort

import "MagistraturaLabsASD/labs/comparingSorts/compareFuncs"

type ShellSort struct {}

func(ss ShellSort) Sort(array []int, compareFunc compareFuncs.CompareFunc) []int {
	for d := len(array) / 2; d >= 1; d /= 2 {
		for i := 0; i < len(array) - d; i++ {
			for j := i + d; j < len(array); j += d {
				if !compareFunc(array[i], array[j]) {
					swap(&array[i], &array[j])
				}
			}
		}
	}
	return array
}

func swap(a, b *int) {
	t := *b
	*b = *a
	*a = t
}
