package main

import (
	"math"

	"github.com/gen2brain/raylib-go/raylib"
)

type planet struct {
	orb
	*player

	satellites    []*planet
	ships         []*ship
	shipsProduced float32
	size          float32
}

func newPlanet(dist, size, dir float32, vel raylib.Vector2, anchor *raylib.Vector2, player *player) *planet {
	p := &planet{
		orb: orb{
			dist:   dist,
			anchor: anchor,
			vel:    vel,
			dir:    dir,
		},
		player:     player,
		size:       size,
		satellites: []*planet{},
		ships:      make([]*ship, int(size/3)),
	}

	if anchor != nil {
		p.pos.X = anchor.X + dist
	} else {
		p.pos.X = dist
	}

	for i := 0; i < len(p.ships); i++ {
		p.ships[i] = newShip(p, player)
	}
	p.distributeShips()

	return p
}

func (p *planet) rotateAll(dt float32) {
	// Rotate the planet and adjust the position of its satellites.
	dx, dy := p.rotate(dt)
	for i := 0; i < len(p.satellites); i++ {
		p.satellites[i].pos.X += dx
		p.satellites[i].pos.Y += dy
	}
}

func (p *planet) update(dt float32) {
	p.rotateAll(dt)
	// Ship production depends on planet size: production = sqrt(radius)/5
	prod := math.Sqrt(float64(p.size)) * productionFactor
	p.shipsProduced += float32(prod) * dt

	// Add new ships to slice.
	for i := 0; i < int(math.Floor(float64(p.shipsProduced))); i++ {
		added := false
		// Search a free spot and if there is none append.
		nship := newShip(p, p.player)
		for j := 0; j < len(p.ships); j++ {
			if p.ships[i] == nil {
				p.ships[i] = nship
				added = true
			}
		}
		if !added {
			p.ships = append(p.ships, nship)
		}
		p.shipsProduced--
	}

	p.distributeShips()
}

func (p *planet) draw() {
	// Draw a black "background circle" to imitate a shadow.
	raylib.DrawCircleV(p.pos, p.size+2, raylib.Black)
	// And of course draw the planet itself.
	raylib.DrawCircleV(p.pos, p.size, p.color)
	// And all of its satellites are connected via a line
	for _, s := range p.satellites {
		raylib.DrawLineV(p.pos, s.pos, s.color)
	}

	// Draw all ships stationed at this planet.
	for _, s := range p.ships {
		raylib.DrawCircleV(s.pos, 1, s.color)
	}
}

func (p *planet) distributeShips() {
	amount := len(p.ships)
	step := (2 * math.Pi) / float32(amount)

	for i := 0; i < amount; i++ {
		p.ships[i].pos.X = p.pos.X + p.ships[i].dist
		p.ships[i].pos.Y = p.pos.Y

		omega := float32(i) * step
		rotatePoint(&p.pos, &p.ships[i].pos, omega)
	}
}

type ship struct {
	orb
	*player
}

func newShip(planet *planet, player *player) *ship {
	// Test if there is a ship available for recycling.
	var sp = &ship{}
	i := -1
	for i, sp = range recycledShips {
		if sp != nil {
			break
		}
	}
	if sp != nil && i >= 0 {
		// Remove ship from recycled slice.
		recycledShips[i] = nil
	}

	// TODO magic numbers
	sp.dist = planet.size * 2
	sp.anchor = &planet.pos
	sp.vel = raylib.Vector2{X: 5, Y: 5}
	sp.dir = 1
	sp.player = player

	return sp
}

// trash removes a ship from a planet and puts it into the
// recycledShips slice for later reuse.
func (s *ship) trash(p *planet) {
	// Add the ship to recycled first to hold the pointer.
	added := false
	for i := 0; i < len(recycledShips); i++ {
		if recycledShips[i] == nil {
			recycledShips[i] = s
			added = true
		}
	}
	if !added {
		recycledShips = append(recycledShips, s)
	}

	// Now remove the ship from planet ships.
	for i := 0; i < len(p.ships); i++ {
		if p.ships[i] == s {
			p.ships[i] = nil
		}
	}
}
