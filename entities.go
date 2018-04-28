package main

import (
	"math"

	"github.com/gen2brain/raylib-go/raylib"
)

type planet struct {
	orb
	*player

	satellites []*planet
	shipCount  float32
	size       float32
}

func newPlanet(dist, size, dir float32, vel raylib.Vector2, anchor *raylib.Vector2, player *player) *planet {
	p := planet{
		orb: orb{
			dist:   dist,
			anchor: anchor,
			vel:    vel,
			dir:    dir,
		},
		player:     player,
		size:       size,
		satellites: []*planet{},
	}

	if anchor != nil {
		p.pos.X = anchor.X + dist
	} else {
		p.pos.X = dist
	}

	return &p
}

func (p *planet) update(dt float32) {
	// Rotate the planet and adjust the position of its satellites.
	dx, dy := p.rotate(dt)
	for i := 0; i < len(p.satellites); i++ {
		p.satellites[i].pos.X += dx
		p.satellites[i].pos.Y += dy
	}
	// Ship production depends on planet size: production = sqrt(radius)/5
	prod := math.Sqrt(float64(p.size)) / 5
	p.shipCount += float32(prod) * dt
}

func (p *planet) draw() {
	// Draw a black "background circle" to imitate a shadow.
	raylib.DrawCircleV(p.pos, p.size+2, raylib.Black)
	// And of course draw the planet itself.
	raylib.DrawCircleV(p.pos, p.size, p.color)
	// And all of its satellites are connected via a line
	for _, a := range p.satellites {
		raylib.DrawLineV(p.pos, a.pos, a.color)
	}
}

type ship struct {
	orb
	*player
}
