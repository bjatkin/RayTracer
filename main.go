package main

import (
	"fmt"
	"os"
)

const SUB_PIXELS = 3
const DEPTH = 3
const SHADOW_SAMPLES = 5
const REFLECT_RAYS = 5
const TRANS_RAYS = 5
const JITTER = 0.0025
const PathDecay = 0.1
const PathAmbientLight = 0.5
const PathCount = 500
const PathGoRoutine = 20

func main() {
	objs, lights, C := scene1()

	drawType := "ray"
	fileName := "defult"
	if len(os.Args) > 1 {
		drawType = os.Args[1]
		fileName = os.Args[2]
	}
	medSplit := GenerateSplit(BoundingBox(*objs), 2, 5)

	if drawType == "both" || drawType == "ray" {
		fmt.Printf("starting ray trace\n")
		rayImage := NewImage(C.Width, C.Height, fileName)
		rayImage = C.Render(
			medSplit,
			lights,
			rayImage,
		)

		pngFile, err := os.Create("/Users/brandon/go/src/Projects/School/RayTracer/RT_" + fileName + ".png")
		if err != nil {
			fmt.Printf("There was an error: %s", err.Error())
			return
		}

		fmt.Printf("saving ray trace\n")
		rayImage.WritePNG(pngFile)

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
