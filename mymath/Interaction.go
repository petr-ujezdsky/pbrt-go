package mymath

import "pbrt-go/material"

type Interaction struct {
	P               Point3
	Time            float64
	PError, Wo      Vector3
	N               Normal3
	MediumInterface *material.MediumInterface
}

func NewInteraction(p Point3, n Normal3, pError, wo Vector3, time float64, mediumInterface *material.MediumInterface) Interaction {
	return Interaction{P: p, N: n, PError: pError, Wo: wo, Time: time, MediumInterface: mediumInterface}
}

func (i Interaction) IsSurfaceInteraction() bool {
	return i.N != NewNormal3(0, 0, 0)
}
