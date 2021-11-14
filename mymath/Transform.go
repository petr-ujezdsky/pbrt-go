package mymath

import "math"

type Transform struct {
	m, mInv Matrix4x4
}

func NewTransformEmpty() Transform {
	m := Identity()

	return Transform{m, m}
}

func NewTransform(m Matrix4x4) (Transform, error) {
	mInv, err := m.Inverse()
	if err != nil {
		return Transform{}, err
	}

	return Transform{m, mInv}, nil
}

func NewTransformFull(m Matrix4x4, mInv Matrix4x4) Transform {
	return Transform{m, mInv}
}

func NewTransformTranslate(delta Vector3) Transform {
	return Transform{
		NewMatrix4x4All(
			1, 0, 0, float32(delta.X),
			0, 1, 0, float32(delta.Y),
			0, 0, 1, float32(delta.Z),
			0, 0, 0, 1),
		NewMatrix4x4All(
			1, 0, 0, -float32(delta.X),
			0, 1, 0, -float32(delta.Y),
			0, 0, 1, -float32(delta.Z),
			0, 0, 0, 1),
	}
}

func NewTransformScale(x, y, z float32) Transform {
	return Transform{
		NewMatrix4x4All(
			x, 0, 0, 0,
			0, y, 0, 0,
			0, 0, z, 0,
			0, 0, 0, 1),
		NewMatrix4x4All(
			1/x, 0, 0, 0,
			0, 1/y, 0, 0,
			0, 0, 1/z, 0,
			0, 0, 0, 1),
	}
}

func NewTransformRotateX(theta float32) Transform {
	sinTheta := float32(math.Sin(float64(theta)))
	cosTheta := float32(math.Cos(float64(theta)))

	m := NewMatrix4x4All(
		1, 0, 0, 0,
		0, cosTheta, -sinTheta, 0,
		0, sinTheta, cosTheta, 0,
		0, 0, 0, 1)

	return Transform{m, m.Transpose()}
}

func NewTransformRotateY(theta float32) Transform {
	sinTheta := float32(math.Sin(float64(theta)))
	cosTheta := float32(math.Cos(float64(theta)))

	m := NewMatrix4x4All(
		cosTheta, 0, sinTheta, 0,
		0, 1, 0, 0,
		-sinTheta, 0, cosTheta, 0,
		0, 0, 0, 1)

	return Transform{m, m.Transpose()}
}

func NewTransformRotateZ(theta float32) Transform {
	sinTheta := float32(math.Sin(float64(theta)))
	cosTheta := float32(math.Cos(float64(theta)))

	m := NewMatrix4x4All(
		cosTheta, -sinTheta, 0, 0,
		sinTheta, cosTheta, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1)

	return Transform{m, m.Transpose()}
}

func NewTransformRotate(theta float64, axis Vector3) Transform {
	a := axis.Normalize()

	sinTheta := math.Sin(theta)
	cosTheta := math.Cos(theta)

	m := Matrix4x4{}

	// Compute rotation of first basis vector
	m.M[0][0] = float32(a.X*a.X + (1-a.X*a.X)*cosTheta)
	m.M[0][1] = float32(a.X*a.Y*(1-cosTheta) - a.Z*sinTheta)
	m.M[0][2] = float32(a.X*a.Z*(1-cosTheta) + a.Y*sinTheta)
	m.M[0][3] = 0

	// Compute rotations of second and third basis vectors
	m.M[1][0] = float32(a.X*a.Y*(1-cosTheta) + a.Z*sinTheta)
	m.M[1][1] = float32(a.Y*a.Y + (1-a.Y*a.Y)*cosTheta)
	m.M[1][2] = float32(a.Y*a.Z*(1-cosTheta) - a.X*sinTheta)
	m.M[1][3] = 0

	m.M[2][0] = float32(a.X*a.Z*(1-cosTheta) - a.Y*sinTheta)
	m.M[2][1] = float32(a.Y*a.Z*(1-cosTheta) + a.X*sinTheta)
	m.M[2][2] = float32(a.Z*a.Z + (1-a.Z*a.Z)*cosTheta)
	m.M[2][3] = 0

	return Transform{m, m.Transpose()}
}

func (t Transform) ApplyP(p Point3) Point3 {
	return t.m.MultiplyP(p)
}

func (t Transform) ApplyV(v Vector3) Vector3 {
	return t.m.MultiplyV(v)
}

func (t Transform) Inverse() Transform {
	return Transform{t.mInv, t.m}
}

func (t Transform) Transpose() Transform {
	return Transform{t.m.Transpose(), t.mInv.Transpose()}
}

func (t Transform) IsIdentity() bool {
	return t.m.IsIdentity()
}

func (t Transform) HasScale() bool {
	la2 := t.ApplyV(Vector3{1, 0, 0}).LengthSq()
	lb2 := t.ApplyV(Vector3{0, 1, 0}).LengthSq()
	lc2 := t.ApplyV(Vector3{0, 0, 1}).LengthSq()

	notOne := func(x float64) bool {
		return x < 0.999 || x > 1.001
	}

	return notOne(la2) || notOne(lb2) || notOne(lc2)
}
