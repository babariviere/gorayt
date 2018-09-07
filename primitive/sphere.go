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

func (s Sphere) Hit(r Ray, tmin, tmax float64) (bool, Hit) {
	var rec Hit
	oc := NewVector2(r.Origin, s.Center)
	a := r.Direction.Dot(r.Direction)
	b := oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	disc := b*b - a*c
	if disc > 0 {
		sqdisc := math.Sqrt(disc)
		t0 := (-b - sqdisc) / a
		t1 := (-b + sqdisc) / a
		if t1 < t0 {
			t1, t0 = t0, t1
		}
		if t0 <= tmin {
			t0 = t1
			if t0 <= tmin {
				return false, rec
			}
		}
		rec.Distance = t0
		rec.Normal.Origin = r.PointAt(t0)
		rec.Normal.Direction = rec.Normal.Origin.Sub(s.Center).Vec().Div(s.Radius)
		return true, rec
	}
	return false, rec
}
