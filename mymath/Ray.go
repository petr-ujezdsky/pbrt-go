package mymath

import (
	"pbrt-go/material"
)

type Ray struct {
	O      Point3d
	D      Vector3d
	TMax   float64
	Time   float32
	Medium material.Medium
}

func NewRay(o Point3d, d Vector3d, tMax float64, time float32, medium material.Medium) Ray {
	return Ray{o, d, tMax, time, medium}
}
