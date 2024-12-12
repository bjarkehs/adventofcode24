package day12

type SidePosition struct {
	position  int
	direction int
}

func sidePositionForCoordinate(coordinate Coordinate, direction int) SidePosition {
	if direction%2 == 0 {
		return SidePosition{coordinate.y, direction}
	}

	return SidePosition{coordinate.x, direction}
}

func (sidePosition SidePosition) pointForCoordinate(coordinate Coordinate) int {
	if sidePosition.direction%2 == 0 {
		return coordinate.x
	}

	return coordinate.y
}
