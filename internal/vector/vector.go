package vector

import "fmt"

type Vector struct {
	X, Y float32
}

func CreateVector(x, y float32) *Vector {
	return &Vector{X: x, Y: y}
}

func (v *Vector) AddVector(v2 *Vector) *Vector {
	v.X += v2.X
	v.Y += v2.Y
	return v
}

func (v *Vector) SubtractVector(v2 *Vector) *Vector {
	v.X -= v2.X
	v.Y -= v2.Y
	return v
}

func (v *Vector) MultiplyScalar(scalar float32) *Vector {
	fmt.Println("scalar", scalar)
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
