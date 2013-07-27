package main

import "github.com/GGalizzi/gocurses"

var GamePad *gocurses.Window
var ScreenHeight int
var ScreenWidth int
var WorldHeight int
var WorldWidth int

func Init() {
  gocurses.Initscr()
  gocurses.Cbreak()
  gocurses.Noecho()
  gocurses.Stdscr.Keypad(true)
  gocurses.Curs_set(0)

  ScreenHeight,ScreenWidth = gocurses.Getmaxyx()
  ScreenHeight,ScreenWidth = Percent(90,ScreenHeight), Percent(90,ScreenWidth)
}

//Sets the GamePad and WH-WW info to the current area in the game object.
func SetPad(h,w int) {
  GamePad = gocurses.NewPad(h,w)
  WorldHeight,WorldWidth = h,w
}

func End() {
  gocurses.End()
}

func Clear() {
  gocurses.Clear()
}


func Draw(y Coord, x Coord, ch rune) {
  GamePad.Mvaddch(int(y),int(x), ch)
  refreshPad(int(y),int(x))
}

func DrawMap(a *Area) {
  for y := 0; y < a.height; y++ {
    for x := 0; x < a.width; x++ {
      GamePad.Mvaddch(y,x,a.tiles[x+y*a.width].ch)
    }
  }
}


func refreshPad(y int, x int) {
  fromY := Max(0,y-ScreenHeight/2)
  fromX := Max(0,x-ScreenWidth/2)

  //Now snap camera to walls if we are at the edges of world
  if bottomPoint := fromY + ScreenHeight; bottomPoint >= WorldHeight {
    fromY = (WorldHeight - ScreenHeight)
  }
  if rightmostPoint := fromX + ScreenWidth; rightmostPoint >= WorldWidth {
    fromX = (WorldWidth - ScreenWidth)
  }

  GamePad.PRefresh(fromY,fromX, 0,0,ScreenHeight-1,ScreenWidth-1)
}


func Write(y int, x int, s string) {
  gocurses.Mvaddstr(y,x,s)
}

func DebugLog(s string) {
  gocurses.Mvaddstr(ScreenHeight+1,1,"                         ")
  gocurses.Mvaddstr(ScreenHeight+1,1, s)
  gocurses.Refresh()
}

func GetInput() string {
  return string(gocurses.Getch())
}
