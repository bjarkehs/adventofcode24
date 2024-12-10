package day5

import (
	"adventofcode2024-go/helpers"
	"math"
	"strings"
)

func isValidUpdate(numbers []int, rules map[int][][]int) bool {
	for i := 0; i < len(numbers); i++ {
		rulesForNumber := rules[numbers[i]]
		if len(rulesForNumber) == 0 {
			continue
		}

		if !validateRules(numbers, rulesForNumber, numbers[i], i) {
			return false
		}
	}

	return true
}

func getMiddlePage(numbers []int) int {
	if len(numbers) < 1 {
		return 0
	}
	middlePageIndex := int(math.Ceil(float64(len(numbers))/2)) - 1
	return numbers[middlePageIndex]
}

func numbersFromLine(lineText string) []int {
	if lineText == "" {
		return []int{}
	}

	numberStrings := strings.Split(lineText, ",")
	return helpers.MapToInts(numberStrings)
}

func checkIsAfter(numbers []int, number int, index int) bool {
	for _, n := range numbers[index+1:] {
		if n == number {
			return true
		}
	}

	return false
}

func checkIsBefore(numbers []int, number int, index int) bool {
	for _, n := range numbers[:index] {
		if n == number {
			return true
		}
	}

	return false
}

func validateRules(numbers []int, rules [][]int, number int, index int) bool {
	validRules := true
	for _, rule := range rules {
		if rule[0] == number {
			if checkIsBefore(numbers, rule[1], index) {
				validRules = false
			}
		} else {
			if checkIsAfter(numbers, rule[0], index) {
				validRules = false
			}
		}
	}

	return validRules
}
