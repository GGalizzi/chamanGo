package main

type Coord int

type Mob struct {
	Y    Coord
	X    Coord
	ch   rune
	area *Area // Area it's in.

  *stats
}

type stats struct {
  Hp int
  Att int
  Def int
}

//Constructor for pointer allocation
func NewMob(y Coord, x Coord, ch rune, area *Area) *Mob {
	return &Mob{y, x, ch, area, nil}
}

func newStats(hp,att,def int) *stats {
  return &stats{hp,att,def}
}

func NewMobWithStats(y Coord, x Coord, ch rune, area *Area, hp,att,def int) *Mob {
  return &Mob{y,x,ch,area,newStats(hp,att,def)}
}

func (m *Mob) Move(y, x Coord) {
  if blocks, hasMob := m.area.IsBlocking(m.Y+y, m.X+x); !blocks {
		//Draw(m.Y,m.X,' ') // TEMP:Will have to check previous tile.
    if !hasMob {
      m.Y += y
      m.X += x
      return
    }
    MessageLog.log("Mob there")
    return
	}
	MessageLog.log("You bump against a wall.")
}
