package mymath

type RayDifferential struct {
	Ray
	HasDifferentials bool
	RxOrigin         Point3
	RyOrigin         Point3
	RxDirection      Vector3
	RyDirection      Vector3
}

func NewRayDifferentialRay(ray Ray) RayDifferential {
	return RayDifferential{Ray: ray}
}

func (rd *RayDifferential) ScaleDifferentials(s float64) {
	rd.RxOrigin = rd.O.AddV(rd.RxOrigin.SubtractP(rd.O).Multiply(s))
	rd.RyOrigin = rd.O.AddV(rd.RyOrigin.SubtractP(rd.O).Multiply(s))
	rd.RxDirection = rd.D.Add(rd.RxDirection.Subtract(rd.D).Multiply(s))
	rd.RyDirection = rd.D.Add(rd.RyDirection.Subtract(rd.D).Multiply(s))
}
