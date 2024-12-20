package day20

import (
	"adventofcode2024-go/types"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle2() {
	file, err := os.Open("day20/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	raceTrack := make([][]string, 0)
	startPosition := types.Coordinate{X: 0, Y: 0}
	startPositionFound := false
	endPosition := types.Coordinate{X: 0, Y: 0}
	endPositionFound := false
	for scanner.Scan() {
		lineText := scanner.Text()
		raceTrack = append(raceTrack, strings.Split(lineText, ""))

		if !startPositionFound {
			if indexOfStart := strings.Index(lineText, "S"); indexOfStart != -1 {
				startPosition = types.Coordinate{X: indexOfStart, Y: len(raceTrack) - 1}
				startPositionFound = true
			}
		}

		if !endPositionFound {
			if indexOfEnd := strings.Index(lineText, "E"); indexOfEnd != -1 {
				endPosition = types.Coordinate{X: indexOfEnd, Y: len(raceTrack) - 1}
				endPositionFound = true
			}
		}
	}

	result := SolveMaze(raceTrack, startPosition, endPosition, 20, 100)
	fmt.Println(result)
}
