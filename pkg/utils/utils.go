package utils

import "math"

func Clamp(value, min, max float32) float32 {
	return float32(math.Max(float64(min), math.Min(float64(max), float64(value))))
}
