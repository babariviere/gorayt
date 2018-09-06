package primitive

// Ray is a line in space with an origin and a direction
type Ray struct {
	Origin    Point
	Direction Vector
}

// NewRay initialize a new ray with given value
func NewRay(Origin Point, Direction Vector) Ray {
	return Ray{Origin, Direction}
}

// PointAt gives a point along the ray at factor t
func (r Ray) PointAt(t float64) Point {
	return r.Origin.Addv(r.Direction.Mulf(t))
}
