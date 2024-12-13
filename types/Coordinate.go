package types

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) isValid(maxX, maxY int) bool {
	return c.X >= 0 && c.X < maxX && c.Y >= 0 && c.Y < maxY
}

func (c Coordinate) adjacentCoordinates() [4]Coordinate {
	return [4]Coordinate{
		{c.X, c.Y - 1},
		{c.X + 1, c.Y},
		{c.X, c.Y + 1},
		{c.X - 1, c.Y},
	}
}
