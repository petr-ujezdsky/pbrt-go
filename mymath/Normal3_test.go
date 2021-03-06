package mymath_test

import (
	"pbrt-go/mymath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormal3_NewNormal3(t *testing.T) {
	v := mymath.NewNormal3(1, 2, 3)

	assert.Equal(t, 1.0, v.X)
	assert.Equal(t, 2.0, v.Y)
	assert.Equal(t, 3.0, v.Z)
}

func TestNormal3_LengthSq(t *testing.T) {
	v := mymath.NewNormal3(3, 3, 3)

	assert.Equal(t, 27.0, v.LengthSq())
}

func TestNormal3_Length(t *testing.T) {
	v := mymath.NewNormal3(1, 2, 3)

	assert.Equal(t, 3.7416573867739413, v.Length())
}
func TestNormal3_Add(t *testing.T) {
	v := mymath.NewNormal3(1, 2, 3)
	w := mymath.NewNormal3(5, 6, 7)

	res := v.Add(w)

	assert.Equal(t, 6.0, res.X)
	assert.Equal(t, 8.0, res.Y)
	assert.Equal(t, 10.0, res.Z)
}
func TestNormal3_Subtract(t *testing.T) {
	v := mymath.NewNormal3(1, 2, 3)
	w := mymath.NewNormal3(5, 6, 7)

	res := v.Subtract(w)

	assert.Equal(t, -4.0, res.X)
	assert.Equal(t, -4.0, res.Y)
	assert.Equal(t, -4.0, res.Z)
}

func TestNormal3_Multiply(t *testing.T) {
	v := mymath.NewNormal3(1, 2, 3)

	res := v.Multiply(5.0)

	assert.Equal(t, 5.0, res.X)
	assert.Equal(t, 10.0, res.Y)
	assert.Equal(t, 15.0, res.Z)
}

func TestNormal3_Divide(t *testing.T) {
	v := mymath.NewNormal3(2, 4, -8)

	res := v.Divide(2.0)

	assert.Equal(t, 1.0, res.X)
	assert.Equal(t, 2.0, res.Y)
	assert.Equal(t, -4.0, res.Z)
}
func TestNormal3_Negate(t *testing.T) {
	v := mymath.NewNormal3(2, 4, -8)

	res := v.Negate()

	assert.Equal(t, -2.0, res.X)
	assert.Equal(t, -4.0, res.Y)
	assert.Equal(t, 8.0, res.Z)
}

// func TestNormal3_Abs(t *testing.T) {
// 	v := mymath.NewNormal3(2, 4, -8)

// 	res := v.Abs()

// 	assert.Equal(t, 2.0, res.X)
// 	assert.Equal(t, 4.0, res.Y)
// 	assert.Equal(t, 8.0, res.Z)
// }

// func TestNormal3_Normalize(t *testing.T) {
// 	v := mymath.NewNormal3(1, 2, 3)

// 	res := v.Normalize()

// 	assert.Equal(t, 0.2672612419124244, res.X)
// 	assert.Equal(t, 0.5345224838248488, res.Y)
// 	assert.Equal(t, 0.8017837257372732, res.Z)
// }

// func TestNormal3_Dot(t *testing.T) {
// 	v := mymath.NewNormal3(-1, -2, -3)
// 	w := mymath.NewNormal3(4, 5, 6)

// 	res := v.Dot(w)

// 	assert.Equal(t, -32.0, res)
// }

// func TestNormal3_Cross(t *testing.T) {
// 	v := mymath.NewNormal3(1, 2, 3)
// 	w := mymath.NewNormal3(4, 5, 6)

// 	res := v.Cross(w)

// 	assert.Equal(t, -3.0, res.X)
// 	assert.Equal(t, 6.0, res.Y)
// 	assert.Equal(t, -3.0, res.Z)
// }

func TestNormal3_Get(t *testing.T) {
	v := mymath.NewNormal3(1, 2, 3)

	assert.Equal(t, 1.0, v.Get(0))
	assert.Equal(t, 2.0, v.Get(1))
	assert.Equal(t, 3.0, v.Get(2))
}

// func TestNormal3_GetMinComponent(t *testing.T) {
// 	assert.Equal(t, 1.0, mymath.NewNormal3(1, 2, 3).GetMinComponent())
// 	assert.Equal(t, 2.0, mymath.NewNormal3(4, 2, 3).GetMinComponent())
// 	assert.Equal(t, 3.0, mymath.NewNormal3(5, 5, 3).GetMinComponent())
// }

// func TestNormal3_GetMaxComponent(t *testing.T) {
// 	assert.Equal(t, 3.0, mymath.NewNormal3(1, 2, 3).GetMaxComponent())
// 	assert.Equal(t, 4.0, mymath.NewNormal3(4, 2, 3).GetMaxComponent())
// 	assert.Equal(t, 5.0, mymath.NewNormal3(5, 5, 3).GetMaxComponent())
// }

// func TestNormal3_GetMaxDimension(t *testing.T) {
// 	assert.Equal(t, 2, mymath.NewNormal3(1, 2, 3).GetMaxDimension())
// 	assert.Equal(t, 2, mymath.NewNormal3(2, 1, 3).GetMaxDimension())
// 	assert.Equal(t, 0, mymath.NewNormal3(4, 2, 3).GetMaxDimension())
// 	assert.Equal(t, 1, mymath.NewNormal3(4, 5, 3).GetMaxDimension())
// 	assert.Equal(t, 1, mymath.NewNormal3(5, 5, 3).GetMaxDimension())
// }

// func TestNormal3_Min(t *testing.T) {
// 	v := mymath.NewNormal3(1, 8, 3)
// 	w := mymath.NewNormal3(4, 5, 6)

// 	res := v.Min(w)

// 	assert.Equal(t, 1.0, res.X)
// 	assert.Equal(t, 5.0, res.Y)
// 	assert.Equal(t, 3.0, res.Z)
// }

// func TestNormal3_Max(t *testing.T) {
// 	v := mymath.NewNormal3(1, 8, 3)
// 	w := mymath.NewNormal3(4, 5, 6)

// 	res := v.Max(w)

// 	assert.Equal(t, 4.0, res.X)
// 	assert.Equal(t, 8.0, res.Y)
// 	assert.Equal(t, 6.0, res.Z)
// }
// func TestNormal3_Permute(t *testing.T) {
// 	v := mymath.NewNormal3(1, 8, 3)

// 	res := v.Permute(2, 0, 1)

// 	assert.Equal(t, 3.0, res.X)
// 	assert.Equal(t, 1.0, res.Y)
// 	assert.Equal(t, 8.0, res.Z)
// }

//////////////////////////////////////////////////////////////////////////////////////////////////
// BENCHMARKS ////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////
func BenchmarkNormal3_NewNormal3(b *testing.B) {
	var res mymath.Normal3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = mymath.NewNormal3(1, 2, 3)
	}

	assert.NotNil(b, res)
}
func BenchmarkNormal3_LengthSq(b *testing.B) {
	v := mymath.NewNormal3(1, 2, 3)

	var res float64

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = v.LengthSq()
	}

	assert.NotNil(b, res)
}

func BenchmarkNormal3_Length(b *testing.B) {
	v := mymath.NewNormal3(1, 2, 3)

	var res float64

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = v.Length()
	}

	assert.NotNil(b, res)
}

func BenchmarkNormal3_Add(b *testing.B) {
	v := mymath.NewNormal3(1, 2, 3)

	res := v

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Add(v)
	}

	assert.NotNil(b, res)
}

func BenchmarkNormal3_Subtract(b *testing.B) {
	v := mymath.NewNormal3(1, 2, 3)

	res := v

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Subtract(v)
	}

	assert.NotNil(b, res)
}

func BenchmarkNormal3_Multiply(b *testing.B) {
	v := mymath.NewNormal3(1, 2, 3)

	res := v

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Multiply(1.0000003)
	}

	assert.NotNil(b, res)
}

func BenchmarkNormal3_Divide(b *testing.B) {
	v := mymath.NewNormal3(1, 2, 3)

	res := v

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Divide(1.0000003)
	}

	assert.NotNil(b, res)
}

func BenchmarkNormal3_Negate(b *testing.B) {
	v := mymath.NewNormal3(1, 2, 3)

	res := v

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = res.Negate()
	}

	assert.NotNil(b, res)
}

// func BenchmarkNormal3_Abs(b *testing.B) {
// 	v := mymath.NewNormal3(1, 2, 3)

// 	res := v

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		res = res.Abs()
// 	}

// 	assert.NotNil(b, res)
// }

// func BenchmarkNormal3_Normalize(b *testing.B) {
// 	v := mymath.NewNormal3(1, 2, 3)

// 	res := v

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		res = res.Normalize()
// 	}

// 	assert.NotNil(b, res)
// }

// func BenchmarkNormal3_Dot(b *testing.B) {
// 	v1 := mymath.NewNormal3(1, 2, 3)
// 	v2 := mymath.NewNormal3(4, 5, 6)

// 	var res float64

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		res = v1.Dot(v2)
// 	}

// 	assert.NotNil(b, res)
// }

// func BenchmarkNormal3_Cross(b *testing.B) {
// 	v := mymath.NewNormal3(1, 2, 3)

// 	res := v

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		res = res.Cross(v)
// 	}

// 	assert.NotNil(b, res)
// }

func BenchmarkNormal3_Get(b *testing.B) {
	v := mymath.NewNormal3(1, 2, 3)

	var res float64

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = v.Get(0)
	}

	assert.NotNil(b, res)
}

// func BenchmarkNormal3_GetMinComponent(b *testing.B) {
// 	v := mymath.NewNormal3(1, 2, 3)

// 	var res float64

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		res = v.GetMinComponent()
// 	}

// 	assert.NotNil(b, res)
// }

// func BenchmarkNormal3_GetMaxComponent(b *testing.B) {
// 	v := mymath.NewNormal3(1, 2, 3)

// 	var res float64

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		res = v.GetMaxComponent()
// 	}

// 	assert.NotNil(b, res)
// }

// func BenchmarkNormal3_GetMaxDimension(b *testing.B) {
// 	v := mymath.NewNormal3(1, 2, 3)

// 	var res float64

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		res = v.GetMaxComponent()
// 	}

// 	assert.NotNil(b, res)
// }

// func BenchmarkNormal3_Min(b *testing.B) {
// 	v := mymath.NewNormal3(1, 2, 3)

// 	res := v

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		res = v.Min(res)
// 	}

// 	assert.NotNil(b, res)
// }

// func BenchmarkNormal3_Max(b *testing.B) {
// 	v := mymath.NewNormal3(1, 2, 3)

// 	res := v

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		res = v.Max(res)
// 	}

// 	assert.NotNil(b, res)
// }

// func BenchmarkNormal3_Permute(b *testing.B) {
// 	v := mymath.NewNormal3(1, 2, 3)

// 	res := v

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		res = res.Permute(2, 0, 1)
// 	}

// 	assert.NotNil(b, res)
// }
