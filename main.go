package main

import (
	"fmt"
	"os"
)

func main() {
	C := Camera{
		Width:   1080,
		Height:  720,
		FOVx:    120,
		FOVy:    100,
		BGColor: RGB{0, 100, 200},
		Fpoint:  V3{0, 0, 0},
		Lpoint:  V3{0, 0, -1},
		Clip:    5000,
	}

	out := C.Render(
		[]Sphere{
			Sphere{
				Loc:   V3{0, 0, -10},
				Rad:   1,
				Color: RGB{200, 0, 0},
			},
			Sphere{
				Loc:   V3{3, 0, -15},
				Rad:   3,
				Color: RGB{10, 155, 10},
			}},
		[]DirLight{})

	pngFile, err := os.Create("/Users/brandon/go/src/Projects/School/RayTracer/test.png")
	if err != nil {
		fmt.Printf("There was an error: %s", err.Error())
		return
	}

	out.WritePNG(pngFile)
}
