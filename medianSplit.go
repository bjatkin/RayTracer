package main

type split struct {
	bound   boundBox
	objects []*Object
}

func GenerateSplit(objs []*Object, objCount, splitCount int) []split {
	//Find the size of the full box
	ret := []split{}
	min := (*objs[0]).BoundBox().p1
	max := (*objs[0]).BoundBox().p2
	for _, o := range objs {
		bb := (*o).BoundBox()
		//Set new min V3
		if min.x > bb.p1.x {
			min.x = bb.p1.x
		}
		if min.y > bb.p1.y {
			min.y = bb.p1.y
		}
		if min.z > bb.p1.z {
			min.z = bb.p1.z
		}
		if min.x > bb.p2.x {
			min.x = bb.p2.x
		}
		if min.y > bb.p2.y {
			min.y = bb.p2.y
		}
		if min.z > bb.p2.z {
			min.z = bb.p2.z
		}

		//Set new max V3
		if max.x < bb.p1.x {
			max.x = bb.p1.x
		}
		if max.y < bb.p1.y {
			max.y = bb.p1.y
		}
		if max.z < bb.p1.z {
			max.z = bb.p1.z
		}
		if max.x < bb.p2.x {
			max.x = bb.p2.x
		}
		if max.y < bb.p2.y {
			max.y = bb.p2.y
		}
		if max.z < bb.p2.z {
			max.z = bb.p2.z
		}
	}
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

	box1 := boundBox{
		p1: min,
		p2: mid1,
	}
	split1 := split{bound: box1}
	box1Count := 0
	box2 := boundBox{
		p1: mid2,
		p2: max,
	}
	split2 := split{bound: box2}
	box2Count := 0
	//Check how many objects are in b1 and 2
	for _, o := range objs {
		if (*o).BoundBox().Overlap(box1) {
			//Add o to box 1
			split1.objects = append(split1.objects, o)
			box1Count++
		}
		if (*o).BoundBox().Overlap(box2) {
			//Add o to box 1
			split2.objects = append(split2.objects, o)
			box2Count++
		}
	}

	if box1Count < objCount || splitCount == 0 {
		//Add this region to the list of regions
		ret = append(ret, split1)
	} else {
		//Recurse down a level
		splits := GenerateSplit(split1.objects, objCount, splitCount-1)
		for _, s := range splits {
			ret = append(ret, s)
		}
	}

	if box2Count < objCount || splitCount == 0 {
		//Add this region to the list of regions
		ret = append(ret, split2)
	} else {
		//Recures down a level
		splits := GenerateSplit(split2.objects, objCount, splitCount-1)
		for _, s := range splits {
			ret = append(ret, s)
		}
	}

	return ret //The full data structure
}
