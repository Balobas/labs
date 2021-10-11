package sorter

import "MagistraturaLabsASD/labs/comparingSorts/compareFuncs"

type Sorter interface {
	Sort(array []int, compareFunc compareFuncs.CompareFunc) []int
}
