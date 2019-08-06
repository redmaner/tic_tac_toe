package main

type player struct {
	name    string
	sign    string
	machine bool
}

func (g *game) getCurrentPlayer() *player {
	return g.players[g.round%2]
}
