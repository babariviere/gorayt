package primitive

import (
	"fmt"
	"math"
	"math/rand"
)

// Vector is a mathematical representation of a vector.
// It has both a magnitude and a direction
type Vector struct {
	X, Y, Z float64
}

// NewVector initialize a new vector with given value
func NewVector(X, Y, Z float64) Vector {
	return Vector{X, Y, Z}
}

// NewVector2 creates a vector from two points
func NewVector2(a Point, b Point) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Add do the add operation of the vector with an interface
func (v Vector) Add(o interface{}) Vector {
	switch t := o.(type) {
	default:
		panic(fmt.Sprint("unable to add vector with", t))
	case Vector:
		return Vector{v.X + t.X, v.Y + t.Y, v.Z + t.Z}
	case Point:
		return Vector{v.X + t.X, v.Y + t.Y, v.Z + t.Z}
	case float64:
		return Vector{v.X + t, v.Y + t, v.Z + t}
	}
}

// Sub do the substract operation of the vector with an interface
func (v Vector) Sub(o interface{}) Vector {
	switch t := o.(type) {
	default:
		panic(fmt.Sprint("unable to substract vector with", t))
	case Vector:
		return Vector{v.X - t.X, v.Y - t.Y, v.Z - t.Z}
	case Point:
		return Vector{v.X - t.X, v.Y - t.Y, v.Z - t.Z}
	case float64:
		return Vector{v.X - t, v.Y - t, v.Z - t}
	}
}

// Mul do the multiplication operation of the vector with an interface
func (v Vector) Mul(o interface{}) Vector {
	switch t := o.(type) {
	default:
		panic(fmt.Sprint("unable to multiply vector with", t))
	case Vector:
		return Vector{v.X * t.X, v.Y * t.Y, v.Z * t.Z}
	case Point:
		return Vector{v.X * t.X, v.Y * t.Y, v.Z * t.Z}
	case float64:
		return Vector{v.X * t, v.Y * t, v.Z * t}
	}
}

// Div do the division operation on the vector with an interface
func (v Vector) Div(o interface{}) Vector {
	switch t := o.(type) {
	default:
		panic(fmt.Sprint("unable to divide vector with", t))
	case Vector:
		return Vector{v.X / t.X, v.Y / t.Y, v.Z / t.Z}
	case Point:
		return Vector{v.X / t.X, v.Y / t.Y, v.Z / t.Z}
	case float64:
		return Vector{v.X / t, v.Y / t, v.Z / t}
	}
}

// Dot calculates the dot product of two vector
func (v Vector) Dot(o Vector) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

// Len calculates len of vector
func (v Vector) Len() float64 {
	return math.Sqrt(v.Dot(v))
}

// Normalize normalize vector to a scale of one
func (v Vector) Normalize() Vector {
	return v.Div(v.Len())
}

// Cross calculates the cross product of two vector
func (v Vector) Cross(o Vector) Vector {
	return Vector{
		v.Y*o.Z - v.Z*o.Y,
		v.Z*o.X - v.X*o.Z,
		v.X*o.Y - v.Y*o.X,
	}
}

// Color converts a vector to a color
func (v Vector) Color() Color {
	return Color{v.X, v.Y, v.Z}
}

// RandomUnitSphere returns a random point in a unit radius sphere
func RandomUnitSphere() (v Vector) {
	for {
		v = Vector{rand.Float64(), rand.Float64(), rand.Float64()}.Mul(2.0).Sub(Vector{1, 1, 1})
		if v.Dot(v) < 1 {
			return
		}
	}
}
