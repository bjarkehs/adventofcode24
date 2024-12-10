package day2

import (
	"adventofcode2024-go/helpers"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Puzzle2() {
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
		values := helpers.Map(stringValues, func(t string) int {
			value, err := strconv.Atoi(t)
			helpers.Check(err)
			return value
		})

		if len(values) < 2 {
			continue
		}

		if areValuesSafe(values, 0) {
			fmt.Println(safeRows, lineText)
			safeRows++
		}
	}

	fmt.Println(safeRows)
}

func areValuesSafe(values []int, problemsDetected int) bool {
	fmt.Println("Checking", values, problemsDetected)
	isAscending := values[0] < values[1]
	previousValue := values[0]
	for valueIndex, v := range values {
		if valueIndex == 0 {
			continue
		}
		if isAscending != (previousValue < v) {
			return areSubValuesSafe(values, problemsDetected, valueIndex)
		}
		difference := helpers.AbsInt(v - previousValue)
		if difference < 1 || difference > 3 {
			return areSubValuesSafe(values, problemsDetected, valueIndex)
		}
		previousValue = v
	}

	return true
}

func areSubValuesSafe(values []int, problemsDetected int, valueIndex int) bool {
	if problemsDetected > 0 {
		return false
	}

	valuesWithoutPreviousValue := helpers.RemoveIndex(values, valueIndex-1)
	if areValuesSafe(valuesWithoutPreviousValue, problemsDetected+1) {
		fmt.Println("Safe")
		return true
	}
	valuesWithoutCurrentValue := helpers.RemoveIndex(values, valueIndex)
	if areValuesSafe(valuesWithoutCurrentValue, problemsDetected+1) {
		fmt.Println("Safe")
		return true
	}

	if valueIndex > 1 {
		valuesWithout2ndPreviousValue := helpers.RemoveIndex(values, valueIndex-2)
		if areValuesSafe(valuesWithout2ndPreviousValue, problemsDetected+1) {
			fmt.Println("Safe")
			return true
		}
	}
	fmt.Println("Still unsafe")
	return false
}
