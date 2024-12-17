package day17

import (
	"adventofcode2024-go/helpers"
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("day17/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	registerA := 0
	registerB := 0
	registerC := 0
	program := make([]int, 0)
	for scanner.Scan() {
		lineText := scanner.Text()

		if lineText == "" {
			continue
		}

		registerRegex := regexp.MustCompile(`Register (\w): (\d+)`)
		match := registerRegex.FindStringSubmatch(lineText)
		if len(match) > 0 {
			registerValue, _ := strconv.Atoi(match[2])
			if match[1] == "A" {
				registerA = registerValue
			} else if match[1] == "B" {
				registerB = registerValue
			} else if match[1] == "C" {
				registerC = registerValue
			}
		} else {
			programIndex := strings.Index(lineText, " ")
			programAsString := lineText[programIndex+1:]
			program = helpers.MapToInts(strings.Split(programAsString, ","))
		}
	}

	machine := Machine{RegisterA: registerA, RegisterB: registerB, RegisterC: registerC, Program: program}

	machine.Run()
	machine.PrintOutput()
}
