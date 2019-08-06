package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func (g *game) drawPlayBoard() {

	switch {
	case g.players[0] == nil:
		fmt.Println("We need to configure player 1")
	case g.players[1] == nil:
		fmt.Println("We need to configure player 2")
	default:
		if g.getCurrentPlayer().machine {
			fmt.Println("Do specific machine code")
			break
		}

		// Do human code
		g.window.SetContent(g.boardGrid())
		g.window.Show()
	}
}

func (g *game) boardSizeSelect() *fyne.Container {

	label := widget.NewLabel("Select a board size:")
	label.Alignment = fyne.TextAlignLeading
	label.TextStyle.Monospace = true

	return fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		label,
		g.boardSizeButton("3x3"),
		g.boardSizeButton("4x4"),
		g.boardSizeButton("5x5"),
		g.boardSizeButton("6x6"),
		g.boardSizeButton("7x7"),
		g.boardSizeButton("8x8"),
		g.boardSizeButton("9x9"),
		g.boardSizeButton("10x10"),
	)
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
		labelSign = widget.NewLabel(fmt.Sprintf("Player %s has won", g.winner.name))
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

// boardSizeButton
func (g *game) boardSizeButton(char string) *widget.Button {
	action := func() {
		str := strings.Split(char, "x")
		size, _ := strconv.Atoi(str[0])
		g.setBoardSize(size)
	}
	return g.addButton(char, action)
}
