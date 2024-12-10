package day10

import (
	"adventofcode2024-go/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle2() {
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

	fmt.Println(findTrailRatings(hikingMap))
}

func findTrailRatings(hikingMap [][]int) int {
	sum := 0
	for y, row := range hikingMap {
		for x, cell := range row {
			if cell == 0 {
				sum += trailHeadRating(hikingMap, Coordinate{x, y}, 0)
			}
		}
	}

	return sum
}

func trailHeadRating(hikingMap [][]int, coordinate Coordinate, height int) int {
	if height == 9 {
		return 1
	}

	sum := 0
	for direction := 0; direction < 4; direction++ {
		newCoordinate := coordinateForDirection(coordinate, direction)
		if !validCoordinate(newCoordinate, len(hikingMap[0]), len(hikingMap)) {
			continue
		}

		if hikingMap[newCoordinate.y][newCoordinate.x] == height+1 {
			sum += trailHeadRating(hikingMap, newCoordinate, height+1)
		}
	}

	return sum
}
