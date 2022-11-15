package array

import (
	"math/rand"
	"time"
)

// Generate an array composed of random values from 0 to 1 of lenght size.
func RandArrFloat64(size int) []float64 {
	rand.Seed(time.Now().UnixMicro())

	values := []float64{}

	for i := 0; i < size; i++ {
		num := rand.Float64()
		values = append(values, num)
	}
	return values
}

// Linear spacing on interval [start, stop] with num evenly spaced points.
func Linspace(start, stop float64, num int) []float64 {
	step := (stop - start) / float64(num-1)
	values := make([]float64, num)

	for i := 0; i < num; i++ {
		values[i] = start + float64(i)*step
	}
	return values
}

// Make an array from start to stop with elements incremented by step.
func Arange(start, stop, step float64) []float64 {
	length := int((stop - start) / step)
	values := make([]float64, length)

	for i := 0; i < length; i++ {
		values[i] = start + float64(i)*step
	}
	return values
}
