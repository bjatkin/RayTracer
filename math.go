package main

//V3 is a 4D vector
type V3 struct {
	x, y, z float64
}

//V4 is a 4D vector
type V4 struct {
	x, y, z, w float64
}

//Mat is a general matrix
type Mat struct {
	x, y int8
	data []float64
}

//At returns the value at x,y in the matrix
func (m *Mat) At(x, y int8) float64 {
	return m.data[x*m.x+m.y]
}

//NewMat returns a new matrix with the given dimentions and data
func NewMat(x, y int8, data []float64) *Mat {
	return &Mat{
		x:    x,
		y:    y,
		data: data,
	}
}

func main() {}
