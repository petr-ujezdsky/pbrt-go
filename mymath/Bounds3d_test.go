package mymath_test

import (
	"pbrt-go/mymath"
	"pbrt-go/mymath/point3d"
	"pbrt-go/mymath/vector3d"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBounds3dP(t *testing.T) {
	p := point3d.NewPoint3d(1, 2, 3)

	b := mymath.NewBounds3dP(p)

	assert.Equal(t, p, b.PMin)
	assert.Equal(t, p, b.PMax)
}

func TestConstructor(t *testing.T) {
	p1 := point3d.NewPoint3d(0, 0, 0)
	p2 := point3d.NewPoint3d(1, 2, 3)

	b := mymath.NewBounds3d(p1, p2)

	assert.Equal(t, p1, b.PMin)
	assert.Equal(t, p2, b.PMax)
}

func TestConstructorMinMax(t *testing.T) {
	p1 := point3d.NewPoint3d(0, 0, 0)
	p2 := point3d.NewPoint3d(1, 2, 3)

	b := mymath.NewBounds3d(p2, p1)

	assert.Equal(t, p1, b.PMin)
	assert.Equal(t, p2, b.PMax)
}

func TestGet(t *testing.T) {
	p1 := point3d.NewPoint3d(5, 6, 7)
	p2 := point3d.NewPoint3d(1, 2, 3)

	b := mymath.NewBounds3d(p1, p2)

	assert.Equal(t, p2, b.Get(0))
	assert.Equal(t, p1, b.Get(1))
}

func TestCorner(t *testing.T) {
	p1 := point3d.NewPoint3d(0, 0, 0)
	p2 := point3d.NewPoint3d(1, 1, 1)

	b := mymath.NewBounds3d(p1, p2)

	assert.Equal(t, point3d.NewPoint3d(0, 0, 0), b.Corner(0))
	assert.Equal(t, point3d.NewPoint3d(1, 0, 0), b.Corner(1))
	assert.Equal(t, point3d.NewPoint3d(0, 1, 0), b.Corner(2))
	assert.Equal(t, point3d.NewPoint3d(1, 1, 0), b.Corner(3))
	assert.Equal(t, point3d.NewPoint3d(0, 0, 1), b.Corner(4))
	assert.Equal(t, point3d.NewPoint3d(1, 0, 1), b.Corner(5))
	assert.Equal(t, point3d.NewPoint3d(0, 1, 1), b.Corner(6))
	assert.Equal(t, point3d.NewPoint3d(1, 1, 1), b.Corner(7))
}

func TestUnionP(t *testing.T) {
	p1 := point3d.NewPoint3d(0, 0, 0)
	p2 := point3d.NewPoint3d(1, 1, 1)

	b := mymath.NewBounds3d(p1, p2)

	res := b.UnionP(point3d.NewPoint3d(-1, 0, 2))

	assert.Equal(t, point3d.NewPoint3d(-1, 0, 0), res.PMin)
	assert.Equal(t, point3d.NewPoint3d(1, 1, 2), res.PMax)
}

func TestUnionB(t *testing.T) {
	b1 := mymath.NewBounds3d(
		point3d.NewPoint3d(0, 0, 0),
		point3d.NewPoint3d(1, 1, 1))

	b2 := mymath.NewBounds3d(
		point3d.NewPoint3d(0, 0, 0),
		point3d.NewPoint3d(-1, -1, -1))

	res := b1.UnionB(b2)

	assert.Equal(t, point3d.NewPoint3d(-1, -1, -1), res.PMin)
	assert.Equal(t, point3d.NewPoint3d(1, 1, 1), res.PMax)
}

func TestIntersect(t *testing.T) {
	b1 := mymath.NewBounds3d(
		point3d.NewPoint3d(0, 0, 0),
		point3d.NewPoint3d(1, 1, 1))

	b2 := mymath.NewBounds3d(
		point3d.NewPoint3d(0.5, 0.5, 0.5),
		point3d.NewPoint3d(2, 2, 2))

	res := b1.Intersect(b2)

	assert.Equal(t, point3d.NewPoint3d(0.5, 0.5, 0.5), res.PMin)
	assert.Equal(t, point3d.NewPoint3d(1, 1, 1), res.PMax)
}

func TestIntersectNone(t *testing.T) {
	b1 := mymath.NewBounds3d(
		point3d.NewPoint3d(0, 0, 0),
		point3d.NewPoint3d(1, 1, 1))

	b2 := mymath.NewBounds3d(
		point3d.NewPoint3d(-1, -1, -1),
		point3d.NewPoint3d(-2, -2, -2))

	res := b1.Intersect(b2)

	assert.Equal(t, point3d.NewPoint3d(-1, -1, -1), res.PMin)
	assert.Equal(t, point3d.NewPoint3d(0, 0, 0), res.PMax)
}

func TestOverlapsTrue(t *testing.T) {
	b1 := mymath.NewBounds3d(
		point3d.NewPoint3d(0, 0, 0),
		point3d.NewPoint3d(1, 1, 1))

	b2 := mymath.NewBounds3d(
		point3d.NewPoint3d(0.5, 0.5, 0.5),
		point3d.NewPoint3d(2, 2, 2))

	assert.True(t, b1.Overlaps(b2))
}

func TestOverlapsFalse(t *testing.T) {
	b1 := mymath.NewBounds3d(
		point3d.NewPoint3d(0, 0, 0),
		point3d.NewPoint3d(1, 1, 1))

	b2 := mymath.NewBounds3d(
		point3d.NewPoint3d(-1, -1, -1),
		point3d.NewPoint3d(-2, -2, -2))

	assert.False(t, b1.Overlaps(b2))
}

func TestInside(t *testing.T) {
	b := mymath.NewBounds3d(
		point3d.NewPoint3d(0, 0, 0),
		point3d.NewPoint3d(1, 1, 1))

	assert.True(t, b.Inside(point3d.NewPoint3d(0, 0, 0)))
	assert.True(t, b.Inside(point3d.NewPoint3d(1, 1, 1)))
	assert.True(t, b.Inside(point3d.NewPoint3d(0.5, 0.5, 0.5)))
	assert.False(t, b.Inside(point3d.NewPoint3d(2, 2, 2)))
}

func TestInsideExclusive(t *testing.T) {
	b := mymath.NewBounds3d(
		point3d.NewPoint3d(0, 0, 0),
		point3d.NewPoint3d(1, 1, 1))

	assert.True(t, b.InsideExclusive(point3d.NewPoint3d(0, 0, 0)))
	assert.False(t, b.InsideExclusive(point3d.NewPoint3d(1, 1, 1)))
}

func TestExpand(t *testing.T) {
	b := mymath.NewBounds3d(
		point3d.NewPoint3d(0, 0, 0),
		point3d.NewPoint3d(1, 1, 1))

	res := b.Expand(2)

	assert.Equal(t, point3d.NewPoint3d(-2, -2, -2), res.PMin)
	assert.Equal(t, point3d.NewPoint3d(3, 3, 3), res.PMax)
}

func TestDiagonal(t *testing.T) {
	b := mymath.NewBounds3d(
		point3d.NewPoint3d(1, 2, 3),
		point3d.NewPoint3d(5, 6, 7))

	res := b.Diagonal()

	assert.Equal(t, vector3d.NewVector3d(4, 4, 4), res)
}

func TestSurfaceArea(t *testing.T) {
	b := mymath.NewBounds3d(
		point3d.NewPoint3d(1, 2, 3),
		point3d.NewPoint3d(5, 6, 7))

	res := b.SurfaceArea()

	assert.Equal(t, 96.0, res)
}

func TestVolume(t *testing.T) {
	b := mymath.NewBounds3d(
		point3d.NewPoint3d(1, 2, 3),
		point3d.NewPoint3d(5, 6, 7))

	res := b.Volume()

	assert.Equal(t, 64.0, res)
}

func TestMaximumExtent(t *testing.T) {
	assert.Equal(t,
		0,
		mymath.NewBounds3d(
			point3d.NewPoint3d(0, 0, 0),
			point3d.NewPoint3d(5, 1, 1)).MaximumExtent())

	assert.Equal(t,
		1,
		mymath.NewBounds3d(
			point3d.NewPoint3d(0, 0, 0),
			point3d.NewPoint3d(1, 5, 1)).MaximumExtent())

	assert.Equal(t,
		2,
		mymath.NewBounds3d(
			point3d.NewPoint3d(0, 0, 0),
			point3d.NewPoint3d(1, 1, 5)).MaximumExtent())
}

// func TestOverlaps(t *testing.T) {
// 	b1 := mymath.NewBounds3d(
// 		point3d.NewPoint3d(0, 0, 0),
// 		point3d.NewPoint3d(1, 1, 1))

// 	b2 := mymath.NewBounds3d(
// 		point3d.NewPoint3d(0.5, 0.5, 0.5),
// 		point3d.NewPoint3d(2, 2, 2))

// 	res := b1.Intersect(b2)

// 	assert.Equal(t, point3d.NewPoint3d(0.5, 0.5, 0.5), res.PMin)
// 	assert.Equal(t, point3d.NewPoint3d(1, 1, 1), res.PMax)
// }
