package main

import (
	"fmt"

	"pjo018/2dphysics/internal/vector"
	"pjo018/2dphysics/pkg/particle"
	"pjo018/2dphysics/pkg/system"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	SCREEN_WIDTH   = 800
	SCREEN_HEIGHT  = 600
	PARTICLE_COUNT = 1
	FPS            = 30
	FRAME_DELAY    = 1000 / FPS
)

func setup(system *system.System) {
	screenWidth, screenHeight := system.Window.GetSize()
	particle := particle.CreateParticle(float32(screenWidth)/2, float32(screenHeight)/2, 10, 10, sdl.Color{R: 120, G: 255, B: 120, A: 255})
	acc := vector.CreateVector(0, 5)
	vel := vector.CreateVector(0, 1)
	particle.Velocity.AddVector(vel)
	particle.Acceleration.AddVector(acc)
	system.Particles = append(system.Particles, particle)
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

func update(system *system.System) {
	for _, particle := range system.Particles {
		particle.Velocity.AddVector(particle.Acceleration.MultiplyScalarNew(float32(system.DeltaTime) / 1000))
		fmt.Println("particle.Velocity", particle.Velocity.X, particle.Velocity.Y)
		particle.Position.AddVector(&particle.Velocity)
	}
}

func render(system *system.System) {
	system.Renderer.SetDrawColor(0, 0, 0, 255)
	system.Renderer.Clear()

	for _, particle := range system.Particles {
		particle.Draw(system.Renderer)
	}

	system.Renderer.Present()
}

func main() {
	sys, err := system.InitSystem()
	if err != nil {
		fmt.Println("Error initializing system: ", err)
		panic(err)
	}
	defer sys.Destroy()

	setup(sys)

	sys.RunningFlag = true
	for sys.RunningFlag {
		frameStartTime := sdl.GetTicks64()

		processInput(sys)
		update(sys)
		render(sys)

		sys.DeltaTime = sdl.GetTicks64() - frameStartTime

		if FRAME_DELAY > sys.DeltaTime {
			sdl.Delay(uint32(FRAME_DELAY - sys.DeltaTime))
		}
	}
}
