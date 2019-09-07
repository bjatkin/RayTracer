package main

import "fmt"

//Ray is a ray in space
type Ray struct {
	Origin       V3
	Dest         V3
	Children     *[]Ray
	Spheres      *[]Sphere
	Lights       *[]DirLight
	MaxLength    float64
	BGColor      RGB
	AmbientLight RGB
	CameraOrg    V3

	dir    V3
	dirSet bool //cache direction in the case of repeted calls
}

func (r *Ray) String() string {
	return fmt.Sprintf("%s -> %s", r.Origin, r.Dest)
}

//Child sends a child ray out
func (r *Ray) Child(dest V3) {
	children := append(*r.Children, Ray{Origin: r.Dest, Dest: dest})
	r.Children = &children
}

//Scale scales up a ray
func (r *Ray) Scale(s float64) Ray {
	newDir := MulV3(s, r.Dir())
	return Ray{
		Origin: r.Origin,
		Dest:   AddV3(r.Origin, newDir),
	}
}

//Color calculate of the ray
func (r *Ray) Color() RGB {
	//Find the closest ray collision
	hDist := r.MaxLength
	color := r.BGColor
	for _, s := range *r.Spheres {
		dist, hit, success := s.Intersect(r)
		if success && dist < hDist {
			hDist = dist
			ptNormal := Unit(SubV3(hit, s.Loc))
			color = calculateColor(r.AmbientLight, hit, ptNormal, SubV3(r.CameraOrg, hit), s.Mat, r.Lights)
		}
	}

	return color
}

//Dir returns the direction the ray is pointing
func (r *Ray) Dir() V3 {
	if !r.dirSet {
		r.dir = SubV3(r.Dest, r.Origin)
		r.dirSet = true
	}
	return r.dir
}
