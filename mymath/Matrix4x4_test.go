package mymath_test

import (
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrix4x4NewMatrix4x4All(t *testing.T) {
	m := mymath.NewMatrix4x4All(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16)

	// 1. row
	assert.Equal(t, float32(1.0), m.M[0][0])
	assert.Equal(t, float32(2.0), m.M[0][1])
	assert.Equal(t, float32(3.0), m.M[0][2])
	assert.Equal(t, float32(4.0), m.M[0][3])

	// 2. row
	assert.Equal(t, float32(5.0), m.M[1][0])
	assert.Equal(t, float32(6.0), m.M[1][1])
	assert.Equal(t, float32(7.0), m.M[1][2])
	assert.Equal(t, float32(8.0), m.M[1][3])

	// 3. row
	assert.Equal(t, float32(9.0), m.M[2][0])
	assert.Equal(t, float32(10.0), m.M[2][1])
	assert.Equal(t, float32(11.0), m.M[2][2])
	assert.Equal(t, float32(12.0), m.M[2][3])

	// 4. row
	assert.Equal(t, float32(13.0), m.M[3][0])
	assert.Equal(t, float32(14.0), m.M[3][1])
	assert.Equal(t, float32(15.0), m.M[3][2])
	assert.Equal(t, float32(16.0), m.M[3][3])
}

func TestMatrix4x4Identity(t *testing.T) {
	m := mymath.Identity()

	assert.Equal(t, float32(1.0), m.M[0][0])
	assert.Equal(t, float32(0.0), m.M[0][1])
	assert.Equal(t, float32(0.0), m.M[0][2])
	assert.Equal(t, float32(0.0), m.M[0][3])

	assert.Equal(t, float32(0.0), m.M[1][0])
	assert.Equal(t, float32(1.0), m.M[1][1])
	assert.Equal(t, float32(0.0), m.M[1][2])
	assert.Equal(t, float32(0.0), m.M[1][3])

	assert.Equal(t, float32(0.0), m.M[2][0])
	assert.Equal(t, float32(0.0), m.M[2][1])
	assert.Equal(t, float32(1.0), m.M[2][2])
	assert.Equal(t, float32(0.0), m.M[2][3])

	assert.Equal(t, float32(0.0), m.M[3][0])
	assert.Equal(t, float32(0.0), m.M[3][1])
	assert.Equal(t, float32(0.0), m.M[3][2])
	assert.Equal(t, float32(1.0), m.M[3][3])
}

func TestMatrix4x4_Transpose(t *testing.T) {
	m := mymath.NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	res := m.Transpose()

	expected := mymath.NewMatrix4x4All(
		5, 7, 9, 1,
		2, 3, 3, 8,
		8, 5, 2, 3,
		3, 3, 4, 8)

	assert.Equal(t, expected, *res)
}

func TestMatrix4x4_Multiply(t *testing.T) {
	m1 := mymath.NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	m2 := mymath.NewMatrix4x4All(
		3, 9, 9, 3,
		5, 1, 9, 2,
		6, 4, 4, 4,
		7, 9, 2, 1)

	res := m2.Multiply(&m1)

	expected := mymath.NewMatrix4x4All(
		162, 84, 96, 96,
		115, 56, 69, 70,
		98, 68, 88, 78,
		117, 55, 108, 64)

	assert.Equal(t, expected, *res)
}
