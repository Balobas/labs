package lab9

import (
	"MagistraturaLabsASD/labs/structs"
	"fmt"
)

/*
Напечатать в порядке возрастания первые n натуральных чисел,
в разложение которых на простые множители входят только числа 2 и 3
 */

/*
числа добаляются в стек, чтобы избежать повторов и отсортировать, значения так же помещаются в бинарное дерево
 */
func FirstN(n int) {
	fmt.Println("первые n натуральных чисел, \nв разложение которых на простые множители входят только числа 2 и 3\nв порядке возрастания ")
	stack := NewStack()
	tree := structs.NewBinaryTree(2)
	tree.Push(structs.NewBinaryTree(3))

	stack.Push(2)
	stack.Push(3)

	count := 2
	last := 0

	for count <= n {
		current := stack.Pop()
		if current == last {
			continue
		}

		stack.Push(current * 2)
		stack.Push(current * 3)

		if ok := tree.Push(structs.NewBinaryTree(current*2)); ok {
			count++
		}
		if ok := tree.Push(structs.NewBinaryTree(current*3)); ok {
			count++
		}

		last = current
	}

	fmt.Println("n =", n)
	values := tree.GetSortedValues([]int{})
	for _, val := range values {
		fmt.Println(val)
	}
}

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{[]int{}}
}

func(s *Stack) Push(elem int) {
	s.data = append(s.data, elem)
}

func(s *Stack) Pop() int {
	res := s.data[0]
	s.data = s.data[1:]
	return res
}

func(s *Stack) Size() int {
	return len(s.data)
}