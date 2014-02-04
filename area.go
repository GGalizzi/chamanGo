package main

import "math/rand"
import "fmt"

type Coords struct {
	y int
	x int
}

type Tile struct {
	Ch         rune
	BlockMove  bool
	BlockSight bool
}

type Area struct {
	Tiles []Tile
	Mobs  []*Mob
	Items []*Mob

	Height int
	Width  int
}

// Generates a new area, and returns the coordinates the player will spawn in.
func NewArea(h, w int) (*Area, Coord, Coord) {
	nIts := 4 // quantity of iterations, will be a parameter or calculation based on size of area
	var ry, rx Coord
	ry, rx = 0, 0
	t := make([][]Tile, nIts)
	SetPad(h, w) // set engines pad
	for it := 0; it < nIts; it++ {
		t[it] = make([]Tile, w*h)
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {

				// On first iteration, place random tiles.
				if it == 0 {
					t[it][x+y*w] = placeRandomTile()
					continue
				}

				// Otherwise check for wall placement.
				if mapBorders(y, x) || adjacentWalls(y, x, w, t[it-1]) >= 4 {
					t[it][x+y*w] = Tile{'#', true, true}
					continue
				}
				// Or draw ground.
				t[it][x+y*w] = Tile{'.', false, false}

				// If we are at last Iteration
				if it == nIts-1 {
					// Select random tile, and explode it with tiles of same type.
					/*
					   yy,xx := selectRandomTile(h,w)
					   if !t[it][xx+yy*w].BlockMove { explodeTile(yy,xx,w,&t[it]) }
					*/
					// Find the first ground tile, and flood it with a region tile.
					//yy, xx := firstGroundTile(&t[it])
					//floodFill(yy, xx, &t[it])
					//set the spawn coords
					ry = Coord(y)
					rx = Coord(x)
				}
			}
		}
		/* Debug */ /*
		   if it == 0 { GetInput() }
		   DrawMap(&Area{t[it],h,w})
		   Write(50,2,fmt.Sprintf("Iteration: %d", it))
		   RefreshPad(int(ry),int(rx))
		   GetInput()
		   /**/
		Write(51, 2, fmt.Sprint("Done"))
	}
	return &Area{t[nIts-1], nil, nil, h, w}, ry, rx
}

// Returns if the tile in given coords has BlockMove attribute.
// Also checks if there's a mob in that tile, returning
// a pointer to it if there is.
// Otherwise hasMob is nil.
func (a *Area) IsBlocking(y, x Coord) (blocks bool, hasMob *Mob) {
	blocks = a.Tiles[int(x)+int(y)*a.Width].BlockMove
	for _, m := range a.Mobs {
		if m.Hp > 0 && m.X == x && m.Y == y {
			hasMob = m
			return
		}
	}
	return
}

func placeRandomTile() Tile {
	if rand.Intn(100) <= 30 {
		return Tile{'#', true, true}
	}
	return Tile{'.', false, false}
}

// Returns a random set of coordinates.
func selectRandomTile(h, w int) (int, int) {
	y := rand.Intn(h)
	x := rand.Intn(w)
	return y, x
}

// With the tile given as argument make some more of those randomly around it.
func explodeTile(y, x, w int, t *[]Tile) {
	originalTile := (*t)[x+y*w]
	for it := 0; it < 5; it++ {
		ry := rand.Intn(2)
		rx := rand.Intn(2)
		if rand.Intn(100) > 50 {
			ry *= -1
		}
		if rand.Intn(100) > 50 {
			rx *= -1
		}
		defer func() {
			if r := recover(); r != nil {
				ry, rx = ry*-1, rx*-1
			}
		}()
		if !mapBorders(y+ry, x+rx) || !mapBorders(y, x) {
			(*t)[(x+rx)+(y+ry)*w] = originalTile
		}
		y, x = y+ry, x+rx
	}
}

// Searches for the first walkable tile on the map.
func firstGroundTile(t *[]Tile) (int, int) {
	for y := 0; y < WorldHeight; y++ {
		for x := 0; x < WorldWidth; x++ {
			if !(*t)[x+y*WorldWidth].BlockMove {
				return y, x
			}
		}
	}
	return 0, 0
}

func floodFill(y, x int, t *[]Tile) {
	c := make([]Coords, WorldHeight*WorldWidth)
	c[0] = Coords{y, x}
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	for it := 0; it < 9; it++ {
		for i, coord := range c {
			if coord.x >= 0 && coord.x < WorldWidth &&
				coord.y >= 0 && coord.y < WorldHeight &&
				!(*t)[coord.x+coord.y*WorldWidth].BlockMove {
				(*t)[coord.x+coord.y*WorldWidth] = Tile{'1', false, false}
				appendCoords(coord.y, coord.x, &c, t)
				c = append(c[:i], c[i+1:]...)
			}
		}
	}
}

func appendCoords(y, x int, c *[]Coords, t *[]Tile) {
	w := WorldWidth
	dy, dx := y+1, x+1
	if withinBounds(dy, dx) &&
		(*t)[dx+dy*w].Ch == '.' {
		*c = append(*c, Coords{dx, dy})
	}
	dy, dx = y-1, x-1
	if withinBounds(dy, dx) &&
		(*t)[dx+dy*w].Ch == '.' {
		*c = append(*c, Coords{dx, dy})
	}
	dy, dx = y-1, x+1
	if withinBounds(dy, dx) &&
		(*t)[dx+dy*w].Ch == '.' {
		*c = append(*c, Coords{dx, dy})
	}
	dy, dx = y+1, x-1
	if withinBounds(dy, dx) &&
		(*t)[dx+dy*w].Ch == '.' {
		*c = append(*c, Coords{dx, dy})
	}
	dy, dx = y+1, x
	if withinBounds(dy, dx) &&
		(*t)[dx+dy*w].Ch == '.' {
		*c = append(*c, Coords{dx, dy})
	}
	dy, dx = y-1, x
	if withinBounds(dy, dx) &&
		(*t)[dx+dy*w].Ch == '.' {
		*c = append(*c, Coords{dx, dy})
	}
	dy, dx = y, x+1
	if withinBounds(dy, dx) &&
		(*t)[dx+dy*w].Ch == '.' {
		*c = append(*c, Coords{dx, dy})
	}
	dy, dx = y, x-1
	if withinBounds(dy, dx) &&
		(*t)[dx+y*w].Ch == '.' {
		*c = append(*c, Coords{dx, dy})
	}
}

// returns true if the coords are not out of range.
func withinBounds(y, x int) bool {
	return y > 0 && y < WorldHeight &&
		x > 0 && x < WorldWidth
}

// Returns true if any of the adjacent tiles block move
func anyAdjacentWalls(y, x, w int, t []Tile) bool {
	// If this one was a wall.
	return t[x+y*w].BlockMove ||
		// Directly adjacent
		t[(x+1)+(y+1)*w].BlockMove || t[(x-1)+(y-1)*w].BlockMove ||
		t[(x+1)+(y-1)*w].BlockMove || t[(x-1)+(y+1)*w].BlockMove ||
		t[x+(y+1)*w].BlockMove || t[x+(y-1)*w].BlockMove ||
		t[(x+1)+y*w].BlockMove || t[(x-1)+y*w].BlockMove
}

// Returns a value depending on how many walls are adjacent.
// Used for cave-like generation.
// Being a wall counts for 2 for better generation.
func adjacentWalls(y, x, w int, t []Tile) int {
	counter := 0
	if t[x+y*w].BlockMove {
		counter += 2
	}
	if t[(x+1)+(y+1)*w].BlockMove {
		counter++
	}
	if t[(x-1)+(y-1)*w].BlockMove {
		counter++
	}
	if t[(x+1)+(y-1)*w].BlockMove {
		counter++
	}
	if t[x+(y+1)*w].BlockMove {
		counter++
	}
	if t[x+(y-1)*w].BlockMove {
		counter++
	}
	if t[(x+1)+y*w].BlockMove {
		counter++
	}
	if t[(x-1)+y*w].BlockMove {
		counter++
	}
	return counter
}

func mapBorders(y, x int) bool {
	return y == 0 || y == WorldHeight-1 || x == 0 || x == WorldWidth-1
}
