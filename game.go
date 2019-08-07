package main

import (
	"fyne.io/fyne"
)

type game struct {
	board     []string
	boardSize int
	round     int
	winner    *player
	players   [2]*player

	// fyne toolkit
	app    fyne.App
	window fyne.Window
}

func newGame(app fyne.App) *game {
	return &game{
		app: app,
	}
}

func (g *game) start() {
	g.window = g.app.NewWindow("Tic Tac Go")

	// Set the board size
	g.window.SetContent(g.boardSizeSelect())
	g.window.ShowAndRun()
}
