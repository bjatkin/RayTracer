package main

import (
	"math"
)

//Sphere is a sphere that can be rendered
type Sphere struct {
	Loc V3
	Rad float64
}

//Intersect takes a ray and returns the nearist intersection
func (s Sphere) Intersect(ray Ray) (V3, bool) {

	d := SubV3(ray.Origin, s.Loc)

	a := DotV3(ray.Dir, ray.Dir)
	b := 2 * DotV3(ray.Dir, d)
	c := DotV3(d, d) - (s.Rad * s.Rad)

	disc := b*b - 4*a*c

	if disc < 0 {
		return V3{}, false
	}

	i1 := (-b + math.Sqrt(disc)) / (2 * a)
	i2 := (-b - math.Sqrt(disc)) / (2 * a)

	if i1 < i2 {
		return ray.Scale(i1).Dest(), true
	}

	return ray.Scale(i2).Dest(), true
}
