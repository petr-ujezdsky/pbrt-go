package mymath_test

import (
	"pbrt-go/material"
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertAlmostEqualRayDifferential(t *testing.T, expected, actual mymath.RayDifferential, msgAndArgs ...interface{}) {
	assertAlmostEqualRay(t, expected.Ray, actual.Ray, msgAndArgs...)
	assert.Equal(t, expected.HasDifferentials, actual.HasDifferentials)
	assertAlmostEqualPoint3(t, expected.RxOrigin, actual.RxOrigin, msgAndArgs...)
	assertAlmostEqualPoint3(t, expected.RyOrigin, actual.RyOrigin, msgAndArgs...)
	assertAlmostEqualVector3(t, expected.RxDirection, actual.RxDirection, msgAndArgs...)
	assertAlmostEqualVector3(t, expected.RyDirection, actual.RyDirection, msgAndArgs...)
}

func TestRayDifferential_NewRayDifferentialRay(t *testing.T) {
	ray := mymath.NewRay(
		mymath.NewPoint3(1, 2, 3),
		mymath.NewVector3(5, 6, 7),
		9999,
		100,
		material.Medium{})

	rd := mymath.NewRayDifferentialRay(ray)

	assert.Equal(t, mymath.NewPoint3(1, 2, 3), rd.O)
	assert.Equal(t, mymath.NewVector3(5, 6, 7), rd.D)
	assert.Equal(t, 9999.0, rd.TMax)
	assert.Equal(t, float32(100.0), rd.Time)
	assert.Equal(t, material.Medium{}, rd.Medium)

	assert.Equal(t, false, rd.HasDifferentials)
	assert.Equal(t, mymath.NewPoint3(0, 0, 0), rd.RxOrigin)
	assert.Equal(t, mymath.NewPoint3(0, 0, 0), rd.RyOrigin)
	assert.Equal(t, mymath.NewVector3(0, 0, 0), rd.RxDirection)
	assert.Equal(t, mymath.NewVector3(0, 0, 0), rd.RyDirection)
}

func TestRayDifferential_ScaleDifferentials(t *testing.T) {
	ray := mymath.NewRay(
		mymath.NewPoint3(1, 2, 3),
		mymath.NewVector3(5, 6, 7),
		9999,
		100,
		material.Medium{})

	rd := mymath.NewRayDifferentialRay(ray)

	rd.RxOrigin = mymath.NewPoint3(1, 1, 1)
	rd.RyOrigin = mymath.NewPoint3(2, 2, 2)
	rd.RxDirection = mymath.NewVector3(1, 1, 1)
	rd.RyDirection = mymath.NewVector3(2, 2, 2)

	rd.ScaleDifferentials(2)

	assert.Equal(t, mymath.NewPoint3(1, 2, 3), rd.O)
	assert.Equal(t, mymath.NewVector3(5, 6, 7), rd.D)
	assert.Equal(t, 9999.0, rd.TMax)
	assert.Equal(t, float32(100.0), rd.Time)
	assert.Equal(t, material.Medium{}, rd.Medium)

	assert.Equal(t, mymath.NewPoint3(1, 0, -1), rd.RxOrigin)
	assert.Equal(t, mymath.NewPoint3(3, 2, 1), rd.RyOrigin)
	assert.Equal(t, mymath.NewVector3(-3, -4, -5), rd.RxDirection)
	assert.Equal(t, mymath.NewVector3(-1, -2, -3), rd.RyDirection)
}
