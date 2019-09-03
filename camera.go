package main

//Camera is a camera that can be added to the scene
type Camera struct {
	Fpoint V3 //focal point
	Lpoint V3 //look at point
	FOVx   float32
	FOVy   float32
}
