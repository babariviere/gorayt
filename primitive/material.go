package primitive

import (
	"math"
	"math/rand"
)

// Material represents material of an object
type Material interface {
	Scatter(ray Ray, rec Hit) (bool, Ray)
	Color() Color
}

type Matte struct {
	c Color
}

func NewMatte(color Color) Matte {
	return Matte{color}
}

func (m Matte) Color() Color {
	return m.c
}

func (m Matte) Scatter(ray Ray, rec Hit) (bool, Ray) {
	target := rec.Point.Add(rec.Normal).Vec().Add(RandomUnitSphere())
	return true, NewRay(rec.Point, target.Sub(rec.Point))
}

type Metal struct {
	c    Color
	fuzz float64
}

func NewMetal(color Color, fuzz float64) Metal {
	if fuzz > 1 {
		fuzz = 1
	}
	return Metal{color, fuzz}
}

func (m Metal) Color() Color {
	return m.c
}

func (m Metal) Scatter(ray Ray, rec Hit) (bool, Ray) {
	reflected := reflect(ray.Direction.Normalize(), rec.Normal)
	scattered := NewRay(rec.Point, reflected.Add(RandomUnitSphere().Mul(m.fuzz)))
	ok := scattered.Direction.Dot(rec.Normal) > 0
	return ok, scattered
}

type Glass struct {
	refIdx float64
}

func NewGlass(refIdx float64) Glass {
	return Glass{refIdx}
}

func (g Glass) Scatter(ray Ray, rec Hit) (_ bool, scattered Ray) {
	reflected := reflect(ray.Direction, rec.Normal)
	var outwardNorm Vector
	var niOverNt, cosine, reflectProb float64
	if ray.Direction.Dot(rec.Normal) > 0 {
		outwardNorm = rec.Normal.Mul(-1.)
		niOverNt = g.refIdx
		cosine = g.refIdx * ray.Direction.Dot(rec.Normal) / ray.Direction.Len()
	} else {
		outwardNorm = rec.Normal
		niOverNt = 1 / g.refIdx
		cosine = -ray.Direction.Dot(rec.Normal) / ray.Direction.Len()
	}
	var success bool
	var refracted Vector
	if success, refracted = refract(ray.Direction, outwardNorm, niOverNt); success {
		reflectProb = shclick(cosine, g.refIdx)
	} else {
		reflectProb = 1.0
	}
	if rand.Float64() < reflectProb {
		scattered = NewRay(rec.Point, reflected)
	} else {
		scattered = NewRay(rec.Point, refracted)
	}
	return true, scattered
}

func (g Glass) Color() Color {
	return Color{1, 1, 1}
}

func reflect(v Vector, n Vector) Vector {
	return v.Sub(n.Mul(2 * v.Dot(n)))
}

func refract(v Vector, n Vector, niOverNt float64) (bool, Vector) {
	uv := v.Normalize()
	dt := uv.Dot(n)
	disc := 1 - niOverNt*niOverNt*(1-dt*dt)
	if disc > 0 {
		a := uv.Sub(n.Mul(dt))
		b := n.Mul(math.Sqrt(disc))
		return true, a.Mul(niOverNt).Sub(b)
	}
	return false, Vector{}
}

func shclick(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
