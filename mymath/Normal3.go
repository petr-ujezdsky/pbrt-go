package mymath

import "math"

type Normal3 struct {
	X, Y, Z float64
}

func NewNormal3(x, y, z float64) Normal3 {
	return Normal3{x, y, z}
}

func NewNormal3V(v Vector3) Normal3 {
	return Normal3(v)
}

func (v Normal3) LengthSq() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Normal3) Length() float64 {
	return math.Sqrt(v.LengthSq())
}

func (v Normal3) Add(w Normal3) Normal3 {
	return NewNormal3(v.X+w.X, v.Y+w.Y, v.Z+w.Z)
}

func (v Normal3) Subtract(w Normal3) Normal3 {
	return NewNormal3(v.X-w.X, v.Y-w.Y, v.Z-w.Z)
}

func (v Normal3) Multiply(d float64) Normal3 {
	return NewNormal3(v.X*d, v.Y*d, v.Z*d)
}

func (v Normal3) Divide(d float64) Normal3 {
	inv := 1 / d
	return v.Multiply(inv)
}

func (v Normal3) Negate() Normal3 {
	return NewNormal3(-v.X, -v.Y, -v.Z)
}

// func (v Normal3) Abs() Normal3 {
// 	return NewNormal3(math.Abs(v.X), math.Abs(v.Y), math.Abs(v.Z))
// }

// func (v Normal3) Normalize() Normal3 {
// 	lengthInv := 1 / v.Length()
// 	return v.Multiply(lengthInv)
// }

// func (v Normal3) Dot(w Normal3) float64 {
// 	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
// }

// func (v Normal3) Cross(w Normal3) Normal3 {
// 	// always use doubles here!
// 	return NewNormal3(
// 		v.Y*w.Z-v.Z*w.Y,
// 		v.Z*w.X-v.X*w.Z,
// 		v.X*w.Y-v.Y*w.X)
// }

func (v Normal3) Get(component int) float64 {
	switch component {
	case 0:
		return v.X
	case 1:
		return v.Y
	default:
		return v.Z
	}
}

// func (v Normal3) GetMinComponent() float64 {
// 	return math.Min(v.X, math.Min(v.Y, v.Z))
// }

// func (v Normal3) GetMaxComponent() float64 {
// 	return math.Max(v.X, math.Max(v.Y, v.Z))
// }

// func (v Normal3) GetMaxDimension() int {
// 	if v.X > v.Y {
// 		if v.X > v.Z {
// 			return 0
// 		}
// 		return 2
// 	}
// 	if v.Y > v.Z {
// 		return 1
// 	}
// 	return 2
// }

// func (v Normal3) Min(w Normal3) Normal3 {
// 	return NewNormal3(
// 		math.Min(v.X, w.X),
// 		math.Min(v.Y, w.Y),
// 		math.Min(v.Z, w.Z))
// }

// func (v Normal3) Max(w Normal3) Normal3 {
// 	return NewNormal3(
// 		math.Max(v.X, w.X),
// 		math.Max(v.Y, w.Y),
// 		math.Max(v.Z, w.Z))
// }

// func (v Normal3) Permute(x, y, z int) Normal3 {
// 	return NewNormal3(
// 		v.Get(x),
// 		v.Get(y),
// 		v.Get(z))
// }

func (v Normal3) HasNaNs(x, y, z int) bool {
	return math.IsNaN(v.X) || math.IsNaN(v.Y) || math.IsNaN(v.Z)
}
