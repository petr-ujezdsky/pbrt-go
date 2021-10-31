package mymath_test

import (
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVector3NewVector3(t *testing.T) {
	v := mymath.NewVector3(1, 2, 3)

	assert.Equal(t, 1.0, v.X)
	assert.Equal(t, 2.0, v.Y)
	assert.Equal(t, 3.0, v.Z)
}

func TestVector3Vector3LengthSq(t *testing.T) {
	v := mymath.NewVector3(3, 3, 3)

	assert.Equal(t, 27.0, v.LengthSq())
}

func TestVector3Vector3Length(t *testing.T) {
	v := mymath.NewVector3(1, 2, 3)

	assert.Equal(t, 3.7416573867739413, v.Length())
}
func TestVector3Vector3Add(t *testing.T) {
	v := mymath.NewVector3(1, 2, 3)
	w := mymath.NewVector3(5, 6, 7)

	res := v.Add(w)

	assert.Equal(t, 6.0, res.X)
	assert.Equal(t, 8.0, res.Y)
	assert.Equal(t, 10.0, res.Z)
}
func TestVector3Vector3Subtract(t *testing.T) {
	v := mymath.NewVector3(1, 2, 3)
	w := mymath.NewVector3(5, 6, 7)

	res := v.Subtract(w)

	assert.Equal(t, -4.0, res.X)
	assert.Equal(t, -4.0, res.Y)
	assert.Equal(t, -4.0, res.Z)
}

func TestVector3Vector3Multiply(t *testing.T) {
	v := mymath.NewVector3(1, 2, 3)

	res := v.Multiply(5.0)

	assert.Equal(t, 5.0, res.X)
	assert.Equal(t, 10.0, res.Y)
	assert.Equal(t, 15.0, res.Z)
}

func TestVector3Vector3ivide(t *testing.T) {
	v := mymath.NewVector3(2, 4, -8)

	res := v.Divide(2.0)

	assert.Equal(t, 1.0, res.X)
	assert.Equal(t, 2.0, res.Y)
	assert.Equal(t, -4.0, res.Z)
}
func TestVector3Vector3Negate(t *testing.T) {
	v := mymath.NewVector3(2, 4, -8)

	res := v.Negate()

	assert.Equal(t, -2.0, res.X)
	assert.Equal(t, -4.0, res.Y)
	assert.Equal(t, 8.0, res.Z)
}

func TestVector3Vector3Abs(t *testing.T) {
	v := mymath.NewVector3(2, 4, -8)

	res := v.Abs()

	assert.Equal(t, 2.0, res.X)
	assert.Equal(t, 4.0, res.Y)
	assert.Equal(t, 8.0, res.Z)
}

func TestVector3Vector3Normalize(t *testing.T) {
	v := mymath.NewVector3(1, 2, 3)

	res := v.Normalize()

	assert.Equal(t, 0.2672612419124244, res.X)
	assert.Equal(t, 0.5345224838248488, res.Y)
	assert.Equal(t, 0.8017837257372732, res.Z)
}

func TestVector3ot(t *testing.T) {
	v := mymath.NewVector3(-1, -2, -3)
	w := mymath.NewVector3(4, 5, 6)

	res := v.Dot(w)

	assert.Equal(t, -32.0, res)
}

func TestVector3Cross(t *testing.T) {
	v := mymath.NewVector3(1, 2, 3)
	w := mymath.NewVector3(4, 5, 6)

	res := v.Cross(w)

	assert.Equal(t, -3.0, res.X)
	assert.Equal(t, 6.0, res.Y)
	assert.Equal(t, -3.0, res.Z)
}

func TestVector3Get(t *testing.T) {
	v := mymath.NewVector3(1, 2, 3)

	assert.Equal(t, 1.0, v.Get(0))
	assert.Equal(t, 2.0, v.Get(1))
	assert.Equal(t, 3.0, v.Get(2))
}

func TestVector3GetMinComponent(t *testing.T) {
	assert.Equal(t, 1.0, mymath.NewVector3(1, 2, 3).GetMinComponent())
	assert.Equal(t, 2.0, mymath.NewVector3(4, 2, 3).GetMinComponent())
	assert.Equal(t, 3.0, mymath.NewVector3(5, 5, 3).GetMinComponent())
}

func TestVector3GetMaxComponent(t *testing.T) {
	assert.Equal(t, 3.0, mymath.NewVector3(1, 2, 3).GetMaxComponent())
	assert.Equal(t, 4.0, mymath.NewVector3(4, 2, 3).GetMaxComponent())
	assert.Equal(t, 5.0, mymath.NewVector3(5, 5, 3).GetMaxComponent())
}

func TestVector3GetMaxDimension(t *testing.T) {
	assert.Equal(t, 2, mymath.NewVector3(1, 2, 3).GetMaxDimension())
	assert.Equal(t, 2, mymath.NewVector3(2, 1, 3).GetMaxDimension())
	assert.Equal(t, 0, mymath.NewVector3(4, 2, 3).GetMaxDimension())
	assert.Equal(t, 1, mymath.NewVector3(4, 5, 3).GetMaxDimension())
	assert.Equal(t, 1, mymath.NewVector3(5, 5, 3).GetMaxDimension())
}

func TestVector3Min(t *testing.T) {
	v := mymath.NewVector3(1, 8, 3)
	w := mymath.NewVector3(4, 5, 6)

	res := v.Min(w)

	assert.Equal(t, 1.0, res.X)
	assert.Equal(t, 5.0, res.Y)
	assert.Equal(t, 3.0, res.Z)
}

func TestVector3Max(t *testing.T) {
	v := mymath.NewVector3(1, 8, 3)
	w := mymath.NewVector3(4, 5, 6)

	res := v.Max(w)

	assert.Equal(t, 4.0, res.X)
	assert.Equal(t, 8.0, res.Y)
	assert.Equal(t, 6.0, res.Z)
}
func TestVector3Permute(t *testing.T) {
	v := mymath.NewVector3(1, 8, 3)

	res := v.Permute(2, 0, 1)

	assert.Equal(t, 3.0, res.X)
	assert.Equal(t, 1.0, res.Y)
	assert.Equal(t, 8.0, res.Z)
}
