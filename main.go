package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

var (
	ScreenWidth  int32 = 800
	ScreenHeight int32 = 600
	FPS          int32 = 60
	Title              = "Orbiteer"
)

// initGame sets up:
// - window size, fps
func initGame(title string, width, height, fps int32) {
	raylib.InitWindow(width, height, title)
	raylib.SetTargetFPS(fps)
}

// update handles all logic changes in the game. This includes
// moving objects or handling input.
func update(dt float32) {

}

// draw is called after update and just draws everything visible
// to the screen.
func draw() {
	raylib.BeginDrawing()

	raylib.ClearBackground(raylib.Black)

	raylib.DrawCircle(raylib.GetScreenWidth()/2, raylib.GetScreenHeight()/2, 55, raylib.Gold)
	raylib.DrawText(fmt.Sprintf("FPS: %d", int(raylib.GetFPS())), 10, 10, 32, raylib.RayWhite)

	raylib.EndDrawing()
}

func main() {
	initGame(Title, ScreenWidth, ScreenHeight, FPS)

	// The main game loop is here. It periodically calls the update and draw functions.
	for !raylib.WindowShouldClose() {
		dt := raylib.GetFrameTime()
		update(dt)
		draw()
	}

	raylib.CloseWindow()
}
