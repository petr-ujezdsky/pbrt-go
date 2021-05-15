package mymath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	v := NewVector3d(1, 2, 3)

	assert.Equal(t, 1.0, v.X)
	assert.Equal(t, 2.0, v.Y)
	assert.Equal(t, 3.0, v.Z)
}

func TestLengthSq(t *testing.T) {
	v := NewVector3d(3, 3, 3)

	assert.Equal(t, 27.0, v.LengthSq())
}

func TestLength(t *testing.T) {
	v := NewVector3d(1, 2, 3)

	assert.Equal(t, 3.7416573867739413, v.Length())
}
func TestAdd(t *testing.T) {
	v := NewVector3d(1, 2, 3)
	w := NewVector3d(5, 6, 7)

	res := v.Add(w)

	assert.Equal(t, 6.0, res.X)
	assert.Equal(t, 8.0, res.Y)
	assert.Equal(t, 10.0, res.Z)
}
func TestSubtract(t *testing.T) {
	v := NewVector3d(1, 2, 3)
	w := NewVector3d(5, 6, 7)

	res := v.Subtract(w)

	assert.Equal(t, -4.0, res.X)
	assert.Equal(t, -4.0, res.Y)
	assert.Equal(t, -4.0, res.Z)
}

func TestMultiply(t *testing.T) {
	v := NewVector3d(1, 2, 3)

	res := v.Multiply(5.0)

	assert.Equal(t, 5.0, res.X)
	assert.Equal(t, 10.0, res.Y)
	assert.Equal(t, 15.0, res.Z)
}

func TestDivide(t *testing.T) {
	v := NewVector3d(2, 4, -8)

	res := v.Divide(2.0)

	assert.Equal(t, 1.0, res.X)
	assert.Equal(t, 2.0, res.Y)
	assert.Equal(t, -4.0, res.Z)
}
func TestNegate(t *testing.T) {
	v := NewVector3d(2, 4, -8)

	res := v.Negate()

	assert.Equal(t, -2.0, res.X)
	assert.Equal(t, -4.0, res.Y)
	assert.Equal(t, 8.0, res.Z)
}

func TestAbs(t *testing.T) {
	v := NewVector3d(2, 4, -8)

	res := v.Abs()

	assert.Equal(t, 2.0, res.X)
	assert.Equal(t, 4.0, res.Y)
	assert.Equal(t, 8.0, res.Z)
}

func TestNormalize(t *testing.T) {
	v := NewVector3d(1, 2, 3)

	res := v.Normalize()

	assert.Equal(t, 0.2672612419124244, res.X)
	assert.Equal(t, 0.5345224838248488, res.Y)
	assert.Equal(t, 0.8017837257372732, res.Z)
}

func TestDot(t *testing.T) {
	v := NewVector3d(-1, -2, -3)
	w := NewVector3d(4, 5, 6)

	res := v.Dot(w)

	assert.Equal(t, -32.0, res)
}

func TestCross(t *testing.T) {
	v := NewVector3d(1, 2, 3)
	w := NewVector3d(4, 5, 6)

	res := v.Cross(w)

	assert.Equal(t, -3.0, res.X)
	assert.Equal(t, 6.0, res.Y)
	assert.Equal(t, -3.0, res.Z)
}

func TestGet(t *testing.T) {
	v := NewVector3d(1, 2, 3)

	assert.Equal(t, 1.0, v.Get(0))
	assert.Equal(t, 2.0, v.Get(1))
	assert.Equal(t, 3.0, v.Get(2))
}

func TestGetMinComponent(t *testing.T) {
	assert.Equal(t, 1.0, NewVector3d(1, 2, 3).GetMinComponent())
	assert.Equal(t, 2.0, NewVector3d(4, 2, 3).GetMinComponent())
	assert.Equal(t, 3.0, NewVector3d(5, 5, 3).GetMinComponent())
}

func TestGetMaxComponent(t *testing.T) {
	assert.Equal(t, 3.0, NewVector3d(1, 2, 3).GetMaxComponent())
	assert.Equal(t, 4.0, NewVector3d(4, 2, 3).GetMaxComponent())
	assert.Equal(t, 5.0, NewVector3d(5, 5, 3).GetMaxComponent())
}

func TestGetMaxDimension(t *testing.T) {
	assert.Equal(t, 2, NewVector3d(1, 2, 3).GetMaxDimension())
	assert.Equal(t, 0, NewVector3d(4, 2, 3).GetMaxDimension())
	assert.Equal(t, 1, NewVector3d(4, 5, 3).GetMaxDimension())
	assert.Equal(t, 1, NewVector3d(5, 5, 3).GetMaxDimension())
}

func TestMin(t *testing.T) {
	v := NewVector3d(1, 8, 3)
	w := NewVector3d(4, 5, 6)

	res := v.Min(w)

	assert.Equal(t, 1.0, res.X)
	assert.Equal(t, 5.0, res.Y)
	assert.Equal(t, 3.0, res.Z)
}

func TestMax(t *testing.T) {
	v := NewVector3d(1, 8, 3)
	w := NewVector3d(4, 5, 6)

	res := v.Max(w)

	assert.Equal(t, 4.0, res.X)
	assert.Equal(t, 8.0, res.Y)
	assert.Equal(t, 6.0, res.Z)
}
func TestPermute(t *testing.T) {
	v := NewVector3d(1, 8, 3)

	res := v.Permute(2, 0, 1)

	assert.Equal(t, 3.0, res.X)
	assert.Equal(t, 1.0, res.Y)
	assert.Equal(t, 8.0, res.Z)
}
