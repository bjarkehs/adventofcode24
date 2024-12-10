package day2

import (
	"adventofcode2024-go/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeRows := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		stringValues := strings.Split(lineText, " ")
		values := helpers.MapToInts(stringValues)

		if len(values) < 2 {
			continue
		}

		isAscending := values[0] < values[1]
		previousValue := values[0]
		isSafeRow := true
		for _, v := range values[1:] {
			if isAscending != (previousValue < v) {
				isSafeRow = false
				break
			}
			difference := helpers.AbsInt(v - previousValue)
			if difference < 1 || difference > 3 {
				isSafeRow = false
				break
			}
			previousValue = v
		}

		if isSafeRow {
			fmt.Println(safeRows, lineText)
			safeRows++
		}
	}

	fmt.Println(safeRows)
}
