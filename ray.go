package main

import (
	"fmt"
	"math"
)

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
			color = calculateColor(r.AmbientLight, hit, ptNormal, SubV3(hit, r.CameraOrg), s.Mat, r.Lights)
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

func calculateColor(ambLight RGB, point V3, normal V3, toView V3, mat Material, lights *[]DirLight) RGB {
	//Calculate the lighting portion of the lighting equation
	color := HadMulV3(MulV3(mat.AmbCoeff, ambLight.V3()), mat.DiffColor.V3())
	diffCol := MulV3(mat.DiffCoeff, mat.DiffColor.V3())
	specCol := MulV3(mat.SpecCoeff, mat.SpecColor.V3())

	for _, l := range *lights {
		diffDir := DotV3(Unit(normal), Unit(l.Dir))
		diff := V3{}
		if diffDir > 0 {
			diff = MulV3(diffDir, diffCol)
		}

		specDir := DotV3(Unit(ReflectV3(l.Dir, normal)), Unit(toView))
		spec := V3{}
		if specDir > 0 {
			spec = MulV3(math.Pow(specDir, mat.Phong), specCol)
		}

		color = AddV3(color, HadMulV3(l.Color.V3(), AddV3(diff, spec)))
	}

	return color.RGB()
}
