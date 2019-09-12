package main

import (
	"fmt"
	"os"
)

func main() {
	C := Camera{
		Width:        1080,
		Height:       720,
		FOVx:         120,
		FOVy:         100,
		BGColor:      RGB{0, 100, 200},
		Fpoint:       V3{0, 0, 0},
		Lpoint:       V3{0, 0, -1},
		Clip:         5000,
		AmbientLight: RGB{100, 100, 100},
	}

	out := C.Render(
		&[]Sphere{
			Sphere{
				Loc: V3{0, 3, -10},
				Rad: 1,
				Mat: Material{
					AmbCoeff:   0.5,
					DiffCoeff:  0.4,
					DiffColor:  RGB{240, 20, 20},
					SpecCoeff:  0.6,
					SpecColor:  RGB{240, 20, 20},
					TransCoeff: 0.5,
					Phong:      50,
				},
			},
			Sphere{
				Loc: V3{3, 0, -15},
				Rad: 3,
				Mat: Material{
					AmbCoeff:   0.5,
					DiffCoeff:  0.8,
					DiffColor:  RGB{230, 200, 40},
					SpecCoeff:  0.2,
					SpecColor:  RGB{255, 255, 255},
					TransCoeff: 0.5,
					Phong:      1,
				},
			}},
		&[]Light{
			&DirLight{
				Color: RGB{200, 200, 175},
				Dir:   V3{-1, -1, 0},
			},
			&PointLight{
				Color: RGB{200, 30, 0},
				Loc:   V3{0, 0, 0},
			},
		})

	pngFile, err := os.Create("/Users/brandon/go/src/Projects/School/RayTracer/test.png")
	if err != nil {
		fmt.Printf("There was an error: %s", err.Error())
		return
	}

	out.WritePNG(pngFile)
}
