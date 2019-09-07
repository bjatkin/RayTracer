package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"strconv"
	"strings"
)

//Image represents a image file
type Image struct {
	width  int
	height int
	pixels []RGB
}

//NewImage returns a new image object
func NewImage(width, height int) *Image {
	px := make([]RGB, width*height)
	return &Image{
		width:  width,
		height: height,
		pixels: px,
	}
}

//WritePPM writes out the Image as a vaild ppm file to an io.Writer
func (p *Image) WritePPM(dest io.Writer) {
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

//WritePNG wirets out the image as a valid png file to the io.Writer
func (p *Image) WritePNG(dest io.Writer) {
	out := image.NewNRGBA(image.Rect(0, 0, p.width, p.height))

	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			pxl := p.GetPixel(x, y)

			out.Set(x, y, color.NRGBA{
				R: pxl.R, //uint8((x + y) & 255),
				G: pxl.G, //uint8((x + y) << 1 & 255),
				B: pxl.B, //uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}

	png.Encode(dest, out)
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

//V3 converts an RGB to a V3
func (rgb RGB) V3() V3 {
	return V3{
		x: float64(rgb.R) / 255,
		y: float64(rgb.G) / 255,
		z: float64(rgb.B) / 255,
	}
}

//SetPixel colors a pixel at the given x, y coord
func (p *Image) SetPixel(x, y int, rgb RGB) {
	p.pixels[y*p.width+x] = rgb
}

//GetPixel returns the rgb value of the pixel at x, y coord
func (p *Image) GetPixel(x, y int) RGB {
	return p.pixels[y*p.width+x]
}
