package main

import (
	"fmt"
	"math"
	"math/rand"
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

//RGB converts a vector into an rgb object
func (v V3) RGB() RGB {
	if v.x < 0 {
		v.x = 0
	}
	if v.y < 0 {
		v.y = 0
	}
	if v.z < 0 {
		v.z = 0
	}
	if v.x > 1 {
		v.x = 1
	}
	if v.y > 1 {
		v.y = 1
	}
	if v.z > 1 {
		v.z = 1
	}
	return RGB{
		R: v.x * 255,
		G: v.y * 255,
		B: v.z * 255,
	}
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

//HadMulV3 multiples each element by the corsponding element in the other vector
func HadMulV3(v1, v2 V3) V3 {
	return V3{
		x: v1.x * v2.x,
		y: v1.y * v2.y,
		z: v1.z * v2.z,
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

//ReflectV3 is the reflection of one vector 3 across another
func ReflectV3(d, n V3) V3 {
	n = Unit(n)
	d = Unit(MulV3(-1, d))
	return SubV3(MulV3(2*DotV3(d, n), n), d)
}

//MinV3 returns a v3 with all the min components of the two vectors
func MinV3(a V3, b V3) V3 {
	if a.x > b.x {
		a.x = b.x
	}
	if a.y > b.y {
		a.y = b.y
	}
	if a.z > b.z {
		a.z = b.z
	}
	return a
}

//MaxV3 returns a v3 with all the max components of the two vectors
func MaxV3(a V3, b V3) V3 {
	if a.x < b.x {
		a.x = b.x
	}
	if a.y < b.y {
		a.y = b.y
	}
	if a.z < b.z {
		a.z = b.z
	}
	return a
}

//RandV3 returns a random V3 with each component in the range -size,+size
func RandV3(size float64) V3 {
	return V3{
		x: RandGaus() * size,
		y: RandGaus() * size,
		z: RandGaus() * size,
	}
}

//JitterV3 randomly jitters a V3 by the size given
func JitterV3(size float64, start V3) V3 {
	return AddV3(start, RandV3(size))
}

//Rad returns the degrees as radians
func Rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

//Deg returns the radians as degrees
func Deg(rad float64) float64 {
	return rad * (180 / math.Pi)
}

//RandGaus returns a random number between -1 to 1 with a gausian distribution
func RandGaus() float64 {
	return rand.NormFloat64() / math.MaxFloat64
}
