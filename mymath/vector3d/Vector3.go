package vector3d

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

func (v Vector3d) Add(w Vector3d) Vector3d {
	return NewVector3d(v.X+w.X, v.Y+w.Y, v.Z+w.Z)
}

func (v Vector3d) Subtract(w Vector3d) Vector3d {
	return NewVector3d(v.X-w.X, v.Y-w.Y, v.Z-w.Z)
}

func (v Vector3d) Multiply(d float64) Vector3d {
	return NewVector3d(v.X*d, v.Y*d, v.Z*d)
}

func (v Vector3d) Divide(d float64) Vector3d {
	inv := 1 / d
	return v.Multiply(inv)
}

func (v Vector3d) Negate() Vector3d {
	return NewVector3d(-v.X, -v.Y, -v.Z)
}

func (v Vector3d) Abs() Vector3d {
	return NewVector3d(math.Abs(v.X), math.Abs(v.Y), math.Abs(v.Z))
}

func (v Vector3d) Normalize() Vector3d {
	lengthInv := 1 / v.Length()
	return v.Multiply(lengthInv)
}

func (v Vector3d) Dot(w Vector3d) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

func (v Vector3d) Cross(w Vector3d) Vector3d {
	// always use doubles here!
	return NewVector3d(
		v.Y*w.Z-v.Z*w.Y,
		v.Z*w.X-v.X*w.Z,
		v.X*w.Y-v.Y*w.X)
}

func (v Vector3d) Get(component int) float64 {
	switch component {
	case 0:
		return v.X
	case 1:
		return v.Y
	default:
		return v.Z
	}
}

func (v Vector3d) GetMinComponent() float64 {
	return math.Min(v.X, math.Min(v.Y, v.Z))
}

func (v Vector3d) GetMaxComponent() float64 {
	return math.Max(v.X, math.Max(v.Y, v.Z))
}

func (v Vector3d) GetMaxDimension() int {
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

func (v Vector3d) Min(w Vector3d) Vector3d {
	return NewVector3d(
		math.Min(v.X, w.X),
		math.Min(v.Y, w.Y),
		math.Min(v.Z, w.Z))
}

func (v Vector3d) Max(w Vector3d) Vector3d {
	return NewVector3d(
		math.Max(v.X, w.X),
		math.Max(v.Y, w.Y),
		math.Max(v.Z, w.Z))
}

func (v Vector3d) Permute(x, y, z int) Vector3d {
	return NewVector3d(
		v.Get(x),
		v.Get(y),
		v.Get(z))
}
