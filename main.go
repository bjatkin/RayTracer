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
	red := Material{
		AmbCoeff:     0.5,
		DiffCoeff:    0.4,
		DiffColor:    RGB{240, 20, 20},
		SpecCoeff:    0.6,
		SpecColor:    RGB{240, 20, 20},
		TransCoeff:   0.5,
		Phong:        50,
		ReflectCoeff: 0.3,
		RefractCoeff: 1.33,
	}
	yellow := Material{
		AmbCoeff:     0.5,
		DiffCoeff:    0.8,
		DiffColor:    RGB{230, 200, 40},
		SpecCoeff:    0.2,
		SpecColor:    RGB{255, 255, 255},
		TransCoeff:   0,
		Phong:        1,
		ReflectCoeff: 0.5,
		RefractCoeff: 1.33,
	}
	white := Material{
		AmbCoeff:     0.1,
		DiffCoeff:    0.9,
		DiffColor:    RGB{255, 255, 255},
		SpecCoeff:    0.2,
		SpecColor:    RGB{255, 255, 255},
		TransCoeff:   0.5,
		Phong:        1,
		ReflectCoeff: 0,
		RefractCoeff: 1.33,
	}
	blue := Material{
		AmbCoeff:     0.5,
		DiffCoeff:    0.4,
		DiffColor:    RGB{0, 155, 200},
		SpecCoeff:    0.6,
		SpecColor:    RGB{0, 155, 200},
		TransCoeff:   0,
		Phong:        50,
		ReflectCoeff: 0,
		RefractCoeff: 1.33,
	}

	out := C.Render(
		&[]Object{
			&Sphere{
				Loc: V3{5, 0, -10},
				Rad: 1,
				Mat: red,
			},
			&Sphere{
				Loc: V3{-5, 0, -10},
				Rad: 1.2,
				Mat: yellow,
			},
			&Plane{
				Points: [3]V3{V3{-3, -3, -20}, V3{3, -3, -20}, V3{-3, 3, -20}},
				Mat:    blue,
			},
			&Plane{
				Points: [3]V3{V3{3, 3, -20}, V3{-3, 3, -20}, V3{3, -3, -20}},
				Mat:    blue,
			},
			&Sphere{
				Loc: V3{0, 0, -12},
				Rad: .5,
				Mat: white,
			},
		},
		&[]Light{
			&DirLight{
				Color:   RGB{150, 150, 175},
				Dir:     V3{-1, -1, 0},
				MaxDist: 5000,
			},
			// &PointLight{
			// 	Color: RGB{255, 255, 255},
			// 	Loc:   V3{-100, 30, -15},
			// },
			&PointLight{
				Color: RGB{50, 50, 100},
				Loc:   V3{0, 0, -10},
			},
		})

	pngFile, err := os.Create("/Users/brandon/go/src/Projects/School/RayTracer/test.png")
	if err != nil {
		fmt.Printf("There was an error: %s", err.Error())
		return
	}

	out.WritePNG(pngFile)
}
