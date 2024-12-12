package day12

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) isValid(maxX, maxY int) bool {
	return c.x >= 0 && c.x < maxX && c.y >= 0 && c.y < maxY
}

func (c Coordinate) adjacentCoordinates() [4]Coordinate {
	return [4]Coordinate{
		{c.x, c.y - 1},
		{c.x + 1, c.y},
		{c.x, c.y + 1},
		{c.x - 1, c.y},
	}
}
