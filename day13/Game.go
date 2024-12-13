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
