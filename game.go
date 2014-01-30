package main

type GameState string

type Game struct {
	state  GameState
	//Mobs   []*Mob
	Player *Mob
	Area   *Area
}

func (g *Game) Init() {
	var y, x Coord
	g.state = "menu"
	g.Area, y, x = NewArea(240, 250)
	g.Player = NewMobWithStats(y, x, '@', g.Area, 30,30,10,5)
	g.Area.Mobs = append(g.Area.Mobs, g.Player)
  g.Area.Mobs = append(g.Area.Mobs, NewMobWithStats(y-10,x-1, 'O', g.Area,20,30,15,0))
}

func (s GameState) Menuing() bool {
	return s == "menu"
}

func (g *Game) Menu() GameState {
	Write(Percent(25, ConsoleHeight), ConsoleWidth/2, "Welcome to Chamán")
	Write(Percent(25, ConsoleHeight)+1, ConsoleWidth/2, "Press any key to continue, press 'L' to load")
	key := GetInput()
	if key == "L" {
    g.Init()
		g.LoadGame()
    SetPad(g.Area.Height, g.Area.Width)
    DrawMap(g.Area)
    Clear()
    return "playing"
	}
	Clear()
  g.Init()
	return "playing"
}

func (g *Game) Death() {
  Clear()
  Write(Percent(25, ConsoleHeight), ConsoleWidth/2, "Your body has perished")
  Write(Percent(25, ConsoleHeight)+1, ConsoleWidth/2, "You are in the spirit plane now.")
  GetInput()
  g.state = "quit"
}


func (s GameState) Quiting() bool {
	return s == "quit"
}

func (g *Game) Output() {
	DrawMap(g.Area)
  for _, i := range g.Area.Items {
    if i.Hp <= 0 {
      DrawColors(i.Y, i.X, i.ch, 1)
      continue
    }
    //Draw(i.Y, i.X, i.ch)
  }
	for _, m := range g.Area.Mobs {
		Draw(m.Y, m.X, m.ch)
	}
	RefreshPad(int(g.Player.Y), int(g.Player.X))
  g.Player.UpdateStats()
}
func (g *Game) Input() {
	key := GetInput()

	switch key {
	case "8":
		g.Player.Move(-1, 0)
	case "9":
		g.Player.Move(-1, 1)
	case "6":
		g.Player.Move(0, 1)
	case "3":
		g.Player.Move(1, 1)
	case "2":
		g.Player.Move(1, 0)
	case "1":
		g.Player.Move(1, -1)
	case "4":
		g.Player.Move(0, -1)
	case "7":
		g.Player.Move(-1, -1)
	case "S":
    if Confirm("Save and Quit? Y/N") {
      g.SaveGame()
      MessageLog.log("Game Saved")
      g.state = "quit"
    }
	case "Q":
		g.state = "quit"
	}

  g.processAi()
}
