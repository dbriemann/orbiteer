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
	camera = raylib.Camera2D{
		Target:   raylib.Vector2{X: 0, Y: 0},
		Offset:   raylib.Vector2{X: float32(screenWidth) / 2, Y: float32(screenHeight) / 2},
		Rotation: 0,
		Zoom:     1,
	}
}

func initSolarSystem(planetAmount, maxArbiters, minDist, maxDist int) {
	span := maxDist - minDist
	step := span / planetAmount
	current := minDist

	for i := 0; i < planetAmount; i++ {
		p := newPlanet(float32(current), 8, raylib.Vector2{X: 10, Y: 10}, origin, &players[0])
		// Add a little random adjustment to the planet's position to make
		// it look less static.
		p.pos.X += float32(rand.Intn(step/10)*2 - step/10)
		// Random rotation to nicely distribute planets on screen.
		//		rotatePoint(&p.pos, p.anchor, rand.Float32()*math.Pi*2)
		planets = append(planets, p)

		// Add arbiters
		arbs := rand.Intn(maxArbiters + 1)

		for a := 0; a < arbs; a++ {
			size := rand.Float32()*2 + 3
			arb := newPlanet(float32((a+1)*20), size, raylib.Vector2{X: 10, Y: 10}, &p.pos, &players[0])
			arb.rotate(rand.Float32() * arb.dist)
			p.arbiters = append(p.arbiters, arb)
			planets = append(planets, arb)
		}

		p.update(rand.Float32() * p.dist)

		current += step
	}
}

// update handles all logic changes in the game. This includes
// moving objects or handling input.
func update(dt float32) {
	// Keep in mind never to use range loops if you want to alter the objects.
	for i := 0; i < len(planets); i++ {
		planets[i].update(dt)
	}
}

// draw is called after update and just draws everything visible
// to the screen.
func draw() {
	raylib.BeginDrawing()

	raylib.ClearBackground(raylib.Black)

	raylib.Begin2dMode(camera)

	// Draw block that is applied to the camera view.
	raylib.DrawCircle(0, 0, 15, raylib.Gold)

	for _, p := range planets {
		p.draw()
	}

	raylib.End2dMode()

	// Draw block applied outside of camera view (HUD elements etc.)
	raylib.DrawText(fmt.Sprintf("FPS: %d", int(raylib.GetFPS())), 10, 10, 32, raylib.RayWhite)

	raylib.EndDrawing()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	initScreen(title, screenWidth, screenHeight, fps)
	initPlayers("RagingDave", 0)
	initSolarSystem(5, 3, 100, int(screenHeight/2))

	// The main game loop is here. It periodically calls the update and draw functions.
	for !raylib.WindowShouldClose() {
		dt := raylib.GetFrameTime()
		update(dt)
		draw()
	}

	raylib.CloseWindow()
}
