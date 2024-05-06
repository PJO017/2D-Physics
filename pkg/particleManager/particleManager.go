package particlemanager

import (
	"math/rand"
	"sync"

	"pjo018/2dphysics/internal/vector"
	"pjo018/2dphysics/pkg/particle"

	"github.com/veandco/go-sdl2/sdl"
)

type Particlemanager struct {
	Particles []*particle.Particle
	lock      sync.RWMutex
}

func CreateParticleManager() *Particlemanager {
	return &Particlemanager{
		Particles: []*particle.Particle{},
		lock:      sync.RWMutex{},
	}
}

func (pm *Particlemanager) CreateParticle(screenWidth, screenHeight int32, x, y, radius, mass float32) *particle.Particle {
	pm.lock.Lock()
	defer pm.lock.Unlock()
	particle := particle.CreateParticle(x, y, radius, mass, sdl.Color{R: 120, G: 255, B: 120, A: 255})
	pm.Particles = append(pm.Particles, particle)
	return particle
}

func (pm *Particlemanager) CreateRandomParticle(screenWidth, screenHeight int32) *particle.Particle {
	pm.lock.Lock()
	defer pm.lock.Unlock()

	radius := 10 + rand.Float32()*10
	x := 0 + rand.Float32()*(float32(screenWidth))
	// y := 0 + rand.Float32()*(float32(screenHeight))
	y := float32(screenHeight) - radius
	mass := 5 * radius

	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	a := uint8(rand.Intn(255))

	// velX := 5 + rand.Float32()*5
	// velY := 10 + rand.Float32()*10
	velX := float32(15)
	velY := float32(-1)

	particle := particle.CreateParticle(x, y, radius, mass, sdl.Color{R: r, G: g, B: b, A: a})
	particle.Velocity.AddVector(&vector.Vector{X: velX, Y: velY})

	pm.Particles = append(pm.Particles, particle)
	return particle
}
