package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

//PPM represents a ppm file
type PPM struct {
	width  int
	height int
	pixels []RGB
}

//NewPPM returns a new ppm object
func NewPPM(width, height int) *PPM {
	px := make([]RGB, width*height)
	return &PPM{
		width:  width,
		height: height,
		pixels: px,
	}
}

//Write writes out the PPM structur with correct file syntax to an io.Writer
func (p *PPM) Write(dest io.Writer) {
	var pxs []string
	for _, px := range p.pixels {
		pxs = append(pxs, px.String())
	}
	dest.Write([]byte(
		"P3\n" +
			fmt.Sprintf("%d %d\n", p.width, p.height) +
			"255\n" +
			strings.Join(pxs, " "),
	))
}

//RGB is an rgb color with rgb values between 0 and 255
type RGB struct {
	R uint8
	G uint8
	B uint8
}

func (rgb RGB) String() string {
	return strconv.FormatInt(int64(rgb.R), 10) + " " +
		strconv.FormatInt(int64(rgb.G), 10) + " " +
		strconv.FormatInt(int64(rgb.B), 10)
}

//SetPixel colors a pixel at the given x, y coord
func (p *PPM) SetPixel(x, y int, rgb RGB) {
	p.pixels[x*p.width+y] = rgb
}
