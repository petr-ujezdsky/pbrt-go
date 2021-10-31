package mymath_test

import (
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint3NewPoint3d(t *testing.T) {
	p := mymath.NewPoint3d(1, 2, 3)

	assert.Equal(t, 1.0, p.X)
	assert.Equal(t, 2.0, p.Y)
	assert.Equal(t, 3.0, p.Z)
}

func TestPoint3AddP(t *testing.T) {
	p1 := mymath.NewPoint3d(1, 2, 3)
	p2 := mymath.NewPoint3d(5, 6, 7)

	res := p1.AddP(p2)

	assert.Equal(t, 6.0, res.X)
	assert.Equal(t, 8.0, res.Y)
	assert.Equal(t, 10.0, res.Z)
}

func TestPoint3AddV(t *testing.T) {
	p := mymath.NewPoint3d(1, 2, 3)
	v := mymath.NewVector3d(5, 6, 7)

	res := p.AddV(v)

	assert.Equal(t, 6.0, res.X)
	assert.Equal(t, 8.0, res.Y)
	assert.Equal(t, 10.0, res.Z)
}

func TestPoint3SubtractP(t *testing.T) {
	p1 := mymath.NewPoint3d(1, 2, 3)
	p2 := mymath.NewPoint3d(5, 6, 7)

	res := p1.SubtractP(p2)

	assert.Equal(t, -4.0, res.X)
	assert.Equal(t, -4.0, res.Y)
	assert.Equal(t, -4.0, res.Z)
}

func TestPoint3SubtractV(t *testing.T) {
	p := mymath.NewPoint3d(1, 2, 3)
	v := mymath.NewVector3d(5, 6, 7)

	res := p.SubtractV(v)

	assert.Equal(t, -4.0, res.X)
	assert.Equal(t, -4.0, res.Y)
	assert.Equal(t, -4.0, res.Z)
}

func TestPoint3Multiply(t *testing.T) {
	v := mymath.NewPoint3d(1, 2, 3)

	res := v.Multiply(5.0)

	assert.Equal(t, 5.0, res.X)
	assert.Equal(t, 10.0, res.Y)
	assert.Equal(t, 15.0, res.Z)
}

func TestPoint3Distance(t *testing.T) {
	p1 := mymath.NewPoint3d(1, 2, 3)
	p2 := mymath.NewPoint3d(5, 6, 7)

	res := p1.Distance(p2)

	assert.Equal(t, 6.928203230275509, res)
}

func TestPoint3DistanceSq(t *testing.T) {
	p1 := mymath.NewPoint3d(1, 2, 3)
	p2 := mymath.NewPoint3d(5, 6, 7)

	res := p1.DistanceSq(p2)

	assert.Equal(t, 48.0, res)
}

func TestPoint3Lerp(t *testing.T) {
	p := mymath.NewPoint3d(1, 2, 3)

	res := p.Lerp(0.5, mymath.NewPoint3d(5, 6, 7))

	assert.Equal(t, 3.0, res.X)
	assert.Equal(t, 4.0, res.Y)
	assert.Equal(t, 5.0, res.Z)
}

func TestPoint3LerpP(t *testing.T) {
	p := mymath.NewPoint3d(1, 2, 3)

	res := p.LerpP(mymath.NewPoint3d(0, 0.5, 1), mymath.NewPoint3d(5, 6, 7))

	assert.Equal(t, 1.0, res.X)
	assert.Equal(t, 4.0, res.Y)
	assert.Equal(t, 7.0, res.Z)
}

func TestPoint3Min(t *testing.T) {
	v := mymath.NewPoint3d(1, 8, 3)
	w := mymath.NewPoint3d(4, 5, 6)

	res := v.Min(w)

	assert.Equal(t, 1.0, res.X)
	assert.Equal(t, 5.0, res.Y)
	assert.Equal(t, 3.0, res.Z)
}

func TestPoint3Max(t *testing.T) {
	v := mymath.NewPoint3d(1, 8, 3)
	w := mymath.NewPoint3d(4, 5, 6)

	res := v.Max(w)

	assert.Equal(t, 4.0, res.X)
	assert.Equal(t, 8.0, res.Y)
	assert.Equal(t, 6.0, res.Z)
}

func TestPoint3Floor(t *testing.T) {
	v := mymath.NewPoint3d(1.8, 8.9, 3.1)

	res := v.Floor()

	assert.Equal(t, 1.0, res.X)
	assert.Equal(t, 8.0, res.Y)
	assert.Equal(t, 3.0, res.Z)
}

func TestPoint3Ceil(t *testing.T) {
	v := mymath.NewPoint3d(1.8, 8.9, 3.1)

	res := v.Ceil()

	assert.Equal(t, 2.0, res.X)
	assert.Equal(t, 9.0, res.Y)
	assert.Equal(t, 4.0, res.Z)
}

func TestPoint3Abs(t *testing.T) {
	v := mymath.NewPoint3d(-1, -2, -3)

	res := v.Abs()

	assert.Equal(t, 1.0, res.X)
	assert.Equal(t, 2.0, res.Y)
	assert.Equal(t, 3.0, res.Z)
}

func TestPoint3Get(t *testing.T) {
	v := mymath.NewPoint3d(1, 2, 3)

	assert.Equal(t, 1.0, v.Get(0))
	assert.Equal(t, 2.0, v.Get(1))
	assert.Equal(t, 3.0, v.Get(2))
}

func TestPoint3Permute(t *testing.T) {
	v := mymath.NewPoint3d(1, 8, 3)

	res := v.Permute(2, 0, 1)

	assert.Equal(t, 3.0, res.X)
	assert.Equal(t, 1.0, res.Y)
	assert.Equal(t, 8.0, res.Z)
}
