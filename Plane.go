package main

//Plane is a sphere that can be rendered
type Plane struct {
	Points      [3]V3
	Mat         Material
	Flipped     bool
	boundingBox boundBox
	setBBox     bool
}

func (p *Plane) genBBox() {
	dist := V3{0.001, 0.001, 0.001}
	min := p.Points[0]
	max := p.Points[0]
	for _, p := range p.Points {
		if p.x < min.x {
			min.x = p.x
		}
		if p.x > max.x {
			max.x = p.x
		}
		if p.z < min.z {
			min.z = p.z
		}
		if p.z > max.z {
			max.z = p.z
		}
		if p.y < min.y {
			min.y = p.y
		}
		if p.y > max.y {
			max.y = p.y
		}
	}

	//Prevent zero width/depth/height errors
	min = SubV3(min, dist)
	max = AddV3(max, dist)
	p.boundingBox = boundBox{p1: min, p2: max}
	p.setBBox = true
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

func (p Plane) BoundBox() boundBox {
	if !p.setBBox {
		p.genBBox()
	}
	return p.boundingBox
}

//Intersect takes a ray and returns the nearist intersection
func (p Plane) Intersect(ray *Ray) (float64, V3, bool) {
	//check if we intersect the bounding box
	if !p.BoundBox().Intersect(ray) {
		return 0, V3{}, false
	}

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

func (p Plane) IntersectPath(path *path) (float64, bool) {
	//check if we intersect the bounding box
	if !p.BoundBox().IntersectPath(path) {
		return 0, false
	}

	rayDir := path.Unit()
	v0 := p.Points[0]
	v1 := p.Points[1]
	v2 := p.Points[2]

	edge1 := SubV3(v1, v0)
	edge2 := SubV3(v2, v0)

	h := CrossV3(rayDir, edge2)
	a := DotV3(edge1, h)
	if a > -PathEpsilon && a < PathEpsilon {
		return 0, false //this ray is paralell to this triangle
	}
	f := 1.0 / a
	s := SubV3(path.Origin, v0)
	u := f * DotV3(s, h)
	if u < 0.0 || u > 1.0 {
		return 0, false
	}
	q := CrossV3(s, edge1)
	v := f * DotV3(rayDir, q)
	if v < 0.0 || u+v > 1.0 {
		return 0, false
	}

	//Compute t to find our intersection
	t := f * DotV3(edge2, q)
	if t > PathEpsilon {
		return t, true
	}
	return 0, false // we intersected with a line
}
