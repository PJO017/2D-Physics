package main

import (
	"fmt"

	"pjo018/2dphysics/internal/vector"
	"pjo018/2dphysics/pkg/config"
	"pjo018/2dphysics/pkg/particle"
	"pjo018/2dphysics/pkg/particleManager"
	"pjo018/2dphysics/pkg/system"

	"github.com/veandco/go-sdl2/sdl"
)

func setup() *particlemanager.Particlemanager {
	pm := particlemanager.CreateParticleManager()

	for i := 0; i < config.PARTICLE_COUNT; i++ {
		p := pm.CreateRandomParticle(config.SCREEN_WIDTH, config.SCREEN_HEIGHT)
		gravityForce := particle.CreateConstantForce(*vector.CreateVector(0, 9.8*config.SCALE))
		p.AddForce(gravityForce)
		p.ApplyForces()
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
	for idx, p := range pm.Particles {
		p.Update(deltaTime, idx, pm.Particles)
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
	sys, err := system.InitSystem(config.SCREEN_WIDTH, config.SCREEN_HEIGHT)
	if err != nil {
		fmt.Println("Error initializing system: ", err)
		panic(err)
	}
	defer sys.Destroy()

	pm := setup()

	accumulator := 0.0
	previousTime := sdl.GetTicks64()
	sys.RunningFlag = true
	for sys.RunningFlag {
		frameStartTime := sdl.GetTicks64()
		frameTime := float64(frameStartTime-previousTime) / 1000

		accumulator += frameTime
		accumulator = min(accumulator, config.TIME_STEP*2)

		processInput(sys)

		for accumulator >= config.TIME_STEP {
			update(pm, config.TIME_STEP)
			accumulator -= config.TIME_STEP
		}

		elapsedTime := float64(sdl.GetTicks64()-frameStartTime) / 1000
		if config.FRAME_DELAY > elapsedTime {
			sdl.Delay(uint32(config.FRAME_DELAY - elapsedTime))
		}

		previousTime = frameStartTime

		render(pm, sys.Renderer)
	}
}
