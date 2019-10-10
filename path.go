package main

import (
	"math"
	"math/rand"
)

const PathEpsilon = 0.0000001

type path struct {
	Origin          V3
	t               float64
	dest, dir       V3
	destSet, dirSet bool
	child           *path
}

func newDestPath(Origin, Dest V3) path {
	return path{
		Origin:  Origin,
		dest:    Dest,
		destSet: true,
	}
}

func newDirPath(Origin, Dir V3) path {
	return path{
		Origin: Origin,
		dir:    Unit(Dir),
		t:      Dir.Magnitude(),
		dirSet: true,
	}
}

func (p path) Dest() V3 {
	if !p.destSet {
		//set the dest here
		p.dest = AddV3(p.Origin, MulV3(p.t, p.dir))
		p.destSet = true
	}
	return p.dest
}

func (p path) Dir() V3 {
	if !p.dirSet {
		//set the dir here
		d := SubV3(p.dest, p.Origin)
		p.dir = Unit(d)
		p.t = d.Magnitude()
		p.dirSet = true
	}
	return MulV3(p.t, p.dir)
}

func (p path) Unit() V3 {
	//the unit vector repesentation of the path
	if !p.dirSet {
		p.Dir() //set dir
	}
	return p.dir
}

func (p path) Reflect(normal V3, jitter float64) path {
	dir := Unit(ReflectV3(p.Dir(), normal))
	origin := AddV3(p.Dest(), MulV3(PathEpsilon, dir))
	if jitter > 0 {
		dir = JitterV3(dir, jitter)
	}

	return newDirPath(origin, dir)
}

func (p path) Transmit(normal V3, refraction float64, jitter float64) path {
	I := Unit(p.Dir())
	N := Unit(normal)
	cos := DotV3(I, N)
	if cos < 0 {
		cos *= -1
	} else {
		N = MulV3(-1, N)
	}
	nit := refraction

	p1 := MulV3(nit, I)
	p2 := math.Sqrt(1 + nit*nit*(cos*cos-1))
	p3 := nit*cos - p2

	dir := AddV3(p1, MulV3(p3, N))
	origin := AddV3(p.Origin, MulV3(PathEpsilon, dir))
	if jitter > 0 {
		dir = JitterV3(dir, jitter)
	}

	return newDirPath(origin, dir)
}

func (p path) Diffuse() path {
	dir := Unit(JitterV3(p.Unit(), 1-PathEpsilon))
	origin := AddV3(p.Origin, MulV3(PathEpsilon, dir))

	return newDirPath(origin, dir)
}

func (p path) Specular(normal V3, jitter float64) path {
	//Same as reflection but the light calculation will be different?
	dir := Unit(ReflectV3(p.Dir(), normal))
	origin := AddV3(p.Dest(), MulV3(PathEpsilon, dir))
	if jitter > 0 {
		dir = JitterV3(dir, jitter)
	}

	return newDirPath(origin, dir)
}

//TODO finish this method
func (p *path) Color(objects medianSplit, lights []*Light, depth int) RGB {
	//Calculate my color
	cur := p
	for x := 0; x < depth; x++ {
		//Intersect the path against all the objects
		mat := Material{}
		normal := V3{}
		cur = cur.Next(mat, normal, JITTER)
	}

	//Add in my child's color
	return RGB{}
}

func (p *path) Next(mat Material, normal V3, jitter float64) *path {
	//squash all the coefficients
	total := mat.DiffCoeff + mat.ReflectCoeff + mat.TransCoeff + mat.SpecCoeff
	diff := mat.DiffCoeff / total
	spec := mat.SpecCoeff / total
	refl := mat.ReflectCoeff / total

	r := rand.Float64()
	c := path{}
	switch {
	case r < diff:
		c = p.Diffuse()
	case r < spec:
		c = p.Specular(normal, jitter)
	case r < refl:
		c = p.Reflect(normal, jitter)
	default:
		c = p.Transmit(normal, mat.RefractCoeff, jitter)
	}

	p.child = &c
	return &c
}
