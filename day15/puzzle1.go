package day15

import (
	"adventofcode2024-go/types"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("day15/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	area := make([][]string, 0)
	moves := make([]int, 0)
	isParsingArea := true
	foundStartPosition := false
	startPosition := types.Coordinate{X: 0, Y: 0}
	for scanner.Scan() {
		lineText := scanner.Text()

		if lineText == "" {
			isParsingArea = false
			continue
		}

		if !foundStartPosition {
			indexOfRobot := strings.Index(lineText, "@")
			if indexOfRobot != -1 {
				startPosition = types.Coordinate{X: indexOfRobot, Y: len(area)}
				foundStartPosition = true
			}
		}

		characters := strings.Split(lineText, "")
		if isParsingArea {
			area = append(area, characters)
		} else {
			for _, character := range characters {
				moves = append(moves, findMove(character))
			}
		}
	}

	movedArea := moveRobot(area, moves, startPosition)
	coordinates := objectCoordinates(movedArea)
	sum := 0
	for _, coordinate := range coordinates {
		sum += 100*coordinate.Y + coordinate.X
	}

	fmt.Println(sum)
}

func moveToPosition(area [][]string, position types.Coordinate, moving string, direction int) bool {
	if area[position.Y][position.X] == "#" {
		return false
	}

	if area[position.Y][position.X] == "." {
		area[position.Y][position.X] = moving
		return true
	}

	if moveToPosition(area, position.CoordinateForDirection(direction), area[position.Y][position.X], direction) {
		area[position.Y][position.X] = moving
		return true
	}

	return false
}

func moveRobot(area [][]string, moves []int, startPosition types.Coordinate) [][]string {
	currentPosition := startPosition
	for _, move := range moves {
		potentialNewCoordinate := currentPosition.CoordinateForDirection(move)
		if moveToPosition(area, potentialNewCoordinate, "@", move) {
			area[currentPosition.Y][currentPosition.X] = "."
			currentPosition = potentialNewCoordinate
		}
	}

	return area
}

func objectCoordinates(area [][]string) []types.Coordinate {
	coordinates := make([]types.Coordinate, 0)
	for y, row := range area {
		for x, value := range row {
			if value == "O" {
				coordinates = append(coordinates, types.Coordinate{X: x, Y: y})
			}
		}
	}

	return coordinates
}
