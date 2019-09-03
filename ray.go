package main

//Ray is a ray in space
type Ray struct {
	Origin V3
	Dir    V3
}

//Scale scales up a ray
func (r Ray) Scale(s float64) Ray {
	newDir := V3{
		x: r.Dir.x * s,
		y: r.Dir.y * s,
		z: r.Dir.z * s,
	}
	ret := Ray{
		Origin: r.Origin,
		Dir:    newDir,
	}
	return ret
}

//Dest returns the point that the ray ends on
func (r Ray) Dest() V3 {
	return V3{
		x: r.Origin.x + r.Dir.x,
		y: r.Origin.y + r.Dir.y,
		z: r.Origin.z + r.Dir.z,
	}
}
