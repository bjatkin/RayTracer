package main

type medianSplit []Split

func (split medianSplit) itter(p *path) splitItter {
	ret := splitItter{oIndex: -1}
	for _, s := range split {
		if s.bound.IntersectPath(p) {
			ret.splits = append(ret.splits, s)
		}
	}
	return ret
}

type splitItter struct {
	splits []Split
	sIndex int
	oIndex int
}

func splitItterable(splits []Split, r *Ray) splitItter {
	ret := splitItter{oIndex: -1}
	for _, s := range splits {
		if s.bound.Intersect(r) {
			ret.splits = append(ret.splits, s)
		}
	}

	return ret
}

func (si *splitItter) Next() bool {
	if len(si.splits) == 0 {
		return false
	}
	si.oIndex++
	if len(si.splits[si.sIndex].objects) <= si.oIndex {
		si.sIndex++
		si.oIndex = 0
	}
	if len(si.splits) <= si.sIndex {
		return false
	}
	if len(si.splits[si.sIndex].objects) == 0 {
		return false
	}
	return true
}

func (si *splitItter) Obj() Object {
	ret := si.splits[si.sIndex].objects[si.oIndex]
	return ret
}

// Split is an oct-tree structure used to subdivide the
// scean for faster rendering
type Split struct {
	bound   BoundBox
	objects []Object
}

// BoundingBox split the object into nwew smaller bounding boxes
func BoundingBox(objs []Object) Split {
	start := objs[0].BoundBox()
	min := start.p1
	max := start.p2

	for _, o := range objs {
		bb := o.BoundBox()
		min = MinV3(min, bb.p1)
		min = MinV3(min, bb.p2)

		max = MaxV3(max, bb.p1)
		max = MaxV3(max, bb.p2)
	}

	return Split{
		bound: BoundBox{
			p1: min,
			p2: max,
		},
		objects: objs,
	}
}

// SplitBox splits a bounding box, and it's objects, into 2 sub-boxes
func SplitBox(b BoundBox) (BoundBox, BoundBox) {
	min, max := b.p1, b.p2
	magX := max.x - min.x
	magY := max.y - min.y
	magZ := max.z - max.z
	diff := SubV3(max, min)

	//Split the box along it's greatest axis
	if magX > magY && magX > magZ {
		diff.x *= 0.5
	}
	if magY > magX && magY > magZ {
		diff.y *= 0.5
	}
	if magZ > magY && magZ > magX {
		diff.z *= 0.5
	}

	return BoundBox{
			p1: SubV3(max, diff),
			p2: max,
		},
		BoundBox{
			p1: min,
			p2: AddV3(min, diff),
		}
}

// GenerateSplit splits a scean into several smaller
// sections making rendering faster
func GenerateSplit(start Split, objCount, splitCount int) []Split {
	//Split the box along it's greatest axis
	box1, box2 := SplitBox(start.bound)
	split1, split2 := Split{bound: box1}, Split{bound: box2}

	//Check how many objects are in b1 and 2
	for _, o := range start.objects {
		if o.BoundBox().Overlap(box1) {
			//Add o to box 1
			split1.objects = append(split1.objects, o)
		}
		if o.BoundBox().Overlap(box2) {
			//Add o to box 2
			split2.objects = append(split2.objects, o)
		}
	}

	ret := []Split{}
	if len(split1.objects) < objCount || splitCount == 0 {
		//Add this region to the list of regions
		ret = appendSplit(ret, split1)
	} else {
		//Recurse down a level
		splits := GenerateSplit(split1, objCount, splitCount-1)
		for _, s := range splits {
			ret = appendSplit(ret, s)
		}
	}

	if len(split2.objects) < objCount || splitCount == 0 {
		//Add this region to the list of regions
		ret = appendSplit(ret, split2)
	} else {
		//Recures down a level
		splits := GenerateSplit(split2, objCount, splitCount-1)
		for _, s := range splits {
			ret = appendSplit(ret, s)
		}
	}

	return ret //The full data structure
}

func appendSplit(a []Split, b Split) []Split {
	if len(b.objects) > 0 {
		a = append(a, b)
	}
	return a
}
