package main

import (
	"math"

	"github.com/gen2brain/raylib-go/raylib"
)

// rotatePoint rotates point around anchor by omage (rad).
func rotatePoint(anchor, point *raylib.Vector2, omega float32) {
	nx := anchor.X + (point.X-anchor.X)*float32(math.Cos(float64(omega))) - (point.Y-anchor.Y)*float32(math.Sin(float64(omega)))
	ny := anchor.Y + (point.X-anchor.X)*float32(math.Sin(float64(omega))) + (point.Y-anchor.Y)*float32(math.Cos(float64(omega)))

	point.X = nx
	point.Y = ny
}
