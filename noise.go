package main

import (
	"github.com/aquilax/go-perlin"
	"math"
	"math/rand"
	"time"
)

const (
	alpha = 2.
	beta  = 2.
	n     = 3
)

func p5Constrain(val, low, high float64) float64 {
	return math.Max(math.Min(val, high), low)
}

func p5Map(value, start1, stop1, start2, stop2 float64, withinBounds bool) float64 {
	newVal := (value-start1)/(stop1-start1)*(stop2-start2) + start2
	if !withinBounds {
		return newVal
	}

	if start2 < stop2 {
		return p5Constrain(newVal, start2, stop2)
	} else {
		return p5Constrain(newVal, stop2, start2)
	}
}

type NoiseLoop struct {
	diameter float64
	min      float64
	max      float64
	cx       float64
	cy       float64
	noise    *perlin.Perlin
}

func (nl *NoiseLoop) Value(a float64, bounds bool) float64 {
	xOffset := p5Map(math.Cos(a), -1, 1, nl.cx, nl.cx+nl.diameter, bounds)
	yOffset := p5Map(math.Sin(a), -1, 1, nl.cy, nl.cy+nl.diameter, bounds)

	r := nl.noise.Noise2D(xOffset, yOffset)

	return p5Map(r, 0, 1, nl.min, nl.max, bounds)
}

func NewNoiseLoop(diameter, min, max float64) *NoiseLoop {
	var nl NoiseLoop

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	nl.diameter = diameter
	nl.min = min
	nl.max = max
	nl.cx = r.Float64() * 100000
	nl.cy = r.Float64() * 100000
	nl.noise = perlin.NewPerlin(alpha, beta, n, 100)

	return &nl
}
