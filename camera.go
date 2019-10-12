package main

import (
	"math"
)

//Camera is a camera that can be added to the scene
type Camera struct {
	Fpoint       V3 //focal point
	Lpoint       V3 //look at point
	Width        int
	Height       int
	FOVx         float64
	FOVy         float64
	BGColor      RGB
	Clip         float64
	AmbientLight RGB
}

//Render renders the objects using the given camera
func (c Camera) Render(objects []split, lights *[]Light, test *[]Object) *Image {
	out := NewImage(c.Width, c.Height, "") //the output of the render

	upVector, sideVector := c.stepVectors()
	upVector = MulV3(1/float64(SUB_PIXELS), upVector)
	sideVector = MulV3(1/float64(SUB_PIXELS), sideVector)
	progress := progressBar{
		total: c.Width * c.Height,
		len:   70,
	}
	progress.Draw()

	// checkpointFile, err := os.Open("/Users/brandon/go/src/Projects/School/RayTracer/checkpoints/checkpoint_-1_359.png")
	// if err != nil {
	// 	fmt.Printf("THE WAS and error Reading the checkpoint file\n")
	// 	return out
	// }
	// chk, err := png.Decode(checkpointFile)
	// if err != nil {
	// 	fmt.Printf("THE WAS and error Reading the checkpoint file\n")
	// 	return out
	// }

	for x := -c.Width / 2; x < c.Width/2; x++ {
		for y := -c.Height / 2; y < c.Height/2; y++ {
			//Check to see if this point already exsists
			// if x < -3 {
			// 	progress.Update()
			// 	progress.Draw()
			// 	r, g, b, _ := chk.At(x+c.Width/2, y+c.Height/2).RGBA()
			// 	out.SetPixel(x+c.Width/2, y+c.Height/2, RGB{uint8(r), uint8(g), uint8(b)})
			// 	continue
			// }
			//create a new ray pointing at the viewport
			rg := RayGroup{}
			for sx := 0; sx < SUB_PIXELS; sx++ {
				for sy := 0; sy < SUB_PIXELS; sy++ {
					r := Ray{
						Origin:       c.Fpoint,
						Dest:         AddV3(AddV3(c.Lpoint, MulV3(float64(x*SUB_PIXELS+sx), upVector)), MulV3(float64(y*SUB_PIXELS+sy), sideVector)),
						MaxLength:    c.Clip,
						BGColor:      c.BGColor,
						AmbientLight: c.AmbientLight,
						CameraOrg:    c.Fpoint,
						Objects:      objects,
						TestObj:      test,
						Lights:       lights,
					}
					r.Jitter(sideVector.x * 2 / float64(SUB_PIXELS))
					rg = append(rg, &r)
				}
			}

			progress.Update()
			progress.Draw()
			// if progress.current%(c.Width*10) == 0 {
			// 	//Save a checkpoint file
			// 	pngFile, err := os.Create(
			// 		fmt.Sprintf("/Users/brandon/go/src/Projects/School/RayTracer/checkpoints/checkpoint_%d_%d.png", x, y),
			// 	)
			// 	if err != nil {
			// 		fmt.Printf("There was an error creating a checkpoint file: %s", err.Error())
			// 	} else {
			// 		out.WritePNG(pngFile)
			// 	}
			// }

			out.SetPixel(x+c.Width/2, y+c.Height/2, rg.Color(DEPTH))
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
