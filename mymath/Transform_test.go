package mymath

import (
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

	tr := NewTransform(&m)

	mInv := *m.Inverse()

	assert.Equal(t, m, tr.m)
	assert.Equal(t, mInv, tr.mInv)
}

func TestTransform_NewTransformFull(t *testing.T) {
	m := NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	mInv := *m.Inverse()

	tr := NewTransformFull(&m, &mInv)

	assert.Equal(t, m, tr.m)
	assert.Equal(t, mInv, tr.mInv)
}

func TestTransform_NewTransformTranslate(t *testing.T) {
	delta := NewVector3(5, 6, 7)
	tr := NewTransformTranslate(&delta)

	p := NewPoint3(1, 2, 3)
	res := tr.Apply(&p)

	assert.Equal(t, NewPoint3(6, 8, 10), *res)
}

func TestTransform_NewTransformScale(t *testing.T) {
	tr := NewTransformScale(2, 3, 4)

	p := NewPoint3(1, 2, 3)
	res := tr.Apply(&p)

	assert.Equal(t, NewPoint3(2, 6, 12), *res)
}

func TestTransform_Inverse(t *testing.T) {
	m := NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	var tr = NewTransform(&m)

	tr = tr.Inverse()

	mInv := *m.Inverse()

	assert.Equal(t, mInv, tr.m)
	assert.Equal(t, m, tr.mInv)
}

func TestTransform_Transpose(t *testing.T) {
	m := NewMatrix4x4All(
		5, 2, 8, 3,
		7, 3, 5, 3,
		9, 3, 2, 4,
		1, 8, 3, 8)

	var tr = NewTransform(&m)

	tr = tr.Transpose()

	mInv := *m.Inverse()

	assert.Equal(t, *m.Transpose(), tr.m)
	assert.Equal(t, *mInv.Transpose(), tr.mInv)
}

func TestTransform_IsIdentity(t *testing.T) {
	tr := NewTransformEmpty()

	assert.True(t, tr.IsIdentity())
}
