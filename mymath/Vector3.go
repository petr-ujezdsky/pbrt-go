package mymath

import "math"

type Vector3d struct {
	X, Y, Z float64
}

func NewVector3d(x, y, z float64) Vector3d {
	return Vector3d{x, y, z}
}

func (v Vector3d) LengthSq() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector3d) Length() float64 {
	return math.Sqrt(v.LengthSq())
}
