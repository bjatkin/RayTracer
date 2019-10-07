package main

import "math/rand"

//Light is a light in the scean
type Light interface {
	ToLight(V3) V3
	GetColor() RGB
	SampleSize() int
}

//PointLight is a point light object
type PointLight struct {
	Loc   V3
	Color RGB
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
	Area  Plane
	Color RGB
}

//ToLight is a random vector direction to a point on the area light
func (l *AreaLight) ToLight(from V3) V3 {
	p1 := l.Area.Points[0]
	p2 := l.Area.Points[1]
	p3 := l.Area.Points[2]
	v1 := SubV3(p2, p1)
	v2 := SubV3(p3, p1)

	a := rand.Float64()
	b := rand.Float64() * (1 - a)
	if rand.Float64() > 0.5 {
		return AddV3(MulV3(a, v1), MulV3(b, v2))
	}
	return AddV3(MulV3(a, v2), MulV3(b, v1))
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
