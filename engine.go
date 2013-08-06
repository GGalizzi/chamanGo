package main

import (
	"encoding/gob"
	"fmt"
	"github.com/GGalizzi/gocurses"
	"math/rand"
	"os"
	"time"
)

type log struct {
	pad   *gocurses.Window // Pad used to display messages to player.
	line  int              //line where the message will be added
	dline int              //line where we start showing messages.
}

var GamePad *gocurses.Window     // Pad used to display play screen
var debugWindow *gocurses.Window // Window used to show debug information (X,Y,stuff)
var MessageLog log

var StatsWindow *gocurses.Window // Window used to show players stats.

var ConsoleHeight int // height and
var ConsoleWidth int  // width of the whole console the game was opened in

var ScreenHeight int // Height and width
var ScreenWidth int  // of the playable screen GamePad

var WorldHeight int // Height and width of
var WorldWidth int  // the area currently in.

func Init() {
	gocurses.Initscr()
	gocurses.Cbreak()
	gocurses.Noecho()
	gocurses.Stdscr.Keypad(true)
	gocurses.CursSet(0)
	if !gocurses.HasColors() {
		panic("No colors")
	}
	gocurses.StartColor()

	ConsoleHeight, ConsoleWidth = gocurses.Getmaxyx()
	ScreenHeight, ScreenWidth = Percent(85, ConsoleHeight), Percent(70, ConsoleWidth)
  StatsWindow = gocurses.NewWindow(ScreenHeight, ConsoleWidth - ScreenWidth, 0, ScreenWidth+1)

	debugWindow = gocurses.NewWindow(5, ConsoleWidth, ConsoleHeight-1, 1)
	MessageLog.pad = gocurses.NewPad(100, ScreenWidth)

	rand.Seed(time.Now().UnixNano())


  // Color pairs
  gocurses.InitPair(1, gocurses.COLOR_RED, gocurses.COLOR_BLACK)
}

//Sets the GamePad and WH-WW info to the current area in the game object.
func SetPad(h, w int) {
	GamePad = gocurses.NewPad(h, w)
	WorldHeight, WorldWidth = h, w
}

func End() {
	gocurses.End()
}

func Clear() {
	gocurses.Clear()
}

func Draw(y,x Coord, ch rune) {
	GamePad.Mvaddch(int(y), int(x), ch)
	//RefreshPad(int(y),int(x))
}

func DrawColors(y,x Coord, ch rune, col int) {
  GamePad.Attron(gocurses.ColorPair(col))
  GamePad.Mvaddch(int(y), int(x), ch)
  GamePad.Attroff(gocurses.ColorPair(col))
}

func DrawMap(a *Area) {
	for y := 0; y < a.Height; y++ {
		for x := 0; x < a.Width; x++ {
			GamePad.Mvaddch(y, x, a.Tiles[x+y*a.Width].Ch)
      // Debug
      if a.Tiles[x+y*a.Width].BlockMove &&
        a.Tiles[x+y*a.Width].Ch != '#' {
        GamePad.Mvaddch(y,x, '!')
        a.Tiles[x+y*a.Width].BlockMove = false
      }

      if !a.Tiles[x+y*a.Width].BlockMove &&
        a.Tiles[x+y*a.Width].Ch != '.' {
          GamePad.Mvaddch(y,x, '_')
        }

      // Debug
		}
	}
}

func RefreshPad(y int, x int) {
	fromY := Max(0, y-ScreenHeight/2)
	fromX := Max(0, x-ScreenWidth/2)

	//Now snap camera to walls if we are at the edges of world
	if bottomPoint := fromY + ScreenHeight; bottomPoint >= WorldHeight {
		fromY = (WorldHeight - ScreenHeight)
	}
	if rightmostPoint := fromX + ScreenWidth; rightmostPoint >= WorldWidth {
		fromX = (WorldWidth - ScreenWidth)
	}

	GamePad.PnoutRefresh(fromY, fromX, 0, 0, ScreenHeight-1, ScreenWidth-1)
}

func Write(y int, x int, s string) {
	gocurses.Mvaddstr(y, x, s)
}

func DebugLog(s string) {
	debugWindow.Mvaddstr(0, 0, "                         ")
	debugWindow.Mvaddstr(0, 0, s)
	debugWindow.NoutRefresh()
}

func (l *log) log(s string) {
	l.pad.Mvaddstr(l.line, 0, fmt.Sprintf("%v %d", s, l.line))
	l.pad.PnoutRefresh(l.dline, 0, ScreenHeight+1, 0, ConsoleHeight-2, ConsoleWidth)
	// Checks if we need to scroll the window
	if l.line >= ((ConsoleHeight - 2) - (ScreenHeight + 1)) {
		l.dline++
	}
	// Checks if we need to start over on the log. (TEMP)
	if l.line >= 100 {
		l.line = 0
		l.dline = 0
	} else {
		l.line++
	}
}

func (p *Mob) UpdateStats() {
  StatsWindow.Mvaddstr(1,0, fmt.Sprintf("HP: %d/%d", p.Hp, p.MaxHp))
  StatsWindow.NoutRefresh()
}

func GetInput() string {
	gocurses.Doupdate()
	return string(gocurses.Getch())
}

func Confirm(q string) bool {
  Write(ScreenHeight/2,ScreenWidth/2,q)
  key := GetInput()
  if key == "Y" { return true }
  return false
}

func (g *Game) SaveGame() {
	file, err := os.OpenFile("player.sav", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(g)
	if err != nil {
		panic(err)
	}
}

func (g *Game) LoadGame() {
	file, err := os.OpenFile("player.sav", os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(g)
	if err != nil {
		panic(err)
	}
}
