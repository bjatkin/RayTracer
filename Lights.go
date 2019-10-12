package main

import (
	"math"
	"math/rand"
)

//Light is a light in the scean
type Light interface {
	ToLight(V3) V3
	GetColor() RGB
	GetIntensity() float64
	SampleSize() int
	Intersect(*path) (float64, bool)
}

//PointLight is a point light object
type PointLight struct {
	Loc   V3
	Color RGB
}

func (l *PointLight) GetIntensity() float64 {
	return 0
}

//TODO finish this as a sphere
func (l *PointLight) Intersect(p *path) (float64, bool) {
	return 0, false
}

//ToLight is the non normalized vector from the point to the light
func (l *PointLight) ToLight(from V3) V3 {
	return SubV3(l.Loc, from)
}

//SampleSize returns the number of samples needed to
// have an accurate representation of the light source
func (l *PointLight) SampleSize() int {
	return 1
}

//GetColor returns the color of the light
func (l *PointLight) GetColor() RGB {
	return l.Color
}

//DirLight is a directional light source
type DirLight struct {
	Dir     V3
	Color   RGB
	MaxDist float64
}

func (l *DirLight) GetIntensity() float64 {
	return 0
}

//TODO finish this with dot products stuff
func (l *DirLight) Intersect(p *path) (float64, bool) {
	return 0, false
}

//ToLight is the vector direction of the light
func (l *DirLight) ToLight(from V3) V3 {
	return l.Dir
	//return MulV3(l.MaxDist, l.Dir)
}

//SampleSize returns the number of samples needed to
// have an accurate representation of the light source
func (l *DirLight) SampleSize() int {
	return 1
}

//GetColor returns the color of the light
func (l *DirLight) GetColor() RGB {
	return l.Color
}

//AreaLight is an area light source
type AreaLight struct {
	Area      Plane
	Intensity float64
	Color     RGB
}

func (l *AreaLight) GetIntensity() float64 {
	return l.Intensity
}

func (l *AreaLight) Intersect(p *path) (float64, bool) {
	return l.Area.IntersectPath(p)
}

//ToLight is a random vector direction to a point on the area light
func (l *AreaLight) ToLight(from V3) V3 {
	p1 := l.Area.Points[0]
	p2 := l.Area.Points[1]
	p3 := l.Area.Points[2]

	r1 := rand.Float64()
	r2 := rand.Float64()
	A := MulV3((1 - math.Sqrt(r1)), p1)
	B := MulV3(math.Sqrt(r1)*(1-r2), p2)
	C := MulV3(r2*math.Sqrt(r1), p3)
	loc := AddV3(AddV3(A, B), C)
	return SubV3(loc, from)
}

//SampleSize returns the number of samples needed to
// have an accurate representation of the light source
func (l *AreaLight) SampleSize() int {
	return SHADOW_SAMPLES
}

//GetColor returns the color of the light
func (l *AreaLight) GetColor() RGB {
	return l.Color
}
