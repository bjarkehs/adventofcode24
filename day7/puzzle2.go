package day7

import (
	"adventofcode2024-go/helpers"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Puzzle2() {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		lineText := scanner.Text()

		equationParts := strings.Split(lineText, ":")
		if len(equationParts) != 2 {
			continue
		}

		result, resultErr := strconv.Atoi(equationParts[0])
		helpers.Check(resultErr)

		numberStrings := strings.Split(equationParts[1], " ")
		numbers := make([]int, 0)
		for _, numberString := range numberStrings {
			if numberString == "" {
				continue
			}
			number, numberErr := strconv.Atoi(numberString)
			helpers.Check(numberErr)
			numbers = append(numbers, number)
		}

		if isValidPuzzle2(result, numbers, 0) {
			sum += result
		}
	}

	fmt.Println(sum)
}

func isValidPuzzle2(result int, numbers []int, acc int) bool {
	if len(numbers) == 0 {
		return result == acc
	}

	result1 := acc + numbers[0]
	result2 := acc * numbers[0]
	result3String := strconv.Itoa(acc) + strconv.Itoa(numbers[0])
	result3, result3Err := strconv.Atoi(result3String)

	return isValidPuzzle2(result, numbers[1:], result1) || isValidPuzzle2(result, numbers[1:], result2) || (result3Err == nil && isValidPuzzle2(result, numbers[1:], result3))
}
