package main

import "github.com/gen2brain/raylib-go/raylib"

var (
	screenWidth  int32 = 1200
	screenHeight int32 = 800
	fps          int32 = 0
	title              = "Orbiteer"
	spriteCount        = 0

	productionFactor float32 = 0.1

	planets []*planet
	players []player
	camera  raylib.Camera2D

	// origin is an anchor for all planets that are not satellites to another planet.
	origin         = &raylib.Vector2{X: 1, Y: 1}
	planetSizes    = []int{7, 8, 9}
	satelliteSizes = []int{3, 4, 5}
	recycledShips  = []*ship{}

	planetTextures = map[int]raylib.Texture2D{}
	shipTexture    raylib.Texture2D
	sunTexture     raylib.Texture2D
)
