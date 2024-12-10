package day10

type Coordinate struct {
	x int
	y int
}

func validCoordinate(coordinate Coordinate, maxX int, maxY int) bool {
	return coordinate.x >= 0 && coordinate.x < maxX && coordinate.y >= 0 && coordinate.y < maxY
}

func coordinateForDirection(coordinate Coordinate, direction int) Coordinate {
	switch direction {
	case 0:
		return Coordinate{coordinate.x, coordinate.y - 1}
	case 1:
		return Coordinate{coordinate.x + 1, coordinate.y}
	case 2:
		return Coordinate{coordinate.x, coordinate.y + 1}
	case 3:
		return Coordinate{coordinate.x - 1, coordinate.y}
	}

	panic("Invalid direction")
}
