package labs

import (
	"MagistraturaLabsASD/labs/comparingSorts/compareFuncs"
	"MagistraturaLabsASD/labs/comparingSorts/sorts"
	"fmt"
)

type Labs struct {}

func(l Labs) ShowAllLabs() {
	l.ShowLab1()
	l.ShowLab2()
	l.ShowLab3()
	l.ShowLab4()
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
	fmt.Println("Lab 4: merge sort")
	sort, _ := sorts.NewSort(sorts.MergeSort)
	compareFunc := compareFuncs.ASC
	fmt.Println("по возрастанию:")
	sorts.TestSort(sort, compareFunc)

	compareFunc = compareFuncs.DESC
	fmt.Println("по убыванию")
	sorts.TestSort(sort, compareFunc)
}
