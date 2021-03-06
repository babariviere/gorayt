package primitive

import (
	"math"
)

// Color is a RGB color
type Color struct {
	R, G, B float64
}

// NewColor initialize a new color with given values
func NewColor(R, G, B float64) Color {
	return Color{R, G, B}
}

// Add do the add operation on color
func (c Color) Add(o Color) Color {
	return Color{c.R + o.R, c.G + o.G, c.B + o.B}
}

// Mul do the mul operation on color
func (c Color) Mul(o Color) Color {
	return Color{c.R * o.R, c.G * o.G, c.B * o.B}
}

// Div do the div operation on color
func (c Color) Div(o float64) Color {
	return Color{c.R / o, c.G / o, c.B / o}
}

// Vec converts color to a vec
func (c Color) Vec() Vector {
	return Vector{c.R, c.G, c.B}
}

// Gamma fix color with gamma encoding
func (c Color) Gamma() Color {
	return Color{math.Sqrt(c.R), math.Sqrt(c.G), math.Sqrt(c.B)}
}
