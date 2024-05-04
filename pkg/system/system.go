package system

import (
	"pjo018/2dphysics/pkg/particle"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	SCREEN_WIDTH   = 800
	SCREEN_HEIGHT  = 600
	PARTICLE_COUNT = 1
)

type System struct {
	RunningFlag bool
	DeltaTime   uint64
	Renderer    *sdl.Renderer
	Window      *sdl.Window
	Particles   []*particle.Particle
}

func (s *System) Destroy() {
	s.Renderer.Destroy()
	s.Window.Destroy()
	sdl.Quit()
}

func InitSystem() (*System, error) {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		return nil, err
	}

	window, err := sdl.CreateWindow("2D Fiscs", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, SCREEN_WIDTH, SCREEN_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}

	return &System{RunningFlag: true, DeltaTime: 0, Renderer: renderer, Window: window, Particles: []*particle.Particle{}}, nil
}

func (s *System) Update() {
	for _, particle := range s.Particles {
		particle.Position.X += 1
	}
}
