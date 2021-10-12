package labs

import (
	"MagistraturaLabsASD/labs/comparingSorts/compareFuncs"
	"MagistraturaLabsASD/labs/comparingSorts/sorts"
	"MagistraturaLabsASD/labs/lab9"
	"fmt"
)

type Labs struct {}

func(l Labs) ShowAllLabs() {
	l.ShowLab1()
	l.ShowLab2()
	l.ShowLab3()
	l.ShowLab4()
	l.ShowLab5()
	l.ShowLab6()
	l.ShowLab7()
	l.ShowLab8()
	l.ShowLab9()
}

func(l Labs) ShowLab1() {
	fmt.Println("Lab 1: bubble sort")
	sort, _ := sorts.NewSort(sorts.BubbleSort)
	compareFunc := compareFuncs.ASC
	fmt.Println("по возрастанию:")
	sorts.TestSort(sort, compareFunc)

	compareFunc = compareFuncs.DESC
	fmt.Println("по убыванию")
	sorts.TestSort(sort, compareFunc)
}

func(l Labs) ShowLab2() {
	fmt.Println("Lab 2: selection sort")
	sort, _ := sorts.NewSort(sorts.SelectionSort)
	compareFunc := compareFuncs.ASC
	fmt.Println("по возрастанию:")
	sorts.TestSort(sort, compareFunc)

	compareFunc = compareFuncs.DESC
	fmt.Println("по убыванию")
	sorts.TestSort(sort, compareFunc)
}

func(l Labs) ShowLab3() {
	fmt.Println("Lab 3: insertion sort")
	sort, _ := sorts.NewSort(sorts.InsertionSort)
	compareFunc := compareFuncs.ASC
	fmt.Println("по возрастанию:")
	sorts.TestSort(sort, compareFunc)

	compareFunc = compareFuncs.DESC
	fmt.Println("по убыванию")
	sorts.TestSort(sort, compareFunc)
}

func(l Labs) ShowLab4() {
	fmt.Println("Lab 4: merge sort recursive")
	sort, _ := sorts.NewSort(sorts.MergeSort)
	compareFunc := compareFuncs.ASC
	fmt.Println("по возрастанию:")
	sorts.TestSort(sort, compareFunc)

	compareFunc = compareFuncs.DESC
	fmt.Println("по убыванию")
	sorts.TestSort(sort, compareFunc)


	fmt.Println()
	fmt.Println("merge sort stack")
	sort, _ = sorts.NewSort(sorts.MergeSortStack)
	compareFunc = compareFuncs.ASC
	fmt.Println("по возрастанию:")
	sorts.TestSort(sort, compareFunc)

	compareFunc = compareFuncs.DESC
	fmt.Println("по убыванию")
	sorts.TestSort(sort, compareFunc)
}

func(l Labs) ShowLab5() {
	fmt.Println("Lab 5: quick sort median recursive")
	sort, _ := sorts.NewSort(sorts.QuickSortMedian)
	compareFunc := compareFuncs.ASC
	fmt.Println("по возрастанию:")
	sorts.TestSort(sort, compareFunc)

	compareFunc = compareFuncs.DESC
	fmt.Println("по убыванию")
	sorts.TestSort(sort, compareFunc)


	fmt.Println()
	fmt.Println("quick sort median stack")
	sort, _ = sorts.NewSort(sorts.QuickSortMedianStack)
	compareFunc = compareFuncs.ASC
	fmt.Println("по возрастанию:")
	sorts.TestSort(sort, compareFunc)

	compareFunc = compareFuncs.DESC
	fmt.Println("по убыванию")
	sorts.TestSort(sort, compareFunc)
}

func(l Labs) ShowLab6() {
	fmt.Println("Lab 6: quick sort Hoare recursive")
	sort, _ := sorts.NewSort(sorts.QuickSortHoare)
	compareFunc := compareFuncs.ASC
	fmt.Println("по возрастанию:")
	sorts.TestSort(sort, compareFunc)

	compareFunc = compareFuncs.DESC
	fmt.Println("по убыванию")
	sorts.TestSort(sort, compareFunc)


	fmt.Println()
	fmt.Println("quick sort Hoare stack")
	sort, _ = sorts.NewSort(sorts.QuickSortHoareStack)
	compareFunc = compareFuncs.ASC
	fmt.Println("по возрастанию:")
	sorts.TestSort(sort, compareFunc)

	compareFunc = compareFuncs.DESC
	fmt.Println("по убыванию")
	sorts.TestSort(sort, compareFunc)
}

func(l Labs) ShowLab7() {
	fmt.Println("Lab 7: Shell sort")
	sort, _ := sorts.NewSort(sorts.ShellSort)
	compareFunc := compareFuncs.ASC
	fmt.Println("по возрастанию:")
	sorts.TestSort(sort, compareFunc)

	compareFunc = compareFuncs.DESC
	fmt.Println("по убыванию")
	sorts.TestSort(sort, compareFunc)
}

func(l Labs) ShowLab8() {
	fmt.Println("Lab 8: counting sort")
	sort, _ := sorts.NewSort(sorts.CountingSort)
	compareFunc := compareFuncs.ASC
	fmt.Println("по возрастанию:")
	sorts.TestSort(sort, compareFunc)

	compareFunc = compareFuncs.DESC
	fmt.Println("по убыванию")
	sorts.TestSort(sort, compareFunc)
}

func(l Labs) ShowLab9() {
	lab9.FirstN(25)
}
