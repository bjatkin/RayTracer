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
	Objects      *[]Object
	Lights       *[]Light
	MaxLength    float64
	BGColor      RGB
	AmbientLight RGB
	CameraOrg    V3

	//Do some caching so we don't recalculate the same thing over and over
	dir    V3
	dirSet bool
	len    float64
	lenSet bool
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

//Length returns the lenght of the light ray
func (r *Ray) Length() float64 {
	if !r.lenSet {
		r.len = DotV3(r.Dir(), r.Dir())
		r.lenSet = true
	}
	return r.len
}

//Color calculate of the ray
func (r *Ray) Color(depth int) RGB {
	//Find the closest ray collision
	hDist := r.MaxLength
	color := r.BGColor
	hO := (*r.Objects)[0]
	hPoint := V3{}
	for _, o := range *r.Objects {
		dist, hit, success := o.Intersect(r)
		if success && dist < hDist {
			hDist = dist
			hO = o
			hPoint = hit
		}
	}

	if hDist != r.MaxLength {
		color = r.calculateColor(hPoint, hO, depth)
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

func (r *Ray) calculateColor(point V3, o Object, depth int) RGB {
	ambLight := r.AmbientLight
	normal := o.Normal(point)
	mat := o.GetMat()
	toView := SubV3(point, r.CameraOrg)
	//Calculate the lighting portion of the lighting equation
	color := HadMulV3(MulV3(mat.AmbCoeff, ambLight.V3()), mat.DiffColor.V3())
	diffCol := MulV3(mat.DiffCoeff, mat.DiffColor.V3())
	specCol := MulV3(mat.SpecCoeff, mat.SpecColor.V3())

	for _, l := range *r.Lights {
		dir := l.ToLight(point)
		// cast a shadow ray to see if we need to calculate this light
		org := AddV3(point, MulV3(0.00000001, dir))
		sRay := Ray{
			Origin:    org, //shift along dir so we don't collide with the same object
			Dest:      AddV3(org, dir),
			MaxLength: r.MaxLength,
		}
		shadow := false
		for _, o := range *r.Objects {
			dist, _, cross := o.Intersect(&sRay)
			if cross {
				test := Ray{
					Origin: org,
					Dest:   AddV3(org, Unit(sRay.Dir())),
				}
				test = test.Scale(dist)
				if sRay.Length() > test.Length() {
					shadow = true
					break
				}
			}
		}
		if shadow {
			continue
		}
		diff := V3{}
		diffDir := DotV3(Unit(normal), Unit(dir))
		if diffDir > 0 {
			diff = MulV3(diffDir, diffCol)
		}

		spec := V3{}
		specDir := DotV3(Unit(ReflectV3(dir, normal)), Unit(toView))
		if specDir > 0 {
			spec = MulV3(math.Pow(specDir, mat.Phong), specCol)
		}

		color = AddV3(color, HadMulV3(l.GetColor().V3(), AddV3(diff, spec)))
	}

	depth--
	// cast a reflection ray
	reflectColor := V3{}
	if o.GetMat().ReflectCoeff > 0 && depth >= 0 {
		reflectV3 := Unit(ReflectV3(r.Dir(), o.Normal(point)))
		apex := AddV3(point, MulV3(0.000000001, reflectV3))
		reflect := Ray{
			Origin:       apex,
			Dest:         AddV3(apex, reflectV3),
			MaxLength:    r.MaxLength,
			Objects:      r.Objects,
			Lights:       r.Lights,
			BGColor:      r.BGColor,
			AmbientLight: r.AmbientLight,
			CameraOrg:    r.CameraOrg,
		}
		reflectColor = MulV3(o.GetMat().ReflectCoeff, reflect.Color(depth).V3())
	}

	// cast a transmission ray
	transColor := V3{}
	if o.GetMat().TransCoeff > 0 && depth >= 0 {
		I := Unit(r.Dir())
		N := Unit(normal)
		cos := math.Cos(Rad(math.Acos(DotV3(I, N))))
		nit := o.GetMat().RefractCoeff

		p1 := MulV3(nit, I)
		p2 := math.Sqrt(1 + nit*nit*(cos*cos-1))
		p3 := nit*cos - p2
		tdir := AddV3(p1, MulV3(p3, N))
		apex := AddV3(point, MulV3(0.00001, tdir))
		tRay := Ray{
			Origin:       apex,
			Dest:         AddV3(apex, tdir),
			MaxLength:    r.MaxLength,
			Objects:      r.Objects,
			Lights:       r.Lights,
			BGColor:      r.BGColor,
			AmbientLight: r.AmbientLight,
			CameraOrg:    r.CameraOrg,
		}
		transColor = MulV3(o.GetMat().TransCoeff, tRay.Color(depth).V3())
	}

	return AddV3(transColor, AddV3(reflectColor, color)).RGB()
}
