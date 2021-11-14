package mymath

import "math"

var epsilon = math.Nextafter(1, 2) - 1

var Gamma3 = gamma(3)

func Lerp(t, v1, v2 float64) float64 {
	return (1-t)*v1 + t*v2
}

func gamma(n int) float64 {
	ne := float64(n) * epsilon
	return ne / (1 - ne)
}
