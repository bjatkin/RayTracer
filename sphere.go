package main

import (
	"math"
)

//Sphere is a sphere that can be rendered
type Sphere struct {
	Loc V3
	Rad float64
	Mat Material
}

//Intersect takes a ray and returns the nearist intersection
func (s Sphere) Intersect(ray *Ray) (float64, V3, bool) {

	d := SubV3(ray.Origin, s.Loc)

	rayDir := Unit(ray.Dir())
	b := 2 * DotV3(rayDir, d)
	c := DotV3(d, d) - (s.Rad * s.Rad)

	disc := b*b - 4*c

	if disc < 0 {
		return 0, V3{}, false
	}

	disc = math.Sqrt(disc)
	i1 := (-b + disc) / 2
	if i1 > 0 {
		ret := Ray{
			Origin: ray.Origin,
			Dest:   AddV3(ray.Origin, rayDir),
		}
		return i1, ret.Scale(i1).Dest, true
	}

	i2 := (-b - disc) / 2
	if i2 > 0 {
		ret := Ray{
			Origin: ray.Origin,
			Dest:   AddV3(ray.Origin, rayDir),
		}
		return i1, ret.Scale(i1).Dest, true

	}

	return 0, V3{}, false
}
