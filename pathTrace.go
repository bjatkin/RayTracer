package main

import (
	"fmt"
	"os"
)

func (c Camera) pathTrace(objects []split, lights *[]Light) *Image {
	out := NewImage(c.Width, c.Height) //the output of the render

	upVector, sideVector := c.stepVectors()
	upVector = MulV3(1/float64(SUB_PIXELS), upVector)
	sideVector = MulV3(1/float64(SUB_PIXELS), sideVector)
	progress := progressBar{
		total: c.Width * c.Height,
		len:   70,
	}
	progress.Draw()

	for x := -c.Width / 2; x < c.Width/2; x++ {
		for y := -c.Height / 2; y < c.Height/2; y++ {
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
						Lights:       lights,
					}
					r.Jitter(sideVector.x * 2 / float64(SUB_PIXELS))
					rg = append(rg, &r)
				}
			}

			progress.Update()
			progress.Draw()
			if progress.current%(c.Width*10) == 0 {
				//Save a checkpoint file
				pngFile, err := os.Create(
					fmt.Sprintf("/Users/brandon/go/src/Projects/School/RayTracer/checkpoints/checkpoint_%d_%d.png", x, y),
				)
				if err != nil {
					fmt.Printf("There was an error creating a checkpoint file: %s", err.Error())
				} else {
					out.WritePNG(pngFile)
				}
			}

			out.SetPixel(x+c.Width/2, y+c.Height/2, rg.Color(DEPTH))
		}
	}
	return out
}
