package mymath

import (
	"math"
)

// Disk
// see https://github.com/mmp/pbrt-v3/blob/master/src/shapes/disk.h
// see https://github.com/mmp/pbrt-v3/blob/master/src/shapes/disk.cpp
type Disk struct {
	Shape
	Height,
	Radius,
	InnerRadius,
	PhiMax float64
}

func NewDisk(height, radius, innerRadius, phiMax float64, objectToWorld, worldToObject *Transform, reverseOrientation bool) *Disk {
	return &Disk{
		NewShape(objectToWorld, worldToObject, reverseOrientation),
		height,
		radius,
		innerRadius,
		Radians(Clamp(phiMax, 0, 360)),
	}
}

func (disk Disk) ObjectBound() Bounds3 {
	return NewBounds3(
		NewPoint3(-disk.Radius, -disk.Radius, disk.Height),
		NewPoint3(disk.Radius, disk.Radius, disk.Height))
}

// Intersect finds ray-shape collision point and its metadata
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/shapes/disk.cpp#L48
func (disk Disk) Intersect(r Ray, _ bool) (bool, float64, *SurfaceInteraction) {
	// Transform Ray to object space
	ray := disk.WorldToObject.ApplyR(r)

	// Compute plane intersection for disk
	// Reject disk intersections for rays parallel to the disk’s plane
	if ray.D.Z == 0 {
		return false, 0, nil
	}

	tShapeHit := (disk.Height - ray.O.Z) / ray.D.Z
	if tShapeHit <= 0 || tShapeHit >= ray.TMax {
		return false, 0, nil
	}

	// See if hit point is inside disk radii and phi_max
	pHit := ray.Apply(tShapeHit)
	dist2 := pHit.X*pHit.X + pHit.Y*pHit.Y
	if dist2 > disk.Radius*disk.Radius || dist2 < disk.InnerRadius*disk.InnerRadius {
		return false, 0, nil
	}

	// Test disk phi value against phi_max
	phi := math.Atan2(pHit.Y, pHit.X)
	if phi < 0 {
		phi += 2 * math.Pi
	}

	if phi > disk.PhiMax {
		return false, 0, nil
	}

	// Find parametric representation of disk hit
	u := phi / disk.PhiMax
	rHit := math.Sqrt(dist2)
	oneMinusV := (rHit - disk.InnerRadius) / (disk.Radius - disk.InnerRadius)
	v := 1 - oneMinusV
	dpdu := NewVector3(-disk.PhiMax*pHit.Y, disk.PhiMax*pHit.X, 0)
	dpdv := NewVector3(pHit.X, pHit.Y, 0).Multiply((disk.InnerRadius - disk.Radius) / rHit)
	dndu := NewNormal3(0, 0, 0)
	dndv := NewNormal3(0, 0, 0)

	// Refine disk intersection point
	pHit.Z = disk.Height

	// Compute error bounds for disk intersection
	pError := NewVector3(0, 0, 0)

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
		&disk.Shape)

	isect := disk.ObjectToWorld.ApplySI(&si)

	// Update _tHit_ for quadric intersection
	return true, tShapeHit, isect
}

// IntersectP finds if ray collides with this shape
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/shapes/disk.cpp#L98
func (disk Disk) IntersectP(_ Intersecter, r Ray, _ bool) bool {
	// Transform Ray to object space
	ray := disk.WorldToObject.ApplyR(r)

	// Compute plane intersection for disk
	// Reject disk intersections for rays parallel to the disk’s plane
	if ray.D.Z == 0 {
		return false
	}

	tShapeHit := (disk.Height - ray.O.Z) / ray.D.Z
	if tShapeHit <= 0 || tShapeHit >= ray.TMax {
		return false
	}

	// See if hit point is inside disk radii and phi_max
	pHit := ray.Apply(tShapeHit)
	dist2 := pHit.X*pHit.X + pHit.Y*pHit.Y
	if dist2 > disk.Radius*disk.Radius || dist2 < disk.InnerRadius*disk.InnerRadius {
		return false
	}

	// Test disk phi value against phi_max
	phi := math.Atan2(pHit.Y, pHit.X)
	if phi < 0 {
		phi += 2 * math.Pi
	}

	if phi > disk.PhiMax {
		return false
	}

	return true
}

func (disk Disk) Area() float64 {
	return disk.PhiMax * 0.5 * (disk.Radius*disk.Radius - disk.InnerRadius*disk.InnerRadius)
}
