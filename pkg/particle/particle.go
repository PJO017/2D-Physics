package particle

import (
	"fmt"

	"pjo018/2dphysics/internal/vector"
	"pjo018/2dphysics/pkg/config"
	"pjo018/2dphysics/pkg/utils"

	"github.com/veandco/go-sdl2/sdl"
)

type Particle struct {
	Position     vector.Vector
	Velocity     vector.Vector
	Acceleration vector.Vector
	Radius       float32
	Mass         float32
	Color        sdl.Color
	Forces       []Force
}

func CreateParticle(x, y, radius, mass float32, color sdl.Color) *Particle {
	return &Particle{Position: vector.Vector{X: x, Y: y}, Velocity: vector.Vector{X: 0, Y: 0}, Acceleration: vector.Vector{X: 0, Y: 0}, Radius: radius, Mass: mass, Color: color}
}

func (p *Particle) Draw(Renderer *sdl.Renderer) {
	Renderer.SetDrawColor(p.Color.R, p.Color.G, p.Color.B, p.Color.A)
	for i := p.Position.X - p.Radius; i < p.Position.X+p.Radius; i++ {
		for j := p.Position.Y - p.Radius; j < p.Position.Y+p.Radius; j++ {
			if (i-p.Position.X)*(i-p.Position.X)+(j-p.Position.Y)*(j-p.Position.Y) < p.Radius*p.Radius {
				Renderer.DrawPointF(i, j)
			}
		}
	}
}

func (p *Particle) Update(deltaTime float64, idx int, particles []*Particle) {
	p.ApplyForces()
	frictionForce := p.ApplyFriction()

	p.Velocity.AddVector(p.Acceleration.MultiplyScalarNew(float32(deltaTime) / 1000))
	p.Position.AddVector(&p.Velocity)
	p.HandleCollisions(particles, idx)

	if frictionForce != nil {
		p.Acceleration.SubtractVector(frictionForce.MultiplyScalarNew(1 / p.Mass))
	}
}

func ParticlesCollide(p1, p2 *Particle) bool {
	return (p1.Position.X-p2.Position.X)*(p1.Position.X-p2.Position.X)+(p1.Position.Y-p2.Position.Y)*(p1.Position.Y-p2.Position.Y) < (p1.Radius+p2.Radius)*(p1.Radius+p2.Radius)
}

func (p *Particle) HandleCollisions(particles []*Particle, idx int) {
	if p.Position.X-p.Radius < 0 || p.Position.X+p.Radius > float32(config.SCREEN_WIDTH) {
		p.Velocity.X *= -1 * config.DAMPING_FACTOR
		p.Position.X = utils.Clamp(p.Position.X, p.Radius, float32(config.SCREEN_WIDTH)-p.Radius)

	}

	if p.Position.Y-p.Radius < 0 || p.Position.Y+p.Radius > float32(config.SCREEN_HEIGHT) {
		p.Velocity.Y *= -1 * config.DAMPING_FACTOR
		p.Position.Y = utils.Clamp(p.Position.Y, p.Radius, float32(config.SCREEN_HEIGHT)-p.Radius)
	}

	for i := idx + 1; i < len(particles); i++ {
		if ParticlesCollide(p, particles[i]) {
			fmt.Println("Handling collisions")

			p1 := p
			p2 := particles[i]

			// Relative velocity
			relativeVelocity := p1.Velocity.SubtractVectorNew(&p2.Velocity)

			// Center of Mass Velocity
			cmVelocity := p1.Velocity.MultiplyScalarNew(p1.Mass).AddVector(p2.Velocity.MultiplyScalarNew(p2.Mass)).MultiplyScalarNew(1 / (p1.Mass + p2.Mass) * 0.5)

			// Relative momentum
			p1Momentum := relativeVelocity.MultiplyScalarNew(p1.Mass)

			p1FinalMomentum := p1Momentum.MultiplyScalarNew((p1.Mass - p2.Mass) / (p1.Mass + p2.Mass))
			p2FinalMomentum := p1Momentum.MultiplyScalarNew((p2.Mass - p1.Mass) / (p2.Mass + p1.Mass))

			// Final velocities
			p1FinalVelocity := p1FinalMomentum.MultiplyScalarNew(1 / p1.Mass).AddVector(cmVelocity)
			p2FinalVelocity := p2FinalMomentum.MultiplyScalarNew(1 / p2.Mass).AddVector(cmVelocity)

			p1.Velocity = *p1FinalVelocity
			p2.Velocity = *p2FinalVelocity

		}
	}
}

func (p *Particle) AddForce(force Force) {
	p.Forces = append(p.Forces, force)
}

func (p *Particle) ApplyForce(force *vector.Vector) {
	p.Acceleration.AddVector(force.MultiplyScalarNew(1 / p.Mass))
}

func (p *Particle) ApplyFriction() *vector.Vector {
	if p.OnGround() {
		frictionMagnitude := float64(config.FRICTION_COEFFIECIENT * p.Mass * 9.8 * config.SCALE)
		frictionForce := p.Velocity.Normalize().MultiplyScalarNew(-1 * float32(frictionMagnitude))
		p.ApplyForce(frictionForce)
		return frictionForce

	}
	return nil
}

func (p *Particle) OnGround() bool {
	return p.Position.Y-p.Radius <= 0 || p.Position.Y+p.Radius >= float32(config.SCREEN_HEIGHT)
}

func (p *Particle) ApplyForces() {
	for _, force := range p.Forces {
		if force.IsConstant || (force.Condition != nil && force.Condition(p)) {
			p.ApplyForce(&force.Vector)
		}
	}
}
