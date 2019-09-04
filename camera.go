package main

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
func (c Camera) Render(out PPM, spheres ...Sphere) {
	//dist := SubV3(c.Fpoint, c.Lpoint).Magnitude()
	// Swidth := 2 * ((dist * math.Sin(c.FOVx/2)) / (math.Sin(90 - c.FOVx/2)))
	// Sheight := 2 * ((dist * math.Sin(c.FOVy/2)) / (math.Sin(90 - c.FOVy/2)))

	// Wstep := Swidth / float64(c.Width)
	// Hstep := Sheight / float64(c.Height)

	for row := 0; row < c.Width; row++ {
		for col := 0; col < c.Height; col++ {
			//create a new ray
			r := Ray{
				Origin: c.Fpoint,
				Dest:   V3{},
			}

			for _, s := range spheres {
				_, success := s.Intersect(r)
				if success {
					out.SetPixel(row, col, s.Color)
					continue
				}
				out.SetPixel(row, col, c.BGColor)
			}
		}
	}
}
