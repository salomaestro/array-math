package array

func CompositeTrapezoid(f func([]float64) []float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)

	x := Linspace(a, b, n+1)[1:n]

	arrA := []float64{a}
	arrB := []float64{b}

	evalA := f(arrA)
	evalB := f(arrB)

	sumFuncX := Sum(f(x))

	res := h / 2 * (evalA[0] + evalB[0] + 2*sumFuncX)

	return res
}

// Sum of float64 array
func Sum(arr []float64) float64 {
	var res float64 = 0
	for _, elem := range arr {
		res += elem
	}
	return res
}

// Mean of array
func Mean(arr []float64) float64 {
	return Sum(arr) / float64(len(arr))
}

// Apply a function to each element of an array
func ApplyFuncToArr(f func(float64) float64, arr []float64) []float64 {
	res := make([]float64, len(arr))
	for i, elem := range arr {
		res[i] = f(elem)
	}
	return res
}
