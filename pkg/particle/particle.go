package particle

import (
	"fmt"

	"pjo018/2dphysics/internal/vector"
	"pjo018/2dphysics/pkg/utils"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	FRICTION_COEFFIECIENT = 0.5
	SCALE                 = 100
	SCREEN_WIDTH          = 1200
	SCREEN_HEIGHT         = 800
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

func (p *Particle) Update(deltaTime float64) {
	p.ApplyForces()
	frictionForce := p.ApplyFriction()
	fmt.Println("acc", p.Acceleration.X, p.Acceleration.Y)
	p.Velocity.AddVector(p.Acceleration.MultiplyScalarNew(float32(deltaTime) / 1000))
	p.Position.AddVector(&p.Velocity)

	if frictionForce != nil {
		p.Acceleration.SubtractVector(frictionForce.MultiplyScalarNew(1 / p.Mass))
	}
}

func (p *Particle) HandleCollision(screenWidth, screenHeight int, DAMPING_FACTOR float32) {
	if p.Position.X-p.Radius < 0 || p.Position.X+p.Radius > float32(screenWidth) {
		p.Velocity.X *= -1 * DAMPING_FACTOR
		p.Position.X = utils.Clamp(p.Position.X, p.Radius, float32(screenWidth)-p.Radius)

	}

	if p.Position.Y-p.Radius < 0 || p.Position.Y+p.Radius > float32(screenHeight) {
		p.Velocity.Y *= -1 * DAMPING_FACTOR
		p.Position.Y = utils.Clamp(p.Position.Y, p.Radius, float32(screenHeight)-p.Radius)
	}
}

func (p *Particle) AddForce(force Force) {
	p.Forces = append(p.Forces, force)
}

func (p *Particle) ApplyForce(force *vector.Vector) {
	p.Acceleration.AddVector(force.MultiplyScalarNew(1 / p.Mass))
}

func (p *Particle) ApplyFriction() *vector.Vector {
	if p.OnSurface() {
		fmt.Println("On surface")
		frictionMagnitude := float64(FRICTION_COEFFIECIENT * p.Mass * 9.8 * SCALE)
		frictionForce := p.Velocity.Normalize().MultiplyScalarNew(-1 * float32(frictionMagnitude))
		p.ApplyForce(frictionForce)
		return frictionForce

	}
	return nil
}

func (p *Particle) OnSurface() bool {
	return p.Position.X-p.Radius <= 0 ||
		p.Position.X+p.Radius >= float32(SCREEN_WIDTH) ||
		p.Position.Y-p.Radius <= 0 ||
		p.Position.Y+p.Radius >= float32(SCREEN_HEIGHT)
}

func (p *Particle) ApplyForces() {
	for _, force := range p.Forces {
		if force.IsConstant || (force.Condition != nil && force.Condition(p)) {
			fmt.Println("Applying force", force.Vector.X, force.Vector.Y)
			p.ApplyForce(&force.Vector)
		}
	}
}

