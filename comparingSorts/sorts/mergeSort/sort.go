package mergeSort

import "MagistraturaLabsASD/labs/comparingSorts/compareFuncs"

type MergeSort struct {}

func (m MergeSort) Sort(array []int, compareFunc compareFuncs.CompareFunc) []int {
	count := len(array)

	switch {
	case count > 2:
		lb := m.Sort(array[:count/2], compareFunc)
		rb := m.Sort(array[count/2:], compareFunc)

		array = make([]int, len(array))

		i, j := 0, 0
		for i < len(lb) && j < len(rb) {
			if compareFunc(lb[i], rb[j]) {
				array[i+j] = lb[i]
				i++
			} else {
				array[i+j] = rb[j]
				j++
			}
		}

		for ; i < len(lb); i++ {
			array[j+i] = lb[i]
		}

		for ; j < len(rb); j++ {
			array[i+j] = rb[j]
		}
	case count > 1 && !compareFunc(array[0], array[1]):
		tmp := array[0]
		array[0] = array[1]
		array[1] = tmp
	}

	return array
}
