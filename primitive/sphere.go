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

func solveQuadratic(a, b, c float64, t0, t1 *float64) bool {
	disc := b*b - 4*a*c
	if disc < 0 {
		return false
	} else if disc == 0 {
		*t0 = -0.5 * b / a
		*t1 = *t0
	} else {
		var q float64
		if b > 0 {
			q = -0.5 * (b + math.Sqrt(disc))
		} else {
			q = -0.5 * (b - math.Sqrt(disc))
		}
		*t0 = q / a
		*t1 = c / q
	}
	return true
}

func (s Sphere) Hit(r Ray, tmin, tmax float64) (bool, Hit) {
	var rec Hit
	var t0, t1 float64
	l := NewVector2(r.Origin, s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(l)
	c := l.Dot(l) - s.Radius*s.Radius
	if !solveQuadratic(a, b, c, &t0, &t1) {
		return false, rec
	}
	if t1 < t0 {
		t1, t0 = t0, t1
	}
	rec.Distance = t0
	rec.Normal.Origin = r.PointAt(t0)
	rec.Normal.Direction = rec.Normal.Origin.Sub(s.Center).Vec().Div(s.Radius)
	return true, rec
}
