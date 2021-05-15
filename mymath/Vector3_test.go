package mymath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	v := NewVector3d(1, 2, 3)

	assert.Equal(t, 1.0, v.X)
	assert.Equal(t, 2.0, v.Y)
	assert.Equal(t, 3.0, v.Z)
}

func TestLengthSq(t *testing.T) {
	v := NewVector3d(3, 3, 3)

	assert.Equal(t, 27.0, v.LengthSq())
}

func TestLength(t *testing.T) {
	v := NewVector3d(1, 2, 3)

	assert.Equal(t, 3.7416573867739413, v.Length())
}
