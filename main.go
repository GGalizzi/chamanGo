package main

var G Game
func main() {

	Init()      // Start curses stuff
	defer End() // defer endwin

  G.state = "menu"
	//g.Init()

	for !G.state.Quiting() {
		if G.state.Menuing() {
			G.state = G.Menu()
			continue
		}
		G.Output()
		G.Input()
	}

}
