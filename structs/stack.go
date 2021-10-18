package structs

type Stack struct {
	data [][]int
}

func NewStack() *Stack {
	return &Stack{[][]int{}}
}

func(s *Stack) Push(elem []int) {
	s.data = append(s.data, elem)
}

func(s *Stack) Pop() []int {
	res := s.data[0]
	s.data = s.data[1:]
	return res
}

func(s *Stack) Size() int {
	return len(s.data)
}

