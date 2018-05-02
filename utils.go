package main

import (
	"math"

	"github.com/gen2brain/raylib-go/raylib"
)

func sincos(radians float32) (float32, float32) {
	sn64, cs64 := math.Sincos(float64(radians))
	return float32(sn64), float32(cs64)
}

func frac(v float32) float32 {
	x := float64(v)
	return float32(x - math.Trunc(x))
}

func round(v float32) float32 { return float32(math.Round(float64(v))) }
func sqrt(v float32) float32  { return float32(math.Sqrt(float64(v))) }

// rotatePoint rotates point around anchor by omage (rad).
func rotatePoint(anchor, point *raylib.Vector2, omega float32) {
	sn, cs := sincos(omega)
	nx := anchor.X + (point.X-anchor.X)*cs - (point.Y-anchor.Y)*sn
	ny := anchor.Y + (point.X-anchor.X)*sn + (point.Y-anchor.Y)*cs

	point.X = nx
	point.Y = ny
}

func DrawTextureV(texture raylib.Texture2D, position raylib.Vector2, tint raylib.Color) {
	src := raylib.Rectangle{0, 0, texture.Width, texture.Height}
	ix, iy := int32(position.X), int32(position.Y)
	dx, dy := position.X-float32(ix), position.Y-float32(iy)
	dst := raylib.Rectangle{ix, iy, texture.Width, texture.Height}
	org := raylib.Vector2{-dx, -dy}
	raylib.DrawTexturePro(texture, src, dst, org, 0, tint)
}
