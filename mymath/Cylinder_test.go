package mymath_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"pbrt-go/material"
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

func TestCylinder_Intersect(t *testing.T) {
	identity := mymath.NewTransformEmpty()

	c := mymath.NewCylinder(
		16.1,
		-16.1,
		16.1,
		360,
		&identity,
		&identity,
		false)

	ray := mymath.NewRay(
		mymath.NewPoint3(-20, 0, 0),
		mymath.NewVector3(1, 0, 0),
		50,
		0,
		material.Medium{})

	ok, tHit, si := c.Intersect(ray, false)
	assert.Equal(t, true, ok)
	assert.InDelta(t, 3.9, tHit, equalDelta)

	assert.Equal(t, mymath.NewPoint2(0.5, 0.5), si.Uv)
	assert.Equal(t, mymath.NewPoint3(-16.1, 0, 0), si.Interaction.P)
	assert.Equal(t, 0.0, si.Interaction.Time)
	InDeltaVector3(t, mymath.NewVector3(0, 0, 0), si.Interaction.PError)
	assert.Equal(t, mymath.NewVector3(-1, 0, 0), si.Interaction.Wo)
	assert.Equal(t, mymath.NewNormal3(-1, 0, 0), si.Interaction.N)

	InDeltaVector3(t, mymath.NewVector3(0, -101.15928344559134, 0), si.Dpdu)
	InDeltaVector3(t, mymath.NewVector3(0, 0, 32.2), si.Dpdv)
	InDeltaNormal3(t, mymath.NewNormal3(0, -6.283185307179587, 0), si.Dndu)
	InDeltaNormal3(t, mymath.NewNormal3(0, 0, 0), si.Dndv)
}

func TestCylinder_IntersectP(t *testing.T) {
	identity := mymath.NewTransformEmpty()

	s := mymath.NewCylinder(
		16.1,
		-16.1,
		16.1,
		360,
		&identity,
		&identity,
		false)

	ray := mymath.NewRay(
		mymath.NewPoint3(-20, 0, 0),
		mymath.NewVector3(1, 0, 0),
		50,
		0,
		material.Medium{})

	ok := s.IntersectP(s, ray, false)
	assert.Equal(t, true, ok)
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
