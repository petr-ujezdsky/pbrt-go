package mymath_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"pbrt-go/mymath"
	"testing"
)

func TestNewSphere(t *testing.T) {
	identity := mymath.NewTransformEmpty()

	s := mymath.NewSphere(
		16.1,
		-5,
		16.1,
		360,
		&identity,
		&identity,
		false)

	assert.Equal(t, 16.1, s.Radius)
	assert.Equal(t, -5.0, s.ZMin)
	assert.Equal(t, 16.1, s.ZMax)
	assert.Equal(t, 16.1, s.ZMax)
	assert.InDelta(t, 1.8865773873649494, s.ThetaMin, equalDelta)
	assert.Equal(t, 0.0, s.ThetaMax)
}

func TestSphere_ObjectBound(t *testing.T) {
	identity := mymath.NewTransformEmpty()

	s := mymath.NewSphere(
		16.1,
		-5,
		10,
		360,
		&identity,
		&identity,
		false)

	res := s.ObjectBound()
	assert.Equal(
		t,
		mymath.NewBounds3(mymath.NewPoint3(-16.1, -16.1, -5), mymath.NewPoint3(16.1, 16.1, 10)),
		res)
}

func TestSphere_Area(t *testing.T) {
	identity := mymath.NewTransformEmpty()

	s := mymath.NewSphere(
		16.1,
		-16.1,
		16.1,
		360,
		&identity,
		&identity,
		false)

	assert.InDelta(t, 4*math.Pi*16.1*16.1, s.Area(), equalDelta)
}
