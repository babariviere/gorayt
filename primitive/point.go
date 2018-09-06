package primitive

// Point is a point in a 3D space
type Point struct {
	X, Y, Z float64
}

// NewPoint initialize a new point with given value
func NewPoint(X, Y, Z float64) Point {
	return Point{X, Y, Z}
}

/// Addv adds a vector to the point
func (p Point) Addv(v Vector) Point {
	return Point{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}
