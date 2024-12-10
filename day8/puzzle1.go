package day8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	antennaLocations := make(map[string][]Coordinate)
	y := 0
	maxX := 0
	for scanner.Scan() {
		lineText := scanner.Text()

		characters := strings.Split(lineText, "")
		maxX = len(characters)
		for x, char := range characters {
			if char == "." {
				continue
			}
			antennaLocations[char] = append(antennaLocations[char], Coordinate{x, y})
		}
		y++
	}

	uniqueCoordinates := antinodeCoordinates(antennaLocations, maxX, y)
	fmt.Println(uniqueCoordinates)
}

func antinodeCoordinates(antennaLocations map[string][]Coordinate, maxX int, maxY int) int {
	antinodes := make(map[Coordinate]bool)
	uniqueCoordinates := 0
	for _, coordinates := range antennaLocations {
		for _, coordinate := range coordinates {
			for _, otherCoordinate := range coordinates {
				if coordinate == otherCoordinate {
					continue
				}

				antinodeCoordinate := antinodeCoordinateFromTwoCoordinates(coordinate, otherCoordinate)
				if !validCoordinate(antinodeCoordinate, maxX, maxY) {
					continue
				}
				if _, ok := antinodes[antinodeCoordinate]; !ok {
					antinodes[antinodeCoordinate] = true
					uniqueCoordinates++
				}
			}
		}
	}

	return uniqueCoordinates
}

func antinodeCoordinateFromTwoCoordinates(first Coordinate, second Coordinate) Coordinate {
	xDiff := second.x - first.x
	yDiff := second.y - first.y
	return Coordinate{first.x - xDiff, first.y - yDiff}
}
