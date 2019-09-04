package main

import "fmt"

func main() {
	s := Sphere{
		Loc: V3{10, 10, 10},
		Rad: 5,
	}

	r := Ray{
		Origin: V3{2, 2, 2},
		Dest:   V3{5, 4.5, 5.5},
	}

	fmt.Println(s.Intersect(r))
}
