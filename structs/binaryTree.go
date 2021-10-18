package structs

type BinaryTree struct {
	value int
	left *BinaryTree
	right *BinaryTree
}

func NewBinaryTree(value int) *BinaryTree {
	return &BinaryTree{
		value: value,
		left:  nil,
		right: nil,
	}
}
func(t *BinaryTree) GetValue() int {
	return t.value
}

func(t *BinaryTree) Left() *BinaryTree {
	return t.left
}

func(t *BinaryTree) Right() *BinaryTree {
	return t.right
}

func(t *BinaryTree) Push(elem *BinaryTree) bool {
	if elem.value < t.value {
		if t.left == nil {
			t.left = elem
			return true
		}
		return t.left.Push(elem)
	}
	
	if elem.value > t.value {
		if t.right == nil {
			t.right = elem
			return true
		}
		return t.right.Push(elem)
	}
	return false
}

func(t *BinaryTree) GetSortedValues(array []int) []int {
	if t.left == nil && t.right == nil {
		return append(array, t.value)
	}

	var l []int
	if t.left != nil {
		l = append(array, t.left.GetSortedValues(array)...)
	}

	l = append(l, t.value)

	if t.right != nil {
		l = append(l, t.right.GetSortedValues(array)...)
	}
	
	return l
}
