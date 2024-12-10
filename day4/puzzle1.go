package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle1() {
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
			sum += checkForXmas(inputs, x, y)
		}
	}

	fmt.Println(sum)
}

func checkForXmas(inputs [][]string, x int, y int) int {
	word := "XMAS"
	if inputs[y][x] != word[0:1] {
		return 0
	}

	sum := 0
	for i := 0; i < 8; i++ {
		if checkDirection(inputs, x, y, word[1:], i) {
			sum++
		}
	}

	return sum
}

func checkDirection(inputs [][]string, x int, y int, word string, direction int) bool {
	nextX, nextY := coordinatesForDirection(x, y, direction)
	if !validCoordinates(inputs, nextX, nextY) {
		return false
	}

	if inputs[nextY][nextX] != word[0:1] {
		return false
	}

	if len(word) == 1 {
		return true
	}

	return checkDirection(inputs, nextX, nextY, word[1:], direction)
}

func validCoordinates(inputs [][]string, x int, y int) bool {
	if y < 0 || y >= len(inputs) || x < 0 || x >= len(inputs[y]) {
		return false
	}

	return true
}

func coordinatesForDirection(x int, y int, direction int) (int, int) {
	switch direction {
	case 0:
		return x + 1, y
	case 1:
		return x + 1, y + 1
	case 2:
		return x, y + 1
	case 3:
		return x - 1, y + 1
	case 4:
		return x - 1, y
	case 5:
		return x - 1, y - 1
	case 6:
		return x, y - 1
	case 7:
		return x + 1, y - 1
	}

	return 0, 0
}
