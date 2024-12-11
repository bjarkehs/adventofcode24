package day11

import (
	"adventofcode2024-go/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("day11/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	stones := make([]int, 0)
	for scanner.Scan() {
		lineText := scanner.Text()

		characters := strings.Split(lineText, " ")
		stones = helpers.MapToInts(characters)
	}

	amountOfStones := blinkStonesMap(stones, 25)

	fmt.Println(amountOfStones)
}
