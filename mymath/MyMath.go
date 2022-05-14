package mymath

import "math"

var epsilon = math.Nextafter(1, 2) - 1

var Gamma3 = gamma(3)

// Lerp returns value interpolated between v1 and v2 using parameter t.
func Lerp(t, v1, v2 float64) float64 {
	return (1-t)*v1 + t*v2
}

// Clamp returns value inside given interval unchanged and the interval limits if value is outside the interval.
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/core/pbrt.h#L304
func Clamp(val, low, high float64) float64 {
	if val < low {
		return low
	}

	if val > high {
		return high
	}

	return val
}

func Radians(deg float64) float64 {
	return (math.Pi / 180) * deg
}

func Degrees(rad float64) float64 {
	return (180 / math.Pi) * rad
}

func gamma(n int) float64 {
	ne := float64(n) * epsilon
	return ne / (1 - ne)
}

// NextFloatUp
//
// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/pbrt.h#L241
func NextFloatUp(v float64) float64 {
	// Handle infinity and negative zero for _NextFloatDown()_
	if math.IsInf(v, 0) && v > 0 {
		return v
	}

	if v == -0.0 {
		v = 0.0
	}

	// Advance _v_ to next higher float
	ui := math.Float64bits(v)

	if v >= 0 {
		ui++
	} else {
		ui--
	}

	return math.Float64frombits(ui)
}

// NextFloatDown
//
// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/pbrt.h#L255
func NextFloatDown(v float64) float64 {
	// Handle infinity and positive zero for _NextFloatDown()_
	if math.IsInf(v, 0) && v < 0 {
		return v
	}

	if v == 0.0 {
		v = -0.0
	}

	ui := math.Float64bits(v)

	if v > 0 {
		ui--
	} else {
		ui++
	}

	return math.Float64frombits(ui)
}

func CheckLE(v1, v2 float64) {
	if v1 > v2 {
		panic("Error")
	}
}
