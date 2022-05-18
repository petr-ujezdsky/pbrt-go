package mymath_test

import (
	"github.com/stretchr/testify/assert"
	"pbrt-go/mymath"
	"testing"
)

const equalDelta = 0.00001

func InDeltaMatrix4x4(t *testing.T, expected, actual mymath.Matrix4x4, msgAndArgs ...interface{}) {
	// 1. row
	assert.InDelta(t, expected.M[0][0], actual.M[0][0], equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.M[0][1], actual.M[0][1], equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.M[0][2], actual.M[0][2], equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.M[0][3], actual.M[0][3], equalDelta, msgAndArgs...)

	// 2. row
	assert.InDelta(t, expected.M[1][0], actual.M[1][0], equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.M[1][1], actual.M[1][1], equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.M[1][2], actual.M[1][2], equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.M[1][3], actual.M[1][3], equalDelta, msgAndArgs...)

	// 3. row
	assert.InDelta(t, expected.M[2][0], actual.M[2][0], equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.M[2][1], actual.M[2][1], equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.M[2][2], actual.M[2][2], equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.M[2][3], actual.M[2][3], equalDelta, msgAndArgs...)

	// 4. row
	assert.InDelta(t, expected.M[3][0], actual.M[3][0], equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.M[3][1], actual.M[3][1], equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.M[3][2], actual.M[3][2], equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.M[3][3], actual.M[3][3], equalDelta, msgAndArgs...)
}

func InDeltaVector3(t *testing.T, expected, actual mymath.Vector3, msgAndArgs ...interface{}) {
	assert.InDelta(t, expected.X, actual.X, equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.Y, actual.Y, equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.Z, actual.Z, equalDelta, msgAndArgs...)
}

func InDeltaNormal3(t *testing.T, expected, actual mymath.Normal3, msgAndArgs ...interface{}) {
	assert.InDelta(t, expected.X, actual.X, equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.Y, actual.Y, equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.Z, actual.Z, equalDelta, msgAndArgs...)
}

func InDeltaQuaternion(t *testing.T, expected, actual mymath.Quaternion, msgAndArgs ...interface{}) {
	assert.InDelta(t, expected.V.X, actual.V.X, equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.V.Y, actual.V.Y, equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.V.Z, actual.V.Z, equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.W, actual.W, equalDelta, msgAndArgs...)
}

func InDeltaPoint3(t *testing.T, expected, actual mymath.Point3, msgAndArgs ...interface{}) {
	assert.InDelta(t, expected.X, actual.X, equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.Y, actual.Y, equalDelta, msgAndArgs...)
	assert.InDelta(t, expected.Z, actual.Z, equalDelta, msgAndArgs...)
}

func InDeltaRay(t *testing.T, expected, actual mymath.Ray, msgAndArgs ...interface{}) {
	InDeltaPoint3(t, expected.O, actual.O, equalDelta, msgAndArgs)
	InDeltaVector3(t, expected.D, actual.D, equalDelta, msgAndArgs)

	assert.Equal(t, expected.Time, actual.Time, msgAndArgs...)
	assert.Equal(t, expected.TMax, actual.TMax, msgAndArgs...)
	assert.Equal(t, expected.Medium, actual.Medium, msgAndArgs...)
}

func InDeltaRayDifferential(t *testing.T, expected, actual mymath.RayDifferential, msgAndArgs ...interface{}) {
	InDeltaRay(t, expected.Ray, actual.Ray, msgAndArgs...)
	assert.Equal(t, expected.HasDifferentials, actual.HasDifferentials)
	InDeltaPoint3(t, expected.RxOrigin, actual.RxOrigin, msgAndArgs...)
	InDeltaPoint3(t, expected.RyOrigin, actual.RyOrigin, msgAndArgs...)
	InDeltaVector3(t, expected.RxDirection, actual.RxDirection, msgAndArgs...)
	InDeltaVector3(t, expected.RyDirection, actual.RyDirection, msgAndArgs...)
}
