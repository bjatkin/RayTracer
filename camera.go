package main

import (
	"math"
)

//Camera is a camera that can be added to the scene
type Camera struct {
	Fpoint  V3 //focal point
	Lpoint  V3 //look at point
	Width   int
	Height  int
	FOVx    float64
	FOVy    float64
	BGColor RGB
	Clip    float64
}

//Render renders the objects using the given camera
func (c Camera) Render(spheres []Sphere, lights []DirLight) *Image {
	out := NewImage(c.Width, c.Height) //the output of the render

	upVector, sideVector := c.stepVectors()

	for x := -c.Width / 2; x < c.Width/2; x++ {
		for y := -c.Height / 2; y < c.Height/2; y++ {
			//create a new ray pointing at the viewport
			r := Ray{
				Origin: c.Fpoint,
				Dest:   AddV3(AddV3(c.Lpoint, MulV3(float64(x), upVector)), MulV3(float64(y), sideVector)),
			}

			//Find the closest ray collision
			hit := false
			// hitLoc := V3{}
			hDist := c.Clip
			for _, s := range spheres {
				dist, _, success := s.Intersect(r)
				if success && dist < hDist {
					//Color the pixel
					out.SetPixel(x+c.Width/2, y+c.Height/2, s.Color)
					// hitLoc = hit
					hit = true
					hDist = dist
				}
			}

			if !hit {
				out.SetPixel(x+c.Width/2, y+c.Height/2, c.BGColor)
			}
		}
	}

	return out
}

func (c Camera) stepVectors() (V3, V3) {
	look := SubV3(c.Fpoint, c.Lpoint)
	dist := look.Magnitude()

	radX := Rad(c.FOVx / 2)
	radY := Rad(c.FOVy / 2)
	rad90 := Rad(90)

	Vwidth := math.Abs(((dist * math.Sin(radX)) / (math.Sin(rad90 - radX))))
	Vheight := math.Abs(((dist * math.Sin(radY)) / (math.Sin(rad90 - radY))))

	upVector := V3{
		x: 0,
		y: Vheight / float64(c.Height),
		z: (-look.y * (Vheight / float64(c.Height)) / look.z),
	}
	sideVector := V3{
		x: Vwidth / float64(c.Width),
		y: 0,
		z: (-look.x * (Vwidth / float64(c.Width)) / look.z),
	}

	return upVector, sideVector
}
