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
		BGColor:      RGB{100, 200, 255},
		Fpoint:       V3{0, 0, 0},
		Lpoint:       V3{0, 0, -1},
		Clip:         5000,
		AmbientLight: RGB{100, 100, 100},
	}

	out := C.Render(
		&[]Object{
			&Sphere{
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
			&Sphere{
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
			},
			&Plane{
				Points: [3]V3{V3{2, -3, -13}, V3{2, 3, -13}, V3{-1.5, 0, -10}},
				Mat: Material{
					AmbCoeff:   0.5,
					DiffCoeff:  0.4,
					DiffColor:  RGB{0, 155, 200},
					SpecCoeff:  0.6,
					SpecColor:  RGB{0, 155, 200},
					TransCoeff: 0.5,
					Phong:      50,
				},
			},
			&Plane{
				Points: [3]V3{V3{13, 0, 0}, V3{13, 0, -40}, V3{13, -40, 0}},
				Mat: Material{
					AmbCoeff:   0.1,
					DiffCoeff:  0.9,
					DiffColor:  RGB{50, 215, 105},
					SpecCoeff:  0.6,
					SpecColor:  RGB{50, 215, 105},
					TransCoeff: 0.5,
					Phong:      50,
				},
			},
			&Plane{
				Points: [3]V3{V3{13, 0, 0}, V3{13, 0, -40}, V3{13, 40, 0}},
				Mat: Material{
					AmbCoeff:   0.1,
					DiffCoeff:  0.9,
					DiffColor:  RGB{50, 215, 105},
					SpecCoeff:  0.6,
					SpecColor:  RGB{50, 215, 105},
					TransCoeff: 0.5,
					Phong:      50,
				},
				Flipped: true,
			},
			// &Sphere{
			// 	Loc: V3{4, 0, -10},
			// 	Rad: .5,
			// 	Mat: Material{
			// 		AmbCoeff:   0.1,
			// 		DiffCoeff:  0.9,
			// 		DiffColor:  RGB{255, 255, 255},
			// 		SpecCoeff:  0.2,
			// 		SpecColor:  RGB{255, 255, 255},
			// 		TransCoeff: 0.5,
			// 		Phong:      1,
			// 	},
			// },
		},
		&[]Light{
			// &DirLight{
			// 	Color: RGB{50, 50, 75},
			// 	Dir:   V3{-1, -1, 0},
			//  MaxDist: 5000,
			// },
			&PointLight{
				Color: RGB{50, 50, 75},
				Loc:   V3{4, 0, -10},
			},
			// &PointLight{
			// 	Color: RGB{50, 50, 75},
			// 	Loc:   V3{10, -10, -15},
			// },
		})

	pngFile, err := os.Create("/Users/brandon/go/src/Projects/School/RayTracer/test.png")
	if err != nil {
		fmt.Printf("There was an error: %s", err.Error())
		return
	}

	out.WritePNG(pngFile)
}
