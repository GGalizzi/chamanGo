package main

type Coord int

type Mob struct {
  y Coord
  x Coord
  ch rune
}

//Constructor for pointer allocation
func NewMob(y Coord, x Coord, ch rune) *Mob {
  return &Mob{y,x,ch}
}

func (m *Mob) Move(y Coord, x Coord) {
  Draw(m.y,m.x,' ') // TEMP:Will have to check previous tile.
  m.y += y
  m.x += x
}
