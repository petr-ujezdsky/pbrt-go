package mymath_test

import (
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint2_NewPoint2(t *testing.T) {
	p := mymath.NewPoint2(1, 2)

	assert.Equal(t, 1.0, p.X)
	assert.Equal(t, 2.0, p.Y)
}

//func TestPoint2_AddP(t *testing.T) {
//	p1 := mymath.NewPoint3(1, 2, 3)
//	p2 := mymath.NewPoint3(5, 6, 7)
//
//	res := p1.AddP(p2)
//
//	assert.Equal(t, 6.0, res.X)
//	assert.Equal(t, 8.0, res.Y)
//	assert.Equal(t, 10.0, res.Z)
//}
//
//func TestPoint2_AddV(t *testing.T) {
//	p := mymath.NewPoint3(1, 2, 3)
//	v := mymath.NewVector3(5, 6, 7)
//
//	res := p.AddV(v)
//
//	assert.Equal(t, 6.0, res.X)
//	assert.Equal(t, 8.0, res.Y)
//	assert.Equal(t, 10.0, res.Z)
//}
//
//func TestPoint2_SubtractP(t *testing.T) {
//	p1 := mymath.NewPoint3(1, 2, 3)
//	p2 := mymath.NewPoint3(5, 6, 7)
//
//	res := p1.SubtractP(p2)
//
//	assert.Equal(t, -4.0, res.X)
//	assert.Equal(t, -4.0, res.Y)
//	assert.Equal(t, -4.0, res.Z)
//}
//
//func TestPoint2_SubtractV(t *testing.T) {
//	p := mymath.NewPoint3(1, 2, 3)
//	v := mymath.NewVector3(5, 6, 7)
//
//	res := p.SubtractV(v)
//
//	assert.Equal(t, -4.0, res.X)
//	assert.Equal(t, -4.0, res.Y)
//	assert.Equal(t, -4.0, res.Z)
//}
//
//func TestPoint2_Multiply(t *testing.T) {
//	v := mymath.NewPoint3(1, 2, 3)
//
//	res := v.Multiply(5.0)
//
//	assert.Equal(t, 5.0, res.X)
//	assert.Equal(t, 10.0, res.Y)
//	assert.Equal(t, 15.0, res.Z)
//}
//
//func TestPoint2_Distance(t *testing.T) {
//	p1 := mymath.NewPoint3(1, 2, 3)
//	p2 := mymath.NewPoint3(5, 6, 7)
//
//	res := p1.Distance(p2)
//
//	assert.Equal(t, 6.928203230275509, res)
//}
//
//func TestPoint2_DistanceSq(t *testing.T) {
//	p1 := mymath.NewPoint3(1, 2, 3)
//	p2 := mymath.NewPoint3(5, 6, 7)
//
//	res := p1.DistanceSq(p2)
//
//	assert.Equal(t, 48.0, res)
//}
//
//func TestPoint2_Lerp(t *testing.T) {
//	p := mymath.NewPoint3(1, 2, 3)
//
//	res := p.Lerp(0.5, mymath.NewPoint3(5, 6, 7))
//
//	assert.Equal(t, 3.0, res.X)
//	assert.Equal(t, 4.0, res.Y)
//	assert.Equal(t, 5.0, res.Z)
//}
//
//func TestPoint2_LerpP(t *testing.T) {
//	p := mymath.NewPoint3(1, 2, 3)
//
//	res := p.LerpP(mymath.NewPoint3(0, 0.5, 1), mymath.NewPoint3(5, 6, 7))
//
//	assert.Equal(t, 1.0, res.X)
//	assert.Equal(t, 4.0, res.Y)
//	assert.Equal(t, 7.0, res.Z)
//}
//
//func TestPoint2_Min(t *testing.T) {
//	v := mymath.NewPoint3(1, 8, 3)
//	w := mymath.NewPoint3(4, 5, 6)
//
//	res := v.Min(w)
//
//	assert.Equal(t, 1.0, res.X)
//	assert.Equal(t, 5.0, res.Y)
//	assert.Equal(t, 3.0, res.Z)
//}
//
//func TestPoint2_Max(t *testing.T) {
//	v := mymath.NewPoint3(1, 8, 3)
//	w := mymath.NewPoint3(4, 5, 6)
//
//	res := v.Max(w)
//
//	assert.Equal(t, 4.0, res.X)
//	assert.Equal(t, 8.0, res.Y)
//	assert.Equal(t, 6.0, res.Z)
//}
//
//func TestPoint2_Floor(t *testing.T) {
//	v := mymath.NewPoint3(1.8, 8.9, 3.1)
//
//	res := v.Floor()
//
//	assert.Equal(t, 1.0, res.X)
//	assert.Equal(t, 8.0, res.Y)
//	assert.Equal(t, 3.0, res.Z)
//}
//
//func TestPoint2_Ceil(t *testing.T) {
//	v := mymath.NewPoint3(1.8, 8.9, 3.1)
//
//	res := v.Ceil()
//
//	assert.Equal(t, 2.0, res.X)
//	assert.Equal(t, 9.0, res.Y)
//	assert.Equal(t, 4.0, res.Z)
//}
//
//func TestPoint2_Abs(t *testing.T) {
//	v := mymath.NewPoint3(-1, -2, -3)
//
//	res := v.Abs()
//
//	assert.Equal(t, 1.0, res.X)
//	assert.Equal(t, 2.0, res.Y)
//	assert.Equal(t, 3.0, res.Z)
//}
//
//func TestPoint2_Get(t *testing.T) {
//	v := mymath.NewPoint3(1, 2, 3)
//
//	assert.Equal(t, 1.0, v.Get(0))
//	assert.Equal(t, 2.0, v.Get(1))
//	assert.Equal(t, 3.0, v.Get(2))
//}
//
//func TestPoint2_Permute(t *testing.T) {
//	v := mymath.NewPoint3(1, 8, 3)
//
//	res := v.Permute(2, 0, 1)
//
//	assert.Equal(t, 3.0, res.X)
//	assert.Equal(t, 1.0, res.Y)
//	assert.Equal(t, 8.0, res.Z)
//}

//////////////////////////////////////////////////////////////////////////////////////////////////
// BENCHMARKS ////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
func BenchmarkPoint2_NewPoint3(b *testing.B) {
	var res mymath.Point2

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = mymath.NewPoint2(1.0, 2.0)
	}

	assert.NotNil(b, res)
}

//func BenchmarkPoint2_AddP(b *testing.B) {
//	p := mymath.NewPoint3(1, 2, 3)
//
//	res := p
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = res.AddP(p)
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_AddV(b *testing.B) {
//	p := mymath.NewPoint3(1, 2, 3)
//	v := mymath.NewVector3(4, 5, 6)
//
//	res := p
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = res.AddV(v)
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_SubtractP(b *testing.B) {
//	p1 := mymath.NewPoint3(1, 2, 3)
//	p2 := mymath.NewPoint3(4, 5, 6)
//
//	var res mymath.Vector3
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = p1.SubtractP(p2)
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_SubtractV(b *testing.B) {
//	p := mymath.NewPoint3(1, 2, 3)
//	v := mymath.NewVector3(4, 5, 6)
//
//	res := p
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = res.SubtractV(v)
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_Multiply(b *testing.B) {
//	p := mymath.NewPoint3(1, 2, 3)
//
//	res := p
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = res.Multiply(1.0000003)
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_Distance(b *testing.B) {
//	p1 := mymath.NewPoint3(1, 2, 3)
//	p2 := mymath.NewPoint3(4, 5, 6)
//
//	var res float64
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = p1.Distance(p2)
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_DistanceSq(b *testing.B) {
//	p1 := mymath.NewPoint3(1, 2, 3)
//	p2 := mymath.NewPoint3(4, 5, 6)
//
//	var res float64
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = p1.DistanceSq(p2)
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_Lerp(b *testing.B) {
//	p1 := mymath.NewPoint3(1, 2, 3)
//	p2 := mymath.NewPoint3(4, 5, 6)
//
//	res := p1
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = p1.Lerp(0.9, p2)
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_LerpP(b *testing.B) {
//	t := mymath.NewPoint3(0.1, 0.9, 0.5)
//	p1 := mymath.NewPoint3(1, 2, 3)
//	p2 := mymath.NewPoint3(4, 5, 6)
//
//	res := p1
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = p1.LerpP(t, p2)
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_Min(b *testing.B) {
//	p := mymath.NewPoint3(1, 2, 3)
//
//	res := p
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = res.Min(p)
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_Max(b *testing.B) {
//	p := mymath.NewPoint3(1, 2, 3)
//
//	res := p
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = res.Max(p)
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_Floor(b *testing.B) {
//	p := mymath.NewPoint3(1, 2, 3)
//
//	res := p
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = res.Floor()
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_Ceil(b *testing.B) {
//	p := mymath.NewPoint3(1, 2, 3)
//
//	res := p
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = res.Ceil()
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_Abs(b *testing.B) {
//	p := mymath.NewPoint3(1, 2, 3)
//
//	res := p
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = res.Abs()
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_Get(b *testing.B) {
//	p := mymath.NewPoint3(1, 2, 3)
//
//	var res float64
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = p.Get(1)
//	}
//
//	assert.NotNil(b, res)
//}
//
//func BenchmarkPoint2_Permute(b *testing.B) {
//	p := mymath.NewPoint3(1, 2, 3)
//
//	res := p
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		res = p.Permute(2, 0, 1)
//	}
//
//	assert.NotNil(b, res)
//}
