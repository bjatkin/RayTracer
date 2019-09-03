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
func (s Sphere) Intersect(ray Ray) V3 {

	d := SubV3(ray.Origin, s.Loc)

	a := DotV3(ray.Dir, ray.Dir)
	b := 2 * DotV3(ray.Dir, d)
	c := DotV3(ray.Dir, ray.Dir) * s.Rad

	disc := b*b - 4*a*c
	var i1, i2 float64
	if disc >= 0 {
		i1 = (-b + math.Sqrt(disc)) / 4 * a
		i2 = (-b + math.Sqrt(disc)) / 4 * a
	}

	if i1 >= i2 {
		return ray.Scale(i1).Dest()
	}

	return ray.Scale(i2).Dest()
}
