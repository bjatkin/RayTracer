package main

func scene1() (*[]Object, *[]Light, Camera) {
	C := Camera{
		Width:        1080,
		Height:       720,
		FOVx:         120,
		FOVy:         100,
		BGColor:      RGB{0, 0, 0},
		Fpoint:       V3{-7, 0, 15},
		Lpoint:       V3{-7, 0, 14},
		Clip:         5000,
		AmbientLight: RGB{255, 255, 255},
	}
	S1 := Material{
		AmbCoeff:     0.2,
		DiffCoeff:    0.4,
		DiffColor:    RGB{255, 255, 255},
		SpecCoeff:    0.6,
		SpecColor:    RGB{255, 255, 255},
		Phong:        150,
		ReflectCoeff: 0.6,
	}
	S2 := Material{
		AmbCoeff:     0.0,
		DiffCoeff:    0.2,
		DiffColor:    RGB{225, 225, 255},
		SpecCoeff:    0.8,
		SpecColor:    RGB{255, 255, 255},
		TransCoeff:   .9,
		Phong:        200,
		RefractCoeff: 0.9,
	}
	whiteWall := Material{
		AmbCoeff:  0.1,
		DiffCoeff: 0.8,
		DiffColor: RGB{255, 255, 255},
		SpecCoeff: 0.2,
		SpecColor: RGB{255, 255, 255},
		Phong:     10,
	}
	redWall := Material{
		AmbCoeff:  0.1,
		DiffCoeff: 0.8,
		DiffColor: RGB{170, 30, 30},
		SpecCoeff: 0.2,
		SpecColor: RGB{255, 255, 255},
		Phong:     10,
	}
	blueWall := Material{
		AmbCoeff:  0.1,
		DiffCoeff: 0.8,
		DiffColor: RGB{0, 60, 170},
		SpecCoeff: 0.2,
		SpecColor: RGB{255, 255, 255},
		Phong:     10,
	}
	objs := &[]Object{
		&Sphere{
			Loc: V3{-5, -5, -9},
			Rad: 3,
			Mat: S1,
		},
		&Sphere{
			Loc: V3{0, 4, -7},
			Rad: 4,
			Mat: S2,
		},
		&Plane{
			Points: [3]V3{V3{3, 10, -21}, V3{3, -10, -1}, V3{3, 10, -1}},
			Mat:    whiteWall,
		},
		&Plane{
			Points: [3]V3{V3{3, -10, -1}, V3{3, 10, -21}, V3{3, -10, -21}},
			Mat:    whiteWall,
		},
		&Plane{
			Points: [3]V3{V3{3, -10, -21}, V3{3, 10, -21}, V3{-17, 10, -21}},
			Mat:    whiteWall,
		},
		&Plane{
			Points: [3]V3{V3{-17, -10, -21}, V3{3, -10, -21}, V3{-17, 10, -21}},
			Mat:    whiteWall,
		},
		&Plane{
			Points: [3]V3{V3{3, -10, -1}, V3{3, -10, -21}, V3{-17, -10, -1}},
			Mat:    blueWall,
		},
		&Plane{
			Points: [3]V3{V3{-17, -10, -1}, V3{3, -10, -21}, V3{-17, -10, -21}},
			Mat:    blueWall,
		},
		&Plane{
			Points: [3]V3{V3{3, 10, -21}, V3{3, 10, -1}, V3{-17, 10, -1}},
			Mat:    redWall,
		},
		&Plane{
			Points: [3]V3{V3{3, 10, -21}, V3{-17, 10, -1}, V3{-17, 10, -21}},
			Mat:    redWall,
		},
		&Plane{
			Points: [3]V3{V3{-17, -10, -1}, V3{-17, 10, -21}, V3{-17, 10, -1}},
			Mat:    whiteWall,
		},
		&Plane{
			Points: [3]V3{V3{-17, 10, -21}, V3{-17, -10, -1}, V3{-17, -10, -21}},
			Mat:    whiteWall,
		},
	}
	lights := &[]Light{
		&AreaLight{
			Color: RGB{255, 255, 255},
			Area:  Plane{Points: [3]V3{V3{-16, 3, -10}, V3{-16, -3, -10}, V3{-16, -3, -15}}},
		},
		&AreaLight{
			Color: RGB{255, 255, 255},
			Area:  Plane{Points: [3]V3{V3{-16, 3, -15}, V3{-16, -3, -15}, V3{-16, 3, -10}}},
		},
	}

	return objs, lights, C
}

func scene2() (*[]Object, *[]Light, Camera) {
	C := Camera{
		Width:        1080,
		Height:       720,
		FOVx:         120,
		FOVy:         100,
		BGColor:      RGB{5, 15, 15},
		Fpoint:       V3{-7, 0, 15},
		Lpoint:       V3{-7, 0, 14},
		Clip:         5000,
		AmbientLight: RGB{255, 255, 255},
	}
	whiteTile := Material{
		AmbCoeff:     0.2,
		DiffCoeff:    0.6,
		DiffColor:    RGB{250, 250, 250},
		SpecCoeff:    0.4,
		SpecColor:    RGB{255, 255, 255},
		Phong:        150,
		ReflectCoeff: 0.5,
	}
	blackTile := Material{
		AmbCoeff:     0.2,
		DiffCoeff:    0.6,
		DiffColor:    RGB{25, 25, 25},
		SpecCoeff:    0.4,
		SpecColor:    RGB{255, 255, 255},
		Phong:        150,
		ReflectCoeff: 0.7,
	}
	S1 := Material{
		AmbCoeff:     0.2,
		DiffCoeff:    0.4,
		DiffColor:    RGB{255, 5, 55},
		SpecCoeff:    0.6,
		SpecColor:    RGB{255, 255, 255},
		Phong:        150,
		ReflectCoeff: 0.6,
	}
	S2 := Material{
		AmbCoeff:     0.0,
		DiffCoeff:    0.2,
		DiffColor:    RGB{200, 200, 255},
		SpecCoeff:    0.8,
		SpecColor:    RGB{255, 255, 255},
		TransCoeff:   .9,
		Phong:        200,
		RefractCoeff: 0.9,
	}
	S3 := Material{
		AmbCoeff:  0.2,
		DiffCoeff: 1,
		DiffColor: RGB{5, 255, 255},
		SpecCoeff: 0,
		SpecColor: RGB{255, 255, 255},
		Phong:     20,
	}

	objs := &[]Object{
		&Sphere{
			Loc: V3{2, -5, -10},
			Rad: 3,
			Mat: S1,
		},
		&Sphere{
			Loc: V3{0, 4, -7},
			Rad: 4,
			Mat: S2,
		},
		&Sphere{
			Loc: V3{-2, 0, -18},
			Rad: 7,
			Mat: S3,
		},

		//Row 1
		&Plane{
			Points: [3]V3{V3{5, -18, -4}, V3{5, -6, -4}, V3{5, -18, -10}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -10}, V3{5, -18, -10}, V3{5, -6, -4}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -4}, V3{5, 6, -4}, V3{5, 6, -10}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -4}, V3{5, 6, -10}, V3{5, -6, -10}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 6, -4}, V3{5, 18, -4}, V3{5, 18, -10}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 18, -10}, V3{5, 6, -10}, V3{5, 6, -4}},
			Mat:    blackTile,
		},

		//Row 2
		&Plane{
			Points: [3]V3{V3{5, -18, -10}, V3{5, -6, -10}, V3{5, -18, -16}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -16}, V3{5, -18, -16}, V3{5, -6, -10}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -10}, V3{5, 6, -10}, V3{5, 6, -16}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -10}, V3{5, 6, -16}, V3{5, -6, -16}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 6, -10}, V3{5, 18, -10}, V3{5, 18, -16}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 18, -16}, V3{5, 6, -16}, V3{5, 6, -10}},
			Mat:    whiteTile,
		},

		//Row 3
		&Plane{
			Points: [3]V3{V3{5, -18, -16}, V3{5, -6, -16}, V3{5, -18, -22}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -22}, V3{5, -18, -22}, V3{5, -6, -16}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -16}, V3{5, 6, -16}, V3{5, 6, -22}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -16}, V3{5, 6, -22}, V3{5, -6, -22}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 6, -16}, V3{5, 18, -16}, V3{5, 18, -22}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 18, -22}, V3{5, 6, -22}, V3{5, 6, -16}},
			Mat:    blackTile,
		},

		//Row 4
		&Plane{
			Points: [3]V3{V3{5, -18, -22}, V3{5, -6, -22}, V3{5, -18, -28}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -28}, V3{5, -18, -28}, V3{5, -6, -22}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -22}, V3{5, 6, -22}, V3{5, 6, -28}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -22}, V3{5, 6, -28}, V3{5, -6, -28}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 6, -22}, V3{5, 18, -22}, V3{5, 18, -28}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 18, -28}, V3{5, 6, -28}, V3{5, 6, -22}},
			Mat:    whiteTile,
		},
	}

	lights := &[]Light{
		&AreaLight{
			Color: RGB{200, 200, 200},
			Area:  Plane{Points: [3]V3{V3{-16, 3, -10}, V3{-16, -3, -10}, V3{-16, -3, -15}}},
		},
		&AreaLight{
			Color: RGB{200, 200, 200},
			Area:  Plane{Points: [3]V3{V3{-16, 3, -15}, V3{-16, -3, -15}, V3{-16, 3, -10}}},
		},
	}

	return objs, lights, C
}

func scene3() (*[]Object, *[]Light, Camera) {
	C := Camera{
		Width:        300,
		Height:       300,
		FOVx:         100,
		FOVy:         100,
		BGColor:      RGB{5, 5, 5},
		Fpoint:       V3{-7, 0, 15},
		Lpoint:       V3{-7, 0, 14},
		Clip:         5000,
		AmbientLight: RGB{255, 255, 255},
	}

	whiteTile := Material{
		AmbCoeff:     0.2,
		DiffCoeff:    0.6,
		DiffColor:    RGB{250, 250, 250},
		SpecCoeff:    0.4,
		SpecColor:    RGB{255, 255, 255},
		Phong:        150,
		ReflectCoeff: 0.5,
	}
	blackTile := Material{
		AmbCoeff:     0.0,
		DiffCoeff:    0.6,
		DiffColor:    RGB{2, 2, 2},
		SpecCoeff:    0.4,
		SpecColor:    RGB{255, 255, 255},
		Phong:        150,
		ReflectCoeff: 0.4,
	}
	S1 := Material{
		AmbCoeff:     0.2,
		DiffCoeff:    0.4,
		DiffColor:    RGB{150, 255, 0},
		SpecCoeff:    0.6,
		SpecColor:    RGB{255, 255, 255},
		Phong:        150,
		ReflectCoeff: 0.6,
	}
	S2 := Material{
		AmbCoeff:     0.1,
		DiffCoeff:    0.2,
		DiffColor:    RGB{200, 100, 0},
		SpecCoeff:    0.8,
		SpecColor:    RGB{255, 255, 255},
		TransCoeff:   0.5,
		Phong:        200,
		RefractCoeff: 0.8,
	}
	S3 := Material{
		AmbCoeff:  0.2,
		DiffCoeff: 0.7,
		DiffColor: RGB{5, 255, 255},
		SpecCoeff: 0.3,
		SpecColor: RGB{255, 255, 255},
		Phong:     20,
	}

	objs := &[]Object{
		&Sphere{
			Loc: V3{2, -5, -10},
			Rad: 3,
			Mat: S1,
		},
		&Sphere{
			Loc: V3{0, 4, -7},
			Rad: 4,
			Mat: S2,
		},
		&Sphere{
			Loc: V3{-2, 0, -18},
			Rad: 7,
			Mat: S3,
		},

		//Row 1
		&Plane{
			Points: [3]V3{V3{5, -18, -4}, V3{5, -6, -4}, V3{5, -18, -10}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -10}, V3{5, -18, -10}, V3{5, -6, -4}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -4}, V3{5, 6, -4}, V3{5, 6, -10}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -4}, V3{5, 6, -10}, V3{5, -6, -10}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 6, -4}, V3{5, 18, -4}, V3{5, 18, -10}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 18, -10}, V3{5, 6, -10}, V3{5, 6, -4}},
			Mat:    blackTile,
		},

		//Row 2
		&Plane{
			Points: [3]V3{V3{5, -18, -10}, V3{5, -6, -10}, V3{5, -18, -16}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -16}, V3{5, -18, -16}, V3{5, -6, -10}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -10}, V3{5, 6, -10}, V3{5, 6, -16}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -10}, V3{5, 6, -16}, V3{5, -6, -16}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 6, -10}, V3{5, 18, -10}, V3{5, 18, -16}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 18, -16}, V3{5, 6, -16}, V3{5, 6, -10}},
			Mat:    whiteTile,
		},

		//Row 3
		&Plane{
			Points: [3]V3{V3{5, -18, -16}, V3{5, -6, -16}, V3{5, -18, -22}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -22}, V3{5, -18, -22}, V3{5, -6, -16}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -16}, V3{5, 6, -16}, V3{5, 6, -22}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -16}, V3{5, 6, -22}, V3{5, -6, -22}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 6, -16}, V3{5, 18, -16}, V3{5, 18, -22}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 18, -22}, V3{5, 6, -22}, V3{5, 6, -16}},
			Mat:    blackTile,
		},

		//Row 4
		&Plane{
			Points: [3]V3{V3{5, -18, -22}, V3{5, -6, -22}, V3{5, -18, -28}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -28}, V3{5, -18, -28}, V3{5, -6, -22}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -22}, V3{5, 6, -22}, V3{5, 6, -28}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, -6, -22}, V3{5, 6, -28}, V3{5, -6, -28}},
			Mat:    blackTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 6, -22}, V3{5, 18, -22}, V3{5, 18, -28}},
			Mat:    whiteTile,
		},
		&Plane{
			Points: [3]V3{V3{5, 18, -28}, V3{5, 6, -28}, V3{5, 6, -22}},
			Mat:    whiteTile,
		},
	}

	lights := &[]Light{
		&AreaLight{
			Color: RGB{100, 100, 100},
			Area:  Plane{Points: [3]V3{V3{-12, 10, -7}, V3{-12, -10, -7}, V3{-12, -10, -19}}},
		},
		&AreaLight{
			Color: RGB{100, 100, 100},
			Area:  Plane{Points: [3]V3{V3{-12, 10, -19}, V3{-12, -10, -19}, V3{-12, 10, -7}}},
		},
	}

	return objs, lights, C
}
