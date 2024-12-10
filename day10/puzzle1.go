package day10

import (
	"adventofcode2024-go/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("day10/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hikingMap [][]int
	for scanner.Scan() {
		lineText := scanner.Text()
		characters := strings.Split(lineText, "")

		numbers := helpers.MapToInts(characters)

		hikingMap = append(hikingMap, numbers)
	}

	fmt.Println(findTrailHeadScores(hikingMap))
}

func findTrailHeadScores(hikingMap [][]int) int {
	sum := 0
	for y, row := range hikingMap {
		for x, cell := range row {
			if cell == 0 {
				coordinates := trailHeadScore(hikingMap, Coordinate{x, y}, 0)
				coordinateMap := make(map[Coordinate]bool)
				for _, coordinate := range coordinates {
					coordinateMap[coordinate] = true
				}
				sum += len(coordinateMap)
			}
		}
	}

	return sum
}

func trailHeadScore(hikingMap [][]int, coordinate Coordinate, height int) []Coordinate {
	if height == 9 {
		return []Coordinate{coordinate}
	}

	coordinates := make([]Coordinate, 0)
	for direction := 0; direction < 4; direction++ {
		newCoordinate := coordinateForDirection(coordinate, direction)
		if !validCoordinate(newCoordinate, len(hikingMap[0]), len(hikingMap)) {
			continue
		}

		if hikingMap[newCoordinate.y][newCoordinate.x] == height+1 {
			coordinates = append(coordinates, trailHeadScore(hikingMap, newCoordinate, height+1)...)
		}
	}

	return coordinates
}
