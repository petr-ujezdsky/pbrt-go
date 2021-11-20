package mymath_test

import (
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertAlmostEqualRay(t *testing.T, expected, actual mymath.Ray, msgAndArgs ...interface{}) {
	assertAlmostEqualPoint3(t, expected.O, actual.O, equalDelta, msgAndArgs)
	assertAlmostEqualVector3(t, expected.D, actual.D, equalDelta, msgAndArgs)

	assert.Equal(t, expected.Time, actual.Time, msgAndArgs...)
	assert.Equal(t, expected.TMax, actual.TMax, msgAndArgs...)
	assert.Equal(t, expected.Medium, actual.Medium, msgAndArgs...)
}
