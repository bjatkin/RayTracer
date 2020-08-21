package main

import (
	"math"
)

//Object is anything that can be intersected and drawn by a ray
type Object interface {
	Intersect(*Ray) (float64, V3, bool)
	IntersectPath(*path) (float64, bool)
	GetMat() Material
	Normal(V3) V3
	BoundBox() BoundBox
}

//Sphere is a sphere that can be rendered
type Sphere struct {
	Loc         V3
	Rad         float64
	Mat         Material
	boundingBox BoundBox
	setBBox     bool
}

func (s *Sphere) genBBox() {
	min := V3{
		x: s.Loc.x - s.Rad,
		y: s.Loc.y - s.Rad,
		z: s.Loc.z - s.Rad,
	}

	max := V3{
		x: s.Loc.x + s.Rad,
		y: s.Loc.y + s.Rad,
		z: s.Loc.z + s.Rad,
	}

	s.boundingBox = BoundBox{p1: min, p2: max}
	s.setBBox = true
}

// GetMat returns the material of the sphere
func (s Sphere) GetMat() Material {
	return s.Mat
}

// Normal returns the normal of the sphere at given point
func (s Sphere) Normal(pt V3) V3 {
	return Unit(SubV3(pt, s.Loc))
}

// BoundBox returns a bounding box the encloses the full sphere
func (s Sphere) BoundBox() BoundBox {
	if !s.setBBox {
		s.genBBox()
	}
	return s.boundingBox
}

//Intersect takes a ray and returns the nearist intersection
func (s Sphere) Intersect(ray *Ray) (float64, V3, bool) {
	//check if we intersect the bounding box
	if !s.BoundBox().Intersect(ray) {
		return 0, V3{}, false
	}

	d := SubV3(ray.Origin, s.Loc)

	rayDir := Unit(ray.Dir())
	b := 2 * DotV3(rayDir, d)
	c := DotV3(d, d) - (s.Rad * s.Rad)

	disc := b*b - 4*c

	if disc < 0 {
		return 0, V3{}, false
	}

	disc = math.Sqrt(disc)
	i2 := (-b - disc) / 2
	if i2 > 0 {
		ret := Ray{
			Origin: ray.Origin,
			Dest:   AddV3(ray.Origin, rayDir),
		}
		return i2, ret.Scale(i2).Dest, true

	}

	i1 := (-b + disc) / 2
	if i1 > 0 {
		ret := Ray{
			Origin: ray.Origin,
			Dest:   AddV3(ray.Origin, rayDir),
		}
		return i1, ret.Scale(i1).Dest, true
	}

	return 0, V3{}, false
}

// IntersectPath returns the nearist intersection of the path and the sphere
func (s Sphere) IntersectPath(path *path) (float64, bool) {
	//check if we intersect the bounding box
	if !s.BoundBox().IntersectPath(path) {
		return 0, false
	}

	d := SubV3(path.Origin, s.Loc)

	rayDir := Unit(path.Dir())
	b := 2 * DotV3(rayDir, d)
	c := DotV3(d, d) - (s.Rad * s.Rad)

	disc := b*b - 4*c

	if disc < 0 {
		return 0, false
	}

	disc = math.Sqrt(disc)
	i2 := (-b - disc) / 2
	if i2 > 0 {
		return i2, true
	}

	i1 := (-b + disc) / 2
	if i1 > 0 {
		return i1, true
	}

	return 0, false
}
