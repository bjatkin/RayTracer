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
		BGColor:      RGB{0, 0, 0},
		Fpoint:       V3{-7, 0, 15},
		Lpoint:       V3{-7, 0, 14},
		Clip:         5000,
		AmbientLight: RGB{255, 255, 255},
	}
	S1 := Material{
		AmbCoeff:     0.2,
		DiffCoeff:    0.4,
		DiffColor:    RGB{255, 255, 255},
		SpecCoeff:    0.6,
		SpecColor:    RGB{255, 255, 255},
		TransCoeff:   0.0,
		Phong:        150,
		ReflectCoeff: 0.6,
		RefractCoeff: 1.1,
	}
	S2 := Material{
		AmbCoeff:     0.0,
		DiffCoeff:    0.2,
		DiffColor:    RGB{255, 255, 255},
		SpecCoeff:    0.8,
		SpecColor:    RGB{255, 255, 255},
		TransCoeff:   1,
		Phong:        200,
		ReflectCoeff: 0.0,
		RefractCoeff: 0.965,
	}
	whiteWall := Material{
		AmbCoeff:     0.1,
		DiffCoeff:    0.8,
		DiffColor:    RGB{255, 255, 255},
		SpecCoeff:    0.2,
		SpecColor:    RGB{255, 255, 255},
		TransCoeff:   0,
		Phong:        10,
		ReflectCoeff: 0,
	}
	redWall := Material{
		AmbCoeff:     0.1,
		DiffCoeff:    0.8,
		DiffColor:    RGB{170, 30, 30},
		SpecCoeff:    0.2,
		SpecColor:    RGB{255, 255, 255},
		TransCoeff:   0,
		Phong:        10,
		ReflectCoeff: 0,
	}
	blueWall := Material{
		AmbCoeff:     0.1,
		DiffCoeff:    0.8,
		DiffColor:    RGB{0, 60, 170},
		SpecCoeff:    0.2,
		SpecColor:    RGB{255, 255, 255},
		TransCoeff:   0,
		Phong:        10,
		ReflectCoeff: 0,
	}

	out := C.Render(
		&[]Object{
			&Sphere{
				Loc: V3{-5, -5, -9},
				Rad: 3,
				Mat: S1,
			},
			&Sphere{
				Loc: V3{0, 5, -7},
				Rad: 3,
				Mat: S2,
			},
			&Plane{
				Points: [3]V3{V3{3, 10, -21}, V3{3, -10, -1}, V3{3, 10, -1}},
				Mat:    whiteWall,
			},
			&Plane{
				Points: [3]V3{V3{3, -10, -1}, V3{3, 10, -21}, V3{3, -10, -21}},
				Mat:    whiteWall,
			},
			&Plane{
				Points: [3]V3{V3{3, -10, -21}, V3{3, 10, -21}, V3{-17, 10, -21}},
				Mat:    whiteWall,
			},
			&Plane{
				Points: [3]V3{V3{-17, -10, -21}, V3{3, -10, -21}, V3{-17, 10, -21}},
				Mat:    whiteWall,
			},
			&Plane{
				Points: [3]V3{V3{3, -10, -1}, V3{3, -10, -21}, V3{-17, -10, -1}},
				Mat:    blueWall,
			},
			&Plane{
				Points: [3]V3{V3{-17, -10, -1}, V3{3, -10, -21}, V3{-17, -10, -21}},
				Mat:    blueWall,
			},
			&Plane{
				Points: [3]V3{V3{3, 10, -21}, V3{3, 10, -1}, V3{-17, 10, -1}},
				Mat:    redWall,
			},
			&Plane{
				Points: [3]V3{V3{3, 10, -21}, V3{-17, 10, -1}, V3{-17, 10, -21}},
				Mat:    redWall,
			},
			&Plane{
				Points: [3]V3{V3{-17, -10, -1}, V3{-17, 10, -21}, V3{-17, 10, -1}},
				Mat:    whiteWall,
			},
			&Plane{
				Points: [3]V3{V3{-17, 10, -21}, V3{-17, -10, -1}, V3{-17, -10, -21}},
				Mat:    whiteWall,
			},
		},
		&[]Light{
			// &DirLight{
			// 	Color:   RGB{255, 255, 255},
			// 	Dir:     V3{-1, -1, 0},
			// 	MaxDist: 5000,
			// },
			&PointLight{
				Color: RGB{255, 255, 240},
				Loc:   V3{-15, 0, -10},
			},
			// &PointLight{
			// 	Color: RGB{125, 175, 175},
			// 	Loc:   V3{-15, 2, -10},
			// },
			// &PointLight{
			// 	Color: RGB{125, 175, 175},
			// 	Loc:   V3{-15, -2, -10},
			// },
		}, 5)

	pngFile, err := os.Create("/Users/brandon/go/src/Projects/School/RayTracer/test.png")
	if err != nil {
		fmt.Printf("There was an error: %s", err.Error())
		return
	}

	out.WritePNG(pngFile)
}
