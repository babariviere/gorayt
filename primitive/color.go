package primitive

// Color is a RGB color
type Color struct {
	R, G, B float64
}

// NewColor initialize a new color with given values
func NewColor(R, G, B float64) Color {
	return Color{R, G, B}
}
