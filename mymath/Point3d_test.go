package mymath_test

import (
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	p := mymath.NewPoint3d(1, 2, 3)

	assert.Equal(t, 1.0, p.X)
	assert.Equal(t, 2.0, p.Y)
	assert.Equal(t, 3.0, p.Z)
}

func TestAddP(t *testing.T) {
	p1 := mymath.NewPoint3d(1, 2, 3)
	p2 := mymath.NewPoint3d(5, 6, 7)

	res := p1.AddP(p2)

	assert.Equal(t, 6.0, res.X)
	assert.Equal(t, 8.0, res.Y)
	assert.Equal(t, 10.0, res.Z)
}

func TestAddV(t *testing.T) {
	p := mymath.NewPoint3d(1, 2, 3)
	v := mymath.NewVector3d(5, 6, 7)

	res := p.AddV(v)

	assert.Equal(t, 6.0, res.X)
	assert.Equal(t, 8.0, res.Y)
	assert.Equal(t, 10.0, res.Z)
}

func TestSubtractP(t *testing.T) {
	p1 := mymath.NewPoint3d(1, 2, 3)
	p2 := mymath.NewPoint3d(5, 6, 7)

	res := p1.SubtractP(p2)

	assert.Equal(t, -4.0, res.X)
	assert.Equal(t, -4.0, res.Y)
	assert.Equal(t, -4.0, res.Z)
}

func TestSubtractV(t *testing.T) {
	p := mymath.NewPoint3d(1, 2, 3)
	v := mymath.NewVector3d(5, 6, 7)

	res := p.SubtractV(v)

	assert.Equal(t, -4.0, res.X)
	assert.Equal(t, -4.0, res.Y)
	assert.Equal(t, -4.0, res.Z)
}

func TestMultiply(t *testing.T) {
	v := mymath.NewPoint3d(1, 2, 3)

	res := v.Multiply(5.0)

	assert.Equal(t, 5.0, res.X)
	assert.Equal(t, 10.0, res.Y)
	assert.Equal(t, 15.0, res.Z)
}

func TestDistance(t *testing.T) {
	p1 := mymath.NewPoint3d(1, 2, 3)
	p2 := mymath.NewPoint3d(5, 6, 7)

	res := p1.Distance(p2)

	assert.Equal(t, 6.928203230275509, res)
}

func TestDistanceSq(t *testing.T) {
	p1 := mymath.NewPoint3d(1, 2, 3)
	p2 := mymath.NewPoint3d(5, 6, 7)

	res := p1.DistanceSq(p2)

	assert.Equal(t, 48.0, res)
}

func TestLerp(t *testing.T) {
	p := mymath.NewPoint3d(1, 2, 3)

	res := p.Lerp(0.5, mymath.NewPoint3d(5, 6, 7))

	assert.Equal(t, 3.0, res.X)
	assert.Equal(t, 4.0, res.Y)
	assert.Equal(t, 5.0, res.Z)
}

func TestMin(t *testing.T) {
	v := mymath.NewPoint3d(1, 8, 3)
	w := mymath.NewPoint3d(4, 5, 6)

	res := v.Min(w)

	assert.Equal(t, 1.0, res.X)
	assert.Equal(t, 5.0, res.Y)
	assert.Equal(t, 3.0, res.Z)
}

func TestMax(t *testing.T) {
	v := mymath.NewPoint3d(1, 8, 3)
	w := mymath.NewPoint3d(4, 5, 6)

	res := v.Max(w)

	assert.Equal(t, 4.0, res.X)
	assert.Equal(t, 8.0, res.Y)
	assert.Equal(t, 6.0, res.Z)
}

func TestFloor(t *testing.T) {
	v := mymath.NewPoint3d(1.8, 8.9, 3.1)

	res := v.Floor()

	assert.Equal(t, 1.0, res.X)
	assert.Equal(t, 8.0, res.Y)
	assert.Equal(t, 3.0, res.Z)
}

func TestCeil(t *testing.T) {
	v := mymath.NewPoint3d(1.8, 8.9, 3.1)

	res := v.Ceil()

	assert.Equal(t, 2.0, res.X)
	assert.Equal(t, 9.0, res.Y)
	assert.Equal(t, 4.0, res.Z)
}

func TestAbs(t *testing.T) {
	v := mymath.NewPoint3d(-1, -2, -3)

	res := v.Abs()

	assert.Equal(t, 1.0, res.X)
	assert.Equal(t, 2.0, res.Y)
	assert.Equal(t, 3.0, res.Z)
}

func TestGet(t *testing.T) {
	v := mymath.NewPoint3d(1, 2, 3)

	assert.Equal(t, 1.0, v.Get(0))
	assert.Equal(t, 2.0, v.Get(1))
	assert.Equal(t, 3.0, v.Get(2))
}

func TestPermute(t *testing.T) {
	v := mymath.NewPoint3d(1, 8, 3)

	res := v.Permute(2, 0, 1)

	assert.Equal(t, 3.0, res.X)
	assert.Equal(t, 1.0, res.Y)
	assert.Equal(t, 8.0, res.Z)
}
