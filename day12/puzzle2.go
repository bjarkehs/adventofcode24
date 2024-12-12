package day12

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle2() {
	file, err := os.Open("day12/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	areaMap := make([][]string, 0)
	for scanner.Scan() {
		lineText := scanner.Text()

		characters := strings.Split(lineText, "")
		areaMap = append(areaMap, characters)
	}

	regions := findRegionsInMap(areaMap)
	price := 0
	for _, region := range regions {
		price += len(region.coordinates) * region.amountOfSides()
	}

	fmt.Println(price)
}
