package compareFuncs

type CompareFunc func(a, b int) bool

var (

	// возвращает true, если два элемента в правильной последовательности, false в противном случае

	ASC = func(a, b int) bool {
		return a < b
	}

	DESC = func(a, b int) bool {
		return a > b
	}
)
