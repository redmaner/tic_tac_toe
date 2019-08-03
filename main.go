package main

import "fyne.io/fyne/app"

func main() {

	// We use the Fyne toolkit for the GUI of the Tic Tac Toe game
	// Fyne always starts with an app
	app := app.New()

	g := newGame(app)
	g.start()
}
