package mymath

type Bounds3 struct {
	PMin Point3
	PMax Point3
}

func NewBounds3P(p Point3) Bounds3 {
	return Bounds3{p, p}
}

func NewBounds3(p1 Point3, p2 Point3) Bounds3 {
	return Bounds3{p1.Min(p2), p1.Max(p2)}
}

func (b Bounds3) Get(component int) Point3 {
	if component == 0 {
		return b.PMin
	}

	return b.PMax
}

func (b Bounds3) Corner(corner int) Point3 {
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

func (b Bounds3) UnionP(p Point3) Bounds3 {
	return NewBounds3(
		b.PMin.Min(p),
		b.PMax.Max(p))
}

func (b1 Bounds3) UnionB(b2 Bounds3) Bounds3 {
	return NewBounds3(
		b1.PMin.Min(b2.PMin),
		b1.PMax.Max(b2.PMax))
}

func (b1 Bounds3) Intersect(b2 Bounds3) Bounds3 {
	return NewBounds3(
		b1.PMin.Max(b2.PMin),
		b1.PMax.Min(b2.PMax))
}

func (b1 Bounds3) Overlaps(b2 Bounds3) bool {
	x := (b1.PMax.X >= b2.PMin.X) && (b1.PMin.X <= b2.PMax.X)
	y := (b1.PMax.Y >= b2.PMin.Y) && (b1.PMin.Y <= b2.PMax.Y)
	z := (b1.PMax.Z >= b2.PMin.Z) && (b1.PMin.Z <= b2.PMax.Z)

	return (x && y && z)
}

func (b Bounds3) Inside(p Point3) bool {
	return (p.X >= b.PMin.X && p.X <= b.PMax.X &&
		p.Y >= b.PMin.Y && p.Y <= b.PMax.Y &&
		p.Z >= b.PMin.Z && p.Z <= b.PMax.Z)
}

func (b Bounds3) InsideExclusive(p Point3) bool {
	return (p.X >= b.PMin.X && p.X < b.PMax.X &&
		p.Y >= b.PMin.Y && p.Y < b.PMax.Y &&
		p.Z >= b.PMin.Z && p.Z < b.PMax.Z)
}

func (b Bounds3) Expand(delta float64) Bounds3 {
	vDelta := NewVector3(delta, delta, delta)

	return NewBounds3(
		b.PMin.SubtractV(vDelta),
		b.PMax.AddV(vDelta))
}

func (b Bounds3) Diagonal() Vector3 {
	return b.PMax.SubtractP(b.PMin)
}

func (b Bounds3) SurfaceArea() float64 {
	d := b.Diagonal()
	return 2 * (d.X*d.Y + d.X*d.Z + d.Y*d.Z)
}

func (b Bounds3) Volume() float64 {
	d := b.Diagonal()
	return d.X * d.Y * d.Z
}

func (b Bounds3) MaximumExtent() int {
	d := b.Diagonal()

	if d.X > d.Y && d.X > d.Z {
		return 0
	} else if d.Y > d.Z {
		return 1
	}

	return 2
}

func (b Bounds3) Lerp(t Point3) Point3 {
	return b.PMin.LerpP(t, b.PMax)
}

// Offset returns the continuous position of a point relative to the corners of the box,
// where a point at the minimum corner has offset (0,0,0), a point at the maximum corner has offset (1,1,1), and so forth.
func (b Bounds3) Offset(p Point3) Vector3 {
	o := p.SubtractP(b.PMin)

	if b.PMax.X > b.PMin.X {
		o.X /= b.PMax.X - b.PMin.X
	}

	if b.PMax.Y > b.PMin.Y {
		o.Y /= b.PMax.Y - b.PMin.Y
	}

	if b.PMax.Z > b.PMin.Z {
		o.Z /= b.PMax.Z - b.PMin.Z
	}

	return o
}

func (b Bounds3) BoundingSphere() BoundingSphere {
	center := (b.PMin.AddP(b.PMax).Multiply(0.5))

	var radius float64

	if b.Inside(center) {
		radius = center.Distance(b.PMax)
	} else {
		radius = 0
	}

	return NewBoundingSphere(center, radius)
}

// IntersectP see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/geometry.h#L1388
func (b Bounds3) IntersectP(ray Ray) (bool, float64, float64) {
	t0 := 0.0
	t1 := ray.TMax

	for i := 0; i < 3; i++ {
		// Update interval for ith bounding box slab
		invRayDir := 1.0 / ray.D.Get(i)
		tNear := (b.PMin.Get(i) - ray.O.Get(i)) * invRayDir
		tFar := (b.PMax.Get(i) - ray.O.Get(i)) * invRayDir

		// Update parametric interval from slab intersection  values
		if tNear > tFar {
			tFar, tNear = tNear, tFar
		}

		// Update tFar to ensure robust ray–bounds intersection
		tFar *= 1 + 2*Gamma3

		// both ifs take NaN into account
		if tNear > t0 {
			t0 = tNear
		}

		if tFar < t1 {
			t1 = tFar
		}

		if t0 > t1 {
			return false, 0, 0
		}
	}

	return true, t0, t1
}

// IntersectPPrecomputed see https://github.com/mmp/pbrt-v3/blob/aaa552a4b9cbf9dccb71450f47b268e0ed6370e2/src/core/geometry.h#L1412
func (b Bounds3) IntersectPPrecomputed(ray Ray, invDir Vector3, dirIsNeg [3]int) bool {
	// Check for ray intersection against $x$ and $y$ slabs
	tMin := (b.Get(dirIsNeg[0]).X - ray.O.X) * invDir.X
	tMax := (b.Get(1-dirIsNeg[0]).X - ray.O.X) * invDir.X
	tyMin := (b.Get(dirIsNeg[1]).Y - ray.O.Y) * invDir.Y
	tyMax := (b.Get(1-dirIsNeg[1]).Y - ray.O.Y) * invDir.Y

	// Update _tMax_ and _tyMax_ to ensure robust bounds intersection
	tMax *= 1 + 2*Gamma3
	tyMax *= 1 + 2*Gamma3

	if tMin > tyMax || tyMin > tMax {
		return false
	}

	if tyMin > tMin {
		tMin = tyMin
	}

	if tyMax < tMax {
		tMax = tyMax
	}

	// Check for ray intersection against $z$ slab
	tzMin := (b.Get(dirIsNeg[2]).Z - ray.O.Z) * invDir.Z
	tzMax := (b.Get(1-dirIsNeg[2]).Z - ray.O.Z) * invDir.Z

	// Update _tzMax_ to ensure robust bounds intersection
	tzMax *= 1 + 2*Gamma3

	if tMin > tzMax || tzMin > tMax {
		return false
	}

	if tzMin > tMin {
		tMin = tzMin
	}

	if tzMax < tMax {
		tMax = tzMax
	}

	return (tMin < ray.TMax) && (tMax > 0)
}
