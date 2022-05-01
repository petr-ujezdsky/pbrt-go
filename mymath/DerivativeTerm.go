package mymath

// DerivativeTerm see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/transform.h#L439
type DerivativeTerm struct {
	Kc, Kx, Ky, Kz float64
}

func NewDerivativeTerm(kc, kx, ky, kz float64) DerivativeTerm {
	return DerivativeTerm{kc, kx, ky, kz}
}

func (dt DerivativeTerm) Eval(p Point3) float64 {
	return dt.Kc + dt.Kx*p.X + dt.Ky*p.Y + dt.Kz*p.Z
}
