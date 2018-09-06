package primitive

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

/// Add do the add operation on two factor
func (v Vector) Add(o Vector) Vector {
	return Vector{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

// Sub do the substract operation on two vector
func (v Vector) Sub(o Vector) Vector {
	return Vector{v.X - o.X, v.Y + o.Y, v.Z + o.Z}
}

// Mul do the multiplication operation on two vector
func (v Vector) Mul(o Vector) Vector {
	return Vector{v.X * o.X, v.Y * o.Y, v.Z * o.Z}
}

// Mulf do the multiplication operation with one vector and a float
func (v Vector) Mulf(t float64) Vector {
	return Vector{v.X * t, v.Y * t, v.Z * t}
}

// Div do the division operation on two vector
func (v Vector) Div(o Vector) Vector {
	return Vector{v.X / o.X, v.Y / o.Y, v.Z / o.Z}
}
