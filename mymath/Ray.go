package mymath

import (
	"pbrt-go/material"
)

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
