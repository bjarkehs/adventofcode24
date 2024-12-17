package types

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) IsValid(maxX, maxY int) bool {
	return c.X >= 0 && c.X < maxX && c.Y >= 0 && c.Y < maxY
}

func (c Coordinate) AdjacentCoordinates() [4]Coordinate {
	return [4]Coordinate{
		{c.X, c.Y - 1},
		{c.X + 1, c.Y},
		{c.X, c.Y + 1},
		{c.X - 1, c.Y},
	}
}

func (c Coordinate) CoordinateForDirection(direction int) Coordinate {
	return c.CoordinateForDirectionWithOffset(direction, 1)
}

func (c Coordinate) CoordinateForDirectionWithOffset(direction int, offset int) Coordinate {
	switch direction {
	case 0:
		return Coordinate{c.X, c.Y - offset}
	case 1:
		return Coordinate{c.X + offset, c.Y}
	case 2:
		return Coordinate{c.X, c.Y + offset}
	case 3:
		return Coordinate{c.X - offset, c.Y}
	default:
		panic("Invalid direction")
	}
}
