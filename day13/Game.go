package day13

import "adventofcode2024-go/types"

type Move struct {
	coordinate types.Coordinate
	cost       int
}

type Game struct {
	moveA         Move
	moveB         Move
	prizePosition types.Coordinate
}

func (game Game) solve() int {
	a1 := float64(game.moveA.coordinate.X)
	a2 := float64(game.moveA.coordinate.Y)
	b1 := float64(game.moveB.coordinate.X)
	b2 := float64(game.moveB.coordinate.Y)
	p1 := float64(game.prizePosition.X)
	p2 := float64(game.prizePosition.Y)

	determinant := a1*b2 - a2*b1
	x := (p1*b2 - p2*b1) / determinant
	y := (a1*p2 - a2*p1) / determinant

	if x >= 0 && y >= 0 && isInteger(x) && isInteger(y) {
		return game.moveA.cost*int(x) + game.moveB.cost*int(y)
	}

	return 0
}
