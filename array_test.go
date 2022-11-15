package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	start  float64 = 2
	stop   float64 = 3
	length int     = 5
	step   float64 = 0.1
)

func TestLinspace(t *testing.T) {
	got := Linspace(start, stop, length)

	assert.Equal(t, start, got[0])
	assert.Equal(t, stop, got[len(got)-1])
	assert.Equal(t, length, len(got))
}

func TestArange(t *testing.T) {
	got := Arange(start, stop, step)

	assert.Equal(t, start, got[0])
	assert.Equal(t, stop-step, got[len(got)-1])
	assert.Equal(t, int((stop-start)/step), len(got))
}

func TestRandArrFloat64(t *testing.T) {
	got := RandArrFloat64(length)

	assert.Equal(t, length, len(got))
}
