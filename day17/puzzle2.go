package day17

import (
	"adventofcode2024-go/helpers"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Puzzle2() {
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

	findCopiedProgram(registerA, registerB, registerC, program)
}

func findCopiedProgram(registerA, registerB, registerC int, program []int) {
	options := make([]int, 0)
	options = append(options, 0)
	for _, output := range slices.Backward(program) {
		newOptions := make([]int, 0)
		for _, current := range options {
			for i := 0; i < 8; i++ {
				tested := (current << 3) + i
				machine := Machine{RegisterA: tested, RegisterB: registerB, RegisterC: registerC, Program: program}
				machine.Run()
				if machine.Output[0] == output {
					newOptions = append(newOptions, tested)
				}
			}
		}
		if len(newOptions) == 0 {
			fmt.Println("No solution found")
			return
		}
		options = newOptions
	}

	lowest := math.MaxInt
	for _, option := range options {
		lowest = min(lowest, option)
	}

	fmt.Println(lowest)
}
