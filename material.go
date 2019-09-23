package main

//Material represents a material of an object
type Material struct {
	AmbCoeff     float64
	SpecCoeff    float64
	SpecColor    RGB
	DiffCoeff    float64
	DiffColor    RGB
	TransCoeff   float64
	Phong        float64
	ReflectCoeff float64
	RefractCoeff float64
}
