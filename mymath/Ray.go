package mymath

import (
	"pbrt-go/material"
)

// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/geometry.h#L869
type Ray struct {
	O      Point3
	D      Vector3
	TMax   float64
	Time   float32
	Medium material.Medium
}

func NewRay(o Point3, d Vector3, tMax float64, time float32, medium material.Medium) Ray {
	return Ray{o, d, tMax, time, medium}
}

// Apply computes point along the ray given the parameter t
func (r Ray) Apply(t float64) Point3 {
	return r.O.AddV(r.D.Multiply(t))
}
