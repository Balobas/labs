package selectionSort

import (
	"MagistraturaLabsASD/labs/comparingSorts/compareFuncs"
)

type SelectionSort struct {}

func(s SelectionSort) Sort(array []int, compareFunc compareFuncs.CompareFunc) []int {
	resultArray := array

	for i := 0; i < len(resultArray) - 1; i++ {
		index := i
		for j := i; j < len(resultArray); j ++ {
			if !compareFunc(resultArray[index], resultArray[j]) {
				index = j
			}
		}

		if index != i {
			tmp := resultArray[i]
			resultArray[i] = resultArray[index]
			resultArray[index] = tmp
		}
	}

	return resultArray
}
