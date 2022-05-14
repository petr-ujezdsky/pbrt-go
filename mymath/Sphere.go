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

func (s Sphere) ObjectBound() Bounds3 {
	return NewBounds3(
		NewPoint3(-s.Radius, -s.Radius, s.ZMin),
		NewPoint3(s.Radius, s.Radius, s.ZMax))
}

func (s Sphere) Intersect(r Ray, testAlphaTexture bool) (bool, float64, SurfaceInteraction) {
	//Float phi;
	//Point3f pHit;

	// Transform Ray to object space
	ray, oErr, dErr := s.WorldToObject.ApplyRError(r)

	// Compute quadratic sphere coefficients
	// Initialize EFloat ray coordinate values
	ox := NewEFloatErr(ray.O.X, oErr.X)
	oy := NewEFloatErr(ray.O.Y, oErr.Y)
	oz := NewEFloatErr(ray.O.Z, oErr.Z)

	dx := NewEFloatErr(ray.D.X, dErr.X)
	dy := NewEFloatErr(ray.D.Y, dErr.Y)
	dz := NewEFloatErr(ray.D.Z, dErr.Z)

	a := dx.Multiply(dx).Add(dy.Multiply(dy)).Add(dz.Multiply(dz))
	b := NewEFloat(2).Multiply((dx.Multiply(ox).Add(dy.Multiply(oy)).Add(dz.Multiply(oz)))
	c := ox.Multiply(ox).Add(oy.Multiply(oy)).Add(oz.Multiply(oz)).Subtract(NewEFloat(s.Radius).Multiply(NewEFloat(s.Radius)))
}
