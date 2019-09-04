package main

import "fmt"

//Ray is a ray in space
type Ray struct {
	Origin V3
	Dir    V3
}

func (r Ray) String() string {
	return fmt.Sprintf("%s -> %s", r.Origin, r.Dest)
}

//Scale scales up a ray
func (r Ray) Scale(s float64) Ray {
	return Ray{
		Origin: r.Origin,
		Dir:    MulV3(s, r.Dir),
	}
}

//Dest returns the final destination of the ray
func (r Ray) Dest() V3 {
	return AddV3(r.Origin, r.Dir)
}
