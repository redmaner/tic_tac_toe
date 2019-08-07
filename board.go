package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func (g *game) drawPlayBoard() {

	switch {
	case g.players[0] == nil:
		g.window.SetContent(g.selectPlayer(0, "Player 1"))
	case g.players[1] == nil:
		g.window.SetContent(g.selectPlayer(1, "Player 2"))
	default:

		if g.players[0].sign == g.players[1].sign {
			g.players[1].sign = "%"
		}

		if g.getCurrentPlayer().machine {
			fmt.Println("Do specific machine code")
			return
		}

		// Do human code
		g.window.SetContent(g.boardGrid())
		g.window.Show()
	}
}

// BoardGrid is function that creates a container with the Tic Tac Toe board
// This function is dynamic, which means it can automatically adjust the size of the
// board depending on the board size of the game.
func (g *game) boardGrid() *fyne.Container {

	rows := make([]fyne.CanvasObject, g.boardSize+2)
	rowButtons := make([][]fyne.CanvasObject, g.boardSize)

	// Add the round to the view
	labelRound := widget.NewLabel(fmt.Sprintf("Round: %d", g.round))
	labelRound.Alignment = fyne.TextAlignLeading
	labelRound.TextStyle.Monospace = true
	rows[0] = labelRound

	var position int
	for row := 1; row <= g.boardSize; row++ {
		for pos := 1; pos <= g.boardSize; pos++ {
			position++
			rowButtons[row-1] = append(rowButtons[row-1], g.gameButton(g.board[position-1]))
		}
	}

	for rowNumber, row := range rowButtons {
		rows[rowNumber+1] = fyne.NewContainerWithLayout(layout.NewGridLayout(g.boardSize), row...)
	}

	var labelSign *widget.Label
	switch {
	case g.winner != nil:
		// playerSign
		labelSign = widget.NewLabel(fmt.Sprintf("%s has won!", g.winner.name))
	case g.round > g.boardSize*g.boardSize:
		labelSign = widget.NewLabel("There are no winners")
	default:
		// playerSign
		labelSign = widget.NewLabel(fmt.Sprintf("Player: %s", g.getCurrentPlayer().name))
	}
	labelSign.Alignment = fyne.TextAlignLeading
	labelSign.TextStyle.Monospace = true
	rows[g.boardSize+1] = labelSign

	return fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		rows...,
	)
}

// addButton returns a button
func (g *game) addButton(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	if text == g.players[0].sign || text == g.players[1].sign || g.winner != nil {
		button.Disable()
	}

	return button
}

// gameButton contains the game logic
func (g *game) gameButton(char string) *widget.Button {
	action := func() {
		g.playRound(char)
	}
	return g.addButton(char, action)
}

func (g *game) boardSizeSelect() *fyne.Container {

	label := widget.NewLabel("Select a board size:")
	label.Alignment = fyne.TextAlignLeading
	label.TextStyle.Monospace = true

	return fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		label,
		g.boardSizeButton("3x3", 3),
		g.boardSizeButton("4x4", 4),
		g.boardSizeButton("5x5", 5),
		g.boardSizeButton("6x6", 6),
		g.boardSizeButton("7x7", 7),
		g.boardSizeButton("8x8", 8),
		g.boardSizeButton("9x9", 9),
		g.boardSizeButton("10x10", 10),
	)
}

// boardSizeButton
func (g *game) boardSizeButton(char string, size int) *widget.Button {
	action := func() {
		g.setBoardSize(size)
	}
	return widget.NewButton(char, action)
}

// setBoardSize sets the board size
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
