package primitive

import (
	"math"
	"math/rand"
	"time"
)

// Camera is the camera position in world
type Camera struct {
	LowerLeft  Vector
	Horizontal Vector
	Vertical   Vector
	Origin     Point
	Width      int
	Height     int
	LensRadius float64
	u, v, w    Vector
}

// DefaultCamera creates a default camera
func DefaultCamera() Camera {
	return Camera{
		LowerLeft:  NewVector(-2, -1, -1),
		Horizontal: NewVector(4, 0, 0),
		Vertical:   NewVector(0, 2, 0),
		Origin:     NewPoint(0, 0, 0),
		Width:      1920,
		Height:     1080,
		LensRadius: 0.5,
	}
}

// NewCamera creates a new camera with parameters
func NewCamera(lookFrom, lookAt Point, vUp Vector, vfov float64, width, height int, aperture float64) (c Camera) {
	c.LensRadius = aperture / 2
	c.Width = width
	c.Height = height
	aspect := float64(width) / float64(height)
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := halfHeight * aspect
	c.Origin = lookFrom
	w := lookFrom.Vec().Sub(lookAt).Normalize()
	u := vUp.Cross(w).Normalize()
	v := w.Cross(u)
	c.u, c.v, c.w = u, v, w
	focusDist := lookFrom.Sub(lookAt).Vec().Len()
	c.LowerLeft = c.Origin.Vec().Sub(u.Mul(halfWidth * focusDist)).Sub(v.Mul(halfHeight * focusDist)).Sub(w.Mul(focusDist))
	c.Horizontal = u.Mul(2 * halfWidth * focusDist)
	c.Vertical = v.Mul(2 * halfHeight * focusDist)
	return
}

func randUnitDisk() (p Vector) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		p = Vector{rnd.Float64(), rnd.Float64(), 0}.Mul(2.0).Sub(Vector{1.0, 1.0, 0.0})
		if p.Dot(p) < 1.0 {
			return
		}
	}
}

// GetRay gets ray at position x, y
func (c Camera) GetRay(x, y float64) Ray {
	s := x / float64(c.Width)
	t := y / float64(c.Height)
	hor := c.Horizontal.Mul(s)
	ver := c.Vertical.Mul(t)
	horver := hor.Add(ver)
	rd := randUnitDisk().Mul(c.LensRadius)
	offset := c.u.Mul(rd.X).Add(c.v.Mul(rd.Y))
	return NewRay(c.Origin, c.LowerLeft.Add(horver).Sub(c.Origin).Sub(offset))
}
