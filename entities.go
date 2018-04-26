package main

import "github.com/gen2brain/raylib-go/raylib"

type planet struct {
	orb
	*player

	arbiters []*planet
	size     float32
}

func newPlanet(dist, size float32, vel raylib.Vector2, anchor *raylib.Vector2, player *player) *planet {
	p := planet{
		orb: orb{
			dist:   dist,
			anchor: anchor,
			vel:    vel,
		},
		player:   player,
		size:     size,
		arbiters: []*planet{},
	}

	if anchor != nil {
		p.pos.X = anchor.X + dist
	} else {
		p.pos.X = dist
	}

	return &p
}

func (p *planet) update(dt float32) {
	dx, dy := p.rotate(dt)
	for i := 0; i < len(p.arbiters); i++ {
		p.arbiters[i].pos.X += dx
		p.arbiters[i].pos.Y += dy
	}
}

func (p *planet) draw() {
	raylib.DrawCircleV(p.pos, p.size, p.color)
}

type ship struct {
	orb
	*player
}
