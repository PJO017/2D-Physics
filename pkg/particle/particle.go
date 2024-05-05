package particle

import (
	"pjo018/2dphysics/internal/vector"
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
	p.ApplyForce(vector.CreateVector(0, 9.8*60))
	p.Velocity.AddVector(p.Acceleration.MultiplyScalarNew(float32(deltaTime) / 1000))
	p.Position.AddVector(&p.Velocity)
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

func (p *Particle) ApplyForce(force *vector.Vector) {
	p.Acceleration.AddVector(force.MultiplyScalarNew(1 / p.Mass))
}
