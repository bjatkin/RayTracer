package main

import (
	"fmt"
	"math"
)

//Object is anything that can be intersected and drawn by a ray
type Object interface {
	Intersect(*Ray) (float64, V3, bool)
	GetMat() Material
	Normal(V3) V3
	BoundBox() boundBox
}

//Sphere is a sphere that can be rendered
type Sphere struct {
	Loc         V3
	Rad         float64
	Mat         Material
	boundingBox boundBox
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

	s.boundingBox = boundBox{p1: min, p2: max}
	s.setBBox = true
}

func (s Sphere) GetMat() Material {
	return s.Mat
}

func (s Sphere) Normal(pt V3) V3 {
	return Unit(SubV3(pt, s.Loc))
}

func (s Sphere) BoundBox() boundBox {
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
		if (SubV3(ret.Scale(i2).Dest, s.Loc)).Magnitude()-s.Rad > 0.0001 {
			fmt.Printf("ERROR Dist: %f, Rad: %f!\n", DotV3((SubV3(ret.Scale(i2).Dest, s.Loc)), SubV3(ret.Scale(i2).Dest, s.Loc)), s.Rad)
		}

		return i2, ret.Scale(i2).Dest, true

	}

	i1 := (-b + disc) / 2
	if i1 > 0 {
		ret := Ray{
			Origin: ray.Origin,
			Dest:   AddV3(ray.Origin, rayDir),
		}
		if (SubV3(ret.Scale(i2).Dest, s.Loc)).Magnitude()-s.Rad > 0.0001 {
			fmt.Printf("ERROR Dist: %f, Rad: %f!\n", DotV3((SubV3(ret.Scale(i2).Dest, s.Loc)), SubV3(ret.Scale(i2).Dest, s.Loc)), s.Rad)
		}
		return i1, ret.Scale(i1).Dest, true
	}

	return 0, V3{}, false
}
