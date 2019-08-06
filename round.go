package main

import "strconv"

func (g *game) playRound(pos string) {

	// make integer from pos
	posInt, _ := strconv.Atoi(pos)

	// Change board
	if g.board[posInt-1] != g.players[0].sign && g.board[posInt-1] != g.players[1].sign {
		g.board[posInt-1] = g.getCurrentPlayer().sign
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
