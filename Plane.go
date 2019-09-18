package main

//Plane is a sphere that can be rendered
type Plane struct {
	Points  [3]V3
	Mat     Material
	Flipped bool
}

func (p Plane) GetMat() Material {
	return p.Mat
}

func (p Plane) Normal(pt V3) V3 {
	v0 := p.Points[0]
	v1 := p.Points[1]
	v2 := p.Points[2]

	edge1 := SubV3(v1, v0)
	edge2 := SubV3(v2, v0)
	if p.Flipped {
		return Unit(CrossV3(edge2, edge1))
	}
	return Unit(CrossV3(edge1, edge2))
}

//Intersect takes a ray and returns the nearist intersection
func (p Plane) Intersect(ray *Ray) (float64, V3, bool) {
	const eps = 0.0000001
	rayDir := Unit(ray.Dir())
	v0 := p.Points[0]
	v1 := p.Points[1]
	v2 := p.Points[2]

	edge1 := SubV3(v1, v0)
	edge2 := SubV3(v2, v0)

	h := CrossV3(rayDir, edge2)
	a := DotV3(edge1, h)
	if a > -eps && a < eps {
		return 0.0, V3{}, false //this ray is paralell to this triangle
	}
	f := 1.0 / a
	s := SubV3(ray.Origin, v0)
	u := f * DotV3(s, h)
	if u < 0.0 || u > 1.0 {
		return 0.0, V3{}, false
	}
	q := CrossV3(s, edge1)
	v := f * DotV3(rayDir, q)
	if v < 0.0 || u+v > 1.0 {
		return 0.0, V3{}, false
	}

	//Compute t to find our intersection
	t := f * DotV3(edge2, q)
	if t > eps {
		ret := Ray{
			Origin: ray.Origin,
			Dest:   AddV3(ray.Origin, rayDir),
		}
		return t, ret.Scale(t).Dest, true
	} else {
		return 0.0, V3{}, false // we intersected with a line
	}
}
