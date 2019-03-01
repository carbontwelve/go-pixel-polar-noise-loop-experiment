package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math"
)

const pCount = 100
const TwoPi = math.Pi * 2

func setup() {
	cfg := pixelgl.WindowConfig{
		Title:  "Particle Noise",
		Bounds: pixel.R(0, 0, 330, 240),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	var particles [pCount]*Particle

	for i := 0; i < pCount; i++ {
		particles[i] = NewParticle(win)
	}

	var counter, totalFrames, percent, angle, radians float64

	counter = 0
	totalFrames = 480

	for !win.Closed() {
		win.Clear(colornames.Skyblue)
		percent = counter / totalFrames * 100

		angle = 360 / 100 * percent
		radians = angle / math.Pi / 180

		for i := 0; i < pCount; i++ {
			//fmt.Printf("%0.0f\t%0.4f\n", counter, percent)
			particles[i].Draw(radians * TwoPi)
		}

		//if counter >= totalFrames - 1 {
		//	counter = 0
		//} else {
		//	counter++
		//}

		//if percent >= 100 {
		//counter = 0
		//} else {
		counter++
		//}

		win.Update()
	}
}

func main() {
	pixelgl.Run(setup)
}
