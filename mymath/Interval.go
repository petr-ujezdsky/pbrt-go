package mymath

import "math"

// Interval see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/transform.cpp#L314
type Interval struct {
	Low, High float64
}

func NewIntervalSingle(v float64) Interval {
	return NewInterval(v, v)
}

func NewInterval(low, high float64) Interval {
	return Interval{math.Min(low, high), math.Max(high, low)}
}

func (i1 Interval) Add(i2 Interval) Interval {
	return NewInterval(i1.Low+i2.Low, i1.High+i2.High)
}

func (i1 Interval) Subtract(i2 Interval) Interval {
	return NewInterval(i1.Low-i2.High, i1.High-i2.Low)
}

func (i1 Interval) Multiply(i2 Interval) Interval {
	return NewInterval(
		math.Min(
			math.Min(i1.Low*i2.Low, i1.High*i2.Low),
			math.Min(i1.Low*i2.High, i1.High*i2.High)),
		math.Max(
			math.Max(i1.Low*i2.Low, i1.High*i2.Low),
			math.Max(i1.Low*i2.High, i1.High*i2.High)))
}

// Sin see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/transform.cpp#L335
func Sin(i1 Interval) Interval {
	sinLow := math.Sin(i1.Low)
	sinHigh := math.Sin(i1.High)

	if sinLow > sinHigh {
		sinLow, sinHigh = sinHigh, sinLow
	}

	if i1.Low < math.Pi/2.0 && i1.High > math.Pi/2 {
		sinHigh = 1
	}

	if i1.Low < (3.0/2.0)*math.Pi && i1.High > (3.0/2.0)*math.Pi {
		sinLow = -1
	}

	return NewInterval(sinLow, sinHigh)
}

// Cos see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/transform.cpp#L345
func Cos(i1 Interval) Interval {
	cosLow := math.Cos(i1.Low)
	cosHigh := math.Cos(i1.High)

	if cosLow > cosHigh {
		cosLow, cosHigh = cosHigh, cosLow
	}

	if i1.Low < math.Pi && i1.High > math.Pi {
		cosLow = -1
	}

	return NewInterval(cosLow, cosHigh)
}
