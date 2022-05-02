package mymath

type Point2 struct {
	X, Y float64
}

func NewPoint2(x, y float64) Point2 {
	return Point2{x, y}
}

//
//func (p1 Point3) AddP(p2 Point3) Point3 {
//	return NewPoint3(p1.X+p2.X, p1.Y+p2.Y, p1.Z+p2.Z)
//}
//
//func (p Point3) AddV(v Vector3) Point3 {
//	return NewPoint3(p.X+v.X, p.Y+v.Y, p.Z+v.Z)
//}
//
//func (p1 Point3) SubtractP(p2 Point3) Vector3 {
//	return NewVector3(p1.X-p2.X, p1.Y-p2.Y, p1.Z-p2.Z)
//}
//
//func (p Point3) SubtractV(v Vector3) Point3 {
//	return NewPoint3(p.X-v.X, p.Y-v.Y, p.Z-v.Z)
//}
//
//func (p Point3) Multiply(d float64) Point3 {
//	return NewPoint3(p.X*d, p.Y*d, p.Z*d)
//}
//
//func (p1 Point3) Distance(p2 Point3) float64 {
//	return p1.SubtractP(p2).Length()
//}
//
//func (p1 Point3) DistanceSq(p2 Point3) float64 {
//	return p1.SubtractP(p2).LengthSq()
//}
//
//func (p1 Point3) Lerp(t float64, p2 Point3) Point3 {
//	return p1.Multiply(1 - t).AddP(p2.Multiply(t))
//}
//
//func (p1 Point3) LerpP(t Point3, p2 Point3) Point3 {
//	return NewPoint3(
//		Lerp(t.X, p1.X, p2.X),
//		Lerp(t.Y, p1.Y, p2.Y),
//		Lerp(t.Z, p1.Z, p2.Z))
//}
//
//func (p Point3) Min(w Point3) Point3 {
//	return NewPoint3(
//		math.Min(p.X, w.X),
//		math.Min(p.Y, w.Y),
//		math.Min(p.Z, w.Z))
//}
//
//func (p Point3) Max(w Point3) Point3 {
//	return NewPoint3(
//		math.Max(p.X, w.X),
//		math.Max(p.Y, w.Y),
//		math.Max(p.Z, w.Z))
//}
//
//func (p Point3) Floor() Point3 {
//	return NewPoint3(
//		math.Floor(p.X),
//		math.Floor(p.Y),
//		math.Floor(p.Z))
//}
//
//func (p Point3) Ceil() Point3 {
//	return NewPoint3(
//		math.Ceil(p.X),
//		math.Ceil(p.Y),
//		math.Ceil(p.Z))
//}
//
//func (p Point3) Abs() Point3 {
//	return NewPoint3(
//		math.Abs(p.X),
//		math.Abs(p.Y),
//		math.Abs(p.Z))
//}
//
//func (p Point3) Get(component int) float64 {
//	switch component {
//	case 0:
//		return p.X
//	case 1:
//		return p.Y
//	default:
//		return p.Z
//	}
//}
//
//func (p Point3) Permute(x, y, z int) Point3 {
//	return NewPoint3(
//		p.Get(x),
//		p.Get(y),
//		p.Get(z))
//}
