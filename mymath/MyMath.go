package mymath

func Lerp(t, v1, v2 float64) float64 {
	return (1-t)*v1 + t*v2
}