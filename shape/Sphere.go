package shape

import (
	"math"
	"pbrt-go/mymath"
)

// Sphere see https://github.com/mmp/pbrt-v3/blob/master/src/shapes/sphere.h, https://github.com/mmp/pbrt-v3/blob/master/src/shapes/sphere.cpp
type Sphere struct {
	Shape
	Radius,
	ZMin, ZMax,
	ThetaMin, ThetaMax, PhiMax float64
}

func NewSphere(radius, zMin, zMax, phiMax float64, objectToWorld, worldToObject *mymath.Transform, reverseOrientation bool) *Sphere {
	return &Sphere{
		NewShape(objectToWorld, worldToObject, reverseOrientation),
		radius,
		zMin,
		zMax,
		math.Acos(mymath.Clamp(zMin/radius, -1, 1)),
		math.Acos(mymath.Clamp(zMax/radius, -1, 1)),
		mymath.Radians(mymath.Clamp(phiMax, 0, 360)),
	}
}
