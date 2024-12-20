package day15

import (
	"adventofcode2024-go/types"
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func Puzzle2() {
	file, err := os.Open("day15/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	area := make([][]string, 0)
	moves := make([]int, 0)
	isParsingArea := true
	for scanner.Scan() {
		lineText := scanner.Text()

		if lineText == "" {
			isParsingArea = false
			continue
		}

		characters := strings.Split(lineText, "")
		if isParsingArea {
			mappedCharacters := make([]string, len(lineText)*2)
			for index, character := range characters {
				if character == "@" {
					mappedCharacters[index*2] = character
					mappedCharacters[index*2+1] = "."
				} else if character == "O" {
					mappedCharacters[index*2] = "["
					mappedCharacters[index*2+1] = "]"
				} else {
					mappedCharacters[index*2] = character
					mappedCharacters[index*2+1] = character
				}
			}
			area = append(area, mappedCharacters)
		} else {
			for _, character := range characters {
				moves = append(moves, findMove(character))
			}
		}
	}

	startPosition := robotStartPosition(area)
	movedArea := moveRobot2(area, moves, startPosition)
	coordinates := objectCoordinates2(movedArea)
	sum := 0
	for _, coordinate := range coordinates {
		sum += 100*coordinate.Y + coordinate.X
	}

	fmt.Println(sum)
}

func moveRobot2(area [][]string, moves []int, startPosition types.Coordinate) [][]string {
	currentPosition := startPosition
	for _, move := range moves {
		potentialNewCoordinate := currentPosition.CoordinateForDirection(move)
		if canMove, renders := canMoveToCoordinate(area, potentialNewCoordinate, move); canMove {
			for coordinate, render := range renders {
				area[coordinate.Y][coordinate.X] = render
			}
			area[currentPosition.Y][currentPosition.X] = "."
			currentPosition = potentialNewCoordinate
		}
	}

	return area
}

func printArea(area [][]string) {
	fmt.Println()
	for _, row := range area {
		fmt.Println(strings.Join(row, ""))
	}
}

func canMoveToCoordinate(area [][]string, potentialNewCoordinate types.Coordinate, move int) (bool, map[types.Coordinate]string) {
	visited := make(map[types.Coordinate]bool)
	renders := make(map[types.Coordinate]string)

	queue := list.New()
	queue.PushBack(potentialNewCoordinate)
	renders[potentialNewCoordinate] = "@"

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		coordinate := element.Value.(types.Coordinate)

		if _, ok := visited[coordinate]; ok {
			continue
		}

		if area[coordinate.Y][coordinate.X] == "#" {
			return false, renders
		}

		if area[coordinate.Y][coordinate.X] == "." {
			visited[coordinate] = true
			continue
		}

		objectPart := area[coordinate.Y][coordinate.X]
		visitCoordinate := func(coordinate types.Coordinate, objectPart string) {
			visited[coordinate] = true
			nextCoordinate := coordinate.CoordinateForDirection(move)
			renders[nextCoordinate] = objectPart
			queue.PushBack(nextCoordinate)
			emptyCoordinate(renders, coordinate)
		}
		if move%2 == 0 {
			// Moving vertically
			visitCoordinate(coordinate, objectPart)
			offset := objectPartCoordinateOffset(objectPart)
			otherCoordinate := types.Coordinate{X: coordinate.X + offset, Y: coordinate.Y}
			visitCoordinate(otherCoordinate, otherObjectPart(objectPart))
		} else {
			// Moving horizontally
			visitCoordinate(coordinate, objectPart)
		}
	}

	return true, renders
}

func emptyCoordinate(renders map[types.Coordinate]string, coordinate types.Coordinate) {
	if _, ok := renders[coordinate]; !ok {
		renders[coordinate] = "."
	}
}

func objectCoordinates2(area [][]string) []types.Coordinate {
	coordinates := make([]types.Coordinate, 0)
	for y, row := range area {
		for x, value := range row {
			if value == "[" {
				coordinates = append(coordinates, types.Coordinate{X: x, Y: y})
			}
		}
	}

	return coordinates
}

func robotStartPosition(area [][]string) types.Coordinate {
	for y, row := range area {
		for x, value := range row {
			if value == "@" {
				return types.Coordinate{X: x, Y: y}
			}
		}
	}

	return types.Coordinate{X: 0, Y: 0}
}

func objectPartCoordinateOffset(objectPart string) int {
	if objectPart == "[" {
		return 1
	}

	return -1
}

func otherObjectPart(objectPart string) string {
	if objectPart == "[" {
		return "]"
	}

	return "["
}
