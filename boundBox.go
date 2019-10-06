package main

type boundBox struct {
	p1, p2 V3
}

func (bb boundBox) Overlap(ob boundBox) bool {
	return bb.p1.x < ob.p2.x && ob.p1.x < bb.p2.x &&
		bb.p1.y < ob.p2.y && ob.p1.y < bb.p2.y &&
		bb.p1.z < ob.p2.z && ob.p1.z < bb.p2.z
}

func (bb boundBox) Intersect(r *Ray) bool {
	tmin := (bb.p1.x - r.Origin.x) / r.Dir().x
	tmax := (bb.p2.x - r.Origin.x) / r.Dir().x

	if tmin > tmax {
		tmin, tmax = tmax, tmin //Swap
	}

	tymin := (bb.p1.y - r.Origin.y) / r.Dir().y
	tymax := (bb.p2.y - r.Origin.y) / r.Dir().y

	if tymin > tymax {
		tymin, tymax = tymax, tymin //Swap
	}

	if (tmin > tymax) || (tymin > tmax) {
		return false
	}

	if tymin > tmin {
		tmin = tymin
	}

	if tymax < tmax {
		tmax = tymax
	}

	tzmin := (bb.p1.z - r.Origin.z) / r.Dir().z
	tzmax := (bb.p2.z - r.Origin.z) / r.Dir().z

	if tzmin > tzmax {
		tzmin, tzmax = tzmax, tzmin
	}

	if (tmin > tzmax) || (tzmin > tmax) {
		return false
	}

	if tzmin > tmin {
		tmin = tzmin
	}

	if tzmax < tmax {
		tmax = tzmax
	}

	return true
}
