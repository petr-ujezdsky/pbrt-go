package mymath

import (
	"errors"
	"math"
)

type Matrix4x4 struct {
	// arrays of rows of columns
	M [4][4]float32
}

func NewMatrix4x4All(
	t00, t01, t02, t03,
	t10, t11, t12, t13,
	t20, t21, t22, t23,
	t30, t31, t32, t33 float32) Matrix4x4 {

	return Matrix4x4{[4][4]float32{
		{t00, t01, t02, t03},
		{t10, t11, t12, t13},
		{t20, t21, t22, t23},
		{t30, t31, t32, t33}}}
}

func Identity() Matrix4x4 {
	return NewMatrix4x4All(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1)
}

func (m Matrix4x4) Transpose() Matrix4x4 {
	M := m.M
	return NewMatrix4x4All(
		M[0][0], M[1][0], M[2][0], M[3][0],
		M[0][1], M[1][1], M[2][1], M[3][1],
		M[0][2], M[1][2], M[2][2], M[3][2],
		M[0][3], M[1][3], M[2][3], M[3][3],
	)
}

func (m1 Matrix4x4) Multiply(m2 Matrix4x4) Matrix4x4 {
	r := Matrix4x4{}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			r.M[i][j] = m1.M[i][0]*m2.M[0][j] +
				m1.M[i][1]*m2.M[1][j] +
				m1.M[i][2]*m2.M[2][j] +
				m1.M[i][3]*m2.M[3][j]
		}
	}

	return r
}

func (m Matrix4x4) MultiplyP(p Point3) Point3 {
	xp := float64(m.M[0][0])*p.X + float64(m.M[0][1])*p.Y + float64(m.M[0][2])*p.Z + float64(m.M[0][3])
	yp := float64(m.M[1][0])*p.X + float64(m.M[1][1])*p.Y + float64(m.M[1][2])*p.Z + float64(m.M[1][3])
	zp := float64(m.M[2][0])*p.X + float64(m.M[2][1])*p.Y + float64(m.M[2][2])*p.Z + float64(m.M[2][3])
	wp := float64(m.M[3][0])*p.X + float64(m.M[3][1])*p.Y + float64(m.M[3][2])*p.Z + float64(m.M[3][3])

	if wp == 1 {
		return NewPoint3(xp, yp, zp)
	}
	return NewPoint3(xp/wp, yp/wp, zp/wp)
}

func (m Matrix4x4) MultiplyV(v Vector3) Vector3 {
	xv := float64(m.M[0][0])*v.X + float64(m.M[0][1])*v.Y + float64(m.M[0][2])*v.Z
	yv := float64(m.M[1][0])*v.X + float64(m.M[1][1])*v.Y + float64(m.M[1][2])*v.Z
	zv := float64(m.M[2][0])*v.X + float64(m.M[2][1])*v.Y + float64(m.M[2][2])*v.Z

	return NewVector3(xv, yv, zv)
}

func (m Matrix4x4) IsIdentity() bool {
	return m.M[0][0] == 1.0 && m.M[0][1] == 0.0 && m.M[0][2] == 0.0 && m.M[0][3] == 0.0 &&
		m.M[1][0] == 0.0 && m.M[1][1] == 1.0 && m.M[1][2] == 0.0 && m.M[1][3] == 0.0 &&
		m.M[2][0] == 0.0 && m.M[2][1] == 0.0 && m.M[2][2] == 1.0 && m.M[2][3] == 0.0 &&
		m.M[3][0] == 0.0 && m.M[3][1] == 0.0 && m.M[3][2] == 0.0 && m.M[3][3] == 1.0
}

// Numerically stable Gauss–Jordan elimination routine to compute the inverse
// See https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.cpp#L82
func (m Matrix4x4) Inverse() (Matrix4x4, error) {
	var indxc, indxr [4]int
	ipiv := [4]int{0, 0, 0, 0}
	// copy matrix arrays
	minv := m.M

	for i := range minv {
		irow := 0
		icol := 0
		big := 0.0
		// Choose pivot
		for j := range ipiv {
			if ipiv[j] != 1 {
				for k := range ipiv {
					if ipiv[k] == 0 {
						if math.Abs(float64(minv[j][k])) >= big {
							big = math.Abs(float64(minv[j][k]))
							irow = j
							icol = k
						}
					} else if ipiv[k] > 1 {
						return Matrix4x4{}, errors.New("singular matrix in MatrixInvert")
					}
				}
			}
		}
		ipiv[icol]++
		// Swap rows _irow_ and _icol_ for pivot
		if irow != icol {
			for k := 0; k < 4; k++ {
				// swap
				minv[irow][k], minv[icol][k] = minv[icol][k], minv[irow][k]
			}
		}
		indxr[i] = irow
		indxc[i] = icol
		if minv[icol][icol] == 0.0 {
			return Matrix4x4{}, errors.New("singular matrix in MatrixInvert")
		}

		// Set $m[icol][icol]$ to one by scaling row _icol_ appropriately
		pivinv := 1.0 / minv[icol][icol]
		minv[icol][icol] = 1.
		for j := 0; j < 4; j++ {
			minv[icol][j] *= pivinv
		}

		// Subtract this row from others to zero out their columns
		for j := 0; j < 4; j++ {
			if j != icol {
				save := minv[j][icol]
				minv[j][icol] = 0
				for k := 0; k < 4; k++ {
					minv[j][k] -= minv[icol][k] * save
				}
			}
		}
	}
	// Swap columns to reflect permutation
	for j := 3; j >= 0; j-- {
		if indxr[j] != indxc[j] {
			for k := 0; k < 4; k++ {
				// swap
				minv[k][indxr[j]], minv[k][indxc[j]] = minv[k][indxc[j]], minv[k][indxr[j]]
			}
		}
	}
	return Matrix4x4{minv}, nil
}
