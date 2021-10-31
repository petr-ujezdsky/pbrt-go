package mymath

type Bounds3d struct {
	PMin Point3
	PMax Point3
}

func NewBounds3dP(p Point3) Bounds3d {
	return Bounds3d{p, p}
}

func NewBounds3d(p1 Point3, p2 Point3) Bounds3d {
	return Bounds3d{p1.Min(p2), p1.Max(p2)}
}

func (b Bounds3d) Get(component int) Point3 {
	if component == 0 {
		return b.PMin
	}

	return b.PMax
}

func (b Bounds3d) Corner(corner int) Point3 {
	var c1, c2, c3 int
	c1 = corner & 1

	if corner&2 != 0 {
		c2 = 1
	} else {
		c2 = 0
	}

	if corner&4 != 0 {
		c3 = 1
	} else {
		c3 = 0
	}

	return NewPoint3(
		b.Get(c1).X,
		b.Get(c2).Y,
		b.Get(c3).Z)
}

func (b Bounds3d) UnionP(p Point3) Bounds3d {
	return NewBounds3d(
		b.PMin.Min(p),
		b.PMax.Max(p))
}

func (b1 Bounds3d) UnionB(b2 Bounds3d) Bounds3d {
	return NewBounds3d(
		b1.PMin.Min(b2.PMin),
		b1.PMax.Max(b2.PMax))
}

func (b1 Bounds3d) Intersect(b2 Bounds3d) Bounds3d {
	return NewBounds3d(
		b1.PMin.Max(b2.PMin),
		b1.PMax.Min(b2.PMax))
}

func (b1 Bounds3d) Overlaps(b2 Bounds3d) bool {
	x := (b1.PMax.X >= b2.PMin.X) && (b1.PMin.X <= b2.PMax.X)
	y := (b1.PMax.Y >= b2.PMin.Y) && (b1.PMin.Y <= b2.PMax.Y)
	z := (b1.PMax.Z >= b2.PMin.Z) && (b1.PMin.Z <= b2.PMax.Z)

	return (x && y && z)
}

func (b Bounds3d) Inside(p Point3) bool {
	return (p.X >= b.PMin.X && p.X <= b.PMax.X &&
		p.Y >= b.PMin.Y && p.Y <= b.PMax.Y &&
		p.Z >= b.PMin.Z && p.Z <= b.PMax.Z)
}

func (b Bounds3d) InsideExclusive(p Point3) bool {
	return (p.X >= b.PMin.X && p.X < b.PMax.X &&
		p.Y >= b.PMin.Y && p.Y < b.PMax.Y &&
		p.Z >= b.PMin.Z && p.Z < b.PMax.Z)
}

func (b Bounds3d) Expand(delta float64) Bounds3d {
	vDelta := NewVector3(delta, delta, delta)

	return NewBounds3d(
		b.PMin.SubtractV(vDelta),
		b.PMax.AddV(vDelta))
}

func (b Bounds3d) Diagonal() Vector3 {
	return b.PMax.SubtractP(b.PMin)
}

func (b Bounds3d) SurfaceArea() float64 {
	d := b.Diagonal()
	return 2 * (d.X*d.Y + d.X*d.Z + d.Y*d.Z)
}

func (b Bounds3d) Volume() float64 {
	d := b.Diagonal()
	return d.X * d.Y * d.Z
}

func (b Bounds3d) MaximumExtent() int {
	d := b.Diagonal()

	if d.X > d.Y && d.X > d.Z {
		return 0
	} else if d.Y > d.Z {
		return 1
	}

	return 2
}

func (b Bounds3d) Lerp(t Point3) Point3 {
	return b.PMin.LerpP(t, b.PMax)
}
