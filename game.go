package main

import "fmt"

type GameState string

type Game struct {
  state GameState
  Mobs []*Mob
  Player *Mob
  Area *Area
}

func (g *Game) Init() {
  var y,x Coord
  g.state = "menu"
  g.Area,y,x = NewArea(130,150)
  g.Player = NewMob(y,x,'@',g.Area)
  g.Mobs = append(g.Mobs, g.Player)
}

func (s GameState) Menuing() bool {
  return s == "menu"
}

func (g *Game) Menu() GameState {
  Write(Percent(25,ConsoleHeight),ConsoleWidth/2,"Welcome to Grogue (name in progress)")
  Write(Percent(25,ConsoleHeight)+1,ConsoleWidth/2,"Press any key to continue, press 'L' to load")
  key := GetInput()
  if key == "L" {
    g.LoadGame()
  }
  Clear()
  return "playing"
}

func (s GameState) Quiting() bool {
  return s == "quit"
}

func (g *Game) Output() {
  DrawMap(g.Area)
  for _,m := range g.Mobs {
    Draw(m.Y,m.X,m.ch)
  }
  RefreshPad(int(g.Player.Y),int(g.Player.X))
  DebugLog(fmt.Sprintf("X: %d, Y: %d", g.Player.X,g.Player.Y))
}
func (g *Game) Input() {
  key := GetInput()

  switch key {
  case "8":
    g.Player.Move(-1,0)
  case "9":
    g.Player.Move(-1,1)
  case "6":
    g.Player.Move(0,1)
  case "3":
    g.Player.Move(1,1)
  case "2":
    g.Player.Move(1,0)
  case "1":
    g.Player.Move(1,-1)
  case "4":
    g.Player.Move(0,-1)
  case "7":
    g.Player.Move(-1,-1)
  case "S":
    g.SaveGame()
    MessageLog.log("Game Saved")
  case "Q":
    g.state = "quit"
  }
}
