package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func (c Camera) pathTrace(objects []split, lights *[]Light, destImage *Image) *Image {
	upVector, sideVector := c.stepVectors()

	type destination struct {
		pos V3
		x   int
		y   int
	}
	dests := []destination{}
	for x := -c.Width / 2; x < c.Width/2; x++ {
		for y := -c.Height / 2; y < c.Height/2; y++ {
			//create a new ray pointing at the viewport
			shiftUp := MulV3(float64(x), upVector)
			shiftSide := MulV3(float64(y), sideVector)
			dest := AddV3(AddV3(c.Lpoint, shiftUp), shiftSide)
			dests = append(dests,
				destination{
					pos: dest,
					x:   x + c.Width/2,
					y:   y + c.Height/2,
				},
			)
		}
	}

	//Track our progress
	progress := progressBar{
		total: c.Width * c.Height * pathCount,
		len:   70,
	}
	progress.Draw()

	//Shuffle the points
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for count, i := range r.Perm(len(dests)) {
		dest := dests[i]
		if destImage.GetPixel(dest.x, dest.y) != Black {
			continue //don't recalculate pixels that are filled in
		}

		color := V3{}
		colorChan := make(chan V3, pathGoRoutine)
		chanCount := 0
		for pCount := 0; pCount < pathCount; pCount++ {
			chanCount++
			go func() {
				path := newDestPath(c.Fpoint, JitterV3(jitter, dest.pos), c.Clip)
				col := path.Color(objects, lights, c.BGColor, depth)
				colorChan <- col.V3()
			}()

			//Pause before we run too many GoRoutines
			if chanCount >= pathGoRoutine {
				select {
				case c := <-colorChan:
					color = AddV3(color, c)
					chanCount--
					progress.Update()
					progress.Draw()
				}
			}
		}

		//Finish reading from the GoRoutines
		for chanCount > 0 {
			select {
			case c := <-colorChan:
				color = AddV3(color, c)
				chanCount--
				progress.Update()
				progress.Draw()
			}
		}

		colorAvg := MulV3(1/float64(pathCount), color)
		destImage.SetPixel(dest.x, dest.y, colorAvg.RGB())

		//Save a checkpoint
		if count%1000 == 0 {
			file, err := os.Create("checkpoints/" + destImage.name + "_PT_checkpoint.png")
			if err != nil {
				fmt.Printf("There was an error opening the checkpoint file, Not saving checkpoint\n")
			}
			destImage.WritePNG(file)
		}
	}
	return destImage
}
