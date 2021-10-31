package point3d

import (
	"math"
	"pbrt-go/mymath/vector3d"
)

type Point3d struct {
	X, Y, Z float64
}

var EMPTY = NewPoint3d(0, 0, 0)

func NewPoint3d(x, y, z float64) Point3d {
	return Point3d{x, y, z}
}

func (p1 Point3d) AddP(p2 Point3d) Point3d {
	return NewPoint3d(p1.X+p2.X, p1.Y+p2.Y, p1.Z+p2.Z)
}

func (p Point3d) AddV(v vector3d.Vector3d) Point3d {
	return NewPoint3d(p.X+v.X, p.Y+v.Y, p.Z+v.Z)
}

func (p1 Point3d) SubtractP(p2 Point3d) vector3d.Vector3d {
	return vector3d.NewVector3d(p1.X-p2.X, p1.Y-p2.Y, p1.Z-p2.Z)
}

func (p Point3d) SubtractV(v vector3d.Vector3d) Point3d {
	return NewPoint3d(p.X-v.X, p.Y-v.Y, p.Z-v.Z)
}

func (p Point3d) Multiply(d float64) Point3d {
	return NewPoint3d(p.X*d, p.Y*d, p.Z*d)
}

func (p1 Point3d) Distance(p2 Point3d) float64 {
	return p1.SubtractP(p2).Length()
}

func (p1 Point3d) DistanceSq(p2 Point3d) float64 {
	return p1.SubtractP(p2).LengthSq()
}

func (p1 Point3d) Lerp(t float64, p2 Point3d) Point3d {
	return p1.Multiply(1 - t).AddP(p2.Multiply(t))
}

func lerp(t, v1, v2 float64) float64 {
	return (1-t)*v1 + t*v2
}

func (p1 Point3d) LerpP(t Point3d, p2 Point3d) Point3d {
	return NewPoint3d(
		lerp(t.X, p1.X, p2.X),
		lerp(t.Y, p1.Y, p2.Y),
		lerp(t.Z, p1.Z, p2.Z))
}

func (p Point3d) Min(w Point3d) Point3d {
	return NewPoint3d(
		math.Min(p.X, w.X),
		math.Min(p.Y, w.Y),
		math.Min(p.Z, w.Z))
}

func (p Point3d) Max(w Point3d) Point3d {
	return NewPoint3d(
		math.Max(p.X, w.X),
		math.Max(p.Y, w.Y),
		math.Max(p.Z, w.Z))
}

func (p Point3d) Floor() Point3d {
	return NewPoint3d(
		math.Floor(p.X),
		math.Floor(p.Y),
		math.Floor(p.Z))
}

func (p Point3d) Ceil() Point3d {
	return NewPoint3d(
		math.Ceil(p.X),
		math.Ceil(p.Y),
		math.Ceil(p.Z))
}

func (p Point3d) Abs() Point3d {
	return NewPoint3d(
		math.Abs(p.X),
		math.Abs(p.Y),
		math.Abs(p.Z))
}

func (p Point3d) Get(component int) float64 {
	switch component {
	case 0:
		return p.X
	case 1:
		return p.Y
	default:
		return p.Z
	}
}

func (p Point3d) Permute(x, y, z int) Point3d {
	return NewPoint3d(
		p.Get(x),
		p.Get(y),
		p.Get(z))
}
