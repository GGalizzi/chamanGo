package main

import "fmt"

type GameState string

type Game struct {
  state GameState
  mobs []*Mob
  player *Mob
  area *Area
}

func (g *Game) Init() {
  g.state = "menu"
  g.player = NewMob(5,5,'@')
  g.mobs = append(g.mobs, g.player)
  g.area = NewArea(80,200)
}

func (s GameState) Menuing() bool {
  return s == "menu"
}

func Menu() GameState {
  Write(5,5,"Welcome to Grogue (name in progress)")
  Write(6,5,"Press any key to continue")
  GetInput()
  Clear()
  return "playing"
}

func (s GameState) Quiting() bool {
  return s == "quit"
}

func (g *Game) Output() {
  DrawMap(g.area)
  for _,m := range g.mobs {
    Draw(m.y,m.x,m.ch)
  }
  DebugLog(fmt.Sprintf("X: %d, Y: %d", g.player.x,g.player.y))
}
func (g *Game) Input() {
  key := GetInput()

  switch key {
  case "8":
    g.player.Move(-1,0)
  case "9":
    g.player.Move(-1,1)
  case "6":
    g.player.Move(0,1)
  case "3":
    g.player.Move(1,1)
  case "2":
    g.player.Move(1,0)
  case "1":
    g.player.Move(1,-1)
  case "4":
    g.player.Move(0,-1)
  case "7":
    g.player.Move(-1,-1)
  case "Q":
    g.state = "quit"
  }
}
