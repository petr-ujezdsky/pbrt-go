package mymath_test

import (
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLerp(t *testing.T) {
	assert.Equal(t, 0.0, mymath.Lerp(-1.0, 10.0, 20.0))
	assert.Equal(t, 10.0, mymath.Lerp(0.0, 10.0, 20.0))
	assert.Equal(t, 15.0, mymath.Lerp(0.5, 10.0, 20.0))
	assert.Equal(t, 20.0, mymath.Lerp(1.0, 10.0, 20.0))
	assert.Equal(t, 40.0, mymath.Lerp(3.0, 10.0, 20.0))
}
