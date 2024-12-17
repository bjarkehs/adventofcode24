package day16

import (
	"adventofcode2024-go/helpers"
	"adventofcode2024-go/types"
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("day16/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	maze := make([][]string, 0)
	endPosition := types.Coordinate{X: 0, Y: 0}
	endPositionFound := false

	startPosition := types.Coordinate{X: 0, Y: 0}
	startPositionFound := false
	for scanner.Scan() {
		lineText := scanner.Text()
		maze = append(maze, strings.Split(lineText, ""))
		if !endPositionFound {
			if indexOfEnd := strings.Index(lineText, "E"); indexOfEnd != -1 {
				endPosition = types.Coordinate{X: indexOfEnd, Y: len(maze) - 1}
				endPositionFound = true
			}
		}

		if !startPositionFound {
			if indexOfStart := strings.Index(lineText, "S"); indexOfStart != -1 {
				startPosition = types.Coordinate{X: indexOfStart, Y: len(maze) - 1}
				startPositionFound = true
			}
		}
	}

	cost, _ := dijkstra2(maze, startPosition, endPosition, 1)
	//printMaze(maze, path)
	//fmt.Println(path)
	fmt.Println(cost)
	//fmt.Println(recalculateCostOfPath(path))
}

func aStar(maze [][]string, startPosition, endPosition types.Coordinate, initialDirection int) (int, []Move) {
	startMove := Move{coordinate: startPosition, direction: initialDirection}
	openSet := make(map[Move]bool)

	openSet[startMove] = true

	gScore := make(map[Move]int)
	gScore[startMove] = 0

	fScore := make(map[Move]int)
	fScore[startMove] = scoreOfMove(startMove, endPosition)

	cameFrom := make(map[Move]Move)

	for len(openSet) > 0 {
		current := Move{}
		currentFScore := 0
		for move := range openSet {
			moveFScore, moveHasFScore := fScore[move]
			if currentFScore == 0 || (moveHasFScore && moveFScore < currentFScore) {
				current = move
				currentFScore = fScore[move]
			}
		}

		if current.coordinate.X == endPosition.X && current.coordinate.Y == endPosition.Y {
			reconstructedPath := reconstructPath(cameFrom, current)
			slices.Reverse(reconstructedPath)
			return gScore[current], reconstructedPath
		}

		delete(openSet, current)
		neighbors := current.coordinate.AdjacentCoordinates()
		for _, neighbor := range neighbors {
			if maze[neighbor.Y][neighbor.X] == "#" {
				continue
			}
			neighborDirection := current.DirectionTowards(neighbor)
			if neighborDirection == -1 {
				continue
			}
			neighborMove := Move{coordinate: neighbor, direction: current.DirectionTowards(neighbor)}
			tentativeGScore := gScore[current] + scoreOfMove(current, neighbor)

			if neighborGScore, neighborHasGScore := gScore[neighborMove]; !neighborHasGScore || tentativeGScore < neighborGScore {
				cameFrom[neighborMove] = current
				gScore[neighborMove] = tentativeGScore
				fScore[neighborMove] = tentativeGScore + scoreOfMove(neighborMove, endPosition)
				openSet[neighborMove] = true
			}
		}
	}

	return -1, []Move{}
}

func dijkstra(maze [][]string, startPosition, endPosition types.Coordinate, initialDirection int) (int, []Move) {
	cost := make(map[Move]int)
	cameFrom := make(map[Move]Move)

	startMove := Move{coordinate: startPosition, direction: initialDirection}
	cost[startMove] = 0

	queue := list.New()
	queue.PushBack(startMove)
	for queue.Len() > 0 {
		current := queue.Front()
		queue.Remove(current)
		currentMove := current.Value.(Move)

		neighbors := currentMove.coordinate.AdjacentCoordinates()
		for _, neighbor := range neighbors {
			if maze[neighbor.Y][neighbor.X] == "#" {
				continue
			}
			neighborDirection := currentMove.DirectionTowards(neighbor)
			if neighborDirection == -1 {
				continue
			}
			neighborMove := Move{coordinate: neighbor, direction: currentMove.DirectionTowards(neighbor)}
			costOfVisitingNeighbor := cost[currentMove] + scoreOfMove(currentMove, neighbor)
			neighborCost, hasVisitedNeighbor := cost[neighborMove]
			if !hasVisitedNeighbor || costOfVisitingNeighbor < neighborCost {
				cost[neighborMove] = costOfVisitingNeighbor
				cameFrom[neighborMove] = currentMove
				queue.PushBack(neighborMove)
			}
		}
	}

	// TODO: Simplify this
	possibleEndMoves := []Move{
		{coordinate: types.Coordinate{X: endPosition.X, Y: endPosition.Y}, direction: 0},
		{coordinate: types.Coordinate{X: endPosition.X, Y: endPosition.Y}, direction: 1},
		{coordinate: types.Coordinate{X: endPosition.X, Y: endPosition.Y}, direction: 2},
		{coordinate: types.Coordinate{X: endPosition.X, Y: endPosition.Y}, direction: 3},
	}

	bestEndMove := Move{}
	bestEndMoveCost := math.MaxInt
	for _, possibleEndMove := range possibleEndMoves {
		if moveCost, ok := cost[possibleEndMove]; ok {
			if moveCost < bestEndMoveCost {
				bestEndMoveCost = moveCost
				bestEndMove = possibleEndMove
			}
		}
	}

	reconstructedPath := reconstructPath(cameFrom, bestEndMove)
	slices.Reverse(reconstructedPath)

	return cost[bestEndMove], reconstructedPath
}

func dijkstra2(maze [][]string, startPosition, endPosition types.Coordinate, initialDirection int) (int, []Move) {
	cost := make(map[Move]int)
	cameFrom := make(map[Move]Move)

	startMove := Move{coordinate: startPosition, direction: initialDirection}
	cost[startMove] = 0

	queue := list.New()
	queue.PushBack(startMove)
	for queue.Len() > 0 {
		current := queue.Front()
		queue.Remove(current)
		currentMove := current.Value.(Move)

		neighborMoves := currentMove.AdjacentMoves()
		for _, neighborMove := range neighborMoves {
			if maze[neighborMove.coordinate.Y][neighborMove.coordinate.X] == "#" {
				continue
			}
			costOfVisitingNeighbor := cost[currentMove] + scoreOfMove2(currentMove, neighborMove)
			neighborCost, hasVisitedNeighbor := cost[neighborMove]
			if !hasVisitedNeighbor || costOfVisitingNeighbor < neighborCost {
				cost[neighborMove] = costOfVisitingNeighbor
				cameFrom[neighborMove] = currentMove
				queue.PushBack(neighborMove)
			}
		}
	}

	// TODO: Simplify this
	possibleEndMoves := []Move{
		{coordinate: types.Coordinate{X: endPosition.X, Y: endPosition.Y}, direction: 0},
		{coordinate: types.Coordinate{X: endPosition.X, Y: endPosition.Y}, direction: 1},
		{coordinate: types.Coordinate{X: endPosition.X, Y: endPosition.Y}, direction: 2},
		{coordinate: types.Coordinate{X: endPosition.X, Y: endPosition.Y}, direction: 3},
	}

	bestEndMove := Move{}
	bestEndMoveCost := math.MaxInt
	for _, possibleEndMove := range possibleEndMoves {
		if moveCost, ok := cost[possibleEndMove]; ok {
			if moveCost < bestEndMoveCost {
				bestEndMoveCost = moveCost
				bestEndMove = possibleEndMove
			}
		}
	}

	reconstructedPath := reconstructPath(cameFrom, bestEndMove)
	slices.Reverse(reconstructedPath)

	return cost[bestEndMove], reconstructedPath
}

func scoreOfMove(currentMove Move, nextPosition types.Coordinate) int {
	currentPosition := currentMove.coordinate
	absX := helpers.AbsInt(currentPosition.X - nextPosition.X)
	absY := helpers.AbsInt(currentPosition.Y - nextPosition.Y)
	directionChanges := currentMove.DirectionChangesTo(nextPosition)
	return absX + absY + directionChanges*1000
}

func scoreOfMove2(currentMove Move, nextMove Move) int {
	currentPosition := currentMove.coordinate
	nextPosition := nextMove.coordinate
	absX := helpers.AbsInt(currentPosition.X - nextPosition.X)
	absY := helpers.AbsInt(currentPosition.Y - nextPosition.Y)
	directionChanges := currentMove.DirectionChangesToMove(nextMove)
	return absX + absY + directionChanges*1000
}

func recalculateCostOfPath(path []Move) int {
	totalCost := 0
	for i := 0; i < len(path)-1; i++ {
		currentMove := path[i]
		nextMove := path[i+1]
		cost := scoreOfMove(currentMove, nextMove.coordinate)
		totalCost += cost
	}
	return totalCost
}

func reconstructPath(cameFrom map[Move]Move, current Move) []Move {
	totalPath := make([]Move, 0)
	totalPath = append(totalPath, current)
	for {
		if previous, ok := cameFrom[current]; ok {
			totalPath = append(totalPath, previous)
			current = previous
		} else {
			break
		}
	}

	return totalPath
}

func printMaze(maze [][]string, path []Move) {
	coordinatesMap := make(map[types.Coordinate]int)
	for _, move := range path {
		coordinatesMap[move.coordinate] = move.direction
	}
	for y, row := range maze {
		for x, value := range row {
			coordinate := types.Coordinate{X: x, Y: y}
			if direction, ok := coordinatesMap[coordinate]; ok {
				switch direction {
				case 0:
					fmt.Print("^")
				case 1:
					fmt.Print(">")
				case 2:
					fmt.Print("v")
				case 3:
					fmt.Print("<")
				}
			} else {
				fmt.Print(value)
			}
		}
		fmt.Println()
	}
}
