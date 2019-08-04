package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func (g *game) drawPlayBoard() {

	// When the round is zero, we ask the player to select the board size
	switch g.round {
	case 0:
		g.window.SetContent(g.boardSizeSelect())
	default:
		g.window.SetContent(g.boardGrid())
	}
	g.window.Show()
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
	case g.winner != "":
		// playerSign
		labelSign = widget.NewLabel(fmt.Sprintf("Player %s has won", g.winner))
	case g.round > g.boardSize*g.boardSize:
		labelSign = widget.NewLabel("There are no winners")
	default:
		// playerSign
		labelSign = widget.NewLabel(fmt.Sprintf("Player: %s", g.getPlayerSign()))
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
	if text == "X" || text == "O" || g.winner != "" {
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
		switch char {
		case "3x3":
			g.setBoardSize(3)
		case "4x4":
			g.setBoardSize(4)
		case "5x5":
			g.setBoardSize(5)
		case "6x6":
			g.setBoardSize(6)
		case "7x7":
			g.setBoardSize(7)
		}
	}
	return g.addButton(char, action)
}
