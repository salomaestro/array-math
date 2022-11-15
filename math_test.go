package array

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	arr = Arange(0, 3, 1)
)

func square(x float64) float64 {
	return x * x
}

func squareArr(arr []float64) []float64 {
	res := make([]float64, len(arr))
	for i, elem := range arr {
		res[i] = elem * elem
	}
	return res
}

func TestSum(t *testing.T) {
	got := Sum(arr)

	assert.Equal(t, float64(3), got)
}

func TestMean(t *testing.T) {
	got := Mean(arr)

	assert.Equal(t, float64(1), got)
}

func TestApplyFuncToArr(t *testing.T) {
	got := ApplyFuncToArr(
		square,
		arr,
	)

	assert.Equal(t, []float64{0, 1, 4}, got)
}

func TestCompositeTrapezoid(t *testing.T) {
	got := CompositeTrapezoid(squareArr, float64(0), float64(2), 1000)

	actual := float64(8) / float64(3)

	assert.True(t, math.Abs(got-actual) < 1e-4)
}

func TestConvertToArrFunc(t *testing.T) {
	got1 := ConvertToArrFunc(square)
	got2 := got1(arr)
	actual := squareArr(arr)

	assert.IsType(t, squareArr, got1)
	assert.Equal(t, actual, got2)
}
