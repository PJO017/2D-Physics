package particle

import (
	"pjo018/2dphysics/internal/vector"
)

type ForceCondition func(p *Particle) bool

type Force struct {
	Vector     vector.Vector
	IsConstant bool
	Condition  ForceCondition
}

func CreateConstantForce(v vector.Vector) Force {
	return Force{Vector: v, IsConstant: true, Condition: nil}
}

func CreateConditionalForce(v vector.Vector, condition ForceCondition) Force {
	return Force{Vector: v, IsConstant: false, Condition: condition}
}
