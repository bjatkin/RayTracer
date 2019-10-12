package main

import (
	"math"
	"math/rand"
)

const PathEpsilon = 0.000001

const pTypeOrigin = 0
const pTypeSpec = 1
const pTypeDiff = 2
const pTypeTrans = 3
const pTypeRefl = 4

type path struct {
	Origin          V3
	maxDist         float64
	t               float64
	dest, dir       V3
	destSet, dirSet bool
	pType           int
	child           *path
	parent          *path
}

func newDestPath(Origin, Dest V3, maxDist float64) path {
	return path{
		Origin:  Origin,
		dest:    Dest,
		destSet: true,
		maxDist: maxDist,
		pType:   pTypeOrigin,
	}
}

func newDirPath(Origin, Dir V3, maxDist float64) path {
	return path{
		Origin:  Origin,
		dir:     Unit(Dir),
		t:       Dir.Magnitude(),
		dirSet:  true,
		maxDist: maxDist,
		pType:   pTypeOrigin,
	}
}

func (p *path) T(t float64) {
	p.t = t
	p.destSet = false
}

func (p *path) Dest() V3 {
	if !p.destSet {
		//set the dest here
		p.dest = AddV3(p.Origin, MulV3(p.t, p.dir))
		p.destSet = true
	}
	return p.dest
}

func (p *path) Dir() V3 {
	if !p.dirSet {
		//set the dir here
		d := SubV3(p.dest, p.Origin)
		p.dir = Unit(d)
		p.t = d.Magnitude()
		p.dirSet = true
	}
	return MulV3(p.t, p.dir)
}

func (p *path) Unit() V3 {
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
		dir = JitterV3(jitter, dir)
	}

	ret := newDirPath(origin, dir, p.maxDist)
	ret.pType = pTypeRefl
	return ret
}

func (p path) Transmit(normal V3, refraction float64, jitter float64) path {
	I := p.Unit()
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
	origin := AddV3(p.Dest(), MulV3(PathEpsilon, dir))
	if jitter > 0 {
		dir = JitterV3(jitter, dir)
	}

	ret := newDirPath(origin, dir, p.maxDist)
	ret.pType = pTypeTrans
	return ret
}

func (p path) Diffuse(normal V3) path {
	dir := Unit(JitterV3(1-PathEpsilon, Unit(normal)))
	origin := AddV3(p.Dest(), MulV3(PathEpsilon, dir))

	ret := newDirPath(origin, dir, p.maxDist)
	ret.pType = pTypeDiff
	return ret
}

func (p path) Specular(normal V3, jitter float64) path {
	//Same as reflection but the light calculation will be different
	ret := p.Reflect(normal, jitter)
	ret.pType = pTypeSpec
	return ret
}

func (p *path) Color(objects medianSplit, lights *[]Light, background RGB, depth int) RGB {
	//Calculate my color
	var mat Material
	var normal V3
	t := p.maxDist

	itter := objects.itter(p)
	for itter.Next() {
		obj := itter.Obj()
		nextT, intersect := obj.IntersectPath(p)
		if intersect && nextT < t {
			t = nextT
			p.T(t)
			mat = obj.GetMat()
			normal = obj.Normal(p.Dest())
		}
	}

	light := false
	lightColor := RGB{}
	if p.pType != pTypeOrigin {
		for _, l := range *lights {
			nextT, intersect := l.Intersect(p)
			if intersect && nextT < t {
				t = nextT
				p.T(t)
				light = true
				lightColor = MulRGB(l.GetIntensity(), l.GetColor())
			}
		}
	}

	if light {
		return lightColor
	}

	//calculate lighting depending on the type of ray that I am
	if t == p.maxDist {
		if p.pType == pTypeOrigin || p.pType == pTypeTrans || p.pType == pTypeRefl {
			return background
		}

		return MulRGB(PathAmbientLight, White)
	}

	child := p.Next(mat, normal, JITTER)
	color := mat.DiffColor
	if child.pType == pTypeSpec || child.pType == pTypeRefl {
		color = mat.SpecColor
	}

	if depth == 0 {
		return MixRGB(MulRGB(PathAmbientLight, White), color)
	}

	cColor := child.Color(objects, lights, background, depth-1)
	if child.pType == pTypeTrans || child.pType == pTypeRefl {
		return cColor
	}

	return MixRGB(color, cColor)
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
		c = p.Diffuse(normal)
	case r < diff+spec:
		c = p.Specular(normal, jitter)
	case r < diff+spec+refl:
		c = p.Reflect(normal, jitter)
	default:
		c = p.Transmit(normal, mat.RefractCoeff, jitter)
	}

	p.child = &c
	c.parent = p
	return &c
}
