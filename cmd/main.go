package main

import (
	"fmt"

	"pjo018/2dphysics/pkg/particleManager"
	"pjo018/2dphysics/pkg/system"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	SCREEN_WIDTH   = 800
	SCREEN_HEIGHT  = 600
	PARTICLE_COUNT = 1
	FPS            = 120
	TIME_STEP      = 1.0 / FPS
	FRAME_DELAY    = 1000 / FPS
	DAMPING_FACTOR = 0.80
)

func setup(system *system.System) *particlemanager.Particlemanager {
	screenWidth, screenHeight := system.Window.GetSize()
	pm := particlemanager.CreateParticleManager()
	for i := 0; i < PARTICLE_COUNT; i++ {
		pm.CreateParticle(screenWidth, screenHeight)
	}
	return pm
}

func processInput(system *system.System) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			system.RunningFlag = false
		}
	}
}

func update(pm *particlemanager.Particlemanager, deltaTime float64) {
	for _, p := range pm.Particles {
		p.Update(deltaTime)
		p.HandleCollision(SCREEN_WIDTH, SCREEN_HEIGHT, DAMPING_FACTOR)
	}
}

func render(pm *particlemanager.Particlemanager, renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	for _, particle := range pm.Particles {
		particle.Draw(renderer)
	}

	renderer.Present()
}

func main() {
	sys, err := system.InitSystem()
	if err != nil {
		fmt.Println("Error initializing system: ", err)
		panic(err)
	}
	defer sys.Destroy()

	pm := setup(sys)

	accumulator := 0.0
	previousTime := sdl.GetTicks64()
	sys.RunningFlag = true
	for sys.RunningFlag {
		frameStartTime := sdl.GetTicks64()
		frameTime := float64(frameStartTime-previousTime) / 1000
		previousTime = frameStartTime

		accumulator += frameTime

		processInput(sys)

		for accumulator >= TIME_STEP {
			update(pm, TIME_STEP)
			accumulator -= TIME_STEP
		}

		render(pm, sys.Renderer)

		elapsedTime := float64(sdl.GetTicks64()-frameStartTime) / 1000
		if FRAME_DELAY > elapsedTime {
			sdl.Delay(uint32(FRAME_DELAY - elapsedTime))
		}
	}
}
