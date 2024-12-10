package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Puzzle1() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		tester := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

		matches := tester.FindAllStringSubmatch(lineText, -1)

		if len(matches) == 0 {
			continue
		}

		for _, match := range matches {
			if len(match) != 3 {
				continue
			}

			value1, value1Err := strconv.Atoi(match[1])
			value2, value2Err := strconv.Atoi(match[2])

			if value1Err != nil || value2Err != nil {
				continue
			}

			sum += value1 * value2
		}
	}

	fmt.Println(sum)
}
