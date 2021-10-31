package mymath

type BoundingSphere struct {
	Center Point3d

	Radius float64
}

func NewBoundingSphere(center Point3d, radius float64) BoundingSphere {
	return BoundingSphere{center, radius}
}
