package mymath_test

import (
	"math"
	"pbrt-go/material"
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransform_NewTransformEmpty(t *testing.T) {
	tr := mymath.NewTransformEmpty()

	assert.Equal(t, mymath.Identity(), tr.M)
	assert.Equal(t, mymath.Identity(), tr.MInv)
}

func TestTransform_NewTransform(t *testing.T) {
	m := mymath.NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	tr, err := mymath.NewTransform(m)
	assert.Nil(t, err)

	mInv, err := m.Inverse()

	assert.Nil(t, err)
	assert.Equal(t, m, tr.M)
	assert.Equal(t, mInv, tr.MInv)
}

func TestTransform_NewTransformFull(t *testing.T) {
	m := mymath.NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	mInv, err := m.Inverse()

	tr := mymath.NewTransformFull(m, mInv)

	assert.Nil(t, err)
	assert.Equal(t, m, tr.M)
	assert.Equal(t, mInv, tr.MInv)
}

func TestTransform_NewTransformTranslate(t *testing.T) {
	delta := mymath.NewVector3(5, 6, 7)
	tr := mymath.NewTransformTranslate(delta)

	p := mymath.NewPoint3(1, 2, 3)
	res := tr.ApplyP(p)

	assert.Equal(t, mymath.NewPoint3(6, 8, 10), res)
}

func TestTransform_NewTransformScale(t *testing.T) {
	tr := mymath.NewTransformScale(2, 3, 4)

	p := mymath.NewPoint3(1, 2, 3)
	res := tr.ApplyP(p)

	assert.Equal(t, mymath.NewPoint3(2, 6, 12), res)
}

func TestTransform_RotateX(t *testing.T) {
	// rotate 90 degrees along X axis
	tr := mymath.NewTransformRotateX(math.Pi / 2)

	p := mymath.NewPoint3(10, 2, 3)
	res := tr.ApplyP(p)

	// assert.Equal(t, mymath.NewPoint3(10, -3, 2), res)
	assert.Equal(t, mymath.NewPoint3(10, -3.0000000874227766, 1.9999998688658351), res)
}

func TestTransform_RotateY(t *testing.T) {
	// rotate 90 degrees along Y axis
	tr := mymath.NewTransformRotateY(math.Pi / 2)

	p := mymath.NewPoint3(1, 20, 3)
	res := tr.ApplyP(p)

	// assert.Equal(t, mymath.NewPoint3(3, 20, -1), res)
	assert.Equal(t, mymath.NewPoint3(2.9999999562886117, 20, -1.0000001311341649), res)
}

func TestTransform_RotateZ(t *testing.T) {
	// rotate 90 degrees along Z axis
	tr := mymath.NewTransformRotateZ(math.Pi / 2)

	p := mymath.NewPoint3(1, 2, 30)
	res := tr.ApplyP(p)

	// assert.Equal(t, mymath.NewPoint3(-2, 1, 30), res)
	assert.Equal(t, mymath.NewPoint3(-2.0000000437113883, 0.9999999125772234, 30), res)
}

func TestTransform_Rotate(t *testing.T) {
	// X axis
	var axis = mymath.NewVector3(1, 0, 0)
	var tr = mymath.NewTransformRotate(math.Pi/2, axis)

	var p = mymath.NewPoint3(10, 2, 3)
	var res = tr.ApplyP(p)

	assert.Equal(t, mymath.NewPoint3(10, -3, 2), res)

	// Y axis
	axis = mymath.NewVector3(0, 1, 0)
	tr = mymath.NewTransformRotate(math.Pi/2, axis)

	p = mymath.NewPoint3(1, 20, 3)
	res = tr.ApplyP(p)

	// assert.Equal(t, mymath.NewPoint3(3, 20, -1), res)
	assert.Equal(t, mymath.NewPoint3(3, 20, -0.9999999999999998), res)

	// Z axis
	axis = mymath.NewVector3(0, 0, 1)
	tr = mymath.NewTransformRotate(math.Pi/2, axis)

	p = mymath.NewPoint3(1, 2, 30)
	res = tr.ApplyP(p)

	// assert.Equal(t, mymath.NewPoint3(-2, 1, 30), res)
	assert.Equal(t, mymath.NewPoint3(-2, 1.0000000000000002, 30), res)
}

func TestTransform_NewTransformLookAt(t *testing.T) {
	pos := mymath.NewPoint3(0, 0, 0)
	look := mymath.NewPoint3(0, 0, 10)
	up := mymath.NewVector3(10, 0, 0)

	tr, err := mymath.NewTransformLookAt(pos, look, up)

	assert.Nil(t, err)

	p := mymath.NewPoint3(1, 2, 30)
	res := tr.ApplyP(p)

	assert.Equal(t, mymath.NewPoint3(-2, 1, 30), res)
}

func TestTransform_ApplyB(t *testing.T) {
	tr := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	bounds := mymath.NewBounds3(
		mymath.NewPoint3(0, 0, 0),
		mymath.NewPoint3(1, 1, 1))

	res := tr.ApplyB(bounds)

	expected := mymath.NewBounds3(
		mymath.NewPoint3(1, 2, 3),
		mymath.NewPoint3(2, 3, 4))

	assert.Equal(t, expected, res)
}

func TestTransform_ApplyT(t *testing.T) {
	tr := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	res := tr.ApplyT(tr)

	p := mymath.NewPoint3(0, 0, 0)
	pT := res.ApplyP(p)
	assert.Equal(t, mymath.NewPoint3(2, 4, 6), pT)

	assert.Equal(t, mymath.NewPoint3(2, 4, 6), tr.ApplyP(tr.ApplyP(p)))
}

func TestTransform_Inverse(t *testing.T) {
	m := mymath.NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	tr, err := mymath.NewTransform(m)
	assert.Nil(t, err)

	tr = tr.Inverse()

	mInv, err := m.Inverse()

	assert.Nil(t, err)
	assert.Equal(t, mInv, tr.M)
	assert.Equal(t, m, tr.MInv)
}

func TestTransform_Transpose(t *testing.T) {
	m := mymath.NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	tr, err := mymath.NewTransform(m)
	assert.Nil(t, err)

	tr = tr.Transpose()

	mInv, err := m.Inverse()

	assert.Nil(t, err)
	assert.Equal(t, m.Transpose(), tr.M)
	assert.Equal(t, mInv.Transpose(), tr.MInv)
}

func TestTransform_IsIdentity(t *testing.T) {
	tr := mymath.NewTransformEmpty()

	assert.True(t, tr.IsIdentity())
}

func TestTransform_HasScale(t *testing.T) {
	var tr = mymath.NewTransformEmpty()
	assert.False(t, tr.HasScale())

	tr = mymath.NewTransformTranslate(mymath.Vector3{1, 2, 3})
	assert.False(t, tr.HasScale())

	tr = mymath.NewTransformScale(1, 1, 1)
	assert.False(t, tr.HasScale())

	tr = mymath.NewTransformScale(2, 2, 2)
	assert.True(t, tr.HasScale())
}

//////////////////////////////////////////////////////////////////////////////////////////////////
// BENCHMARKS ////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
func BenchmarkTransform_NewTransformEmpty(b *testing.B) {
	var res mymath.Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = mymath.NewTransformEmpty()
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_NewTransform(b *testing.B) {
	m := mymath.NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	var res mymath.Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _ = mymath.NewTransform(m)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransformNewTransformFull(b *testing.B) {
	m := mymath.NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	mInv, _ := m.Inverse()

	var res mymath.Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = mymath.NewTransformFull(m, mInv)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransformNewTransformTranslate(b *testing.B) {
	delta := mymath.NewVector3(1, 2, 3)

	var res mymath.Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = mymath.NewTransformTranslate(delta)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransformNewTransformScale(b *testing.B) {
	var res mymath.Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = mymath.NewTransformScale(1, 2, 3)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransformNewTransformRotateX(b *testing.B) {
	theta := float32(math.Pi / 2)
	var res mymath.Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = mymath.NewTransformRotateX(theta)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransformNewTransformRotateY(b *testing.B) {
	theta := float32(math.Pi / 2)
	var res mymath.Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = mymath.NewTransformRotateY(theta)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransformNewTransformRotateZ(b *testing.B) {
	theta := float32(math.Pi / 2)
	var res mymath.Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = mymath.NewTransformRotateZ(theta)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransformNewTransformRotate(b *testing.B) {
	axis := mymath.NewVector3(1, 2, 3)
	theta := math.Pi / 2
	var res mymath.Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = mymath.NewTransformRotate(theta, axis)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransformNewTransformLookAt(b *testing.B) {
	pos := mymath.NewPoint3(0, 0, 0)
	look := mymath.NewPoint3(0, 0, 10)
	up := mymath.NewVector3(10, 0, 0)

	var tr mymath.Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tr, _ = mymath.NewTransformLookAt(pos, look, up)
	}

	assert.NotNil(b, tr)
}

func BenchmarkTransform_ApplyP(b *testing.B) {
	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))
	p := mymath.NewPoint3(1, 2, 3)

	var res mymath.Point3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.ApplyP(p)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_ApplyV(b *testing.B) {
	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))
	v := mymath.NewVector3(1, 2, 3)

	var res mymath.Vector3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.ApplyV(v)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_ApplyN(b *testing.B) {
	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))
	n := mymath.NewNormal3(0, 1, 0)

	var res mymath.Normal3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.ApplyN(n)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_ApplyR(b *testing.B) {
	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))
	r := mymath.NewRay(
		mymath.NewPoint3(0, 0, 0),
		mymath.NewVector3(0, 0, 1),
		99,
		0,
		material.Medium{})

	var res mymath.Ray

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.ApplyR(r)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_ApplyRD(b *testing.B) {
	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))
	r := mymath.NewRay(
		mymath.NewPoint3(0, 0, 0),
		mymath.NewVector3(0, 0, 1),
		99,
		0,
		material.Medium{})

	rd := mymath.NewRayDifferentialRay(r)

	var res mymath.RayDifferential

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.ApplyRD(rd)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_ApplyB(b *testing.B) {
	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))
	bounds := mymath.NewBounds3(
		mymath.NewPoint3(0, 0, 0),
		mymath.NewPoint3(1, 1, 1))

	var res mymath.Bounds3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.ApplyB(bounds)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_ApplyT(b *testing.B) {
	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	var res mymath.Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.ApplyT(t)
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_Inverse(b *testing.B) {
	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	res := t

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Inverse()
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_Transpose(b *testing.B) {
	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	res := t

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Transpose()
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_IsIdentity(b *testing.B) {
	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	var res bool

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.IsIdentity()
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_HasScale(b *testing.B) {
	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	var res bool

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.HasScale()
	}

	assert.NotNil(b, res)
}

func BenchmarkTransform_SwapsHandedness(b *testing.B) {
	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	var res bool

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = t.SwapsHandedness()
	}

	assert.NotNil(b, res)
}
