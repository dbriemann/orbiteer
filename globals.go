package main

import "github.com/gen2brain/raylib-go/raylib"

var (
	screenWidth  int32 = 1200
	screenHeight int32 = 800
	fps          int32 = 60
	title              = "Orbiteer"

	planets []*planet
	players []player
	camera  raylib.Camera2D
	// origin is an anchor for all planets that are not arbiters to another planet.
	origin = &raylib.Vector2{X: 1, Y: 1}
)
