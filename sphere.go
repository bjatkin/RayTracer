package main

import (
	"math"
)

//Sphere is a sphere that can be rendered
type Sphere struct {
	Loc   V3
	Rad   float64
	Color RGB
}

//Intersect takes a ray and returns the nearist intersection
func (s Sphere) Intersect(ray Ray) (V3, bool) {

	d := SubV3(ray.Origin, s.Loc)

	rayDir := Unit(ray.Dir)
	b := 2 * DotV3(rayDir, d)
	c := DotV3(d, d) - (s.Rad * s.Rad)

	disc := b*b - 4*c

	if disc < 0 {
		return V3{}, false
	}

	i1 := (-b + math.Sqrt(disc)) / 2
	i2 := (-b - math.Sqrt(disc)) / 2
	ret := Ray{
		Origin: ray.Origin,
		Dir:    rayDir,
	}

	if i1 < i2 {
		return ret.Scale(i1).Dest(), true
	}

	return ret.Scale(i2).Dest(), true
}
