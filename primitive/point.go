package primitive

import "fmt"

// Point is a point in a 3D space
type Point struct {
	X, Y, Z float64
}

// NewPoint initialize a new point with given value
func NewPoint(X, Y, Z float64) Point {
	return Point{X, Y, Z}
}

// Addv adds a vector to the point
func (p Point) Add(o interface{}) Point {
	switch t := o.(type) {
	default:
		panic(fmt.Sprint("unable to add point with", t))
	case Vector:
		return Point{p.X + t.X, p.Y + t.Y, p.Z + t.Z}
	case Point:
		return Point{p.X + t.X, p.Y + t.Y, p.Z + t.Z}
	case float64:
		return Point{p.X + t, p.Y + t, p.Z + t}
	}
}

// Sub substract two points and give a new one
func (p Point) Sub(o Point) Point {
	return Point{p.X - o.X, p.Y - o.Y, p.Z - o.Z}
}

// Vec converts point to a vector
func (p Point) Vec() Vector {
	return Vector{p.X, p.Y, p.Z}
}
