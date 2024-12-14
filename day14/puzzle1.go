package day14

import (
	"adventofcode2024-go/helpers"
	"adventofcode2024-go/types"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Puzzle1() {
	file, err := os.Open("day14/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	tester := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	robots := make([]Robot, 0)
	for scanner.Scan() {
		lineText := scanner.Text()

		match := tester.FindStringSubmatch(lineText)
		if len(match) != 5 {
			continue
		}

		coordinates := helpers.MapToInts(match[1:])

		position := types.Coordinate{X: coordinates[0], Y: coordinates[1]}
		velocity := types.Coordinate{X: coordinates[2], Y: coordinates[3]}

		robot := Robot{Position: position, Velocity: velocity}
		robots = append(robots, robot)
	}

	fmt.Println(calculateSafetyFactor(robots, 100, 101, 103))
}

func quadrant(position types.Coordinate, middleX, middleY int) (int, bool) {
	if position.X < middleX && position.Y < middleY {
		return 0, true
	}
	if position.X > middleX && position.Y < middleY {
		return 1, true
	}
	if position.X > middleX && position.Y > middleY {
		return 2, true
	}
	if position.X < middleX && position.Y > middleY {
		return 3, true
	}

	return -1, false
}

func calculateSafetyFactor(robots []Robot, moves int, maxX, maxY int) int {
	newCoordinates := moveRobots(robots, moves, maxX, maxY)
	middleX := maxX / 2
	middleY := maxY / 2
	quadrants := make(map[int]int)
	for _, coordinate := range newCoordinates {
		quadrant, shouldCount := quadrant(coordinate, middleX, middleY)
		if !shouldCount {
			continue
		}
		quadrants[quadrant]++
	}

	product := 1
	for _, value := range quadrants {
		product *= value
	}
	return product
}
