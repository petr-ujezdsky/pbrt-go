package mymath

import "math"

// EFloat
//
// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/efloat.h
type EFloat struct {
	v, low, high float64
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
		ef.v + ef2.v,
		NextFloatDown(ef.low + ef2.low),
		NextFloatUp(ef.high + ef2.high),
	}

	r.check()

	return r
}

// Subtract
//
// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/efloat.h#L101
func (ef EFloat) Subtract(ef2 EFloat) EFloat {
	r := EFloat{
		ef.v - ef2.v,
		NextFloatDown(ef.low - ef2.high),
		NextFloatUp(ef.high - ef2.low),
	}

	r.check()

	return r
}

// Multiply
//
// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/efloat.h#L112
func (ef EFloat) Multiply(ef2 EFloat) EFloat {
	prod := [4]float64{
		ef.low * ef2.low, ef.high * ef2.low,
		ef.low * ef2.high, ef.high * ef2.high,
	}

	r := EFloat{
		ef.v * ef2.v,
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

	if ef2.low < 0 && ef2.high > 0 {
		// Bah. The interval we're dividing by straddles zero, so just
		// return an interval of everything.
		r = EFloat{
			ef.v / ef2.v,
			math.Inf(-1),
			math.Inf(1),
		}
	} else {
		div := [4]float64{
			ef.low / ef2.low, ef.high / ef2.low,
			ef.low / ef2.high, ef.high / ef2.high,
		}

		r = EFloat{
			ef.v / ef2.v,
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
	return ef.v
}

func (ef EFloat) Double() float64 {
	return ef.v
}

// GetAbsoluteError
//
// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/efloat.h#L91
func (ef EFloat) GetAbsoluteError() float64 {
	return NextFloatUp(math.Max(math.Abs(ef.high-ef.v), math.Abs(ef.v-ef.low)))
}

func (ef EFloat) check() {
	if !math.IsInf(ef.low, 0) && !math.IsNaN(ef.low) && !math.IsInf(ef.high, 0) &&
		!math.IsNaN(ef.high) {
		CheckLE(ef.low, ef.high)
	}
}
