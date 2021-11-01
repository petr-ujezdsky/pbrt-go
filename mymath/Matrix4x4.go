package mymath

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

func (m *Matrix4x4) Transpose() *Matrix4x4 {
	M := m.M
	r := NewMatrix4x4All(
		M[0][0], M[1][0], M[2][0], M[3][0],
		M[0][1], M[1][1], M[2][1], M[3][1],
		M[0][2], M[1][2], M[2][2], M[3][2],
		M[0][3], M[1][3], M[2][3], M[3][3],
	)

	return &r
}

func (m1 *Matrix4x4) Multiply(m2 *Matrix4x4) *Matrix4x4 {
	r := Matrix4x4{}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			r.M[i][j] = m1.M[i][0]*m2.M[0][j] +
				m1.M[i][1]*m2.M[1][j] +
				m1.M[i][2]*m2.M[2][j] +
				m1.M[i][3]*m2.M[3][j]
		}
	}

	return &r
}
