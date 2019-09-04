package main

import (
	"fmt"
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
}

//Render renders the objects using the given camera
func (c Camera) Render(spheres ...Sphere) *PPM {
	out := NewPPM(c.Width, c.Height)
	dist := SubV3(c.Fpoint, c.Lpoint).Magnitude()
	degX := c.FOVx / 2
	radX := Rad(c.FOVx / 2)
	degY := c.FOVy / 2
	radY := Rad(c.FOVy / 2)
	Vwidth := math.Abs(((dist * math.Sin(radX)) / (math.Sin(Rad(90 - degX)))))
	Vheight := math.Abs(((dist * math.Sin(radY)) / (math.Sin(Rad(90 - degY)))))
	fmt.Println(Vwidth, Vheight)

	for x := -c.Width / 2; x < c.Width/2; x++ {
		for y := -c.Height / 2; y < c.Height/2; y++ {
			//create a new ray
			r := Ray{
				Origin: c.Fpoint,
				Dir: V3{
					x: float64(x) * (Vwidth / float64(c.Width)),
					y: float64(y) * (Vheight / float64(c.Height)),
					z: dist,
				},
			}

			hit := false
			for _, s := range spheres {
				_, success := s.Intersect(r)
				if success {
					out.SetPixel(y+c.Height/2, x+c.Width/2, s.Color)
					hit = true
					break
				}
			}
			if !hit {
				out.SetPixel(y+c.Height/2, x+c.Width/2, c.BGColor)
			}
		}
	}

	return out
}
