package bubbleSort

import (
	"MagistraturaLabsASD/labs/comparingSorts/compareFuncs"
)

type BubbleSort struct {}

func(b BubbleSort) Sort(array []int, compareFunc compareFuncs.CompareFunc) []int {
	resultArray := array

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if !compareFunc(resultArray[i], resultArray[j]) {
				tmp := resultArray[i]
				resultArray[i] = resultArray[j]
				resultArray[j] = tmp
			}
		}
	}

	return resultArray
}
