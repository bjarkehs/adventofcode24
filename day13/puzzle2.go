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
			tokens += findTokensEquations(currentGame)
			currentGame = Game{}
			continue
		}
	}

	fmt.Println(tokens)
}

func findTokensEquations(game Game) int {
	a1 := float64(game.moveA.coordinate.X)
	a2 := float64(game.moveA.coordinate.Y)
	b1 := float64(game.moveB.coordinate.X)
	b2 := float64(game.moveB.coordinate.Y)
	p1 := float64(game.prizePosition.X)
	p2 := float64(game.prizePosition.Y)

	determinant := a1*b2 - a2*b1
	x := (p1*b2 - p2*b1) / determinant
	y := (a1*p2 - a2*p1) / determinant

	if x >= 0 && y >= 0 && isInteger(x) && isInteger(y) {
		return game.moveA.cost*int(x) + game.moveB.cost*int(y)
	}

	return 0
}

func isInteger(f float64) bool {
	return f == math.Floor(f)
}
