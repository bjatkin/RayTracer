package main

import (
	"fmt"
	"os"
)

const (
	subPixles        = 3
	depth            = 3
	shadowSamples    = 5
	reflectRays      = 5
	transparentRays  = 5
	jitter           = 0.0025
	pathDecay        = 0.1
	pathAmbientLight = 0.5
	pathCount        = 500
	pathGoRoutine    = 20
)

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

		pngFile, err := os.Create("RT_" + fileName + ".png")
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

		pngFile, err := os.Create("PT_" + fileName + ".png")
		if err != nil {
			fmt.Printf("There was an error: %s", err.Error())
			return
		}

		fmt.Printf("saving path trace\n")
		pathFile.WritePNG(pngFile)

	}

}
