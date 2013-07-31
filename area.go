package main

type Tile struct {
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
  SetPad(h,w) // set engines pad
  for y:=0; y < h; y++ {
    for x:=0; x < w; x++ {
      if y == 0 || x == 0 || y == h-1 || x == w-1 {
        t[x+y*w] = Tile{'#',true,true}
        continue
      }
      t[x+y*w] = Tile{'.',false,false}
    }
  }
  return &Area{t,h,w}
}

func (a *Area) IsBlocking(y,x Coord) bool {
  return a.tiles[int(x)+int(y)*a.width].blockMove
}
