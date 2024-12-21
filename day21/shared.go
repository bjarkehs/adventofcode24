package day21

import (
	"adventofcode2024-go/types"
	"container/list"
	"math"
	"sort"
	"strconv"
)

type Keypad map[string]types.Coordinate

func NewKeypad(input [][]string) Keypad {
	keypad := make(Keypad)
	for y, row := range input {
		for x, value := range row {
			keypad[value] = types.Coordinate{X: x, Y: y}
		}
	}

	return keypad
}

type Command struct {
	coordinate types.Coordinate
	path       string
}

func commandForDirection(direction int) string {
	switch direction {
	case 0:
		return "^"
	case 1:
		return ">"
	case 2:
		return "v"
	case 3:
		return "<"
	}

	return ""
}

func getCommand(keypad Keypad, from string, to string) []string {
	queue := list.New()
	queue.PushBack(Command{coordinate: keypad[from], path: ""})

	invalidCoordinate := keypad["X"]

	distances := make(map[types.Coordinate]int)
	distances[keypad[from]] = 0

	allPaths := make([]string, 0)

	if from == to {
		return []string{"A"}
	}

	for queue.Len() > 0 {
		current := queue.Front()
		currentCommand := current.Value.(Command)
		queue.Remove(current)

		if currentCommand.coordinate == keypad[to] {
			allPaths = append(allPaths, currentCommand.path+"A")
		}

		if distances[currentCommand.coordinate] != 0 && distances[currentCommand.coordinate] < len(currentCommand.path) {
			continue
		}

		neighbors := currentCommand.coordinate.AdjacentCoordinates()
		for direction, neighbor := range neighbors {
			if neighbor == invalidCoordinate {
				continue
			}

			for _, position := range keypad {
				if neighbor != position {
					continue
				}

				newPath := currentCommand.path + commandForDirection(direction)
				existingDistance, hasDistance := distances[neighbor]
				if !hasDistance || len(newPath) <= existingDistance {
					distances[neighbor] = len(newPath)
					queue.PushBack(Command{coordinate: neighbor, path: newPath})
				}
			}
		}
	}

	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	return allPaths
}

func getKeyPresses(input Keypad, directional Keypad, code string, robot int, cache map[string]int) int {
	key := code + strconv.Itoa(robot)
	if val, exists := cache[key]; exists {
		return val
	}

	current := "A"
	length := 0
	for i := 0; i < len(code); i++ {
		moves := getCommand(input, current, string(code[i]))
		if robot == 0 {
			length += len(moves[0])
		} else {
			minLength := math.MaxInt
			for _, move := range moves {
				keyPressesForMove := getKeyPresses(directional, directional, move, robot-1, cache)
				if keyPressesForMove < minLength {
					minLength = keyPressesForMove
				}
			}
			length += minLength
		}

		current = string(code[i])
	}

	cache[key] = length
	return length
}
