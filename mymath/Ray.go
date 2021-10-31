package mymath

import (
	"pbrt-go/material"
	"pbrt-go/mymath/point3d"
	"pbrt-go/mymath/vector3d"
)

type Ray struct {
	O      point3d.Point3d
	D      vector3d.Vector3d
	TMax   float64
	Time   float32
	Medium material.Medium
}

func NewRay(o point3d.Point3d, d vector3d.Vector3d, tMax float64, time float32, medium material.Medium) Ray {
	return Ray{o, d, tMax, time, medium}
}
