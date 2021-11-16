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

// see https://github.com/mmp/pbrt-v3/blob/master/src/core/quaternion.cpp#L61
func NewQuaternionTransform(t Transform) Quaternion {
	quat := NewQuaternionEmpty()
	trace := float64(t.M.M[0][0] + t.M.M[1][1] + t.M.M[2][2])
	if trace > 0 {
		// Compute w from matrix trace, then xyz
		// 4w^2 = m[0][0] + m[1][1] + m[2][2] + m[3][3] (but m[3][3] == 1)
		s := math.Sqrt(trace + 1.0)
		quat.W = s / 2.0
		s = 0.5 / s
		quat.V.X = float64(t.M.M[2][1]-t.M.M[1][2]) * s
		quat.V.Y = float64(t.M.M[0][2]-t.M.M[2][0]) * s
		quat.V.Z = float64(t.M.M[1][0]-t.M.M[0][1]) * s
	} else {
		// Compute largest of $x$, $y$, or $z$, then remaining components
		nxt := [3]int{1, 2, 0}
		q := [3]float64{}
		i := 0
		if t.M.M[1][1] > t.M.M[0][0] {
			i = 1
		}
		if t.M.M[2][2] > t.M.M[i][i] {
			i = 2
		}
		j := nxt[i]
		k := nxt[j]
		s := math.Sqrt(float64(t.M.M[i][i]-(t.M.M[j][j]+t.M.M[k][k])) + 1.0)
		q[i] = s * 0.5
		if s != 0.0 {
			s = 0.5 / s
		}
		quat.W = float64(t.M.M[k][j]-t.M.M[j][k]) * s
		q[j] = float64(t.M.M[j][i]+t.M.M[i][j]) * s
		q[k] = float64(t.M.M[k][i]+t.M.M[i][k]) * s
		quat.V.X = q[0]
		quat.V.Y = q[1]
		quat.V.Z = q[2]
	}

	return quat
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

// see https://github.com/mmp/pbrt-v3/blob/master/src/core/quaternion.cpp#L94
func (q1 Quaternion) Slerp(t float64, q2 Quaternion) Quaternion {
	cosTheta := q1.Dot(q2)
	if cosTheta > 0.9995 {
		return q1.Multiply(1 - t).Add(q2.Multiply(t)).Normalize()
	} else {
		theta := math.Acos(Clamp(cosTheta, -1, 1))
		thetap := theta * t
		qperp := q2.Subtract(q1.Multiply(cosTheta)).Normalize()
		return q1.Multiply(math.Cos(thetap)).Add(qperp.Multiply(math.Sin(thetap)))
	}
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
