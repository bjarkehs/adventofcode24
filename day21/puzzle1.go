package day21

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("day21/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	directionalKeypad := NewKeypad([][]string{
		{"X", "^", "A"},
		{"<", "v", ">"},
	})

	numericKeypad := NewKeypad([][]string{
		{"7", "8", "9"},
		{"4", "5", "6"},
		{"1", "2", "3"},
		{"X", "0", "A"},
	})

	sum := 0
	cache := make(map[string]int)
	for scanner.Scan() {
		lineText := scanner.Text()

		keyPresses := getKeyPresses(numericKeypad, directionalKeypad, lineText, 2, cache)
		numericPart := strings.ReplaceAll(lineText, "A", "")
		value, _ := strconv.Atoi(numericPart)
		sum += keyPresses * value
	}

	fmt.Println(sum)
}
