package quickSortHoareStack

import (
	"MagistraturaLabsASD/labs/comparingSorts/compareFuncs"
	"MagistraturaLabsASD/labs/comparingSorts/sorts/quickSortHoare"
	"MagistraturaLabsASD/labs/structs"
)

type QuickSortHoareStack struct {}

func(qshs QuickSortHoareStack) Sort(array []int, compareFunc compareFuncs.CompareFunc) []int {
	stack := structs.NewStack()

	arr := array

	stack.Push(arr)

	for stack.Size() > 0 {
		a := stack.Pop()

		if len(a) == 1 {
			continue
		}

		p := quickSortHoare.Partition(a, compareFunc)

		stack.Push(a[:p])
		stack.Push(a[p:])
	}

	return arr
}
