package mymath

import "math"

type Vector3 struct {
	X, Y, Z float64
}

var EMPTY = NewVector3(0, 0, 0)

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{x, y, z}
}

func (v Vector3) LengthSq() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector3) Length() float64 {
	return math.Sqrt(v.LengthSq())
}

func (v Vector3) Add(w Vector3) Vector3 {
	return NewVector3(v.X+w.X, v.Y+w.Y, v.Z+w.Z)
}

func (v Vector3) Subtract(w Vector3) Vector3 {
	return NewVector3(v.X-w.X, v.Y-w.Y, v.Z-w.Z)
}

func (v Vector3) Multiply(d float64) Vector3 {
	return NewVector3(v.X*d, v.Y*d, v.Z*d)
}

func (v Vector3) Divide(d float64) Vector3 {
	inv := 1 / d
	return v.Multiply(inv)
}

func (v Vector3) Negate() Vector3 {
	return NewVector3(-v.X, -v.Y, -v.Z)
}

func (v Vector3) Abs() Vector3 {
	return NewVector3(math.Abs(v.X), math.Abs(v.Y), math.Abs(v.Z))
}

func (v Vector3) Normalize() Vector3 {
	lengthInv := 1 / v.Length()
	return v.Multiply(lengthInv)
}

func (v Vector3) Dot(w Vector3) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

func (v Vector3) Cross(w Vector3) Vector3 {
	// always use doubles here!
	return NewVector3(
		v.Y*w.Z-v.Z*w.Y,
		v.Z*w.X-v.X*w.Z,
		v.X*w.Y-v.Y*w.X)
}

func (v Vector3) Get(component int) float64 {
	switch component {
	case 0:
		return v.X
	case 1:
		return v.Y
	default:
		return v.Z
	}
}

func (v Vector3) GetMinComponent() float64 {
	return math.Min(v.X, math.Min(v.Y, v.Z))
}

func (v Vector3) GetMaxComponent() float64 {
	return math.Max(v.X, math.Max(v.Y, v.Z))
}

func (v Vector3) GetMaxDimension() int {
	if v.X > v.Y {
		if v.X > v.Z {
			return 0
		}
		return 2
	}
	if v.Y > v.Z {
		return 1
	}
	return 2
}

func (v Vector3) Min(w Vector3) Vector3 {
	return NewVector3(
		math.Min(v.X, w.X),
		math.Min(v.Y, w.Y),
		math.Min(v.Z, w.Z))
}

func (v Vector3) Max(w Vector3) Vector3 {
	return NewVector3(
		math.Max(v.X, w.X),
		math.Max(v.Y, w.Y),
		math.Max(v.Z, w.Z))
}

func (v Vector3) Permute(x, y, z int) Vector3 {
	return NewVector3(
		v.Get(x),
		v.Get(y),
		v.Get(z))
}
