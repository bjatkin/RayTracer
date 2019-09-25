package main

//Light is a light in the scean
type Light interface {
	ToLight(V3) V3
	GetColor() RGB
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

//GetColor returns the color of the light
func (l *DirLight) GetColor() RGB {
	return l.Color
}
