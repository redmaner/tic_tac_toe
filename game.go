package main

import (
	"strconv"

	"fyne.io/fyne"
)

type game struct {
	board     []string
	boardSize int
	round     int
	winner    string

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
	g.drawPlayBoard()
	g.app.Run()
}
