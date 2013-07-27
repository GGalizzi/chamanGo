package main

func main() {

  Init() // Start curses stuff
  defer End() // defer endwin

  var g Game
  g.Init()

  for !g.state.Quiting() {
    if g.state.Menuing() {
      g.state = Menu()
    }
    g.Output()
    g.Input()
  }

}
