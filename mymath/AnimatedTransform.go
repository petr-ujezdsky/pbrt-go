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
