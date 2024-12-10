package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Puzzle2() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := make(map[int][][]int)
	scanningRules := true
	sum := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		if lineText == "" {
			scanningRules = false
			continue
		}

		if scanningRules {
			pages := strings.Split(lineText, "|")
			page1, page1Err := strconv.Atoi(pages[0])
			page2, page2Err := strconv.Atoi(pages[1])
			if page1Err != nil || page2Err != nil {
				continue
			}

			rules[page1] = append(rules[page1], []int{page1, page2})
			rules[page2] = append(rules[page2], []int{page1, page2})
		} else {
			numbers := numbersFromLine(lineText)

			if !isValidUpdate(numbers, rules) {
				fixedNumbers := fixOrderOfUpdate(numbers, rules)
				sum += getMiddlePage(fixedNumbers)
			}
		}
	}

	fmt.Println(sum)
}

func indexOfNumber(numbers []int, number int) int {
	for i, n := range numbers {
		if n == number {
			return i
		}
	}

	return -1
}

func fixOrderOfUpdate(numbers []int, rules map[int][][]int) []int {
	for i := 0; i < len(numbers); i++ {
		number := numbers[i]
		rulesForNumber := rules[number]
		if len(rulesForNumber) == 0 {
			continue
		}

		for _, rule := range rulesForNumber {
			var numberToCheck int
			if rule[0] == number {
				numberToCheck = rule[1]
			} else {
				numberToCheck = rule[0]
			}

			indexOfNumberToCheck := indexOfNumber(numbers, numberToCheck)
			if indexOfNumberToCheck == -1 {
				continue
			}

			if rule[0] == number && indexOfNumberToCheck < i {
				numbers[i] = numberToCheck
				numbers[indexOfNumberToCheck] = number
				return fixOrderOfUpdate(numbers, rules)
			} else if rule[1] == number && indexOfNumberToCheck > i {
				numbers[i] = numberToCheck
				numbers[indexOfNumberToCheck] = number
				return fixOrderOfUpdate(numbers, rules)
			}
		}
	}

	return numbers
}
