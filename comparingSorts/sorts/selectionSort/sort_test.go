package selectionSort

import (
	"MagistraturaLabsASD/labs/comparingSorts/compareFuncs"
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	testArray := []int{5, 10, 8, 1, 4, 2, 7, 3, 9, 6}

	compareFunc := compareFuncs.DESC

	sortedArray := SelectionSort{}.Sort(testArray, compareFunc)

	for i := 0; i < len(testArray) - 1; i++ {
		if !compareFunc(sortedArray[i], sortedArray[i+1]) {
			fmt.Println(sortedArray[i], sortedArray[i+1])
			t.Error("unsorted")
			fmt.Println(sortedArray)
			t.FailNow()
		}
	}

	t.Log("sorted")
	fmt.Println(sortedArray)
}
