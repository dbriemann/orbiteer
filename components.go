package main

import (
	"math"

	"github.com/gen2brain/raylib-go/raylib"
	"github.com/gen2brain/raylib-go/raymath"
)

type player struct {
	id    int
	name  string
	ai    bool
	color raylib.Color
}

type orb struct {
	anchor *raylib.Vector2
	vel    raylib.Vector2
	dir    float32
	pos    raylib.Vector2
	dist   float32
}

func (o *orb) rotate(dt float32) (float32, float32) {
	if o.anchor == nil {
		return 0, 0 // no rotation without anchor
	}

	len := float32(math.Sqrt(float64(raymath.Vector2LenSqr(o.vel))))
	omega := o.dir * len * dt / o.dist
	nx := o.anchor.X + (o.pos.X-o.anchor.X)*float32(math.Cos(float64(omega))) - (o.pos.Y-o.anchor.Y)*float32(math.Sin(float64(omega)))
	ny := o.anchor.Y + (o.pos.X-o.anchor.X)*float32(math.Sin(float64(omega))) + (o.pos.Y-o.anchor.Y)*float32(math.Cos(float64(omega)))

	dx := nx - o.pos.X
	dy := ny - o.pos.Y

	o.pos.X = nx
	o.pos.Y = ny

	return dx, dy
}
