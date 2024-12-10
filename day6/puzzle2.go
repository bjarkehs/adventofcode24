package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle2() {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var obstacleMap [][]bool
	var guardStartPosition Coordinate
	var direction int

	maxY := 0
	maxX := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		characters := strings.Split(lineText, "")

		obstacleMap = append(obstacleMap, make([]bool, len(characters)))

		if maxX == 0 {
			maxX = len(characters)
		}

		for x, char := range characters {
			if char == "." {
				continue
			}

			if char == "#" {
				obstacleMap[maxY][x] = true
				continue
			}

			if char == "^" {
				guardStartPosition = Coordinate{x, maxY}
				direction = 0
				continue
			}
		}

		maxY++
	}

	possibleNewObstacles := 0
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if checkNewObstacle(guardStartPosition, direction, obstacleMap, maxX, maxY, Coordinate{x, y}) {
				possibleNewObstacles++
			}
		}
	}

	fmt.Println(possibleNewObstacles)
}

func checkNewObstacle(guardStartPosition Coordinate, direction int, obstacleMap [][]bool, maxX int, maxY int, newCoordinate Coordinate) bool {
	if obstacleMap[newCoordinate.y][newCoordinate.x] || (guardStartPosition.x == newCoordinate.x && guardStartPosition.y == newCoordinate.y) {
		return false
	}

	positions := make(map[Coordinate]int)
	isInsideMap := true
	tweakedObstacleMap := make([][]bool, len(obstacleMap))
	for i, row := range obstacleMap {
		tweakedObstacleMap[i] = make([]bool, len(row))
		copy(tweakedObstacleMap[i], row)
	}
	tweakedObstacleMap[newCoordinate.y][newCoordinate.x] = true

	for isInsideMap {
		isInsideMap, guardStartPosition, direction = moveGuardV2(guardStartPosition, direction, tweakedObstacleMap, maxX, maxY)
		if isInsideMap {
			if positions[guardStartPosition]&(0x1<<direction) != 0 {
				// We have already walked here in this direction
				return true
			}
			positions[guardStartPosition] = positions[guardStartPosition] | 0x1<<direction
		}
	}
	return false
}
