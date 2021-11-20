package mymath

import "math"

// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.h#L412
type AnimatedTransform struct {
	StartTransform, EndTransform Transform
	startTime, endTime           float64
	actuallyAnimated             bool
	T                            [2]Vector3
	R                            [2]Quaternion
	S                            [2]Matrix4x4
	HasRotation                  bool
}

// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.cpp#L396
func NewAnimatedTransform(startTransform Transform, startTime float64, endTransform Transform, endTime float64) (AnimatedTransform, error) {
	T0, R0, S0, err := Decompose(startTransform.M)

	if err != nil {
		return AnimatedTransform{}, err
	}

	T1, R1, S1, err := Decompose(endTransform.M)

	if err != nil {
		return AnimatedTransform{}, err
	}

	dot := R0.Dot(R1)
	// Flip _R[1]_ if needed to select shortest path
	if dot < 0 {
		R1 = R1.Negate()
	}

	return AnimatedTransform{
		StartTransform:   startTransform,
		EndTransform:     endTransform,
		startTime:        startTime,
		endTime:          endTime,
		actuallyAnimated: startTransform != endTransform,
		T:                [2]Vector3{T0, T1},
		R:                [2]Quaternion{R0, R1},
		S:                [2]Matrix4x4{S0, S1},
	}, nil

	// TODO
	// hasRotation := dot < 0.9995

	// // Compute terms of motion derivative function
	// if hasRotation {

	// }
}

// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.cpp#L1103
func Decompose(m Matrix4x4) (Vector3, Quaternion, Matrix4x4, error) {
	T := Vector3{}

	// Extract translation _T_ from transformation matrix
	T.X = float64(m.M[0][3])
	T.Y = float64(m.M[1][3])
	T.Z = float64(m.M[2][3])

	// Compute new transformation matrix _M_ without translation
	M := m
	for i := 0; i < 3; i++ {
		M.M[i][3] = 0.0
		M.M[3][i] = 0.0
	}
	M.M[3][3] = 1.0

	// Extract rotation _R_ from transformation matrix
	var norm float64
	count := 0
	R := M
	for ok := true; ok; ok = count < 100 && norm > 0.0001 {
		// Compute next matrix _Rnext_ in series
		Rnext := Matrix4x4{}
		Rit, err := R.Transpose().Inverse()

		if err != nil {
			return Vector3{}, Quaternion{}, Matrix4x4{}, err
		}

		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				Rnext.M[i][j] = 0.5 * (R.M[i][j] + Rit.M[i][j])
			}
		}

		// Compute norm of difference between _R_ and _Rnext_
		norm = 0
		for i := 0; i < 3; i++ {
			n := math.Abs(float64(R.M[i][0]-Rnext.M[i][0])) +
				math.Abs(float64(R.M[i][1]-Rnext.M[i][1])) +
				math.Abs(float64(R.M[i][2]-Rnext.M[i][2]))
			norm = math.Max(norm, n)
		}
		R = Rnext
		count++
	}
	// XXX TODO FIXME deal with flip...
	Rquat := NewQuaternionMatrix(R)

	// Compute scale _S_ using rotation and original matrix
	RInv, err := R.Inverse()

	if err != nil {
		return Vector3{}, Quaternion{}, Matrix4x4{}, err
	}

	S := RInv.Multiply(M)

	return T, Rquat, S, nil
}

// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.cpp#L1144
func (at AnimatedTransform) Interpolate(time float64) (Transform, error) {
	// Handle boundary conditions for matrix interpolation
	if !at.actuallyAnimated || time <= at.startTime {
		return at.StartTransform, nil
	}

	if time >= at.endTime {
		return at.EndTransform, nil
	}

	// 0 <= dt <= 1
	dt := (time - at.startTime) / (at.endTime - at.startTime)

	// Interpolate translation at _dt_
	trans := at.T[0].Multiply(1 - dt).Add(at.T[1].Multiply(dt))

	// Interpolate rotation at _dt_
	rotate := at.R[0].Slerp(dt, at.R[1])

	// Interpolate scale at _dt_
	scale := Matrix4x4{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			scale.M[i][j] = float32(Lerp(dt, float64(at.S[0].M[i][j]), float64(at.S[1].M[i][j])))
		}
	}
	scale.M[3][3] = 1

	S, err := NewTransform(scale)

	if err != nil {
		return NewTransformEmpty(), err
	}

	// Compute interpolated matrix as product of interpolated components
	return NewTransformTranslate(trans).ApplyT(rotate.ToTransform()).ApplyT(S), nil
}

func (at AnimatedTransform) ApplyR(r Ray) (Ray, error) {
	if !at.actuallyAnimated || float64(r.Time) <= at.startTime {
		return at.StartTransform.ApplyR(r), nil
	}

	if float64(r.Time) >= at.endTime {
		return at.EndTransform.ApplyR(r), nil
	}

	t, err := at.Interpolate(float64(r.Time))
	return t.ApplyR(r), err
}

func (at AnimatedTransform) ApplyRD(rd RayDifferential) (RayDifferential, error) {
	if !at.actuallyAnimated || float64(rd.Time) <= at.startTime {
		return at.StartTransform.ApplyRD(rd), nil
	}

	if float64(rd.Time) >= at.endTime {
		return at.EndTransform.ApplyRD(rd), nil
	}

	t, err := at.Interpolate(float64(rd.Time))
	return t.ApplyRD(rd), err
}
