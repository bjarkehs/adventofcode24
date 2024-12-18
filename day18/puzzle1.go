package day18

import (
	"adventofcode2024-go/types"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("day18/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fallingBytes := make([]types.Coordinate, 0)
	for scanner.Scan() {
		lineText := scanner.Text()

		fallingByteAsString := strings.Split(lineText, ",")
		if len(fallingByteAsString) != 2 {
			panic("Invalid input")
		}
		xPosition, _ := strconv.Atoi(fallingByteAsString[0])
		yPosition, _ := strconv.Atoi(fallingByteAsString[1])
		fallingBytes = append(fallingBytes, types.Coordinate{X: xPosition, Y: yPosition})
	}

	maze, startPosition, endPosition := createMaze(fallingBytes, 70, 70, 1024)
	cost := aStar(maze, startPosition, endPosition)
	fmt.Println(cost)
}
