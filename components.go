package main

import (
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

	len := raymath.Vector2Length(o.vel)
	omega := o.dir * len * dt / o.dist
	sn, cs := sincos(omega)
	nx := o.anchor.X + (o.pos.X-o.anchor.X)*cs - (o.pos.Y-o.anchor.Y)*sn
	ny := o.anchor.Y + (o.pos.X-o.anchor.X)*sn + (o.pos.Y-o.anchor.Y)*cs

	dx := nx - o.pos.X
	dy := ny - o.pos.Y

	o.pos.X = nx
	o.pos.Y = ny

	return dx, dy
}
