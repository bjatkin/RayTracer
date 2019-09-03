package main

import (
	"io"
	"strconv"
)

//PPM represents a ppm file
type PPM struct {
	width  int
	height int
	pixels []rgb
}

//NewPPM returns a new ppm object
func NewPPM(width, height int) *PPM {
	px := make([]rgb, width*height)
	return &PPM{
		width:  width,
		height: height,
		pixels: px,
	}
}

//Write writes out the PPM structur with correct file syntax to an io.Writer
func (p *PPM) Write(dest io.Writer) {

}

type rgb struct {
	R uint8
	G uint8
	B uint8
}

func (rgb rgb) string() string {
	return strconv.FormatInt(int64(rgb.R), 10) + " " +
		strconv.FormatInt(int64(rgb.G), 10) + " " +
		strconv.FormatInt(int64(rgb.B), 10)
}

//SetPixel colors a pixel at the given x, y coord
func (p *PPM) SetPixel(x, y int, R, G, B uint8) {
	p.pixels[x*p.width+y] = rgb{R, G, B}
}
