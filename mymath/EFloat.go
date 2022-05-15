package mymath

import "C"
import "math"

// EFloat
//
// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/efloat.h
type EFloat struct {
	V, Low, High float64
}

func NewEFloat(v float64) EFloat {
	return NewEFloatErr(v, 0)
}

func NewEFloatErr(v, err float64) EFloat {
	var r EFloat

	if err == 0.0 {
		r = EFloat{v, v, v}
	} else {
		// Compute conservative bounds by rounding the endpoints away
		// from the middle. Note that this will be over-conservative in
		// cases where v-err or v+err are exactly representable in
		// floating-point, but it's probably not worth the trouble of
		// checking this case.
		r = EFloat{
			v,
			NextFloatDown(v - err),
			NextFloatUp(v + err),
		}
	}

	r.check()

	return r
}

// Add
//
// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/efloat.h#L76
func (ef EFloat) Add(ef2 EFloat) EFloat {
	r := EFloat{
		ef.V + ef2.V,
		NextFloatDown(ef.Low + ef2.Low),
		NextFloatUp(ef.High + ef2.High),
	}

	r.check()

	return r
}

// Subtract
//
// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/efloat.h#L101
func (ef EFloat) Subtract(ef2 EFloat) EFloat {
	r := EFloat{
		ef.V - ef2.V,
		NextFloatDown(ef.Low - ef2.High),
		NextFloatUp(ef.High - ef2.Low),
	}

	r.check()

	return r
}

// Multiply
//
// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/efloat.h#L112
func (ef EFloat) Multiply(ef2 EFloat) EFloat {
	prod := [4]float64{
		ef.Low * ef2.Low, ef.High * ef2.Low,
		ef.Low * ef2.High, ef.High * ef2.High,
	}

	r := EFloat{
		ef.V * ef2.V,
		NextFloatDown(
			math.Min(math.Min(prod[0], prod[1]), math.Min(prod[2], prod[3]))),
		NextFloatUp(
			math.Max(math.Max(prod[0], prod[1]), math.Max(prod[2], prod[3]))),
	}

	r.check()

	return r
}

// Divide
//
// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/efloat.h#L128
func (ef EFloat) Divide(ef2 EFloat) EFloat {
	var r EFloat

	if ef2.Low < 0 && ef2.High > 0 {
		// Bah. The interval we're dividing by straddles zero, so just
		// return an interval of everything.
		r = EFloat{
			ef.V / ef2.V,
			math.Inf(-1),
			math.Inf(1),
		}
	} else {
		div := [4]float64{
			ef.Low / ef2.Low, ef.High / ef2.Low,
			ef.Low / ef2.High, ef.High / ef2.High,
		}

		r = EFloat{
			ef.V / ef2.V,
			NextFloatDown(
				math.Min(math.Min(div[0], div[1]), math.Min(div[2], div[3]))),
			NextFloatUp(
				math.Max(math.Max(div[0], div[1]), math.Max(div[2], div[3]))),
		}
	}

	r.check()

	return r
}

func (ef EFloat) Float() float64 {
	return ef.V
}

func (ef EFloat) Double() float64 {
	return ef.V
}

// GetAbsoluteError
//
// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/efloat.h#L91
func (ef EFloat) GetAbsoluteError() float64 {
	return NextFloatUp(math.Max(math.Abs(ef.High-ef.V), math.Abs(ef.V-ef.Low)))
}

func (ef EFloat) check() {
	if !math.IsInf(ef.Low, 0) && !math.IsNaN(ef.Low) && !math.IsInf(ef.High, 0) &&
		!math.IsNaN(ef.High) {
		CheckLE(ef.Low, ef.High)
	}
}

// Quadratic solves quadratic equation
//
// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/efloat.h#L268
func Quadratic(a, b, c EFloat) (bool, EFloat, EFloat) {
	// Find quadratic discriminant
	discrim := b.V*b.V - 4.*a.V*c.V

	if discrim < 0.0 {
		return false, EFloat{}, EFloat{}
	}

	rootDiscrim := math.Sqrt(discrim)

	floatRootDiscrim := NewEFloatErr(rootDiscrim, epsilon*rootDiscrim)

	// Compute quadratic _t_ values
	var q EFloat
	if b.V < 0.0 {
		q = NewEFloat(-0.5).Multiply(b.Subtract(floatRootDiscrim))
	} else {
		q = NewEFloat(-0.5).Multiply(b.Add(floatRootDiscrim))
	}

	t0 := q.Divide(a)
	t1 := c.Divide(q)

	if t0.V > t1.V {
		t0, t1 = t1, t0
	}

	return true, t0, t1
}
