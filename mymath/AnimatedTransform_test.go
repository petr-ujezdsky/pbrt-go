package mymath_test

import (
	"math"
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnimatedTransform_Decompose_translate(t *testing.T) {
	T := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	Tvec, Rquat, S, err := mymath.Decompose(T.M)

	assert.Nil(t, err)
	assertAlmostEqualVector3(t, mymath.NewVector3(1, 2, 3), Tvec)
	assertAlmostEqual(t, mymath.NewQuaternionEmpty(), Rquat)
	assertAlmostEqualMatrix4x4(t, mymath.Identity(), S)
}

func TestAnimatedTransform_Decompose_rotate(t *testing.T) {
	// rotate 180 degrees along Z axis
	T := mymath.NewTransformRotateZ(math.Pi)

	Tvec, Rquat, S, err := mymath.Decompose(T.M)

	assert.Nil(t, err)
	assertAlmostEqualVector3(t, mymath.NewVector3(0, 0, 0), Tvec)
	assertAlmostEqual(t, mymath.NewQuaternionFull(0, 0, 1, 0), Rquat)
	assertAlmostEqualMatrix4x4(t, mymath.Identity(), S)
}

func TestAnimatedTransform_Decompose_scale(t *testing.T) {
	T := mymath.NewTransformScale(2, 2, 2)

	Tvec, Rquat, S, err := mymath.Decompose(T.M)

	assert.Nil(t, err)
	assertAlmostEqualVector3(t, mymath.NewVector3(0, 0, 0), Tvec)
	assertAlmostEqual(t, mymath.NewQuaternionEmpty(), Rquat)
	assertAlmostEqualMatrix4x4(t, mymath.NewMatrix4x4All(
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
	assertAlmostEqualVector3(t, mymath.NewVector3(1, 2, 3), Tvec)
	assertAlmostEqual(t, mymath.NewQuaternionFull(0, 0, 1, 0), Rquat)
	assertAlmostEqualMatrix4x4(t, mymath.NewMatrix4x4All(
		2, 0, 0, 0,
		0, 2, 0, 0,
		0, 0, 2, 0,
		0, 0, 0, 1), S)
}

//////////////////////////////////////////////////////////////////////////////////////////////////
// BENCHMARKS ////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
func BenchmarkAnimatedTransform_Decompose(b *testing.B) {
	T := mymath.NewTransformTranslate(mymath.NewVector3(1, 2, 3))

	var res mymath.Vector3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _, _, _ = mymath.Decompose(T.M)
	}

	assert.NotNil(b, res)
}
