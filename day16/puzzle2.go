package day16

import (
	"adventofcode2024-go/types"
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"strings"
)

func Puzzle2() {
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

	_, path := dijkstra3(maze, startPosition, endPosition, 1)
	fmt.Println(len(path))
}

func dijkstra3(maze [][]string, startPosition, endPosition types.Coordinate, initialDirection int) (int, []types.Coordinate) {
	cost := make(map[Move]int)
	cameFrom := make(map[Move][]Move)

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
				cameFrom[neighborMove] = []Move{currentMove}
				queue.PushBack(neighborMove)
			} else if costOfVisitingNeighbor == neighborCost {
				cameFrom[neighborMove] = append(cameFrom[neighborMove], currentMove)
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

	allPossiblePositions := possiblePositions(cameFrom, bestEndMove)

	return cost[bestEndMove], allPossiblePositions
}

func possiblePositions(cameFrom map[Move][]Move, current Move) []types.Coordinate {
	possiblePositionsMap := make(map[types.Coordinate]bool)
	visited := make(map[Move]bool)

	queue := list.New()
	queue.PushBack(current)
	for queue.Len() > 0 {
		next := queue.Front()
		queue.Remove(next)
		nextMove := next.Value.(Move)
		possiblePositionsMap[nextMove.coordinate] = true
		previousMoves, ok := cameFrom[nextMove]
		if !ok {
			continue
		}
		for _, move := range previousMoves {
			if _, ok := visited[move]; !ok {
				queue.PushBack(move)
				visited[move] = true
			}
		}
	}

	allPossiblePositions := make([]types.Coordinate, 0)
	for possiblePosition := range possiblePositionsMap {
		allPossiblePositions = append(allPossiblePositions, possiblePosition)
	}

	return allPossiblePositions
}
