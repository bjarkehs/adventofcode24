package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Puzzle1() {
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

			if isValidUpdate(numbers, rules) {
				sum += getMiddlePage(numbers)
			}
		}
	}

	fmt.Println(sum)
}
