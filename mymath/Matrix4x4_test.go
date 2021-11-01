package mymath_test

import (
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrix4x4_NewMatrix4x4All(t *testing.T) {
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

func TestMatrix4x4_Identity(t *testing.T) {
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

// data from https://docs.microsoft.com/en-us/archive/msdn-magazine/2016/july/test-run-matrix-inversion-using-csharp
func TestMatrix4x4_Inverse(t *testing.T) {
	m := mymath.NewMatrix4x4All(
		3.0, 7.0, 2.0, 5.0,
		1.0, 8.0, 4.0, 2.0,
		2.0, 1.0, 9.0, 3.0,
		5.0, 4.0, 7.0, 1.0)

	res := m.Inverse()

	// original data from web
	// expected := mymath.NewMatrix4x4All(
	// 	0.097, -0.183, -0.115, 0.224,
	// 	-0.019, 0.146, -0.068, 0.010,
	// 	-0.087, 0.064, 0.103, -0.002,
	// 	0.204, -0.120, 0.123, -0.147)

	// actual data with minor floating point errors
	expected := mymath.NewMatrix4x4All(
		0.09708738, -0.18270081, -0.11473962, 0.22418359,
		-0.019417476, 0.14563109, -0.06796117, 0.009708739,
		-0.08737864, 0.06443072, 0.10326566, -0.0017652243,
		0.20388348, -0.1200353, 0.12268313, -0.1473963)

	assert.Equal(t, expected, *res)
}

// try to create identity matrix by M*M'
func TestMatrix4x4_Inverse_ToIdentity(t *testing.T) {
	m := mymath.NewMatrix4x4All(
		3.0, 7.0, 2.0, 5.0,
		1.0, 8.0, 4.0, 2.0,
		2.0, 1.0, 9.0, 3.0,
		5.0, 4.0, 7.0, 1.0)

	res := m.Multiply(m.Inverse())

	expected := mymath.NewMatrix4x4All(
		0.99999994, 1.1920929e-07, -5.9604645e-08, 0,
		-2.9802322e-08, 1.0000002, -7.4505806e-08, 2.9802322e-08,
		0, 2.9802322e-08, 0.9999999, 0,
		-4.4703484e-08, 5.2154064e-08, -3.7252903e-08, 1)

	assert.Equal(t, expected, *res)
}
