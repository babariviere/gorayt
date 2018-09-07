package primitive

// Camera is the camera position in world
type Camera struct {
	LowerLeft  Vector
	Horizontal Vector
	Vertical   Vector
	Origin     Point
	Width      int
	Height     int
}

// DefaultCamera creates a default camera
func DefaultCamera() Camera {
	return Camera{
		LowerLeft:  NewVector(-2, -1, -1),
		Horizontal: NewVector(4, 0, 0),
		Vertical:   NewVector(0, 2, 0),
		Origin:     NewPoint(0, 0, 0),
		Width:      200,
		Height:     100,
	}
}

// GetRay gets ray at position x, y
func (c Camera) GetRay(x, y int) Ray {
	u := float64(x) / float64(c.Width)
	v := float64(y) / float64(c.Height)
	hor := c.Horizontal.Mul(u)
	ver := c.Vertical.Mul(v)
	horver := hor.Add(ver)
	return NewRay(c.Origin, c.LowerLeft.Add(horver))
}
