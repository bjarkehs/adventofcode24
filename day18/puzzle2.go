package day18

import (
	"adventofcode2024-go/types"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Puzzle2() {
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

	blockCoordinate := checkBlockingBinarySearch(fallingBytes, 70, 70)
	fmt.Println(fmt.Sprintf("%d,%d", blockCoordinate.X, blockCoordinate.Y))
}

func checkBlockingIteration(fallingBytes []types.Coordinate, maxX, maxY int, lowestAmountOfBytes int) *types.Coordinate {
	for amountOfBytes := lowestAmountOfBytes; amountOfBytes < len(fallingBytes); amountOfBytes++ {
		maze, startPosition, endPosition := createMaze(fallingBytes, maxX, maxY, amountOfBytes+1)
		cost := aStar(maze, startPosition, endPosition)
		if cost == -1 {
			problematicFall := fallingBytes[amountOfBytes]
			return &types.Coordinate{X: problematicFall.X, Y: problematicFall.Y}
		}
	}
	return nil
}

func checkBlockingBinarySearch(fallingBytes []types.Coordinate, maxX, maxY int) *types.Coordinate {
	lowestAmountOfBytes := 0
	highestAmountOfBytes := len(fallingBytes)
	for lowestAmountOfBytes < highestAmountOfBytes {
		mid := (lowestAmountOfBytes + highestAmountOfBytes) / 2
		fmt.Println(mid)
		maze, startPosition, endPosition := createMaze(fallingBytes, maxX, maxY, mid+1)
		cost := aStar(maze, startPosition, endPosition)
		if cost == -1 {
			if highestAmountOfBytes-lowestAmountOfBytes == 1 {
				problematicFall := fallingBytes[mid]
				return &types.Coordinate{X: problematicFall.X, Y: problematicFall.Y}
			}
			highestAmountOfBytes = mid
		} else {
			lowestAmountOfBytes = mid + 1
		}
	}

	return nil
}
