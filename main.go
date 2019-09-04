package main

import "fmt"

func main() {
	s := Sphere{
		Loc: V3{10, 10, 10},
		Rad: 5,
	}

	r := Ray{
		Origin: V3{0, 0, 0},
		Dir:    V3{10, 10, 10},
	}

	fmt.Println(s.Intersect(r))
}
