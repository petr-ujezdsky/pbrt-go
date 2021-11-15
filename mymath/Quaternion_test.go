package mymath_test

import (
	"math"
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuaternion_NewQuaternionEmpty(t *testing.T) {
	q := mymath.NewQuaternionEmpty()

	assert.Equal(t, 0.0, q.V.X)
	assert.Equal(t, 0.0, q.V.Y)
	assert.Equal(t, 0.0, q.V.Z)
	assert.Equal(t, 1.0, q.W)
}

func TestQuaternion_NewQuaternion(t *testing.T) {
	q := mymath.NewQuaternion(
		mymath.NewVector3(1, 2, 3),
		4)

	assert.Equal(t, 1.0, q.V.X)
	assert.Equal(t, 2.0, q.V.Y)
	assert.Equal(t, 3.0, q.V.Z)
	assert.Equal(t, 4.0, q.W)
}

func TestQuaternion_NewQuaternionFull(t *testing.T) {
	q := mymath.NewQuaternionFull(1, 2, 3, 4)

	assert.Equal(t, 1.0, q.V.X)
	assert.Equal(t, 2.0, q.V.Y)
	assert.Equal(t, 3.0, q.V.Z)
	assert.Equal(t, 4.0, q.W)
}

func TestQuaternion_Add(t *testing.T) {
	q1 := mymath.NewQuaternionFull(1, 2, 3, 4)
	q2 := mymath.NewQuaternionFull(5, 6, 7, 8)

	res := q1.Add(q2)

	assert.Equal(t, mymath.NewQuaternionFull(6, 8, 10, 12), res)
}
func TestQuaternion_Subtract(t *testing.T) {
	q1 := mymath.NewQuaternionFull(1, 2, 3, 4)
	q2 := mymath.NewQuaternionFull(5, 6, 7, 8)

	res := q1.Subtract(q2)

	assert.Equal(t, mymath.NewQuaternionFull(-4, -4, -4, -4), res)
}

func TestQuaternion_Multiply(t *testing.T) {
	q := mymath.NewQuaternionFull(1, 2, 3, 4)

	res := q.Multiply(5.0)

	assert.Equal(t, mymath.NewQuaternionFull(5, 10, 15, 20), res)
}

func TestQuaternion_Divide(t *testing.T) {
	q := mymath.NewQuaternionFull(2, 4, -8, 16)

	res := q.Divide(2.0)

	assert.Equal(t, mymath.NewQuaternionFull(1, 2, -4, 8), res)
}

func TestQuaternion_Negate(t *testing.T) {
	q := mymath.NewQuaternionFull(2, 4, -8, 16)

	res := q.Negate()

	assert.Equal(t, mymath.NewQuaternionFull(-2, -4, 8, -16), res)
}

func TestQuaternion_Normalize(t *testing.T) {
	q := mymath.NewQuaternionFull(1, 2, 3, 4)

	res := q.Normalize()

	assert.Equal(t, 0.18257418583505536, res.V.X)
	assert.Equal(t, 0.3651483716701107, res.V.Y)
	assert.Equal(t, 0.5477225575051661, res.V.Z)
	assert.Equal(t, 0.7302967433402214, res.W)

	lengthSq := res.Dot(res)
	assert.Equal(t, 0.9999999999999998, lengthSq)
	assert.Equal(t, 0.9999999999999999, math.Sqrt(lengthSq))
}

func TestQuaternion_Dot(t *testing.T) {
	q1 := mymath.NewQuaternionFull(-1, -2, -3, -4)
	q2 := mymath.NewQuaternionFull(4, 5, 6, 7)

	res := q1.Dot(q2)

	assert.Equal(t, -60.0, res)
}

//////////////////////////////////////////////////////////////////////////////////////////////////
// BENCHMARKS ////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
func BenchmarkQuaternion_NewQuaternionEmpty(b *testing.B) {
	var res mymath.Quaternion

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = mymath.NewQuaternionEmpty()
	}

	assert.NotNil(b, res)
}

func BenchmarkQuaternion_NewQuaternion(b *testing.B) {
	v := mymath.NewVector3(1, 2, 3)
	w := 4.0

	var res mymath.Quaternion

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = mymath.NewQuaternion(v, w)
	}

	assert.NotNil(b, res)
}

func BenchmarkQuaternion_NewQuaternionFull(b *testing.B) {
	var res mymath.Quaternion

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = mymath.NewQuaternionFull(1, 2, 3, 4)
	}

	assert.NotNil(b, res)
}

func BenchmarkQuaternion_Add(b *testing.B) {
	q := mymath.NewQuaternionFull(1, 2, 3, 4)

	res := q

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Add(q)
	}

	assert.NotNil(b, res)
}

func BenchmarkQuaternion_Subtract(b *testing.B) {
	q := mymath.NewQuaternionFull(1, 2, 3, 4)

	res := q

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Subtract(q)
	}

	assert.NotNil(b, res)
}

func BenchmarkQuaternion_Multiply(b *testing.B) {
	q := mymath.NewQuaternionFull(1, 2, 3, 4)

	res := q

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Multiply(1.0000003)
	}

	assert.NotNil(b, res)
}

func BenchmarkQuaternion_Divide(b *testing.B) {
	q := mymath.NewQuaternionFull(1, 2, 3, 4)

	res := q

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Divide(1.0000003)
	}

	assert.NotNil(b, res)
}

func BenchmarkQuaternion_Negate(b *testing.B) {
	q := mymath.NewQuaternionFull(1, 2, 3, 4)

	res := q

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Negate()
	}

	assert.NotNil(b, res)
}

func BenchmarkQuaternion_Normalize(b *testing.B) {
	q := mymath.NewQuaternionFull(1, 2, 3, 4)

	res := q

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Normalize()
	}

	assert.NotNil(b, res)
}

func BenchmarkQuaternion_Dot(b *testing.B) {
	q1 := mymath.NewQuaternionFull(1, 2, 3, 4)
	q2 := mymath.NewQuaternionFull(4, 5, 6, 7)

	var res float64

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = q1.Dot(q2)
	}

	assert.NotNil(b, res)
}
