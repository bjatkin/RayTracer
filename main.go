package main

import (
	"fmt"
	"os"
)

const SUB_PIXELS = 1
const DEPTH = 5
const SHADOW_SAMPLES = 1
const REFLECT_RAYS = 1
const TRANS_RAYS = 1

func main() {
	objs, lights, C := scene2() //scene1()

	medSplit := GenerateSplit(BoundingBox(*objs), 2, 1)
	out := C.Render(
		medSplit,
		lights,
		objs,
	)

	pngFile, err := os.Create("/Users/brandon/go/src/Projects/School/RayTracer/test.png")
	if err != nil {
		fmt.Printf("There was an error: %s", err.Error())
		return
	}

	out.WritePNG(pngFile)
}
