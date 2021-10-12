package quickSortMedianStack

import (
	"MagistraturaLabsASD/labs/comparingSorts/compareFuncs"
	"MagistraturaLabsASD/labs/comparingSorts/sorts/mergeSortStack"
	"MagistraturaLabsASD/labs/comparingSorts/sorts/quickSortMedian"
)

type QuickSortMedianStack struct {}

func(qsms QuickSortMedianStack) Sort(array []int, compareFunc compareFuncs.CompareFunc) []int {
	stack := mergeSortStack.NewStack()
	arr := array

	stack.Push(arr)

	for stack.Size() > 0 {
		a := stack.Pop()

		if len(a) == 1 {
			continue
		}

		p := quickSortMedian.Partition(a, compareFunc)

		stack.Push(a[:p])
		stack.Push(a[p:])
	}

	return arr
}
