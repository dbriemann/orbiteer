package main

import (
	"fmt"
	"math"

	"github.com/gen2brain/raylib-go/raylib"
	"github.com/kr/pretty"
)

func rotatePoint(point, origin *raylib.Vector2, angle float32) {
	fmt.Println("--ROTATE--")
	pretty.Println(point)
	pretty.Println(origin)
	nx := origin.X + (point.X-origin.X)*float32(math.Cos(float64(angle))) - (point.Y-origin.Y)*float32(math.Sin(float64(angle)))
	ny := origin.Y + (point.X-origin.X)*float32(math.Sin(float64(angle))) + (point.Y-origin.Y)*float32(math.Cos(float64(angle)))
	point.X = nx
	point.Y = ny
}
