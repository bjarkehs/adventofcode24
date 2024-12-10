package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle2() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var inputs [][]string
	for scanner.Scan() {
		lineText := scanner.Text()
		inputs = append(inputs, strings.Split(lineText, ""))
	}

	sum := 0
	for y := 0; y < len(inputs); y++ {
		for x := 0; x < len(inputs[y]); x++ {
			if checkForCrossMas(inputs, x, y) {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func checkForCrossMas(inputs [][]string, x int, y int) bool {
	if inputs[y][x] != "A" {
		return false
	}

	return checkAngle(inputs, x, y, true) && checkAngle(inputs, x, y, false)
}

func checkAngle(inputs [][]string, x int, y int, leftToRight bool) bool {
	var firstX int
	var firstY int
	var secondX int
	var secondY int

	if leftToRight {
		firstX = x - 1
		firstY = y - 1
		secondX = x + 1
		secondY = y + 1
	} else {
		firstX = x + 1
		firstY = y - 1
		secondX = x - 1
		secondY = y + 1
	}

	if !validCoordinates(inputs, firstX, firstY) || !validCoordinates(inputs, secondX, secondY) {
		return false
	}

	if inputs[firstY][firstX] == "M" && inputs[secondY][secondX] == "S" {
		return true
	}

	if inputs[firstY][firstX] == "S" && inputs[secondY][secondX] == "M" {
		return true
	}

	return false
}
