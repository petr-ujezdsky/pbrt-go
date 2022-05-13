package shape

import "pbrt-go/mymath"

// Shape see https://github.com/mmp/pbrt-v3/blob/master/src/core/shape.h, https://github.com/mmp/pbrt-v3/blob/master/src/core/shape.cpp
type Shape struct {
	ObjectToWorld, WorldToObject                 *mymath.Transform
	ReverseOrientation, TransformSwapsHandedness bool
}

type ObjectBounder interface {
	ObjectBound() mymath.Bounds3
}

type WorldBounder interface {
	WorldBound(ob ObjectBounder) mymath.Bounds3
}

type Intersecter interface {
	Intersect(ray mymath.Ray, testAlphaTexture bool) (bool, float64, mymath.SurfaceInteraction)
}

type IntersectPer interface {
	IntersectP(i Intersecter, ray mymath.Ray, testAlphaTexture bool) bool
}

type Areaer interface {
	Area() float64
}

type IShape interface {
	ObjectBounder
	WorldBounder
	Intersecter
	IntersectPer
	Areaer
}

func NewShape(objectToWorld, worldToObject *mymath.Transform, reverseOrientation bool) *Shape {
	return &Shape{
		objectToWorld,
		worldToObject,
		reverseOrientation,
		objectToWorld.SwapsHandedness(),
	}
}

func (s Shape) WorldBound(ob ObjectBounder) mymath.Bounds3 {
	return s.ObjectToWorld.ApplyB(ob.ObjectBound())
}

func (s Shape) IntersectP(i Intersecter, ray mymath.Ray, testAlphaTexture bool) bool {
	intersects, _, _ := i.Intersect(ray, testAlphaTexture)
	return intersects
}
