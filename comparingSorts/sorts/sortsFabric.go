package sorts

import (
	"MagistraturaLabsASD/labs/comparingSorts/compareFuncs"
	"MagistraturaLabsASD/labs/comparingSorts/sorter"
	"MagistraturaLabsASD/labs/comparingSorts/sorts/bubbleSort"
	"MagistraturaLabsASD/labs/comparingSorts/sorts/countingSort"
	"MagistraturaLabsASD/labs/comparingSorts/sorts/insertionSort"
	"MagistraturaLabsASD/labs/comparingSorts/sorts/mergeSort"
	"MagistraturaLabsASD/labs/comparingSorts/sorts/mergeSortStack"
	"MagistraturaLabsASD/labs/comparingSorts/sorts/quickSortHoare"
	"MagistraturaLabsASD/labs/comparingSorts/sorts/quickSortHoareStack"
	"MagistraturaLabsASD/labs/comparingSorts/sorts/quickSortMedian"
	"MagistraturaLabsASD/labs/comparingSorts/sorts/quickSortMedianStack"
	"MagistraturaLabsASD/labs/comparingSorts/sorts/selectionSort"
	"MagistraturaLabsASD/labs/comparingSorts/sorts/shellSort"
	"fmt"
	"github.com/pkg/errors"
)

type SortType string

const (
	BubbleSort = SortType("bubble sort")
	SelectionSort = SortType("selection sort")
	InsertionSort = SortType("insertion sort")
	MergeSort = SortType("merge sort")
	MergeSortStack = SortType("merge sort stack")
	QuickSortMedian = SortType("quick sort median")
	QuickSortMedianStack = SortType("quick sort median stack")
	QuickSortHoare = SortType("quick sort Hoare")
	QuickSortHoareStack = SortType("quick sort Hoare stack")
	ShellSort = SortType("Shell sort")
	CountingSort = SortType("counting sort")
)

func NewSort(sortType SortType) (sorter.Sorter, error) {
	switch sortType {
	case BubbleSort:
		return newBubbleSorter(), nil
	case SelectionSort:
		return newSelectionSort(), nil
	case InsertionSort:
		return newInsertionSort(), nil
	case MergeSort:
		return newMergeSort(), nil
	case MergeSortStack:
		return newMergeSortStack(), nil
	case QuickSortMedian:
		return newQuickSortMedian(), nil
	case QuickSortMedianStack:
		return newQuickSortMedianStack(), nil
	case QuickSortHoare:
		return newQuickSortHoare(), nil
	case QuickSortHoareStack:
		return newQuickSortHoareStack(), nil
	case ShellSort:
		return newShellSort(), nil
	case CountingSort:
		return newCountingSort(), nil
	}
	return nil, errors.Errorf("invalid sort type %s", sortType)
}

func newBubbleSorter() sorter.Sorter {
	return bubbleSort.BubbleSort{}
}

func newSelectionSort() sorter.Sorter {
	return selectionSort.SelectionSort{}
}

func newInsertionSort() sorter.Sorter {
	return insertionSort.InsertionSort{}
}

func newMergeSort() sorter.Sorter {
	return mergeSort.MergeSort{}
}

func newMergeSortStack() sorter.Sorter {
	return mergeSortStack.MergeSortStack{}
}

func newQuickSortMedian() sorter.Sorter {
	return quickSortMedian.QuickSortMedian{}
}

func newQuickSortMedianStack() sorter.Sorter {
	return quickSortMedianStack.QuickSortMedianStack{}
}

func newQuickSortHoare() sorter.Sorter {
	return quickSortHoare.QuickSortHoare{}
}

func newQuickSortHoareStack() sorter.Sorter {
	return quickSortHoareStack.QuickSortHoareStack{}
}

func newShellSort() sorter.Sorter {
	return shellSort.ShellSort{}
}

func newCountingSort() sorter.Sorter {
	return countingSort.CountingSort{}
}

func TestSort(sorter sorter.Sorter, compareFunc compareFuncs.CompareFunc) {
	testArray := []int{5, 10, 8, 1, 4, 2, 7, 3, 9, 6}

	sortedArray := sorter.Sort(testArray, compareFunc)

	for i := 0; i < len(testArray) - 1; i++ {
		if !compareFunc(sortedArray[i], sortedArray[i+1]) {
			fmt.Println(sortedArray[i], sortedArray[i+1])
			fmt.Println("unsorted")
			fmt.Println(sortedArray)
			fmt.Println("FAIL")
			return
		}
	}

	fmt.Println(sortedArray)
	fmt.Println("PASS: sorted\n")
}
