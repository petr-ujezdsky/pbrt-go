package mymath

type BoundingSphere struct {
	Center Point3

	Radius float64
}

func NewBoundingSphere(center Point3, radius float64) BoundingSphere {
	return BoundingSphere{center, radius}
}
