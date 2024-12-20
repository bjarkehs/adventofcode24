package day20

import (
	"adventofcode2024-go/helpers"
	"adventofcode2024-go/types"
	"container/list"
)

func BFS(maze [][]string, start types.Coordinate) map[types.Coordinate]int {
	queue := list.New()
	visited := make(map[types.Coordinate]bool)
	distances := make(map[types.Coordinate]int)

	distances[start] = 0
	queue.PushBack(start)

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		coordinate := element.Value.(types.Coordinate)

		if visited[coordinate] {
			continue
		}

		visited[coordinate] = true

		if maze[coordinate.Y][coordinate.X] == "#" {
			continue
		}

		neighbors := coordinate.AdjacentCoordinates()
		for _, neighbor := range neighbors {
			newDistance := distances[coordinate] + 1
			existingDistance, hasDistance := distances[neighbor]
			if !hasDistance || newDistance < existingDistance {
				distances[neighbor] = newDistance
				queue.PushBack(neighbor)
			}
		}
	}

	return distances
}

type Cheat struct {
	coordinate types.Coordinate
	steps      int
}

func CheatCells(maze [][]string, location types.Coordinate, maxCheatSteps int, endPosition types.Coordinate) []types.Coordinate {
	maxX := len(maze[0])
	maxY := len(maze)
	queue := list.New()

	visited := make(map[types.Coordinate]bool)
	queue.PushBack(Cheat{coordinate: location, steps: 0})

	cheatPositions := make([]types.Coordinate, 0)

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		cheat := element.Value.(Cheat)

		if visited[cheat.coordinate] {
			continue
		}

		visited[cheat.coordinate] = true

		if (cheat.coordinate == endPosition || maze[cheat.coordinate.Y][cheat.coordinate.X] == ".") && cheat.steps > 0 {
			cheatPositions = append(cheatPositions, cheat.coordinate)
		}

		if cheat.steps >= maxCheatSteps {
			continue
		}

		neighbors := cheat.coordinate.AdjacentCoordinates()
		for _, neighbor := range neighbors {
			if !neighbor.IsValid(maxX, maxY) {
				continue
			}
			queue.PushBack(Cheat{coordinate: neighbor, steps: cheat.steps + 1})
		}
	}

	return cheatPositions
}

func SolveMaze(maze [][]string, startPosition, endPosition types.Coordinate, maxCheatSteps int, minSave int) int {
	distancesFromStart := BFS(maze, startPosition)
	distancesFromEnd := BFS(maze, endPosition)

	cheats := 0
	for y, row := range maze {
		for x, cell := range row {
			if cell == "#" {
				continue
			}
			currentCoordinate := types.Coordinate{X: x, Y: y}
			possibleCheatLocations := CheatCells(maze, currentCoordinate, maxCheatSteps, endPosition)
			for _, cheatLocation := range possibleCheatLocations {
				newDistance := distancesFromStart[currentCoordinate] + distancesFromEnd[cheatLocation] + helpers.ManhattanDistanceCoordinate(currentCoordinate, cheatLocation)
				difference := distancesFromStart[endPosition] - newDistance
				if difference >= minSave {
					cheats++
				}
			}
		}
	}

	return cheats
}
