package main

import "strconv"

func (g *game) playRound(pos string) {

	// make integer from pos
	posInt, _ := strconv.Atoi(pos)

	// Change board
	if g.board[posInt-1] != "X" && g.board[posInt-1] != "O" {
		g.board[posInt-1] = g.getPlayerSign()
	}

	// Increase round
	g.round++

	// Check winner
	if winner, sign := g.checkWinner(); winner {
		g.winner = sign
	}

	// draw board
	g.drawPlayBoard()
}

func (g *game) getPlayerSign() string {
	switch g.round % 2 {
	case 0:
		return "X"
	default:
		return "O"
	}
}
