package main

import "fmt"

//Ray is a ray in space
type Ray struct {
	Origin   V3
	Dest     V3
	Children *[]Ray
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

//Dir returns the direction the ray is pointing
func (r *Ray) Dir() V3 {
	return SubV3(r.Dest, r.Origin)
}
