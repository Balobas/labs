package mergeSortStack

import (
	"MagistraturaLabsASD/labs/comparingSorts/compareFuncs"
	"MagistraturaLabsASD/labs/structs"
)

type MergeSortStack struct {}


func(mss MergeSortStack) Sort(array []int, compareFunc compareFuncs.CompareFunc) []int {
	stack := structs.NewStack()
	finalStack := structs.NewStack()

	stack.Push(array)
	var current []int
	for stack.Size() > 0 {
		current = stack.Pop()
		a1 := current[:len(current)/2]
		a2 := current[len(current)/2:]

		switch len(a1) {
		case 1:
			finalStack.Push(a1)
		case 2:
			if compareFunc(a1[0], a1[1]) {
				finalStack.Push(a1)
			} else {
				finalStack.Push([]int{a1[1], a1[0]})
			}
		default:
			stack.Push(a1)
		}

		switch len(a2) {
		case 1:
			finalStack.Push(a2)
		case 2:
			if compareFunc(a2[0], a2[1]) {
				finalStack.Push(a2)
			} else {
				finalStack.Push([]int{a2[1], a2[0]})
			}
		default:
			stack.Push(a2)
		}
	}

	var a1, a2 []int

	for finalStack.Size() > 1 {

		a1 = finalStack.Pop()
		a2 = finalStack.Pop()

		arr := make([]int, len(a1) + len(a2))
		i, j := 0, 0
		for i < len(a1) && j < len(a2) {
			if compareFunc(a1[i], a2[j]) {
				arr[i+j] = a1[i]
				i++
			} else {
				arr[i+j] = a2[j]
				j++
			}
		}

		for ; i < len(a1); i++ {
			arr[j+i] = a1[i]
		}

		for ; j < len(a2); j++ {
			arr[i+j] = a2[j]
		}

		finalStack.Push(arr)
	}

	return finalStack.Pop()
}
