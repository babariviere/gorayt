package primitive

import (
	"math"
)

// Sphere represents a sphere in a 3D space
type Sphere struct {
	Center Point
	Radius float64
}

// NewSphere initialize a new sphere with given values
func NewSphere(Center Point, Radius float64) Sphere {
	return Sphere{Center, Radius}
}

// see: http://www.scratchapixel.com/lessons/3d-basic-rendering/minimal-ray-tracer-rendering-simple-shapes/ray-sphere-intersection
func (s Sphere) Hit(r Ray, tmin, tmax float64) (_ bool, rec Hit) {
	l := NewVector2(r.Origin, s.Center)
	a := r.Direction.Dot(r.Direction)
	b := r.Direction.Dot(l) * 2
	c := l.Dot(l) - s.Radius*s.Radius
	disc := b*b - 4*a*c
	var t0, t1 float64
	if disc < 0 {
		return false, rec
	} else if disc == 0 {
		t0 = -0.5 * b / a
		t1 = t0
	} else {
		sqdisc := math.Sqrt(disc)
		if b <= 0 {
			sqdisc = -sqdisc
		}
		q := -0.5 * (b + sqdisc)
		t0 = q / a
		t1 = c / q
	}
	if t0 > t1 {
		t1, t0 = t0, t1
	}
	if t0 < tmin {
		t0 = t1
		if t0 < tmin {
			return false, rec
		}
	}
	if t0 > tmax {
		return false, rec
	}
	rec.Distance = t0
	rec.Normal.Origin = r.PointAt(t0)
	rec.Normal.Direction = rec.Normal.Origin.Sub(s.Center).Vec().Div(s.Radius)
	return true, rec
}
