package main

import (
	"fmt"
	"math"
)

//V3 is a 4D vector
type V3 struct {
	x, y, z float64
}

func (v V3) String() string {
	return fmt.Sprintf("<%f, %f, %f>", v.x, v.y, v.z)
}

//Magnitude is the magnitude of the vector
func (v V3) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

//Unit returns the unit vector of the given V3
func Unit(v V3) V3 {
	return DivV3(v.Magnitude(), v)
}

//AddV3 adds 2 V3's
func AddV3(v1, v2 V3) V3 {
	return V3{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
		z: v1.z + v2.z,
	}
}

//SubV3 subtracts 2 V3's
func SubV3(v1, v2 V3) V3 {
	return V3{
		x: v1.x - v2.x,
		y: v1.y - v2.y,
		z: v1.z - v2.z,
	}
}

//MulV3 multiplies a V3 by a scaler
func MulV3(s float64, v1 V3) V3 {
	return V3{
		x: v1.x * s,
		y: v1.y * s,
		z: v1.z * s,
	}
}

//DivV3 divides a V3 by a scaler
func DivV3(s float64, v1 V3) V3 {
	return V3{
		x: v1.x / s,
		y: v1.y / s,
		z: v1.z / s,
	}
}

//CrossV3 is the cross product of 2 V3's
func CrossV3(v1, v2 V3) V3 {
	return V3{
		x: v1.y*v2.z - v1.z*v2.y,
		y: v1.z*v2.x - v1.x*v2.z,
		z: v1.x*v2.y - v1.y*v2.x,
	}
}

//DotV3 is the dot product of 2 v3's
func DotV3(v1, v2 V3) float64 {
	return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
}

//V4 is a 4D vector
type V4 struct {
	x, y, z, w float64
}

//AddV4 adds 2 V4's
func AddV4(v1, v2 V4) V4 {
	return V4{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
		z: v1.z + v2.z,
		w: v1.w + v2.w,
	}
}

//Mat is a general matrix
type Mat struct {
	x, y int8
	data []float64
}

//At returns the value at x,y in the matrix
func (m *Mat) At(x, y int8) float64 {
	return m.data[x*m.x+m.y]
}

//NewMat returns a new matrix with the given dimentions and data
func NewMat(x, y int8, data []float64) *Mat {
	return &Mat{
		x:    x,
		y:    y,
		data: data,
	}
}

//Rad returns the degrees as radians
func Rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
