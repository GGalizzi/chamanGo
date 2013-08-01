package main

type Coord int

type Mob struct {
  Y Coord
  X Coord
  ch rune
  area *Area
}

//Constructor for pointer allocation
func NewMob(y Coord, x Coord, ch rune, area *Area) *Mob {
  return &Mob{y,x,ch,area}
}

func (m *Mob) Move(y,x Coord) {
  if !m.area.IsBlocking(m.Y+y,m.X+x) {
    Draw(m.Y,m.X,' ') // TEMP:Will have to check previous tile.
    m.Y += y
    m.X += x
    return
  }
  MessageLog.log("You bump against a wall.")
}
