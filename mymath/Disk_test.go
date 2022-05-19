package mymath_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"pbrt-go/material"
	"pbrt-go/mymath"
	"testing"
)

func TestNewDisk(t *testing.T) {
	identity := mymath.NewTransformEmpty()

	d := mymath.NewDisk(
		5,
		30,
		20,
		360,
		&identity,
		&identity,
		false)

	assert.Equal(t, 30.0, d.Radius)
	assert.Equal(t, 20.0, d.InnerRadius)
	assert.InDelta(t, 2*math.Pi, d.PhiMax, equalDelta)
}

func TestDisk_ObjectBound(t *testing.T) {
	identity := mymath.NewTransformEmpty()

	d := mymath.NewDisk(
		5,
		30,
		20,
		360,
		&identity,
		&identity,
		false)

	res := d.ObjectBound()
	assert.Equal(
		t,
		mymath.NewBounds3(mymath.NewPoint3(-30, -30, 5), mymath.NewPoint3(30, 30, 5)),
		res)
}

func TestDisk_Intersect(t *testing.T) {
	identity := mymath.NewTransformEmpty()

	d := mymath.NewDisk(
		5,
		30,
		20,
		360,
		&identity,
		&identity,
		false)

	ray := mymath.NewRay(
		mymath.NewPoint3(25, 0, 15),
		mymath.NewVector3(0, 0, -1),
		50,
		0,
		material.Medium{})

	ok, tHit, si := d.Intersect(ray, false)
	assert.Equal(t, true, ok)
	assert.InDelta(t, 10.0, tHit, equalDelta)

	assert.Equal(t, mymath.NewPoint2(0, 0.5), si.Uv)
	assert.Equal(t, mymath.NewPoint3(25, 0, 5), si.Interaction.P)
	assert.Equal(t, 0.0, si.Interaction.Time)
	InDeltaVector3(t, mymath.NewVector3(0, 0, 0), si.Interaction.PError)
	assert.Equal(t, mymath.NewVector3(0, 0, 1), si.Interaction.Wo)
	assert.Equal(t, mymath.NewNormal3(0, 0, 1), si.Interaction.N)

	InDeltaVector3(t, mymath.NewVector3(0, 157.07963267948966, 0), si.Dpdu)
	InDeltaVector3(t, mymath.NewVector3(-10, 0, 0), si.Dpdv)
	InDeltaNormal3(t, mymath.NewNormal3(0, 0, 0), si.Dndu)
	InDeltaNormal3(t, mymath.NewNormal3(0, 0, 0), si.Dndv)
}

func TestDisk_IntersectP(t *testing.T) {
	identity := mymath.NewTransformEmpty()

	d := mymath.NewDisk(
		5,
		30,
		20,
		360,
		&identity,
		&identity,
		false)

	ray := mymath.NewRay(
		mymath.NewPoint3(25, 0, 15),
		mymath.NewVector3(0, 0, -1),
		50,
		0,
		material.Medium{})

	ok := d.IntersectP(d, ray, false)
	assert.Equal(t, true, ok)
}

func TestDisk_Area(t *testing.T) {
	identity := mymath.NewTransformEmpty()

	d := mymath.NewDisk(
		5,
		30,
		20,
		360,
		&identity,
		&identity,
		false)

	assert.InDelta(t, 2*math.Pi*0.5*(30*30-20*20), d.Area(), equalDelta)
}
