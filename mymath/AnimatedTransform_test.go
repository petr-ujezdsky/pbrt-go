package mymath_test

import (
	"math"
	"pbrt-go/material"
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnimatedTransform_Decompose_translate(t *testing.T) {
	T := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	Tvec, Rquat, S, err := mymath.Decompose(T.M)

	assert.Nil(t, err)
	InDeltaVector3(t, mymath.NewVector3(1, 2, 3), Tvec)
	InDeltaQuaternion(t, mymath.NewQuaternionEmpty(), Rquat)
	InDeltaMatrix4x4(t, mymath.Identity(), S)
}

func TestAnimatedTransform_Decompose_rotate(t *testing.T) {
	// rotate 180 degrees along Z axis
	T := mymath.NewTransformRotateZ(math.Pi)

	Tvec, Rquat, S, err := mymath.Decompose(T.M)

	assert.Nil(t, err)
	InDeltaVector3(t, mymath.NewVector3(0, 0, 0), Tvec)
	InDeltaQuaternion(t, mymath.NewQuaternionFull(0, 0, 1, 0), Rquat)
	InDeltaMatrix4x4(t, mymath.Identity(), S)
}

func TestAnimatedTransform_Decompose_scale(t *testing.T) {
	T := mymath.NewTransformScale(2, 2, 2)

	Tvec, Rquat, S, err := mymath.Decompose(T.M)

	assert.Nil(t, err)
	InDeltaVector3(t, mymath.NewVector3(0, 0, 0), Tvec)
	InDeltaQuaternion(t, mymath.NewQuaternionEmpty(), Rquat)
	InDeltaMatrix4x4(t, mymath.NewMatrix4x4All(
		2, 0, 0, 0,
		0, 2, 0, 0,
		0, 0, 2, 0,
		0, 0, 0, 1), S)
}

func TestAnimatedTransform_Decompose(t *testing.T) {
	// T * R * S
	T := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3)).ApplyT(mymath.NewTransformRotateZ(math.Pi)).ApplyT(mymath.NewTransformScale(2, 2, 2))

	Tvec, Rquat, S, err := mymath.Decompose(T.M)

	assert.Nil(t, err)
	InDeltaVector3(t, mymath.NewVector3(1, 2, 3), Tvec)
	InDeltaQuaternion(t, mymath.NewQuaternionFull(0, 0, 1, 0), Rquat)
	InDeltaMatrix4x4(t, mymath.NewMatrix4x4All(
		2, 0, 0, 0,
		0, 2, 0, 0,
		0, 0, 2, 0,
		0, 0, 0, 1), S)
}

func TestAnimatedTransform_Interpolate_translate(t *testing.T) {
	t0 := mymath.NewTransformEmpty()
	t1 := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, err := mymath.NewAnimatedTransform(t0, 0, t1, 10)

	assert.Nil(t, err)

	// interpolate at 1/2 the time
	res, err := at.Interpolate(5)

	assert.Nil(t, err)

	expected := mymath.NewTransformTranslate(mymath.NewVector3(0.5, 1, 1.5))

	InDeltaMatrix4x4(t, expected.M, res.M)
}

func TestAnimatedTransform_Interpolate_rotate(t *testing.T) {
	t0 := mymath.NewTransformEmpty()
	// rotate 90 degrees along Z axis
	t1 := mymath.NewTransformRotateZ(math.Pi / 2)

	at, err := mymath.NewAnimatedTransform(t0, 0, t1, 10)

	assert.Nil(t, err)

	// interpolate at 1/2 the time
	res, err := at.Interpolate(5)

	assert.Nil(t, err)
	// rotation 45 degrees along Z axis
	expected := mymath.NewTransformRotateZ(math.Pi / 4)

	InDeltaMatrix4x4(t, expected.M, res.M)
}

func TestAnimatedTransform_Interpolate_scale(t *testing.T) {
	t0 := mymath.NewTransformEmpty()
	// rotate 90 degrees along Z axis
	t1 := mymath.NewTransformScale(3, 3, 3)

	at, err := mymath.NewAnimatedTransform(t0, 0, t1, 10)

	assert.Nil(t, err)

	// interpolate at 1/2 the time
	res, err := at.Interpolate(5)

	assert.Nil(t, err)

	expected := mymath.NewTransformScale(2, 2, 2)

	InDeltaMatrix4x4(t, expected.M, res.M)
}

func TestAnimatedTransform_Interpolate(t *testing.T) {
	t0 := mymath.NewTransformEmpty()
	// T * R * S
	t1 := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3)).ApplyT(mymath.NewTransformRotateZ(math.Pi / 2)).ApplyT(mymath.NewTransformScale(3, 3, 3))

	at, err := mymath.NewAnimatedTransform(t0, 0, t1, 10)

	assert.Nil(t, err)

	// interpolate at 1/2 the time
	res, err := at.Interpolate(5)

	assert.Nil(t, err)

	expected := mymath.NewTransformTranslate(mymath.NewVector3(0.5, 1, 1.5)).ApplyT(mymath.NewTransformRotateZ(math.Pi / 4)).ApplyT(mymath.NewTransformScale(2, 2, 2))

	InDeltaMatrix4x4(t, expected.M, res.M)
}

func TestAnimatedTransform_ApplyR(t *testing.T) {
	r := mymath.NewRay(mymath.NewPoint3(0, 0, 0), mymath.NewVector3(0, 0, 1), 9999, 5, material.Medium{})

	t0 := mymath.NewTransformEmpty()
	t1 := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, err := mymath.NewAnimatedTransform(t0, 0, t1, 10)

	assert.Nil(t, err)

	res, err := at.ApplyR(r)

	assert.Nil(t, err)

	expected := mymath.NewRay(mymath.NewPoint3(0.5, 1, 1.5), mymath.NewVector3(0, 0, 1), 9999, 5, material.Medium{})

	InDeltaRay(t, expected, res)
}

func TestAnimatedTransform_ApplyRD(t *testing.T) {
	r := mymath.NewRay(mymath.NewPoint3(0, 0, 0), mymath.NewVector3(0, 0, 1), 9999, 5, material.Medium{})
	rd := mymath.NewRayDifferentialRay(r)

	t0 := mymath.NewTransformEmpty()
	t1 := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, err := mymath.NewAnimatedTransform(t0, 0, t1, 10)

	assert.Nil(t, err)

	res, err := at.ApplyRD(rd)

	assert.Nil(t, err)

	expected := mymath.RayDifferential{
		mymath.NewRay(mymath.NewPoint3(0.5, 1, 1.5), mymath.NewVector3(0, 0, 1), 9999, 5, material.Medium{}),
		false,
		mymath.NewPoint3(0.5, 1, 1.5),
		mymath.NewPoint3(0.5, 1, 1.5),
		mymath.NewVector3(0, 0, 0),
		mymath.NewVector3(0, 0, 0)}

	InDeltaRayDifferential(t, expected, res)
}

func TestAnimatedTransform_ApplyP(t *testing.T) {
	p := mymath.NewPoint3(0, 0, 0)

	t0 := mymath.NewTransformEmpty()
	t1 := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, err := mymath.NewAnimatedTransform(t0, 0, t1, 10)

	assert.Nil(t, err)

	res, err := at.ApplyP(5, p)

	assert.Nil(t, err)

	expected := mymath.NewPoint3(0.5, 1, 1.5)

	InDeltaPoint3(t, expected, res)
}

func TestAnimatedTransform_ApplyV(t *testing.T) {
	v := mymath.NewVector3(0, 0, 0)

	t0 := mymath.NewTransformEmpty()
	t1 := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, err := mymath.NewAnimatedTransform(t0, 0, t1, 10)

	assert.Nil(t, err)

	res, err := at.ApplyV(5, v)

	assert.Nil(t, err)

	expected := mymath.NewVector3(0, 0, 0)

	InDeltaVector3(t, expected, res)
}

//////////////////////////////////////////////////////////////////////////////////////////////////
// BENCHMARKS ////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
func BenchmarkAnimatedTransform_NewAnimatedTransform(b *testing.B) {
	t0 := mymath.NewTransformEmpty()
	t1 := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	var res mymath.AnimatedTransform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _ = mymath.NewAnimatedTransform(t0, 0, t1, 10)
	}

	assert.NotNil(b, res)
}
func BenchmarkAnimatedTransform_Decompose(b *testing.B) {
	T := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	var res mymath.Vector3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _, _, _ = mymath.Decompose(T.M)
	}

	assert.NotNil(b, res)
}

func BenchmarkAnimatedTransform_Interpolate_animated(b *testing.B) {
	t0 := mymath.NewTransformEmpty()
	t1 := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, _ := mymath.NewAnimatedTransform(t0, 0, t1, 10)

	var res mymath.Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _ = at.Interpolate(5)
	}

	assert.NotNil(b, res)
}

func BenchmarkAnimatedTransform_Interpolate_static(b *testing.B) {
	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, _ := mymath.NewAnimatedTransform(t, 0, t, 10)

	var res mymath.Transform

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _ = at.Interpolate(5)
	}

	assert.NotNil(b, res)
}

func BenchmarkAnimatedTransform_ApplyR_animated(b *testing.B) {
	r := mymath.NewRay(mymath.NewPoint3(0, 0, 0), mymath.NewVector3(0, 0, 1), 9999, 5, material.Medium{})

	t0 := mymath.NewTransformEmpty()
	t1 := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, _ := mymath.NewAnimatedTransform(t0, 0, t1, 10)

	var res mymath.Ray

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _ = at.ApplyR(r)
	}

	assert.NotNil(b, res)
}

func BenchmarkAnimatedTransform_ApplyR_static(b *testing.B) {
	r := mymath.NewRay(mymath.NewPoint3(0, 0, 0), mymath.NewVector3(0, 0, 1), 9999, 5, material.Medium{})

	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, _ := mymath.NewAnimatedTransform(t, 0, t, 10)

	var res mymath.Ray

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _ = at.ApplyR(r)
	}

	assert.NotNil(b, res)
}

func BenchmarkAnimatedTransform_ApplyRD_animated(b *testing.B) {
	r := mymath.NewRay(mymath.NewPoint3(0, 0, 0), mymath.NewVector3(0, 0, 1), 9999, 5, material.Medium{})
	rd := mymath.NewRayDifferentialRay(r)

	t0 := mymath.NewTransformEmpty()
	t1 := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, _ := mymath.NewAnimatedTransform(t0, 0, t1, 10)

	var res mymath.RayDifferential

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _ = at.ApplyRD(rd)
	}

	assert.NotNil(b, res)
}

func BenchmarkAnimatedTransform_ApplyRD_static(b *testing.B) {
	r := mymath.NewRay(mymath.NewPoint3(0, 0, 0), mymath.NewVector3(0, 0, 1), 9999, 5, material.Medium{})
	rd := mymath.NewRayDifferentialRay(r)

	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, _ := mymath.NewAnimatedTransform(t, 0, t, 10)

	var res mymath.RayDifferential

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _ = at.ApplyRD(rd)
	}

	assert.NotNil(b, res)
}

func BenchmarkAnimatedTransform_ApplyP_animated(b *testing.B) {
	p := mymath.NewPoint3(0, 0, 0)

	t0 := mymath.NewTransformEmpty()
	t1 := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, _ := mymath.NewAnimatedTransform(t0, 0, t1, 10)

	var res mymath.Point3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _ = at.ApplyP(5.0, p)
	}

	assert.NotNil(b, res)
}

func BenchmarkAnimatedTransform_ApplyP_static(b *testing.B) {
	p := mymath.NewPoint3(0, 0, 0)

	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, _ := mymath.NewAnimatedTransform(t, 0, t, 10)

	var res mymath.Point3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _ = at.ApplyP(5.0, p)
	}

	assert.NotNil(b, res)
}

func BenchmarkAnimatedTransform_ApplyV_animated(b *testing.B) {
	v := mymath.NewVector3(0, 0, 0)

	t0 := mymath.NewTransformEmpty()
	t1 := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, _ := mymath.NewAnimatedTransform(t0, 0, t1, 10)

	var res mymath.Vector3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _ = at.ApplyV(5.0, v)
	}

	assert.NotNil(b, res)
}

func BenchmarkAnimatedTransform_ApplyV_static(b *testing.B) {
	v := mymath.NewVector3(0, 0, 0)

	t := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	at, _ := mymath.NewAnimatedTransform(t, 0, t, 10)

	var res mymath.Vector3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _ = at.ApplyV(5.0, v)
	}

	assert.NotNil(b, res)
}
