package day13

import (
	"adventofcode2024-go/types"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func Puzzle2() {
	file, err := os.Open("day13/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	tokens := 0
	currentGame := Game{}

	buttonTester := regexp.MustCompile(`Button (.+): X\+(\d+), Y\+(\d+)`)
	prizeTester := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
	for scanner.Scan() {
		lineText := scanner.Text()
		if lineText == " " {
			continue
		}

		buttonMatch := buttonTester.FindStringSubmatch(lineText)
		if len(buttonMatch) > 0 {
			xCoordinate, _ := strconv.Atoi(buttonMatch[2])
			yCoordinate, _ := strconv.Atoi(buttonMatch[3])
			coordinate := types.Coordinate{X: xCoordinate, Y: yCoordinate}
			if buttonMatch[1] == "A" {
				currentGame.moveA = Move{coordinate: coordinate, cost: 3}
			}
			if buttonMatch[1] == "B" {
				currentGame.moveB = Move{coordinate: coordinate, cost: 1}
			}
			continue
		}

		prizeMatch := prizeTester.FindStringSubmatch(lineText)
		if len(prizeMatch) > 0 {
			xCoordinate, _ := strconv.Atoi(prizeMatch[1])
			yCoordinate, _ := strconv.Atoi(prizeMatch[2])
			currentGame.prizePosition = types.Coordinate{X: xCoordinate + 10000000000000, Y: yCoordinate + 10000000000000}
			tokens += currentGame.solve()
			currentGame = Game{}
			continue
		}
	}

	fmt.Println(tokens)
}

func isInteger(f float64) bool {
	return f == math.Floor(f)
}
