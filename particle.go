package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Particle struct {
	window *pixelgl.Window
	xNoise *NoiseLoop
	yNoise *NoiseLoop
	rNoise *NoiseLoop
	imd    *imdraw.IMDraw
}

func (p *Particle) Draw(a float64) {
	x := p.xNoise.Value(a, false)
	y := p.yNoise.Value(a, false)
	r := p.rNoise.Value(a, false)

	p.imd.Clear()
	p.imd.Color = pixel.RGB(0, 0, 0)
	p.imd.Push(pixel.V(x, y))
	p.imd.Circle(r, r/2)
	p.imd.Draw(p.window)
}

func NewParticle(win *pixelgl.Window) *Particle {
	var p Particle
	p.window = win
	p.xNoise = NewNoiseLoop(0.1, 150, win.Bounds().W()*2)
	p.yNoise = NewNoiseLoop(0.2, 150, win.Bounds().H()*2)
	p.rNoise = NewNoiseLoop(4, 10, 25)
	p.imd = imdraw.New(nil)
	return &p
}
