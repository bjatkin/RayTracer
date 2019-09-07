package main

import (
	"math"
)

//Camera is a camera that can be added to the scene
type Camera struct {
	Fpoint       V3 //focal point
	Lpoint       V3 //look at point
	Width        int
	Height       int
	FOVx         float64
	FOVy         float64
	BGColor      RGB
	Clip         float64
	AmbientLight RGB
}

//Render renders the objects using the given camera
func (c Camera) Render(spheres *[]Sphere, lights *[]DirLight) *Image {
	out := NewImage(c.Width, c.Height) //the output of the render

	upVector, sideVector := c.stepVectors()

	for x := -c.Width / 2; x < c.Width/2; x++ {
		for y := -c.Height / 2; y < c.Height/2; y++ {
			//create a new ray pointing at the viewport
			r := Ray{
				Origin: c.Fpoint,
				Dest:   AddV3(AddV3(c.Lpoint, MulV3(float64(x), upVector)), MulV3(float64(y), sideVector)),
			}

			//Find the closest ray collision
			hDist := c.Clip
			color := c.BGColor
			for _, s := range *spheres {
				dist, hit, success := s.Intersect(&r)
				if success && dist < hDist {
					hDist = dist
					ptNormal := Unit(SubV3(hit, s.Loc))
					color = calculateColor(c.AmbientLight, hit, ptNormal, SubV3(c.Fpoint, hit), s.Mat, lights)
				}
			}

			out.SetPixel(x+c.Width/2, y+c.Height/2, color)
		}
	}

	return out
}

func (c Camera) stepVectors() (V3, V3) {
	look := SubV3(c.Fpoint, c.Lpoint)
	dist := look.Magnitude()

	radX := Rad(c.FOVx / 2)
	radY := Rad(c.FOVy / 2)
	rad90 := Rad(90)

	Vwidth := math.Abs(((dist * math.Sin(radX)) / (math.Sin(rad90 - radX))))
	Vheight := math.Abs(((dist * math.Sin(radY)) / (math.Sin(rad90 - radY))))

	upVector := V3{
		x: 0,
		y: Vheight / float64(c.Height),
		z: (-look.y * (Vheight / float64(c.Height)) / look.z),
	}
	sideVector := V3{
		x: Vwidth / float64(c.Width),
		y: 0,
		z: (-look.x * (Vwidth / float64(c.Width)) / look.z),
	}

	return upVector, sideVector
}

func calculateColor(ambLight RGB, point V3, normal V3, toView V3, mat Material, lights *[]DirLight) RGB {
	//Calculate the lighting portion of the lighting equation
	color := FlatMulV3(MulV3(mat.AmbCoeff, ambLight.V3()), mat.DiffColor.V3())
	diffCol := MulV3(mat.DiffCoeff, mat.DiffColor.V3())
	specCol := MulV3(mat.SpecCoeff, mat.SpecColor.V3())

	for _, l := range *lights {
		diffDir := DotV3(Unit(normal), Unit(l.Dir))
		diff := V3{}
		if diffDir > 0 {
			diff = MulV3(diffDir, diffCol)
		}

		specDir := DotV3(Unit(ReflectV3(l.Dir, normal)), Unit(toView))
		spec := V3{}
		if specDir > 0 {
			spec = MulV3(math.Pow(specDir, mat.Phong), specCol)
		}

		color = AddV3(color, FlatMulV3(l.Color.V3(), AddV3(diff, spec)))
	}

	return color.RGB()
}
