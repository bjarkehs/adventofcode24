package day8

type Coordinate struct {
	x int
	y int
}

func validCoordinate(coordinate Coordinate, maxX int, maxY int) bool {
	return coordinate.x >= 0 && coordinate.x < maxX && coordinate.y >= 0 && coordinate.y < maxY
}
