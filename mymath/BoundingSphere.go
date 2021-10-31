package mymath

import (
	"pbrt-go/mymath/point3d"
)

type BoundingSphere struct {
	Center point3d.Point3d

	Radius float64
}

func NewBoundingSphere(center point3d.Point3d, radius float64) BoundingSphere {
	return BoundingSphere{center, radius}
}
