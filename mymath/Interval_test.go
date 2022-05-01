package mymath_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"pbrt-go/mymath"
	"testing"
)

func TestInterval_NewIntervalSingle(t *testing.T) {
	i := mymath.NewIntervalSingle(2)

	assert.Equal(t, 2.0, i.Low)
	assert.Equal(t, 2.0, i.High)
}

func TestInterval_NewInterval(t *testing.T) {
	i := mymath.NewInterval(1, 2)

	assert.Equal(t, 1.0, i.Low)
	assert.Equal(t, 2.0, i.High)

	i = mymath.NewInterval(2, 1)

	assert.Equal(t, 1.0, i.Low)
	assert.Equal(t, 2.0, i.High)
}

func TestInterval_Add(t *testing.T) {
	i1 := mymath.NewInterval(1, 2)
	i2 := mymath.NewInterval(10, 20)

	i := i1.Add(i2)
	assert.Equal(t, 11.0, i.Low)
	assert.Equal(t, 22.0, i.High)
}

func TestInterval_Subtract(t *testing.T) {
	i1 := mymath.NewInterval(1, 2)
	i2 := mymath.NewInterval(10, 20)

	i := i1.Subtract(i2)
	assert.Equal(t, -19.0, i.Low)
	assert.Equal(t, -8.0, i.High)
}

func TestInterval_Multiply(t *testing.T) {
	i1 := mymath.NewInterval(-10, 10)
	i2 := mymath.NewInterval(-2, 20)

	i := i1.Multiply(i2)
	assert.Equal(t, -19.0, i.Low)
	assert.Equal(t, -8.0, i.High)
}

func TestInterval_Sin(t *testing.T) {
	i := mymath.Sin(mymath.NewInterval(0, math.Pi))
	assert.Equal(t, 0.0, i.Low)
	assert.Equal(t, 1.0, i.High)

	i = mymath.Sin(mymath.NewInterval(math.Pi/2, 3.0/2.0*math.Pi))
	assert.Equal(t, -1.0, i.Low)
	assert.Equal(t, 1.0, i.High)
}

func TestInterval_Cos(t *testing.T) {
	i := mymath.Cos(mymath.NewInterval(0, math.Pi))
	assert.Equal(t, -1.0, i.Low)
	assert.Equal(t, 1.0, i.High)

	i = mymath.Cos(mymath.NewInterval(math.Pi/2, 3.0/2.0*math.Pi))
	assert.Equal(t, -1.0, i.Low)
	assert.Equal(t, 0.0, i.High)
}
