package day14

import "adventofcode2024-go/types"

func wrapAround(value int, maxValue int) int {
	newValue := value % maxValue
	if newValue < 0 {
		return maxValue + newValue
	}
	return newValue
}

func positionAfterMoves(position types.Coordinate, velocity types.Coordinate, moves int, maxX int, maxY int) types.Coordinate {
	newX := wrapAround(position.X+velocity.X*moves, maxX)
	newY := wrapAround(position.Y+velocity.Y*moves, maxY)

	return types.Coordinate{X: newX, Y: newY}
}

func moveRobots(robots []Robot, moves int, maxX, maxY int) []types.Coordinate {
	coordinates := make([]types.Coordinate, 0)
	for _, robot := range robots {
		newPosition := positionAfterMoves(robot.Position, robot.Velocity, moves, maxX, maxY)
		coordinates = append(coordinates, newPosition)
	}

	return coordinates
}

func moveRobotsInMap(robots []Robot, moves int, maxX, maxY int) map[int]map[int]int {
	coordinatesMap := make(map[int]map[int]int)
	for _, robot := range robots {
		newPosition := positionAfterMoves(robot.Position, robot.Velocity, moves, maxX, maxY)
		if _, ok := coordinatesMap[newPosition.Y]; !ok {
			coordinatesMap[newPosition.Y] = make(map[int]int)
		}
		coordinatesMap[newPosition.Y][newPosition.X] += 1
	}

	return coordinatesMap
}
