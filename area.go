package main

type Tile struct {
  y Coord
  x Coord
  ch rune
  blockMove bool
  blockSight bool
}

type Area struct {
  tiles []Tile
  height int
  width int
}

func NewArea(h,w int) *Area {
  t := make([]Tile, w*h)
  SetPad(h,w)
  for y:=0; y < h; y++ {
    for x:=0; x < w; x++ {
      t[x+y*w].ch = '.'
    }
  }
  return &Area{t,h,w}
}
