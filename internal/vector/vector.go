package vector

import (
	"math"
)

type Vector struct {
	X, Y float32
}

func CreateVector(x, y float32) *Vector {
	return &Vector{X: x, Y: y}
}

func (v *Vector) Distance(other *Vector) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func (v *Vector) DotProduct(other *Vector) float32 {
	return (v.X * other.X) + (v.Y * other.Y)
}

func (v *Vector) AddVector(v2 *Vector) *Vector {
	v.X += v2.X
	v.Y += v2.Y
	return v
}

func (v *Vector) AddVectorNew(v2 *Vector) *Vector {
	return &Vector{X: v.X + v2.X, Y: v.Y + v2.Y}
}

func (v *Vector) SubtractVector(v2 *Vector) *Vector {
	v.X -= v2.X
	v.Y -= v2.Y
	return v
}

func (v *Vector) SubtractVectorNew(v2 *Vector) *Vector {
	return &Vector{X: v.X - v2.X, Y: v.Y - v2.Y}
}

func (v *Vector) SubtractScalar(scalar float32) *Vector {
	v.X += scalar
	v.Y += scalar
	return v
}

func (v *Vector) AddScalar(scalar float32) *Vector {
	v.X += scalar
	v.Y += scalar
	return v
}

func (v *Vector) MultiplyScalar(scalar float32) *Vector {
	v.X *= scalar
	v.Y *= scalar
	return v
}

func (v *Vector) MultiplyScalarNew(scalar float32) *Vector {
	return &Vector{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

func (v *Vector) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v *Vector) Normalize() *Vector {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return &Vector{X: 0, Y: 0}
	}
	return &Vector{
		X: v.X / magnitude,
		Y: v.Y / magnitude,
	}
}
