package main

import "github.com/gen2brain/raylib-go/raylib"

// genPlanetImage produces a procedurally generated planet.
// For now this is just a simple "disc".
func genPlanetImage(size int, density float32, innerColor, outerColor raylib.Color) *raylib.Image {
	im := raylib.GenImageGradientRadial(size, size, density, innerColor, outerColor)

	return im
}
