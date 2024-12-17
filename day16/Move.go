package day16

import (
	"adventofcode2024-go/helpers"
	"adventofcode2024-go/types"
)

type Move struct {
	coordinate types.Coordinate
	direction  int
}

func (move Move) DirectionTowards(otherCoordinate types.Coordinate) int {
	xDiff := otherCoordinate.X - move.coordinate.X
	yDiff := otherCoordinate.Y - move.coordinate.Y
	if xDiff > 0 {
		return 1
	} else if xDiff < 0 {
		return 3
	} else if yDiff > 0 {
		return 2
	} else if yDiff < 0 {
		return 0
	}
	return -1
}

func (move Move) DirectionChangesTo(otherCoordinate types.Coordinate) int {
	directionChanges := 0
	xDiff := otherCoordinate.X - move.coordinate.X
	yDiff := otherCoordinate.Y - move.coordinate.Y
	if xDiff != 0 {
		if move.direction == 0 || move.direction == 2 {
			directionChanges++
		} else if (move.direction > 0) != (xDiff > 0) {
			directionChanges++
		}
	}
	if yDiff != 0 {
		if move.direction == 1 || move.direction == 3 {
			directionChanges++
		} else if (move.direction > 1) != (yDiff > 0) {
			directionChanges++
		}
	}

	return directionChanges
}

func (move Move) DirectionChangesToMove(otherMove Move) int {
	directionDifference := helpers.AbsInt(move.direction - otherMove.direction)
	if directionDifference > 2 {
		return 4 - directionDifference
	}

	return directionDifference
}

func (move Move) AdjacentMoves() [4]Move {
	return [4]Move{
		{coordinate: move.coordinate.CoordinateForDirection(move.direction), direction: move.direction},
		{coordinate: move.coordinate, direction: (move.direction + 1) % 4},
		{coordinate: move.coordinate, direction: (move.direction + 3) % 4},
	}
}
