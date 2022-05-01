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
	C1, C2, C3, C4, C5           [3]DerivativeTerm
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

// MotionBounds see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/transform.cpp#L1215
func (at AnimatedTransform) MotionBounds(b Bounds3) (Bounds3, error) {
	if !at.actuallyAnimated {
		return at.StartTransform.ApplyB(b), nil
	}

	if at.HasRotation {
		return at.StartTransform.ApplyB(b).UnionB(at.EndTransform.ApplyB(b)), nil
	}

	// Return motion bounds accounting for animated rotation
	bounds := Bounds3{}
	for corner := 0; corner < 8; corner++ {
		motionBound, err := at.boundPointMotion(b.Corner(corner))
		if err != nil {
			return Bounds3{}, err
		}

		bounds = bounds.UnionB(motionBound)
	}

	return bounds, nil
}

// BoundPointMotion see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/transform.cpp#L1226
func (at AnimatedTransform) boundPointMotion(p Point3) (Bounds3, error) {
	bounds := NewBounds3(at.StartTransform.ApplyP(p), at.EndTransform.ApplyP(p))
	cosTheta := at.R[0].Dot(at.R[1])
	theta := math.Acos(Clamp(cosTheta, -1, 1))

	for c := 0; c < 3; c++ {
		// Find any motion derivative zeros for the component c
		zeros := [4]float64{}
		nZeros := 0
		intervalFindZeros(at.C1[c].Eval(p), at.C2[c].Eval(p), at.C3[c].Eval(p), at.C4[c].Eval(p), at.C5[c].Eval(p), theta, NewInterval(0.0, 1.0), &zeros, &nZeros, 8)

		// Expand bounding box for any motion derivative zeros found
		for i := 0; i < nZeros; i++ {
			pz, err := at.ApplyP(Lerp(zeros[i], at.startTime, at.endTime), p)
			if err != nil {
				return Bounds3{}, err
			}

			bounds = bounds.UnionP(pz)
		}
	}

	return bounds, nil
}

// intervalFindZeros see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/transform.cpp#L354
func intervalFindZeros(c1, c2, c3, c4, c5, theta float64, tInterval Interval, zeros *[4]float64, zeroCount *int, depth int) {
	// Evaluate motion derivative in interval form, return if no zeros
	span := NewIntervalSingle(c1).Add(
		NewIntervalSingle(c2).Add(NewIntervalSingle(c3).Multiply(tInterval)).Multiply(Cos(NewIntervalSingle(2 * theta).Multiply(tInterval)))).Add(
		NewIntervalSingle(c4).Add(NewIntervalSingle(c5).Multiply(tInterval).Multiply(Sin(NewIntervalSingle(2 * theta).Multiply(tInterval)))))

	if span.Low > 0 || span.High < 0 || span.Low == span.High {
		return
	}

	if depth > 0 {
		// Split tInterval and check both resulting intervals
		mid := (tInterval.Low + tInterval.High) * 0.5
		intervalFindZeros(c1, c2, c3, c4, c5, theta, NewInterval(tInterval.Low, mid), zeros, zeroCount, depth-1)
		intervalFindZeros(c1, c2, c3, c4, c5, theta, NewInterval(mid, tInterval.High), zeros, zeroCount, depth-1)
	} else {
		// Use Newton’s method to refine zero
		tNewton := (tInterval.Low + tInterval.High) * 0.5
		for i := 0; i < 4; i++ {
			fNewton := c1 +
				(c2+c3*tNewton)*math.Cos(2*theta*tNewton) +
				(c4+c5*tNewton)*math.Sin(2*theta*tNewton)

			fPrimeNewton :=
				(c3+2*(c4+c5*tNewton)*theta)*
					math.Cos(2*tNewton*theta) +
					(c5-2*(c2+c3*tNewton)*theta)*
						math.Sin(2*tNewton*theta)

			if fNewton == 0 || fPrimeNewton == 0 {
				break
			}

			tNewton = tNewton - fNewton/fPrimeNewton
		}

		if tNewton >= tInterval.Low-1e-3 && tNewton < tInterval.High+1e-3 {
			zeros[*zeroCount] = tNewton
			*zeroCount++
		}
	}
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

// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.cpp#L1171
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

// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.cpp#L1183
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

// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.cpp#L1195
func (at AnimatedTransform) ApplyP(time float64, p Point3) (Point3, error) {
	if !at.actuallyAnimated || time <= at.startTime {
		return at.StartTransform.ApplyP(p), nil
	}

	if time >= at.endTime {
		return at.EndTransform.ApplyP(p), nil
	}

	t, err := at.Interpolate(time)
	return t.ApplyP(p), err
}

// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.cpp#L1205
func (at AnimatedTransform) ApplyV(time float64, v Vector3) (Vector3, error) {
	if !at.actuallyAnimated || time <= at.startTime {
		return at.StartTransform.ApplyV(v), nil
	}

	if time >= at.endTime {
		return at.EndTransform.ApplyV(v), nil
	}

	t, err := at.Interpolate(time)
	return t.ApplyV(v), err
}
