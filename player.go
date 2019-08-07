package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type player struct {
	name    string
	sign    string
	machine bool
}

func (g *game) getCurrentPlayer() *player {
	return g.players[g.round%2]
}

func (g *game) selectPlayer(p int, n string) *fyne.Container {

	g.players[p] = &player{
		name: n,
	}
	player := g.players[p]

	label := widget.NewLabel("Please enter your name:")
	label.Alignment = fyne.TextAlignLeading
	label.TextStyle.Monospace = true

	name := widget.NewEntry()
	name.SetText(n)
	name.OnChanged = func(n string) {
		player.name = n
	}

	labelSign := widget.NewLabel("Please select a sign")
	label.Alignment = fyne.TextAlignLeading
	label.TextStyle.Monospace = true

	return fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		label,
		name,
		labelSign,
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(4),
			g.playerSignButton(player, "X"),
			g.playerSignButton(player, "O"),
			g.playerSignButton(player, "#"),
			g.playerSignButton(player, "@"),
		),
	)
}

func (g *game) playerSignButton(p *player, sign string) *widget.Button {
	action := func() {
		p.sign = sign
		g.drawPlayBoard()
	}

	return widget.NewButton(sign, action)
}
