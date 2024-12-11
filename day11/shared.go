package day11

import (
	"fmt"
	"strconv"
	"strings"
)

func blinkStonesNaive(stones []int, amountOfBlinks int) int {
	for i := 0; i < amountOfBlinks; i++ {
		stones = moveStones(stones)
		fmt.Println("Blinked {0} times", i)
	}
	return len(stones)
}

func blinkStonesMap(stones []int, amountOfBlinks int) int {
	stoneMap := generateStoneMap(stones)
	cache := make(map[int][]int)
	for i := 0; i < amountOfBlinks; i++ {
		stoneMap = moveStonesMap(stoneMap, cache)
		fmt.Println("Blinked {0} times", i)
	}

	sum := 0
	for _, value := range stoneMap {
		sum += value
	}
	return sum
}

func generateStoneMap(stones []int) map[int]int {
	stoneMap := make(map[int]int)
	for _, stone := range stones {
		stoneMap[stone]++
	}
	return stoneMap
}

func moveStonesMap(stoneMap map[int]int, stoneCache map[int][]int) map[int]int {
	newStoneMap := make(map[int]int)
	for key, value := range stoneMap {
		if value == 0 {
			continue
		}

		if values, ok := stoneCache[key]; ok {
			for _, newValue := range values {
				newStoneMap[newValue] += value
			}
		} else {
			newValues := newStonesForValue(key)
			stoneCache[key] = newValues
			for _, newValue := range newValues {
				newStoneMap[newValue] += value
			}
		}
	}

	return newStoneMap
}

func newStonesForValue(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	if check, splitStones := splitStone(stone); check {
		return splitStones
	}

	return []int{stone * 2024}
}

func moveStones(stones []int) []int {
	newStones := make([]int, 0)
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if check, splitStones := splitStone(stone); check {
			newStones = append(newStones, splitStones...)
		} else {
			newStones = append(newStones, stone*2024)
		}
	}
	return newStones
}

func splitStone(stone int) (bool, []int) {
	digits, stoneAsString := amountOfDigits(stone)
	if digits%2 != 0 {
		return false, []int{}
	}

	left := ""
	right := ""
	for i, digit := range strings.Split(stoneAsString, "") {
		if i < digits/2 {
			left += digit
		} else {
			right += digit
		}
	}

	leftNumber, _ := strconv.Atoi(left)
	rightNumber, _ := strconv.Atoi(right)

	return true, []int{leftNumber, rightNumber}
}

func amountOfDigits(value int) (int, string) {
	valueString := strconv.Itoa(value)
	return len(valueString), valueString
}
