package main

type Coord int

type Mob struct {
  y Coord
  x Coord
  ch rune
  area *Area
}

//Constructor for pointer allocation
func NewMob(y Coord, x Coord, ch rune, area *Area) *Mob {
  return &Mob{y,x,ch,area}
}

func (m *Mob) Move(y,x Coord) {
  if !m.area.IsBlocking(m.y+y,m.x+x) {
    Draw(m.y,m.x,' ') // TEMP:Will have to check previous tile.
    m.y += y
    m.x += x
    return
  }
  MessageLog.log("You bump against a wall.")
}
