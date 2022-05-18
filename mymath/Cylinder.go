package mymath

import (
	"math"
)

// Cylinder
// see https://github.com/mmp/pbrt-v3/blob/master/src/shapes/cylinder.h
// see https://github.com/mmp/pbrt-v3/blob/master/src/shapes/cylinder.cpp
type Cylinder struct {
	Shape
	Radius,
	ZMin, ZMax,
	PhiMax float64
}

func NewCylinder(radius, zMin, zMax, phiMax float64, objectToWorld, worldToObject *Transform, reverseOrientation bool) *Cylinder {
	return &Cylinder{
		NewShape(objectToWorld, worldToObject, reverseOrientation),
		radius,
		math.Min(zMin, zMax),
		math.Max(zMin, zMax),
		Radians(Clamp(phiMax, 0, 360)),
	}
}

func (cyl Cylinder) ObjectBound() Bounds3 {
	return NewBounds3(
		NewPoint3(-cyl.Radius, -cyl.Radius, cyl.ZMin),
		NewPoint3(cyl.Radius, cyl.Radius, cyl.ZMax))
}

// Intersect finds ray-shape collision point and its metadata
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/shapes/cylinder.cpp#L48
func (cyl Cylinder) Intersect(r Ray, _ bool) (bool, float64, *SurfaceInteraction) {
	// Transform Ray to object space
	ray, oErr, dErr := cyl.WorldToObject.ApplyRError(r)

	// Compute quadratic sphere coefficients

	// Initialize EFloat ray coordinate values
	ox := NewEFloatErr(ray.O.X, oErr.X)
	oy := NewEFloatErr(ray.O.Y, oErr.Y)

	dx := NewEFloatErr(ray.D.X, dErr.X)
	dy := NewEFloatErr(ray.D.Y, dErr.Y)

	a := dx.Multiply(dx).Add(dy.Multiply(dy))
	b := NewEFloat(2).Multiply(dx.Multiply(ox).Add(dy.Multiply(oy)))
	c := ox.Multiply(ox).Add(oy.Multiply(oy)).Subtract(NewEFloat(cyl.Radius).Multiply(NewEFloat(cyl.Radius)))

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

	// Compute cylinder hit position and phi
	pHit := ray.Apply(tShapeHit.V)

	// Refine sphere intersection point
	hitRad := math.Sqrt(pHit.X*pHit.X + pHit.Y*pHit.Y)
	pHit.X *= cyl.Radius / hitRad
	pHit.Y *= cyl.Radius / hitRad

	phi := math.Atan2(pHit.Y, pHit.X)
	if phi < 0 {
		phi += 2 * math.Pi
	}

	// Test cylinder intersection against clipping parameters
	if pHit.Z < cyl.ZMin || pHit.Z > cyl.ZMax || phi > cyl.PhiMax {
		if tShapeHit == t1 {
			return false, 0, nil
		}

		tShapeHit = t1
		if t1.High > ray.TMax {
			return false, 0, nil
		}

		// Compute cylinder hit position and phi
		pHit = ray.Apply(tShapeHit.V)

		// Refine sphere intersection point
		hitRad = math.Sqrt(pHit.X*pHit.X + pHit.Y*pHit.Y)
		pHit.X *= cyl.Radius / hitRad
		pHit.Y *= cyl.Radius / hitRad

		phi = math.Atan2(pHit.Y, pHit.X)
		if phi < 0 {
			phi += 2 * math.Pi
		}

		if pHit.Z < cyl.ZMin || pHit.Z > cyl.ZMax || phi > cyl.PhiMax {
			return false, 0, nil
		}
	}

	// Find parametric representation of cylinder hit
	u := phi / cyl.PhiMax
	v := (pHit.Z - cyl.ZMin) / (cyl.ZMax - cyl.ZMin)

	// Compute cylinder dpdu and dpdv
	dpdu := NewVector3(-cyl.PhiMax*pHit.Y, cyl.PhiMax*pHit.X, 0)
	dpdv := NewVector3(0, 0, cyl.ZMax-cyl.ZMin)

	// Compute cylinder dndu and dndv
	d2Pduu := NewVector3(pHit.X, pHit.Y, 0).Multiply(-cyl.PhiMax * cyl.PhiMax)
	d2Pduv := NewVector3(0, 0, 0)
	d2Pdvv := NewVector3(0, 0, 0)

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

	// Compute error bounds for cylinder intersection
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
		&cyl.Shape)

	isect := cyl.ObjectToWorld.ApplySI(&si)

	// Update _tHit_ for quadric intersection
	return true, tShapeHit.V, isect
}

// IntersectP finds if ray collides with this shape
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/shapes/cylinder.cpp#L146
func (cyl Cylinder) IntersectP(i Intersecter, r Ray, _ bool) bool {
	// Transform Ray to object space
	ray, oErr, dErr := cyl.WorldToObject.ApplyRError(r)

	// Compute quadratic sphere coefficients

	// Initialize EFloat ray coordinate values
	ox := NewEFloatErr(ray.O.X, oErr.X)
	oy := NewEFloatErr(ray.O.Y, oErr.Y)

	dx := NewEFloatErr(ray.D.X, dErr.X)
	dy := NewEFloatErr(ray.D.Y, dErr.Y)

	a := dx.Multiply(dx).Add(dy.Multiply(dy))
	b := NewEFloat(2).Multiply(dx.Multiply(ox).Add(dy.Multiply(oy)))
	c := ox.Multiply(ox).Add(oy.Multiply(oy)).Subtract(NewEFloat(cyl.Radius).Multiply(NewEFloat(cyl.Radius)))

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

	// Compute cylinder hit position and phi
	pHit := ray.Apply(tShapeHit.V)

	// Refine sphere intersection point
	hitRad := math.Sqrt(pHit.X*pHit.X + pHit.Y*pHit.Y)
	pHit.X *= cyl.Radius / hitRad
	pHit.Y *= cyl.Radius / hitRad

	phi := math.Atan2(pHit.Y, pHit.X)
	if phi < 0 {
		phi += 2 * math.Pi
	}

	// Test cylinder intersection against clipping parameters
	if pHit.Z < cyl.ZMin || pHit.Z > cyl.ZMax || phi > cyl.PhiMax {
		if tShapeHit == t1 {
			return false
		}

		tShapeHit = t1
		if t1.High > ray.TMax {
			return false
		}

		// Compute cylinder hit position and phi
		pHit = ray.Apply(tShapeHit.V)

		// Refine sphere intersection point
		hitRad = math.Sqrt(pHit.X*pHit.X + pHit.Y*pHit.Y)
		pHit.X *= cyl.Radius / hitRad
		pHit.Y *= cyl.Radius / hitRad

		phi = math.Atan2(pHit.Y, pHit.X)
		if phi < 0 {
			phi += 2 * math.Pi
		}

		if pHit.Z < cyl.ZMin || pHit.Z > cyl.ZMax || phi > cyl.PhiMax {
			return false
		}
	}

	return true
}

func (cyl Cylinder) Area() float64 {
	return cyl.PhiMax * cyl.Radius * (cyl.ZMax - cyl.ZMin)
}
