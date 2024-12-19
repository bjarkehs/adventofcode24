package day19

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle2() {
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

	fmt.Println(findAllDesignVariations(towelPatterns, designs))
}

func findAllDesignVariations(towelPatterns []string, designs []string) int {
	variations := 0
	for _, design := range designs {
		designVariationsCache := make(map[string]int)
		if checkIfDesignIsPossible(design, towelPatterns) {
			variations += variationsForDesign(design, towelPatterns, designVariationsCache)
		}
	}

	return variations
}

func variationsForDesign(design string, towelPatterns []string, designVariationsCache map[string]int) int {
	if design == "" {
		return 1
	}

	if variations, ok := designVariationsCache[design]; ok {
		return variations
	}

	variations := 0
	for _, towelPattern := range towelPatterns {
		if strings.HasPrefix(design, towelPattern) {
			newDesign := design[len(towelPattern):]

			variations += variationsForDesign(newDesign, towelPatterns, designVariationsCache)
		}
	}

	designVariationsCache[design] = variations
	return variations
}
