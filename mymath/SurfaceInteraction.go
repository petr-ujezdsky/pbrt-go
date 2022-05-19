package mymath

// SurfaceInteraction describes local metadata for ray-shape collision point
//
// see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/interaction.cpp
type SurfaceInteraction struct {
	Interaction
	Uv         Point2
	Dpdu, Dpdv Vector3
	Dndu, Dndv Normal3
	Shape      *Shape
	shading    shading
}

type shading struct {
	N          Normal3
	Dpdu, Dpdv Vector3
	Dndu, Dndv Normal3
}

func NewSurfaceInteraction(p Point3, pError Vector3, uv Point2, wo Vector3, dpdu, dpdv Vector3, dndu, dndv Normal3, time float64, shape *Shape) SurfaceInteraction {
	n := NewNormal3V(dpdu.Cross(dpdv).Normalize())

	// Adjust normal based on orientation and handedness
	if shape != nil && (shape.ReverseOrientation != shape.TransformSwapsHandedness) {
		n = n.Negate()
	}

	surfaceInteraction := SurfaceInteraction{
		NewInteraction(p, n, pError, wo, time, nil),
		uv,
		dpdu,
		dpdv,
		dndu,
		dndv,
		shape,
		// Initialize shading geometry from true geometry
		shading{n, dpdu, dpdv, dndu, dndv},
	}

	return surfaceInteraction
}

func (si *SurfaceInteraction) SetShadingGeometry(dpdus, dpdvs Vector3, dndus, dndvs Normal3, orientationIsAuthoritative bool) {
	// Compute shading.n for SurfaceInteraction
	n := NewNormal3V(dpdus.Cross(dpdvs).Normalize())

	// Adjust normal based on orientation and handedness
	if si.Shape != nil && (si.Shape.ReverseOrientation != si.Shape.TransformSwapsHandedness) {
		n = n.Negate()
	}

	si.shading.N = n

	if orientationIsAuthoritative {
		si.N = si.N.FaceForward(si.shading.N)
	} else {
		si.shading.N = si.shading.N.FaceForward(si.N)
	}

	// Initialize shading partial derivative values
	si.shading.Dpdu = dpdus
	si.shading.Dpdv = dpdvs
	si.shading.Dndu = dndus
	si.shading.Dndv = dndvs
}
