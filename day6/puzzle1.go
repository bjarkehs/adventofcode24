package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var obstacles []Coordinate
	var guardStartPosition Coordinate
	var direction int

	y := 0
	maxX := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		characters := strings.Split(lineText, "")

		if maxX == 0 {
			maxX = len(characters)
		}

		for x, char := range characters {
			if char == "." {
				continue
			}

			if char == "#" {
				obstacles = append(obstacles, Coordinate{x, y})
				continue
			}

			if char == "^" {
				guardStartPosition = Coordinate{x, y}
				direction = 0
				continue
			}
		}

		y++
	}

	positions := make(map[Coordinate]bool)
	positions[guardStartPosition] = true
	isInsideMap := true
	for isInsideMap {
		isInsideMap, guardStartPosition, direction = moveGuard(guardStartPosition, direction, obstacles, maxX, y)
		if isInsideMap {
			positions[guardStartPosition] = true
		}
	}
	fmt.Println(len(positions))
}
