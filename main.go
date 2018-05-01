package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

func initPlayers(playerName string, ais int) {
	players = []player{
		// A pseudo player that represents 'no player'.
		player{
			id:    0,
			name:  "not occupied",
			ai:    false,
			color: raylib.RayWhite,
		},
		player{
			id:    1,
			name:  playerName,
			ai:    false,
			color: raylib.SkyBlue,
		},
	}
}

func initScreen(title string, width, height, fps int32) {
	raylib.InitWindow(width, height, title)
	raylib.SetTargetFPS(fps)
	// We set the screen to be centered at 0,0. This means that our screen dimensions are:
	// xrange: -screenWidth/2 to screenWidth/2
	// yrange: -screenHeight/2 to screenHeight/2
	camera = raylib.Camera2D{
		Target:   raylib.Vector2{X: 0, Y: 0},
		Offset:   raylib.Vector2{X: float32(screenWidth) / 2, Y: float32(screenHeight) / 2},
		Rotation: 0,
		Zoom:     1,
	}
}

// genPlanetParameters generates random numbers for all parameters of a planet
// the valid values / ranges are passed in as arrays.
func genPlanetParameters(sizes []int) (size, vel, dir float32) {
	size = float32(sizes[rand.Intn(len(sizes))])
	vel = (rand.Float32() * 7) + 3
	// dir is just 1 or -1, which determines if a planet moves
	// clockwise or counter-clockwise.
	dir = float32(rand.Intn(2)*2 - 1)
	return
}

// We distribute the planets homogeneously on the X axis starting inside of the given range (span).
func initSolarSystem(planetAmount, maxSatellites, minDist, maxDist int) {
	span := maxDist - minDist
	step := span / planetAmount
	current := minDist

	for i := 0; i < planetAmount; i++ {
		size, vel, dir := genPlanetParameters(planetSizes)
		p := newPlanet(float32(current), size, dir, raylib.Vector2{X: vel, Y: vel}, origin, &players[0])
		// Add a little random adjustment to the planet's position to make
		// it look less static.
		shift := float32(rand.Intn(step/3)*2 - step/3)
		p.pos.X += shift
		// The planet is generated. Add it to our global planets slice.
		planets = append(planets, p)

		// Now we do more or less the same again as above. Just this time we are adding satellites
		// which orbit the previously generated planet.
		sats := rand.Intn(maxSatellites + 1)

		for s := 0; s < sats; s++ {
			size, vel, dir := genPlanetParameters(satelliteSizes)
			sat := newPlanet(float32((s+1)*20), size, dir, raylib.Vector2{X: vel, Y: vel}, &p.pos, &players[0])
			sat.rotate(rand.Float32() * sat.dist)
			p.satellites = append(p.satellites, sat)
			planets = append(planets, sat)
		}

		// Now that the planet and its satellites exist we rotate them randomly
		// to achieve a nice distribution "on the clock".
		p.rotateAll(rand.Float32() * p.dist)

		// Next planet please..
		current += step
	}
}

func initTextures() {
	// Generate textures for all planet/satellite sizes.
	for _, size := range planetSizes {
		pim := genPlanetImage(size*2, 0.7, raylib.White, raylib.Blank)
		planetTextures[size*2] = raylib.LoadTextureFromImage(pim)
		raylib.UnloadImage(pim)
	}
	for _, size := range satelliteSizes {
		pim := genPlanetImage(size*2, 0.7, raylib.White, raylib.Blank)
		planetTextures[size*2] = raylib.LoadTextureFromImage(pim)
		raylib.UnloadImage(pim)
	}

	shipIm := raylib.GenImageGradientRadial(3, 3, 0.9, raylib.White, raylib.Blank)
	shipTexture = raylib.LoadTextureFromImage(shipIm)
	raylib.UnloadImage(shipIm)

	sunIm := raylib.GenImageGradientRadial(75, 75, 0.1, raylib.White, raylib.Blank)
	sunTexture = raylib.LoadTextureFromImage(sunIm)
	raylib.UnloadImage(sunIm)
}

// update handles all logic changes in the game. This includes
// moving objects or handling input.
func update(dt float32) {
	spriteCount = len(planets) + 1 // includes the sun
	for i := 0; i < len(planets); i++ {
		planets[i].update(dt)
		for _, s := range planets[i].ships {
			if s != nil {
				spriteCount++
			}
		}
	}
}

// draw is called after update and just draws everything visible
// to the screen.
func draw() {
	raylib.BeginDrawing()

	raylib.ClearBackground(raylib.Black)

	raylib.Begin2dMode(camera)

	// Draw block that is applied to the camera view.
	raylib.DrawTextureV(sunTexture, *origin, raylib.Gold)

	for _, p := range planets {
		p.draw()
	}

	raylib.End2dMode()

	// Draw block applied outside of camera view (HUD elements etc.)
	raylib.DrawText(fmt.Sprintf("FPS: %d", int(raylib.GetFPS())), 10, 10, 32, raylib.RayWhite)
	raylib.DrawText(fmt.Sprintf("Sprites: %d", spriteCount), 10, 50, 32, raylib.RayWhite)

	raylib.EndDrawing()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	initScreen(title, screenWidth, screenHeight, fps)
	initPlayers("RagingDave", 0)
	initSolarSystem(12, 3, 100, int(screenHeight/2))
	initTextures()

	// The main game loop is here. It periodically calls the update and draw functions.
	for !raylib.WindowShouldClose() {
		dt := raylib.GetFrameTime()
		update(dt)
		draw()
	}

	raylib.CloseWindow()
}
