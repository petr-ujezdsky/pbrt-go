package mymath

import "math"

// see https://github.com/mmp/pbrt-v3/blob/master/src/core/quaternion.h#L49
type Quaternion struct {
	V Vector3
	W float64
}

func NewQuaternionEmpty() Quaternion {
	return Quaternion{Vector3{}, 1}
}

func NewQuaternion(v Vector3, w float64) Quaternion {
	return Quaternion{v, w}
}

func NewQuaternionFull(x, y, z, w float64) Quaternion {
	return Quaternion{NewVector3(x, y, z), w}
}

func (q Quaternion) Add(q2 Quaternion) Quaternion {
	return NewQuaternion(q.V.Add(q2.V), q.W+q2.W)
}

func (q Quaternion) Subtract(q2 Quaternion) Quaternion {
	return NewQuaternion(q.V.Subtract(q2.V), q.W-q2.W)
}

func (q Quaternion) Multiply(d float64) Quaternion {
	return NewQuaternion(q.V.Multiply(d), q.W*d)
}

func (q Quaternion) Divide(d float64) Quaternion {
	inv := 1 / d
	return q.Multiply(inv)
}

func (q Quaternion) Negate() Quaternion {
	return NewQuaternion(q.V.Negate(), -q.W)
}

func (q Quaternion) Normalize() Quaternion {
	lengthInv := 1 / math.Sqrt(q.Dot(q))
	return q.Multiply(lengthInv)
}

func (q1 Quaternion) Dot(q2 Quaternion) float64 {
	return q1.V.Dot(q2.V) + q1.W*q2.W
}

// see https://github.com/mmp/pbrt-v3/blob/master/src/core/quaternion.cpp#L41
func (q1 Quaternion) ToTransform() Transform {
	w, x, y, z := q1.W, q1.V.X, q1.V.Y, q1.V.Z

	m := NewMatrix4x4AllF64(
		1-2*y*y-2*z*z, 2*x*y+2*w*z, 2*x*z-2*w*y, 0,
		2*x*y-2*w*z, 1-2*x*x-2*z*z, 2*y*z+2*w*x, 0,
		2*x*z+2*w*y, 2*y*z-2*w*x, 1-2*x*x-2*y*y, 0,
		0, 0, 0, 1)

	return NewTransformFull(m.Transpose(), m)
}
