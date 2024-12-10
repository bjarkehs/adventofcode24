package day6

type Coordinate struct {
	x int
	y int
}

func coordinateInDirection(coordinate Coordinate, direction int) Coordinate {
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

func turnRight(direction int) int {
	return (direction + 1) % 4
}

func moveGuard(guardPosition Coordinate, direction int, obstacles []Coordinate, maxX int, maxY int) (bool, Coordinate, int) {
	nextCoordinate := coordinateInDirection(guardPosition, direction)

	if nextCoordinate.x < 0 || nextCoordinate.y < 0 || nextCoordinate.x >= maxX || nextCoordinate.y >= maxY {
		return false, nextCoordinate, direction
	}

	for _, obstacle := range obstacles {
		if obstacle.x == nextCoordinate.x && obstacle.y == nextCoordinate.y {
			return moveGuard(guardPosition, turnRight(direction), obstacles, maxX, maxY)
		}
	}

	return true, nextCoordinate, direction
}

func moveGuardV2(guardPosition Coordinate, direction int, obstacleMap [][]bool, maxX int, maxY int) (bool, Coordinate, int) {
	nextCoordinate := coordinateInDirection(guardPosition, direction)

	if nextCoordinate.x < 0 || nextCoordinate.y < 0 || nextCoordinate.x >= maxX || nextCoordinate.y >= maxY {
		return false, nextCoordinate, direction
	}

	if obstacleMap[nextCoordinate.y][nextCoordinate.x] {
		return moveGuardV2(guardPosition, turnRight(direction), obstacleMap, maxX, maxY)
	}

	return true, nextCoordinate, direction
}
