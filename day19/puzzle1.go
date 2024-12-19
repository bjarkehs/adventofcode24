package day19

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("day19/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	towelPatterns := make([]string, 0)
	isScanningTowelPatterns := true
	designs := make([]string, 0)
	for scanner.Scan() {
		lineText := scanner.Text()

		if lineText == "" {
			isScanningTowelPatterns = false
			continue
		}

		if isScanningTowelPatterns {
			towelPatterns = append(towelPatterns, strings.Split(lineText, ", ")...)
		} else {
			designs = append(designs, lineText)
		}
	}

	fmt.Println(len(findPossibleDesigns(towelPatterns, designs)))
}

func findPossibleDesigns(towelPatterns []string, designs []string) []string {
	possibleDesigns := make([]string, 0)
	for _, design := range designs {
		if checkIfDesignIsPossible(design, towelPatterns) {
			possibleDesigns = append(possibleDesigns, design)
		}
	}

	return possibleDesigns
}

func checkIfDesignIsPossible(design string, towelPatterns []string) bool {
	if design == "" {
		return true
	}
	for _, towelPattern := range towelPatterns {
		if strings.HasPrefix(design, towelPattern) {
			newDesign := design[len(towelPattern):]
			if checkIfDesignIsPossible(newDesign, towelPatterns) {
				return true
			}
		}
	}

	return false
}
