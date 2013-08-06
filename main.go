package main

func main() {

	Init()      // Start curses stuff
	defer End() // defer endwin

	var g Game
  g.state = "menu"
	//g.Init()

	for !g.state.Quiting() {
		if g.state.Menuing() {
			g.state = g.Menu()
			continue
		}
		g.Output()
		g.Input()
	}

}
