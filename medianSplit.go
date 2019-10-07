package main

type splitItter struct {
	splits []split
	sIndex int
	oIndex int
}

func splitItterable(splits []split, r *Ray) splitItter {
	ret := splitItter{oIndex: -1}
	for _, s := range splits {
		if s.bound.Intersect(r) {
			ret.splits = append(ret.splits, s)
		}
	}

	return ret
}

func (si splitItter) Next() bool {
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
	return true
}

func (si splitItter) Obj() Object {
	return si.splits[si.sIndex].objects[si.oIndex]
}

type split struct {
	bound   boundBox
	objects []Object
}

func BoundingBox(objs []Object) split {
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

	return split{
		bound: boundBox{
			p1: min,
			p2: max,
		},
		objects: objs,
	}
}

func SplitBox(b boundBox) (boundBox, boundBox) {
	min, max := b.p1, b.p2
	magX := max.x - min.x
	magY := max.y - min.y
	magZ := max.z - max.z

	//Split the box along it's greatest axis
	mid1 := min
	mid2 := max
	if magX > magY && magX > magZ {
		mid1.x += magX / 2
		mid2.x -= magX / 2
	}
	if magY > magX && magY > magZ {
		mid1.y += magY / 2
		mid2.y -= magY / 2
	}
	if magZ > magY && magZ > magX {
		mid1.z += magZ / 2
		mid2.z -= magZ / 2
	}

	return boundBox{
			p1: min,
			p2: mid1,
		},
		boundBox{
			p1: mid2,
			p2: max,
		}
}

func GenerateSplit(start split, objCount, splitCount int) []split {
	//Split the box along it's greatest axis
	box1, box2 := SplitBox(start.bound)
	split1, split2 := split{bound: box1}, split{bound: box2}

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

	ret := []split{}
	if len(split1.objects) < objCount || splitCount == 0 {
		//Add this region to the list of regions
		ret = append(ret, split1)
	} else {
		//Recurse down a level
		splits := GenerateSplit(split1, objCount, splitCount-1)
		for _, s := range splits {
			ret = append(ret, s)
		}
	}

	if len(split2.objects) < objCount || splitCount == 0 {
		//Add this region to the list of regions
		ret = append(ret, split2)
	} else {
		//Recures down a level
		splits := GenerateSplit(split2, objCount, splitCount-1)
		for _, s := range splits {
			ret = append(ret, s)
		}
	}

	return ret //The full data structure
}
