package mymath_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"pbrt-go/mymath"
	"testing"
)

func TestNewCylinder(t *testing.T) {
	identity := mymath.NewTransformEmpty()

	c := mymath.NewCylinder(
		16.1,
		-5,
		16.1,
		360,
		&identity,
		&identity,
		false)

	assert.Equal(t, 16.1, c.Radius)
	assert.Equal(t, -5.0, c.ZMin)
	assert.Equal(t, 16.1, c.ZMax)
	assert.Equal(t, 16.1, c.ZMax)
}

func TestCylinder_ObjectBound(t *testing.T) {
	identity := mymath.NewTransformEmpty()

	c := mymath.NewCylinder(
		16.1,
		-5,
		10,
		360,
		&identity,
		&identity,
		false)

	res := c.ObjectBound()
	assert.Equal(
		t,
		mymath.NewBounds3(mymath.NewPoint3(-16.1, -16.1, -5), mymath.NewPoint3(16.1, 16.1, 10)),
		res)
}

func TestCylinder_Area(t *testing.T) {
	identity := mymath.NewTransformEmpty()

	c := mymath.NewCylinder(
		16.1,
		-16.1,
		16.1,
		360,
		&identity,
		&identity,
		false)

	assert.InDelta(t, 4*math.Pi*16.1*16.1, c.Area(), equalDelta)
}
