package mymath

import (
	"math"
	"pbrt-go/material"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransform_NewTransformEmpty(t *testing.T) {
	tr := NewTransformEmpty()

	assert.Equal(t, Identity(), tr.m)
	assert.Equal(t, Identity(), tr.mInv)
}

func TestTransform_NewTransform(t *testing.T) {
	m := NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	tr, err := NewTransform(m)
	assert.Nil(t, err)

	mInv, err := m.Inverse()

	assert.Nil(t, err)
	assert.Equal(t, m, tr.m)
	assert.Equal(t, mInv, tr.mInv)
}

func TestTransform_NewTransformFull(t *testing.T) {
	m := NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	mInv, err := m.Inverse()

	tr := NewTransformFull(m, mInv)

	assert.Nil(t, err)
	assert.Equal(t, m, tr.m)
	assert.Equal(t, mInv, tr.mInv)
}

func TestTransform_NewTransformTranslate(t *testing.T) {
	delta := NewVector3(5, 6, 7)
	tr := NewTransformTranslate(delta)

	p := NewPoint3(1, 2, 3)
	res := tr.ApplyP(p)

	assert.Equal(t, NewPoint3(6, 8, 10), res)
}

func TestTransform_NewTransformScale(t *testing.T) {
	tr := NewTransformScale(2, 3, 4)

	p := NewPoint3(1, 2, 3)
	res := tr.ApplyP(p)

	assert.Equal(t, NewPoint3(2, 6, 12), res)
}

func TestTransform_RotateX(t *testing.T) {
	// rotate 90 degrees along X axis
	tr := NewTransformRotateX(math.Pi / 2)

	p := NewPoint3(10, 2, 3)
	res := tr.ApplyP(p)

	// assert.Equal(t, NewPoint3(10, -3, 2), res)
	assert.Equal(t, NewPoint3(10, -3.0000000874227766, 1.9999998688658351), res)
}

func TestTransform_RotateY(t *testing.T) {
	// rotate 90 degrees along Y axis
	tr := NewTransformRotateY(math.Pi / 2)

	p := NewPoint3(1, 20, 3)
	res := tr.ApplyP(p)

	// assert.Equal(t, NewPoint3(3, 20, -1), res)
	assert.Equal(t, NewPoint3(2.9999999562886117, 20, -1.0000001311341649), res)
}

func TestTransform_RotateZ(t *testing.T) {
	// rotate 90 degrees along Z axis
	tr := NewTransformRotateZ(math.Pi / 2)

	p := NewPoint3(1, 2, 30)
	res := tr.ApplyP(p)

	// assert.Equal(t, NewPoint3(-2, 1, 30), res)
	assert.Equal(t, NewPoint3(-2.0000000437113883, 0.9999999125772234, 30), res)
}

func TestTransform_Rotate(t *testing.T) {
	// X axis
	var axis = NewVector3(1, 0, 0)
	var tr = NewTransformRotate(math.Pi/2, axis)

	var p = NewPoint3(10, 2, 3)
	var res = tr.ApplyP(p)

	assert.Equal(t, NewPoint3(10, -3, 2), res)

	// Y axis
	axis = NewVector3(0, 1, 0)
	tr = NewTransformRotate(math.Pi/2, axis)

	p = NewPoint3(1, 20, 3)
	res = tr.ApplyP(p)

	// assert.Equal(t, NewPoint3(3, 20, -1), res)
	assert.Equal(t, NewPoint3(3, 20, -0.9999999999999998), res)

	// Z axis
	axis = NewVector3(0, 0, 1)
	tr = NewTransformRotate(math.Pi/2, axis)

	p = NewPoint3(1, 2, 30)
	res = tr.ApplyP(p)

	// assert.Equal(t, NewPoint3(-2, 1, 30), res)
	assert.Equal(t, NewPoint3(-2, 1.0000000000000002, 30), res)
}

func TestTransform_NewTransformLookAt(t *testing.T) {
	pos := NewPoint3(0, 0, 0)
	look := NewPoint3(0, 0, 10)
	up := NewVector3(10, 0, 0)

	tr, err := NewTransformLookAt(pos, look, up)

	assert.Nil(t, err)

	p := NewPoint3(1, 2, 30)
	res := tr.ApplyP(p)

	assert.Equal(t, NewPoint3(-2, 1, 30), res)
}

func TestTransform_Inverse(t *testing.T) {
	m := NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	tr, err := NewTransform(m)
	assert.Nil(t, err)

	tr = tr.Inverse()

	mInv, err := m.Inverse()

	assert.Nil(t, err)
	assert.Equal(t, mInv, tr.m)
	assert.Equal(t, m, tr.mInv)
}

func TestTransform_Transpose(t *testing.T) {
	m := NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	tr, err := NewTransform(m)
	assert.Nil(t, err)

	tr = tr.Transpose()

	mInv, err := m.Inverse()

	assert.Nil(t, err)
	assert.Equal(t, m.Transpose(), tr.m)
	assert.Equal(t, mInv.Transpose(), tr.mInv)
}

func TestTransform_IsIdentity(t *testing.T) {
	tr := NewTransformEmpty()

	assert.True(t, tr.IsIdentity())
}

func TestTransform_HasScale(t *testing.T) {
	var tr = NewTransformEmpty()
	assert.False(t, tr.HasScale())

	tr = NewTransformTranslate(Vector3{1, 2, 3})
	assert.False(t, tr.HasScale())

	tr = NewTransformScale(1, 1, 1)
	assert.False(t, tr.HasScale())

	tr = NewTransformScale(2, 2, 2)
	assert.True(t, tr.HasScale())
}

//////////////////////////////////////////////////////////////////////////////////////////////////
// BENCHMARKS ////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
func BenchmarkTransform_NewTransformEmpty(b *testing.B) {
	var res Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = NewTransformEmpty()
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_NewTransform(b *testing.B) {
	m := NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	var res Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _ = NewTransform(m)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_NewTransformFull(b *testing.B) {
	m := NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	mInv, _ := m.Inverse()

	var res Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = NewTransformFull(m, mInv)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_NewTransformTranslate(b *testing.B) {
	delta := NewVector3(1, 2, 3)

	var res Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = NewTransformTranslate(delta)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_NewTransformScale(b *testing.B) {
	var res Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = NewTransformScale(1, 2, 3)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_NewTransformRotateX(b *testing.B) {
	theta := float32(math.Pi / 2)
	var res Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = NewTransformRotateX(theta)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_NewTransformRotateY(b *testing.B) {
	theta := float32(math.Pi / 2)
	var res Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = NewTransformRotateY(theta)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_NewTransformRotateZ(b *testing.B) {
	theta := float32(math.Pi / 2)
	var res Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = NewTransformRotateZ(theta)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_NewTransformRotate(b *testing.B) {
	axis := NewVector3(1, 2, 3)
	theta := math.Pi / 2
	var res Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = NewTransformRotate(theta, axis)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_NewTransformLookAt(b *testing.B) {
	pos := NewPoint3(0, 0, 0)
	look := NewPoint3(0, 0, 10)
	up := NewVector3(10, 0, 0)

	var tr Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tr, _ = NewTransformLookAt(pos, look, up)
	}

	assert.NotNil(b, tr)
}

func BenchmarkTransform_ApplyP(b *testing.B) {
	t := NewTransformTranslate(NewVector3(1, 2, 3))
	p := NewPoint3(1, 2, 3)

	var res Point3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.ApplyP(p)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_ApplyV(b *testing.B) {
	t := NewTransformTranslate(NewVector3(1, 2, 3))
	v := NewVector3(1, 2, 3)

	var res Vector3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.ApplyV(v)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_ApplyN(b *testing.B) {
	t := NewTransformTranslate(NewVector3(1, 2, 3))
	n := NewNormal3(0, 1, 0)

	var res Vector3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.ApplyN(n)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_ApplyR(b *testing.B) {
	t := NewTransformTranslate(NewVector3(1, 2, 3))
	r := NewRay(
		NewPoint3(0, 0, 0),
		NewVector3(0, 0, 1),
		99,
		0,
		material.Medium{})

	var res Ray

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.ApplyR(r)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_ApplyRD(b *testing.B) {
	t := NewTransformTranslate(NewVector3(1, 2, 3))
	r := NewRay(
		NewPoint3(0, 0, 0),
		NewVector3(0, 0, 1),
		99,
		0,
		material.Medium{})

	rd := NewRayDifferentialRay(r)

	var res RayDifferential

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.ApplyRD(rd)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_Inverse(b *testing.B) {
	t := NewTransformTranslate(NewVector3(1, 2, 3))

	res := t

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Inverse()
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_Transpose(b *testing.B) {
	t := NewTransformTranslate(NewVector3(1, 2, 3))

	res := t

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Transpose()
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_IsIdentity(b *testing.B) {
	t := NewTransformTranslate(NewVector3(1, 2, 3))

	var res bool

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.IsIdentity()
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_HasScale(b *testing.B) {
	t := NewTransformTranslate(NewVector3(1, 2, 3))

	var res bool

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.HasScale()
	}

	assert.NotNil(b, res)
}
