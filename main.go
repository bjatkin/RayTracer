package main

import (
	"fmt"
	"os"
)

const SUB_PIXELS = 3
const DEPTH = 5
const SHADOW_SAMPLES = 10
const REFLECT_RAYS = 10
const TRANS_RAYS = 10
const JITTER = 0.0025
const PathAmbientLight = 0.3
const PathCount = 500
const PathGoRoutine = 50

func main() {
	objs, lights, C := scene3() //scene2() //scene1()

	drawType := "ray"
	fileName := "defult"
	if len(os.Args) > 1 {
		drawType = os.Args[1]
		fileName = os.Args[2]
	}
	medSplit := GenerateSplit(BoundingBox(*objs), 2, 5)

	if drawType == "both" || drawType == "ray" {
		fmt.Printf("starting ray trace\n")
		rayImage := C.Render(
			medSplit,
			lights,
			objs,
		)

		pngFile2, err := os.Create("/Users/brandon/go/src/Projects/School/RayTracer/RT_" + fileName + ".png")
		if err != nil {
			fmt.Printf("There was an error: %s", err.Error())
			return
		}

		fmt.Printf("saving ray trace\n")
		rayImage.WritePNG(pngFile2)

	}

	if drawType == "both" || drawType == "path" {
		fmt.Printf("starting path trace\n")
		pathFile := NewImage(C.Width, C.Height, fileName)
		pathFile = C.pathTrace(
			medSplit,
			lights,
			pathFile,
		)

		pngFile, err := os.Create("/Users/brandon/go/src/Projects/School/RayTracer/PT_" + fileName + ".png")
		if err != nil {
			fmt.Printf("There was an error: %s", err.Error())
			return
		}

		fmt.Printf("saving path trace\n")
		pathFile.WritePNG(pngFile)

	}

}
