package day18

import (
	"adventofcode2024-go/helpers"
	"adventofcode2024-go/types"
	"fmt"
)

func createMaze(fallingBytes []types.Coordinate, maxX, maxY int, amountOfBytes int) (maze [][]string, startPosition types.Coordinate, endPosition types.Coordinate) {
	maze = make([][]string, maxY+1)
	startPosition = types.Coordinate{X: 0, Y: 0}
	endPosition = types.Coordinate{X: maxX, Y: maxY}

	for i := 0; i <= maxY; i++ {
		maze[i] = make([]string, maxX+1)
		for j := 0; j <= maxX; j++ {
			maze[i][j] = "."
		}
	}

	for i, fallingByte := range fallingBytes {
		if i >= amountOfBytes {
			break
		}
		maze[fallingByte.Y][fallingByte.X] = "#"
	}

	return
}

func aStar(maze [][]string, startPosition, endPosition types.Coordinate) int {
	openSet := make(map[types.Coordinate]bool)

	openSet[startPosition] = true

	gScore := make(map[types.Coordinate]int)
	gScore[startPosition] = 0

	fScore := make(map[types.Coordinate]int)
	fScore[startPosition] = scoreOfMove(startPosition, endPosition)

	cameFrom := make(map[types.Coordinate]types.Coordinate)

	for len(openSet) > 0 {
		current := types.Coordinate{}
		currentFScore := 0
		for coordinate := range openSet {
			coordinateFScore, coordinateHasFScore := fScore[coordinate]
			if currentFScore == 0 || (coordinateHasFScore && coordinateFScore < currentFScore) {
				current = coordinate
				currentFScore = fScore[coordinate]
			}
		}

		if current.X == endPosition.X && current.Y == endPosition.Y {
			return gScore[current]
		}

		delete(openSet, current)
		neighbors := current.AdjacentCoordinates()
		for _, neighbor := range neighbors {
			if !neighbor.IsValid(len(maze[0]), len(maze)) {
				continue
			}
			if maze[neighbor.Y][neighbor.X] == "#" {
				continue
			}
			tentativeGScore := gScore[current] + scoreOfMove(current, neighbor)

			if neighborGScore, neighborHasGScore := gScore[neighbor]; !neighborHasGScore || tentativeGScore < neighborGScore {
				cameFrom[neighbor] = current
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore + scoreOfMove(neighbor, endPosition)
				openSet[neighbor] = true
			}
		}
	}

	return -1
}

func printMaze(maze [][]string) {
	fmt.Println()
	for _, row := range maze {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func scoreOfMove(currentPosition, endPosition types.Coordinate) int {
	return helpers.ManhattanDistanceCoordinate(currentPosition, endPosition)
}
