package mymath

type Transform struct {
	m, mInv Matrix4x4
}

func NewTransformEmpty() Transform {
	return Transform{Identity(), Identity()}
}

func NewTransform(m *Matrix4x4) Transform {
	return Transform{*m, *m.Inverse()}
}

func NewTransformFull(m *Matrix4x4, mInv *Matrix4x4) Transform {
	return Transform{*m, *mInv}
}

func (t *Transform) Inverse() Transform {
	return Transform{t.mInv, t.m}
}

func (t *Transform) Transpose() Transform {
	return Transform{*t.m.Transpose(), *t.mInv.Transpose()}
}

func (t *Transform) IsIdentity() bool {
	return t.m.IsIdentity()
}
