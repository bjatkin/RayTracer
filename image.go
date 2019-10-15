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
	name   string
}

//NewImage returns a new image object
func NewImage(width, height int, name string) *Image {
	px := make([]RGB, width*height)
	return &Image{
		width:  width,
		height: height,
		pixels: px,
		name:   name,
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

			pxl.Clamp() //prevent wierd blowout
			out.Set(x, y, color.NRGBA{
				R: uint8(pxl.R),
				G: uint8(pxl.G),
				B: uint8(pxl.B),
				A: 255,
			})
		}
	}

	png.Encode(dest, out)
}

//RGB is an rgb color with rgb values between 0 and 255
type RGB struct {
	R float64
	G float64
	B float64
}

var White = RGB{R: 255, G: 255, B: 255}
var Black = RGB{R: 0, G: 0, B: 0}

func (rgb RGB) String() string {
	return strconv.FormatInt(int64(rgb.R), 10) + " " +
		strconv.FormatInt(int64(rgb.G), 10) + " " +
		strconv.FormatInt(int64(rgb.B), 10)
}

func AddRGB(a, b RGB) RGB {
	ret := RGB{}
	ret.R = a.R + b.R
	ret.G = a.G + b.G
	ret.B = a.B + b.B
	// ret.Clamp()
	return ret
}

func MulRGB(scale float64, rgb RGB) RGB {
	ret := RGB{}
	ret.R = rgb.R * scale
	ret.G = rgb.G * scale
	ret.B = rgb.B * scale
	// ret.Clamp()
	return ret
}

//MixRGB mixes colors a and b with the given weight + for a and - for b
func MixRGB(a, b RGB, weight float64) RGB {
	//TODO implement this method
	ret := RGB{}
	ret.R = (a.R / 255.0) * (b.R / 255.0) * 255.0
	ret.G = (a.G / 255.0) * (b.G / 255.0) * 255.0
	ret.B = (a.B / 255.0) * (b.B / 255.0) * 255.0
	// ret.Clamp()
	return ret
}

func (rgb *RGB) Clamp() {
	if rgb.R > 255 {
		rgb.R = 255
	}
	if rgb.R < 0 {
		rgb.R = 0
	}
	if rgb.G > 255 {
		rgb.G = 255
	}
	if rgb.G < 0 {
		rgb.G = 0
	}
	if rgb.B > 255 {
		rgb.B = 255
	}
	if rgb.B < 0 {
		rgb.B = 0
	}
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
