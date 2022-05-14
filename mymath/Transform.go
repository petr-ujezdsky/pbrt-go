package mymath

import (
	"fmt"
	"math"
)

type Transform struct {
	M, MInv Matrix4x4
}

func NewTransformEmpty() Transform {
	m := Identity()

	return Transform{m, m}
}

func NewTransform(m Matrix4x4) (Transform, error) {
	mInv, err := m.Inverse()
	if err != nil {
		return Transform{}, err
	}

	return Transform{m, mInv}, nil
}

func NewTransformFull(m Matrix4x4, mInv Matrix4x4) Transform {
	return Transform{m, mInv}
}

func NewTransformTranslate(delta Vector3) Transform {
	return Transform{
		NewMatrix4x4All(
			1, 0, 0, float32(delta.X),
			0, 1, 0, float32(delta.Y),
			0, 0, 1, float32(delta.Z),
			0, 0, 0, 1),
		NewMatrix4x4All(
			1, 0, 0, -float32(delta.X),
			0, 1, 0, -float32(delta.Y),
			0, 0, 1, -float32(delta.Z),
			0, 0, 0, 1),
	}
}

func NewTransformScale(x, y, z float32) Transform {
	return Transform{
		NewMatrix4x4All(
			x, 0, 0, 0,
			0, y, 0, 0,
			0, 0, z, 0,
			0, 0, 0, 1),
		NewMatrix4x4All(
			1/x, 0, 0, 0,
			0, 1/y, 0, 0,
			0, 0, 1/z, 0,
			0, 0, 0, 1),
	}
}

func NewTransformRotateX(theta float32) Transform {
	sinTheta := float32(math.Sin(float64(theta)))
	cosTheta := float32(math.Cos(float64(theta)))

	m := NewMatrix4x4All(
		1, 0, 0, 0,
		0, cosTheta, -sinTheta, 0,
		0, sinTheta, cosTheta, 0,
		0, 0, 0, 1)

	return Transform{m, m.Transpose()}
}

func NewTransformRotateY(theta float32) Transform {
	sinTheta := float32(math.Sin(float64(theta)))
	cosTheta := float32(math.Cos(float64(theta)))

	m := NewMatrix4x4All(
		cosTheta, 0, sinTheta, 0,
		0, 1, 0, 0,
		-sinTheta, 0, cosTheta, 0,
		0, 0, 0, 1)

	return Transform{m, m.Transpose()}
}

func NewTransformRotateZ(theta float32) Transform {
	sinTheta := float32(math.Sin(float64(theta)))
	cosTheta := float32(math.Cos(float64(theta)))

	m := NewMatrix4x4All(
		cosTheta, -sinTheta, 0, 0,
		sinTheta, cosTheta, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1)

	return Transform{m, m.Transpose()}
}

func NewTransformRotate(theta float64, axis Vector3) Transform {
	a := axis.Normalize()

	sinTheta := math.Sin(theta)
	cosTheta := math.Cos(theta)

	m := Matrix4x4{}

	// Compute rotation of first basis vector
	m.M[0][0] = float32(a.X*a.X + (1-a.X*a.X)*cosTheta)
	m.M[0][1] = float32(a.X*a.Y*(1-cosTheta) - a.Z*sinTheta)
	m.M[0][2] = float32(a.X*a.Z*(1-cosTheta) + a.Y*sinTheta)
	m.M[0][3] = 0

	// Compute rotations of second and third basis vectors
	m.M[1][0] = float32(a.X*a.Y*(1-cosTheta) + a.Z*sinTheta)
	m.M[1][1] = float32(a.Y*a.Y + (1-a.Y*a.Y)*cosTheta)
	m.M[1][2] = float32(a.Y*a.Z*(1-cosTheta) - a.X*sinTheta)
	m.M[1][3] = 0

	m.M[2][0] = float32(a.X*a.Z*(1-cosTheta) - a.Y*sinTheta)
	m.M[2][1] = float32(a.Y*a.Z*(1-cosTheta) + a.X*sinTheta)
	m.M[2][2] = float32(a.Z*a.Z + (1-a.Z*a.Z)*cosTheta)
	m.M[2][3] = 0

	m.M[3][3] = 1

	return Transform{m, m.Transpose()}
}

func NewTransformLookAt(pos, look Point3, up Vector3) (Transform, error) {
	cameraToWorld := Matrix4x4{}

	// Initialize fourth column of viewing matrix
	cameraToWorld.M[0][3] = float32(pos.X)
	cameraToWorld.M[1][3] = float32(pos.Y)
	cameraToWorld.M[2][3] = float32(pos.Z)
	cameraToWorld.M[3][3] = 1

	// Initialize first three columns of viewing matrix
	dir := look.SubtractP(pos).Normalize()
	rightNotNormalized := up.Normalize().Cross(dir)
	if rightNotNormalized.Length() == 0 {
		err := fmt.Errorf(
			"\"up\" vector (%v) and viewing direction (%v) "+
				"passed to LookAt are pointing in the same direction.  Using "+
				"the identity transformation",
			up, dir)
		return NewTransformEmpty(), err
	}

	right := rightNotNormalized.Normalize()
	newUp := dir.Cross(right)

	cameraToWorld.M[0][0] = float32(right.X)
	cameraToWorld.M[1][0] = float32(right.Y)
	cameraToWorld.M[2][0] = float32(right.Z)
	cameraToWorld.M[3][0] = 0
	cameraToWorld.M[0][1] = float32(newUp.X)
	cameraToWorld.M[1][1] = float32(newUp.Y)
	cameraToWorld.M[2][1] = float32(newUp.Z)
	cameraToWorld.M[3][1] = 0
	cameraToWorld.M[0][2] = float32(dir.X)
	cameraToWorld.M[1][2] = float32(dir.Y)
	cameraToWorld.M[2][2] = float32(dir.Z)
	cameraToWorld.M[3][2] = 0

	ctwInv, err := cameraToWorld.Inverse()
	if err != nil {
		return NewTransformEmpty(), err
	}

	return NewTransformFull(ctwInv, cameraToWorld), nil
}

// Applies transformation to Point
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.h#L222
func (t Transform) ApplyP(p Point3) Point3 {
	return t.M.MultiplyP(p)
}

// Applies transformation to Point, also returns error vector
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.h#L278
func (t Transform) ApplyPError(p Point3) (Point3, Vector3) {
	pt := t.M.MultiplyP(p)

	// Compute absolute error for transformed point
	xAbsSum := math.Abs(float64(t.M.M[0][0])*p.X) +
		math.Abs(float64(t.M.M[0][1])*p.Y) +
		math.Abs(float64(t.M.M[0][2])*p.Z) +
		math.Abs(float64(t.M.M[0][3]))

	yAbsSum := math.Abs(float64(t.M.M[1][0])*p.X) +
		math.Abs(float64(t.M.M[1][1])*p.Y) +
		math.Abs(float64(t.M.M[1][2])*p.Z) +
		math.Abs(float64(t.M.M[1][3]))

	zAbsSum := math.Abs(float64(t.M.M[2][0])*p.X) +
		math.Abs(float64(t.M.M[2][1])*p.Y) +
		math.Abs(float64(t.M.M[2][2])*p.Z) +
		math.Abs(float64(t.M.M[2][3]))

	pError := NewVector3(xAbsSum, yAbsSum, zAbsSum).Multiply(Gamma3)

	return pt, pError
}

// Applies transformation to Point, also returns error vector
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.h#L303
func (t Transform) ApplyPPError(p Point3, pError Vector3) (Point3, Vector3) {
	pt := t.M.MultiplyP(p)

	xAbsSum := (Gamma3+1)*(math.Abs(float64(t.M.M[0][0]))*pError.X+
		math.Abs(float64(t.M.M[0][1]))*pError.Y+
		math.Abs(float64(t.M.M[0][2])*pError.Z)) +
		Gamma3*(math.Abs(float64(t.M.M[0][0])*p.X)+
			math.Abs(float64(t.M.M[0][1])*p.Y)+
			math.Abs(float64(t.M.M[0][2])*p.Z)+
			math.Abs(float64(t.M.M[0][3])))

	yAbsSum := (Gamma3+1)*(math.Abs(float64(t.M.M[1][0]))*pError.X+
		math.Abs(float64(t.M.M[1][1]))*pError.Y+
		math.Abs(float64(t.M.M[1][2])*pError.Z)) +
		Gamma3*(math.Abs(float64(t.M.M[1][0])*p.X)+
			math.Abs(float64(t.M.M[1][1])*p.Y)+
			math.Abs(float64(t.M.M[1][2])*p.Z)+
			math.Abs(float64(t.M.M[1][3])))

	zAbsSum := (Gamma3+1)*(math.Abs(float64(t.M.M[2][0]))*pError.X+
		math.Abs(float64(t.M.M[2][1]))*pError.Y+
		math.Abs(float64(t.M.M[2][2])*pError.Z)) +
		Gamma3*(math.Abs(float64(t.M.M[2][0])*p.X)+
			math.Abs(float64(t.M.M[2][1])*p.Y)+
			math.Abs(float64(t.M.M[2][2])*p.Z)+
			math.Abs(float64(t.M.M[2][3])))

	absError := NewVector3(xAbsSum, yAbsSum, zAbsSum)

	return pt, absError
}

// ApplyV applies transformation to Vector
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.h#L236
func (t Transform) ApplyV(v Vector3) Vector3 {
	return t.M.MultiplyV(v)
}

// ApplyVError applies transformation to Vector, also returns error vector
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.h#L337
func (t Transform) ApplyVError(v Vector3) (Vector3, Vector3) {
	vt := t.M.MultiplyV(v)

	// Compute absolute error for transformed vector
	xAbsSum := math.Abs(float64(t.M.M[0][0])*v.X) +
		math.Abs(float64(t.M.M[0][1])*v.Y) +
		math.Abs(float64(t.M.M[0][2])*v.Z)

	yAbsSum := math.Abs(float64(t.M.M[1][0])*v.X) +
		math.Abs(float64(t.M.M[1][1])*v.Y) +
		math.Abs(float64(t.M.M[1][2])*v.Z)

	zAbsSum := math.Abs(float64(t.M.M[2][0])*v.X) +
		math.Abs(float64(t.M.M[2][1])*v.Y) +
		math.Abs(float64(t.M.M[2][2])*v.Z)

	vError := NewVector3(xAbsSum, yAbsSum, zAbsSum).Multiply(Gamma3)

	return vt, vError
}

// Applies transformation to Normal
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.h#L244
func (t Transform) ApplyN(n Normal3) Normal3 {
	return t.M.MultiplyN(n)
}

// Applies transformation to Ray
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.h#L251
func (t Transform) ApplyR(r Ray) Ray {
	o, oError := t.ApplyPError(r.O)
	d := t.ApplyV(r.D)

	// Offset ray origin to edge of error bounds and compute _tMax_
	lengthSquared := d.LengthSq()
	tMax := r.TMax
	if lengthSquared > 0 {
		dt := d.Abs().Dot(oError) / lengthSquared
		o = o.AddV(d.Multiply(dt))
		tMax -= dt
	}

	return NewRay(o, d, tMax, r.Time, r.Medium)
}

// ApplyRError applies transformation to Ray, also returns error vector
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.h#L382
func (t Transform) ApplyRError(r Ray) (Ray, Vector3, Vector3) {
	o, oError := t.ApplyPError(r.O)
	d, dError := t.ApplyVError(r.D)

	tMax := r.TMax
	lengthSq := d.LengthSq()

	if lengthSq > 0 {
		dt := d.Abs().Dot(oError) / lengthSq
		o = o.AddV(d.Multiply(dt))
		//        tMax -= dt;
	}

	return NewRay(o, d, tMax, r.Time, r.Medium), oError, dError
}

// Applies transformation to RayDifferential
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.h#L266
func (t Transform) ApplyRD(r RayDifferential) RayDifferential {
	tr := t.ApplyR(r.Ray)
	ret := NewRayDifferentialRay(tr)

	ret.HasDifferentials = r.HasDifferentials
	ret.RxOrigin = t.ApplyP(r.RxOrigin)
	ret.RyOrigin = t.ApplyP(r.RyOrigin)
	ret.RxDirection = t.ApplyV(r.RxDirection)
	ret.RyDirection = t.ApplyV(r.RyDirection)

	return ret
}

// Applies transformation to Bounds3
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.cpp#L238
func (t Transform) ApplyB(b Bounds3) Bounds3 {
	ret := NewBounds3P(t.ApplyP(b.PMin))

	ret = ret.UnionP(t.ApplyP(NewPoint3(b.PMax.X, b.PMin.Y, b.PMin.Z)))
	ret = ret.UnionP(t.ApplyP(NewPoint3(b.PMin.X, b.PMax.Y, b.PMin.Z)))
	ret = ret.UnionP(t.ApplyP(NewPoint3(b.PMin.X, b.PMin.Y, b.PMax.Z)))
	ret = ret.UnionP(t.ApplyP(NewPoint3(b.PMin.X, b.PMax.Y, b.PMax.Z)))
	ret = ret.UnionP(t.ApplyP(NewPoint3(b.PMax.X, b.PMax.Y, b.PMin.Z)))
	ret = ret.UnionP(t.ApplyP(NewPoint3(b.PMax.X, b.PMin.Y, b.PMax.Z)))
	ret = ret.UnionP(t.ApplyP(NewPoint3(b.PMax.X, b.PMax.Y, b.PMax.Z)))

	return ret
}

// Applies transformation to transformation
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.cpp#L251
func (t1 Transform) ApplyT(t2 Transform) Transform {
	return Transform{
		t1.M.Multiply(t2.M),
		t2.MInv.Multiply(t1.MInv)}
}

func (t1 Transform) ApplySI(si *SurfaceInteraction) *SurfaceInteraction {
	// Transform p and pError in SurfaceInteraction
	p, pError := t1.ApplyPPError(si.P, si.PError)

	// Transform remaining members of SurfaceInteraction
	return &SurfaceInteraction{
		Interaction: NewInteraction(
			p,
			t1.ApplyN(si.N).Normalize(),
			pError,
			t1.ApplyV(si.Wo),
			si.Time,
			si.MediumInterface),
		Uv:    si.Uv,
		Dpdu:  t1.ApplyV(si.Dpdu),
		Dpdv:  t1.ApplyV(si.Dpdv),
		Dndu:  t1.ApplyN(si.Dndu),
		Dndv:  t1.ApplyN(si.Dndv),
		Shape: si.Shape,
		shading: shading{
			t1.ApplyN(si.shading.N).Normalize(),
			t1.ApplyV(si.shading.Dpdu),
			t1.ApplyV(si.shading.Dpdv),
			t1.ApplyN(si.shading.Dndu),
			t1.ApplyN(si.shading.Dndv),
		},
		// TODO in another chapter
		//ret.dudx = si.dudx;
		//ret.dvdx = si.dvdx;
		//ret.dudy = si.dudy;
		//ret.dvdy = si.dvdy;
		//ret.dpdx = t(si.dpdx);
		//ret.dpdy = t(si.dpdy);
		//ret.bsdf = si.bsdf;
		//ret.bssrdf = si.bssrdf;
		//ret.primitive = si.primitive;
		////    ret.n = Faceforward(ret.n, ret.shading.n);
		//ret.shading.n = Faceforward(ret.shading.n, ret.n);
	}
}

func (t Transform) Inverse() Transform {
	return Transform{t.MInv, t.M}
}

func (t Transform) Transpose() Transform {
	return Transform{t.M.Transpose(), t.MInv.Transpose()}
}

func (t Transform) IsIdentity() bool {
	return t.M.IsIdentity()
}

func (t Transform) HasScale() bool {
	la2 := t.ApplyV(Vector3{1, 0, 0}).LengthSq()
	lb2 := t.ApplyV(Vector3{0, 1, 0}).LengthSq()
	lc2 := t.ApplyV(Vector3{0, 0, 1}).LengthSq()

	notOne := func(x float64) bool {
		return x < 0.999 || x > 1.001
	}

	return notOne(la2) || notOne(lb2) || notOne(lc2)
}

// Tells if handedness is changed by a transformation
//
// see https://github.com/mmp/pbrt-v3/blob/master/src/core/transform.cpp#L255
func (t Transform) SwapsHandedness() bool {
	// upper left 3x3 submatrix determinant
	det := t.M.M[0][0]*(t.M.M[1][1]*t.M.M[2][2]-t.M.M[1][2]*t.M.M[2][1]) -
		t.M.M[0][1]*(t.M.M[1][0]*t.M.M[2][2]-t.M.M[1][2]*t.M.M[2][0]) +
		t.M.M[0][2]*(t.M.M[1][0]*t.M.M[2][1]-t.M.M[1][1]*t.M.M[2][0])

	return det < 0
}
