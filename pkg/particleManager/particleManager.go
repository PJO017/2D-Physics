package particlemanager

import (
	"sync"

	"pjo018/2dphysics/pkg/particle"

	"github.com/veandco/go-sdl2/sdl"
)

type Particlemanager struct {
	Particles []*particle.Particle
	lock      sync.RWMutex
}

func CreateParticleManager() *Particlemanager {
	return &Particlemanager{Particles: []*particle.Particle{}}
}

func (pm *Particlemanager) CreateParticle(screenWidth, screenHeight int32) *particle.Particle {
	pm.lock.Lock()
	defer pm.lock.Unlock()
	particle := particle.CreateParticle(float32(screenWidth)/2, float32(screenHeight)/2, 10, 10, sdl.Color{R: 120, G: 255, B: 120, A: 255})
	pm.Particles = append(pm.Particles, particle)
	return particle
}
