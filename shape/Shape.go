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
