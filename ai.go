package main

import "math"
import "fmt"

func (g *Game) processAi() {
	var dy, dx Coord
	for _, m := range g.Area.Mobs {
		if m != g.Player {
			ydist := g.Player.Y - m.Y
			xdist := g.Player.X - m.X
			distance := math.Sqrt(float64(xdist*xdist + ydist*ydist))
			dx, dy = Coord(Round(float64(int(xdist)/Round(distance)))), Coord(Round(float64(int(ydist)/Round(distance))))
			DebugLog(fmt.Sprintf("dx, dy, dist = %d, %d, %g->%d | xdist: %d - ydist: %d    ", dx, dy, distance, Round(distance), xdist, ydist))
			m.Move(dy, dx)
		}
	}
}
