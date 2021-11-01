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

func NewTransformTranslate(delta *Vector3) Transform {
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

func (t *Transform) Apply(p *Point3) *Point3 {
	return t.m.MultiplyP(p)
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
