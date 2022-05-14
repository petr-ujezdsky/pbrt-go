package mymath

import (
	"math"
)

// Sphere see https://github.com/mmp/pbrt-v3/blob/master/src/shapes/sphere.h, https://github.com/mmp/pbrt-v3/blob/master/src/shapes/sphere.cpp
type Sphere struct {
	Shape
	Radius,
	ZMin, ZMax,
	ThetaMin, ThetaMax, PhiMax float64
}

func NewSphere(radius, zMin, zMax, phiMax float64, objectToWorld, worldToObject *Transform, reverseOrientation bool) *Sphere {
	return &Sphere{
		NewShape(objectToWorld, worldToObject, reverseOrientation),
		radius,
		Clamp(math.Min(zMin, zMax), -radius, radius),
		Clamp(math.Max(zMin, zMax), -radius, radius),
		math.Acos(Clamp(zMin/radius, -1, 1)),
		math.Acos(Clamp(zMax/radius, -1, 1)),
		Radians(Clamp(phiMax, 0, 360)),
	}
}
