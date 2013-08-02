package main

import "math/rand"
import "fmt"

type Tile struct {
  Ch rune
  BlockMove bool
  BlockSight bool
}

type Area struct {
  Tiles []Tile
  Height int
  Width int
}

// Generates a new area, and returns the coordinates the player will spawn in.
func NewArea(h,w int) (*Area,Coord,Coord) {
  nIts := 6 // quantity of iterations, will be a parameter or calculation based on size of area
  var ry,rx Coord
  ry,rx=0,0
  t := make([][]Tile, nIts)
  SetPad(h,w) // set engines pad
  for it:=0; it < nIts; it++ {
    t[it] = make([]Tile, w*h)
    for y:=0; y < h; y++ {
      for x:=0; x < w; x++ {

        // On first iteration, place random tiles.
        if it == 0 {
          t[it][x+y*w] = randomTile()
        } else {

          // Otherwise check for wall placement.
          if mapBorders(y,x,w,h) || adjacentWalls(y,x,w,t[it-1]) >= 4 {
            t[it][x+y*w] = Tile{'#',true,true}
          } else {
            t[it][x+y*w] = Tile{'.',false,false}

            // If we are at last Iteration
            if it == nIts-1 {
              //set the spawn coords
              ry = Coord(y); rx = Coord(x)
            }
          }
        }
      }
    }
    /* Debug */
    if it == 0 { GetInput() }
    DrawMap(&Area{t[it],h,w})
    Write(50,2,fmt.Sprintf("Iteration: %d", it))
    RefreshPad(int(ry),int(rx))
    GetInput()
    /**/
  }
  return &Area{t[nIts-1],h,w},ry,rx
}

func (a *Area) IsBlocking(y,x Coord) bool {
  return a.Tiles[int(x)+int(y)*a.Width].BlockMove
}

func randomTile() Tile {
  if rand.Intn(100) <= 30 {
    return Tile{'#',true,true}
  }
  return Tile{'.',false,false}
}

func anyAdjacentWalls(y,x,w int, t []Tile) bool {
         // If this one was a wall.
  return t[x+y*w].BlockMove ||
         // Directly adjacent
         t[(x+1)+(y+1)*w].BlockMove || t[(x-1)+(y-1)*w].BlockMove ||
         t[(x+1)+(y-1)*w].BlockMove || t[(x-1)+(y+1)*w].BlockMove ||
         t[x+(y+1)*w].BlockMove || t[x+(y-1)*w].BlockMove ||
         t[(x+1)+y*w].BlockMove || t[(x-1)+y*w].BlockMove
}

func adjacentWalls(y,x,w int, t []Tile) int {
  counter := 0
  if t[x+y*w].BlockMove { counter += 2 }
  if t[(x+1)+(y+1)*w].BlockMove { counter++ }
  if t[(x-1)+(y-1)*w].BlockMove { counter++ }
  if t[(x+1)+(y-1)*w].BlockMove { counter++ }
  if t[x+(y+1)*w].BlockMove { counter++ }
  if t[x+(y-1)*w].BlockMove { counter++ }
  if t[(x+1)+y*w].BlockMove { counter++ }
  if t[(x-1)+y*w].BlockMove { counter++ }
  return counter
}

func mapBorders(y,x,w,h int) bool {
  return y == 0 || y == h-1 || x == 0 || x == w-1
}
