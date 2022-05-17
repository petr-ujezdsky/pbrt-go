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

// Intersect finds ray-shape collision point and its metadata
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/shapes/sphere.cpp#L49
func (s Sphere) Intersect(r Ray, _ bool) (bool, float64, *SurfaceInteraction) {
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
	b := NewEFloat(2).Multiply(dx.Multiply(ox).Add(dy.Multiply(oy)).Add(dz.Multiply(oz)))
	c := ox.Multiply(ox).Add(oy.Multiply(oy)).Add(oz.Multiply(oz)).Subtract(NewEFloat(s.Radius).Multiply(NewEFloat(s.Radius)))

	// Solve quadratic equation for t values
	ok, t0, t1 := Quadratic(a, b, c)

	if !ok {
		return false, 0, nil
	}

	// Check quadratic shape t0 and t1 for nearest intersection
	if t0.High > ray.TMax || t1.Low <= 0 {
		return false, 0, nil
	}

	tShapeHit := t0
	if tShapeHit.Low <= 0 {
		tShapeHit = t1
		if tShapeHit.High > ray.TMax {
			return false, 0, nil
		}
	}

	// Compute sphere hit position and phi
	pHit := ray.Apply(tShapeHit.V)

	// Refine sphere intersection point
	pHit = pHit.Multiply(s.Radius / pHit.Distance(NewPoint3(0, 0, 0)))

	if pHit.X == 0 && pHit.Y == 0 {
		pHit.X = 1e-5 * s.Radius
	}

	phi := math.Atan2(pHit.Y, pHit.X)
	if phi < 0 {
		phi += 2 * math.Pi
	}

	// Test sphere intersection against clipping parameters
	if (s.ZMin > -s.Radius && pHit.Z < s.ZMin) ||
		(s.ZMax < s.Radius && pHit.Z > s.ZMax) || phi > s.PhiMax {
		if tShapeHit == t1 {
			return false, 0, nil
		}

		if t1.High > ray.TMax {
			return false, 0, nil
		}
		tShapeHit = t1

		// Compute sphere hit position and phi
		pHit := ray.Apply(tShapeHit.V)

		// Refine sphere intersection point
		pHit = pHit.Multiply(s.Radius / pHit.Distance(NewPoint3(0, 0, 0)))

		if pHit.X == 0 && pHit.Y == 0 {
			pHit.X = 1e-5 * s.Radius
		}

		phi := math.Atan2(pHit.Y, pHit.X)
		if phi < 0 {
			phi += 2 * math.Pi
		}

		// Compute sphere hit position and phi
		if (s.ZMin > -s.Radius && pHit.Z < s.ZMin) ||
			(s.ZMax < s.Radius && pHit.Z > s.ZMax) || phi > s.PhiMax {
			return false, 0, nil
		}
	}

	// Find parametric representation of sphere hit
	u := phi / s.PhiMax
	theta := math.Acos(Clamp(pHit.Z/s.Radius, -1, 1))
	v := (theta - s.ThetaMin) / (s.ThetaMax - s.ThetaMin)

	// Compute sphere dpdu and dpdv
	zRadius := math.Sqrt(pHit.X*pHit.X + pHit.Y*pHit.Y)
	invZRadius := 1 / zRadius
	cosPhi := pHit.X * invZRadius
	sinPhi := pHit.Y * invZRadius
	dpdu := NewVector3(-s.PhiMax*pHit.Y, s.PhiMax*pHit.X, 0)
	dpdv := NewVector3(pHit.Z*cosPhi, pHit.Z*sinPhi, -s.Radius*math.Sin(theta)).Multiply(s.ThetaMax - s.ThetaMin)

	// Compute sphere dndu and dndv
	d2Pduu := NewVector3(pHit.X, pHit.Y, 0).Multiply(-s.PhiMax * s.PhiMax)
	d2Pduv := NewVector3(-sinPhi, cosPhi, 0.0).Multiply((s.ThetaMax - s.ThetaMin) * pHit.Z * s.PhiMax)
	d2Pdvv := NewVector3(pHit.X, pHit.Y, pHit.Z).Multiply(-(s.ThetaMax - s.ThetaMin) * (s.ThetaMax - s.ThetaMin))

	// Compute coefficients for fundamental forms
	E := dpdu.Dot(dpdu)
	F := dpdu.Dot(dpdv)
	G := dpdv.Dot(dpdv)
	N := dpdu.Cross(dpdv).Normalize()
	e := N.Dot(d2Pduu)
	f := N.Dot(d2Pduv)
	g := N.Dot(d2Pdvv)

	// Compute dndu and dndv from fundamental form coefficients
	invEGF2 := 1 / (E*G - F*F)
	dndu := NewNormal3V(dpdu.Multiply((f*F - e*G) * invEGF2).Add(dpdv.Multiply((e*F - f*E) * invEGF2)))
	dndv := NewNormal3V(dpdu.Multiply((g*F - f*G) * invEGF2).Add(dpdv.Multiply((f*F - g*E) * invEGF2)))

	// Compute error bounds for sphere intersection
	pError := NewVector3P(pHit).Abs().Multiply(Gamma5)

	// Initialize _SurfaceInteraction_ from parametric information
	si := NewSurfaceInteraction(
		pHit,
		pError,
		NewPoint2(u, v),
		ray.D.Negate(),
		dpdu,
		dpdv,
		dndu,
		dndv,
		float64(ray.Time),
		&s.Shape)

	isect := s.ObjectToWorld.ApplySI(&si)

	// Update _tHit_ for quadric intersection
	return true, tShapeHit.V, isect
}

// IntersectP finds if ray collides with this shape
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/shapes/sphere.cpp#L158
func (s Sphere) IntersectP(i Intersecter, r Ray, _ bool) bool {
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
	b := NewEFloat(2).Multiply(dx.Multiply(ox).Add(dy.Multiply(oy)).Add(dz.Multiply(oz)))
	c := ox.Multiply(ox).Add(oy.Multiply(oy)).Add(oz.Multiply(oz)).Subtract(NewEFloat(s.Radius).Multiply(NewEFloat(s.Radius)))

	// Solve quadratic equation for t values
	ok, t0, t1 := Quadratic(a, b, c)

	if !ok {
		return false
	}

	// Check quadratic shape t0 and t1 for nearest intersection
	if t0.High > ray.TMax || t1.Low <= 0 {
		return false
	}

	tShapeHit := t0
	if tShapeHit.Low <= 0 {
		tShapeHit = t1
		if tShapeHit.High > ray.TMax {
			return false
		}
	}

	// Compute sphere hit position and phi
	pHit := ray.Apply(tShapeHit.V)

	// Refine sphere intersection point
	pHit = pHit.Multiply(s.Radius / pHit.Distance(NewPoint3(0, 0, 0)))

	if pHit.X == 0 && pHit.Y == 0 {
		pHit.X = 1e-5 * s.Radius
	}

	phi := math.Atan2(pHit.Y, pHit.X)
	if phi < 0 {
		phi += 2 * math.Pi
	}

	// Test sphere intersection against clipping parameters
	if (s.ZMin > -s.Radius && pHit.Z < s.ZMin) ||
		(s.ZMax < s.Radius && pHit.Z > s.ZMax) || phi > s.PhiMax {
		if tShapeHit == t1 {
			return false
		}

		if t1.High > ray.TMax {
			return false
		}
		tShapeHit = t1

		// Compute sphere hit position and phi
		pHit := ray.Apply(tShapeHit.V)

		// Refine sphere intersection point
		pHit = pHit.Multiply(s.Radius / pHit.Distance(NewPoint3(0, 0, 0)))

		if pHit.X == 0 && pHit.Y == 0 {
			pHit.X = 1e-5 * s.Radius
		}

		phi := math.Atan2(pHit.Y, pHit.X)
		if phi < 0 {
			phi += 2 * math.Pi
		}

		// Compute sphere hit position and phi
		if (s.ZMin > -s.Radius && pHit.Z < s.ZMin) ||
			(s.ZMax < s.Radius && pHit.Z > s.ZMax) || phi > s.PhiMax {
			return false
		}
	}

	return true
}

func (s Sphere) Area() float64 {
	return s.PhiMax * s.Radius * (s.ZMax - s.ZMin)
}
