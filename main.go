package main

import (
	"fmt"
	"os"
)

// const SUB_PIXELS = 1
// const DEPTH = 3
// const SHADOW_SAMPLES = 1
// const REFLECT_RAYS = 1
// const TRANS_RAYS = 1

const SUB_PIXELS = 3
const DEPTH = 3
const SHADOW_SAMPLES = 8
const REFLECT_RAYS = 8
const TRANS_RAYS = 8
const JITTER = 0.01

func main() {
	objs, lights, C := scene2() //scene1()

	medSplit := GenerateSplit(BoundingBox(*objs), 2, 5)
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
