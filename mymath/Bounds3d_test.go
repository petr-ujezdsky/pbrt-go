package mymath_test

import (
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBounds3NewBounds3dP(t *testing.T) {
	p := mymath.NewPoint3(1, 2, 3)

	b := mymath.NewBounds3dP(p)

	assert.Equal(t, p, b.PMin)
	assert.Equal(t, p, b.PMax)
}

func TestBounds3NewBounds3d(t *testing.T) {
	p1 := mymath.NewPoint3(0, 0, 0)
	p2 := mymath.NewPoint3(1, 2, 3)

	b := mymath.NewBounds3d(p1, p2)

	assert.Equal(t, p1, b.PMin)
	assert.Equal(t, p2, b.PMax)
}

func TestBounds3NewBounds3dMinMax(t *testing.T) {
	p1 := mymath.NewPoint3(0, 0, 0)
	p2 := mymath.NewPoint3(1, 2, 3)

	b := mymath.NewBounds3d(p2, p1)

	assert.Equal(t, p1, b.PMin)
	assert.Equal(t, p2, b.PMax)
}

func TestBounds3Get(t *testing.T) {
	p1 := mymath.NewPoint3(5, 6, 7)
	p2 := mymath.NewPoint3(1, 2, 3)

	b := mymath.NewBounds3d(p1, p2)

	assert.Equal(t, p2, b.Get(0))
	assert.Equal(t, p1, b.Get(1))
}

func TestBounds3Corner(t *testing.T) {
	p1 := mymath.NewPoint3(0, 0, 0)
	p2 := mymath.NewPoint3(1, 1, 1)

	b := mymath.NewBounds3d(p1, p2)

	assert.Equal(t, mymath.NewPoint3(0, 0, 0), b.Corner(0))
	assert.Equal(t, mymath.NewPoint3(1, 0, 0), b.Corner(1))
	assert.Equal(t, mymath.NewPoint3(0, 1, 0), b.Corner(2))
	assert.Equal(t, mymath.NewPoint3(1, 1, 0), b.Corner(3))
	assert.Equal(t, mymath.NewPoint3(0, 0, 1), b.Corner(4))
	assert.Equal(t, mymath.NewPoint3(1, 0, 1), b.Corner(5))
	assert.Equal(t, mymath.NewPoint3(0, 1, 1), b.Corner(6))
	assert.Equal(t, mymath.NewPoint3(1, 1, 1), b.Corner(7))
}

func TestBounds3UnionP(t *testing.T) {
	p1 := mymath.NewPoint3(0, 0, 0)
	p2 := mymath.NewPoint3(1, 1, 1)

	b := mymath.NewBounds3d(p1, p2)

	res := b.UnionP(mymath.NewPoint3(-1, 0, 2))

	assert.Equal(t, mymath.NewPoint3(-1, 0, 0), res.PMin)
	assert.Equal(t, mymath.NewPoint3(1, 1, 2), res.PMax)
}

func TestBounds3UnionB(t *testing.T) {
	b1 := mymath.NewBounds3d(
		mymath.NewPoint3(0, 0, 0),
		mymath.NewPoint3(1, 1, 1))

	b2 := mymath.NewBounds3d(
		mymath.NewPoint3(0, 0, 0),
		mymath.NewPoint3(-1, -1, -1))

	res := b1.UnionB(b2)

	assert.Equal(t, mymath.NewPoint3(-1, -1, -1), res.PMin)
	assert.Equal(t, mymath.NewPoint3(1, 1, 1), res.PMax)
}

func TestBounds3Intersect(t *testing.T) {
	b1 := mymath.NewBounds3d(
		mymath.NewPoint3(0, 0, 0),
		mymath.NewPoint3(1, 1, 1))

	b2 := mymath.NewBounds3d(
		mymath.NewPoint3(0.5, 0.5, 0.5),
		mymath.NewPoint3(2, 2, 2))

	res := b1.Intersect(b2)

	assert.Equal(t, mymath.NewPoint3(0.5, 0.5, 0.5), res.PMin)
	assert.Equal(t, mymath.NewPoint3(1, 1, 1), res.PMax)
}

func TestBounds3IntersectNone(t *testing.T) {
	b1 := mymath.NewBounds3d(
		mymath.NewPoint3(0, 0, 0),
		mymath.NewPoint3(1, 1, 1))

	b2 := mymath.NewBounds3d(
		mymath.NewPoint3(-1, -1, -1),
		mymath.NewPoint3(-2, -2, -2))

	res := b1.Intersect(b2)

	assert.Equal(t, mymath.NewPoint3(-1, -1, -1), res.PMin)
	assert.Equal(t, mymath.NewPoint3(0, 0, 0), res.PMax)
}

func TestBounds3OverlapsTrue(t *testing.T) {
	b1 := mymath.NewBounds3d(
		mymath.NewPoint3(0, 0, 0),
		mymath.NewPoint3(1, 1, 1))

	b2 := mymath.NewBounds3d(
		mymath.NewPoint3(0.5, 0.5, 0.5),
		mymath.NewPoint3(2, 2, 2))

	assert.True(t, b1.Overlaps(b2))
}

func TestBounds3OverlapsFalse(t *testing.T) {
	b1 := mymath.NewBounds3d(
		mymath.NewPoint3(0, 0, 0),
		mymath.NewPoint3(1, 1, 1))

	b2 := mymath.NewBounds3d(
		mymath.NewPoint3(-1, -1, -1),
		mymath.NewPoint3(-2, -2, -2))

	assert.False(t, b1.Overlaps(b2))
}

func TestBounds3Inside(t *testing.T) {
	b := mymath.NewBounds3d(
		mymath.NewPoint3(0, 0, 0),
		mymath.NewPoint3(1, 1, 1))

	assert.True(t, b.Inside(mymath.NewPoint3(0, 0, 0)))
	assert.True(t, b.Inside(mymath.NewPoint3(1, 1, 1)))
	assert.True(t, b.Inside(mymath.NewPoint3(0.5, 0.5, 0.5)))
	assert.False(t, b.Inside(mymath.NewPoint3(2, 2, 2)))
}

func TestBounds3InsideExclusive(t *testing.T) {
	b := mymath.NewBounds3d(
		mymath.NewPoint3(0, 0, 0),
		mymath.NewPoint3(1, 1, 1))

	assert.True(t, b.InsideExclusive(mymath.NewPoint3(0, 0, 0)))
	assert.False(t, b.InsideExclusive(mymath.NewPoint3(1, 1, 1)))
}

func TestBounds3Expand(t *testing.T) {
	b := mymath.NewBounds3d(
		mymath.NewPoint3(0, 0, 0),
		mymath.NewPoint3(1, 1, 1))

	res := b.Expand(2)

	assert.Equal(t, mymath.NewPoint3(-2, -2, -2), res.PMin)
	assert.Equal(t, mymath.NewPoint3(3, 3, 3), res.PMax)
}

func TestBounds3Diagonal(t *testing.T) {
	b := mymath.NewBounds3d(
		mymath.NewPoint3(1, 2, 3),
		mymath.NewPoint3(5, 6, 7))

	res := b.Diagonal()

	assert.Equal(t, mymath.NewVector3d(4, 4, 4), res)
}

func TestBounds3SurfaceArea(t *testing.T) {
	b := mymath.NewBounds3d(
		mymath.NewPoint3(1, 2, 3),
		mymath.NewPoint3(5, 6, 7))

	res := b.SurfaceArea()

	assert.Equal(t, 96.0, res)
}

func TestBounds3Volume(t *testing.T) {
	b := mymath.NewBounds3d(
		mymath.NewPoint3(1, 2, 3),
		mymath.NewPoint3(5, 6, 7))

	res := b.Volume()

	assert.Equal(t, 64.0, res)
}

func TestBounds3MaximumExtent(t *testing.T) {
	assert.Equal(t,
		0,
		mymath.NewBounds3d(
			mymath.NewPoint3(0, 0, 0),
			mymath.NewPoint3(5, 1, 1)).MaximumExtent())

	assert.Equal(t,
		1,
		mymath.NewBounds3d(
			mymath.NewPoint3(0, 0, 0),
			mymath.NewPoint3(1, 5, 1)).MaximumExtent())

	assert.Equal(t,
		2,
		mymath.NewBounds3d(
			mymath.NewPoint3(0, 0, 0),
			mymath.NewPoint3(1, 1, 5)).MaximumExtent())
}

// func TestBounds3Overlaps(t *testing.T) {
// 	b1 := mymath.NewBounds3d(
// 		mymath.NewPoint3(0, 0, 0),
// 		mymath.NewPoint3(1, 1, 1))

// 	b2 := mymath.NewBounds3d(
// 		mymath.NewPoint3(0.5, 0.5, 0.5),
// 		mymath.NewPoint3(2, 2, 2))

// 	res := b1.Intersect(b2)

// 	assert.Equal(t, mymath.NewPoint3(0.5, 0.5, 0.5), res.PMin)
// 	assert.Equal(t, mymath.NewPoint3(1, 1, 1), res.PMax)
// }
