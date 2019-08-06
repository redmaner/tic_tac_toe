package main

import (
	"strconv"

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

func (g *game) setBoardSize(boardSize int) {
	board := make([]string, 0)
	for i := 1; i <= boardSize*boardSize; i++ {
		board = append(board, strconv.Itoa(i))
	}

	g.board = board
	g.boardSize = boardSize
	g.round = 1
	g.drawPlayBoard()
}

func (g *game) start() {
	g.window = g.app.NewWindow("Tic Tac Go")

	g.players[0] = &player{
		name:    "Red",
		sign:    "X",
		machine: false,
	}

	g.players[1] = &player{
		name:    "Maner",
		sign:    "O",
		machine: false,
	}

	// Set the board size
	g.window.SetContent(g.boardSizeSelect())
	g.window.ShowAndRun()
}
