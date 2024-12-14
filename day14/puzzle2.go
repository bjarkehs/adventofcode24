package day14

import (
	"adventofcode2024-go/helpers"
	"adventofcode2024-go/types"
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
)

func Puzzle2() {
	file, err := os.Open("day14/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	tester := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	robots := make([]Robot, 0)
	for scanner.Scan() {
		lineText := scanner.Text()

		match := tester.FindStringSubmatch(lineText)
		if len(match) != 5 {
			continue
		}

		coordinates := helpers.MapToInts(match[1:])

		position := types.Coordinate{X: coordinates[0], Y: coordinates[1]}
		velocity := types.Coordinate{X: coordinates[2], Y: coordinates[3]}

		robot := Robot{Position: position, Velocity: velocity}
		robots = append(robots, robot)
	}

	findChristmasTree(robots, 101, 103)
}

func printMap(robotMap map[int]map[int]int, maxX, maxY int) {
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			if robotMap[i][j] == 0 {
				fmt.Print(" ")
				continue
			}
			fmt.Print("X")
		}
		fmt.Println()
	}
}

func saveMapImage(robotMap map[int]map[int]int, maxX, maxY int, imageNumber int) {
	img := image.NewRGBA(image.Rect(0, 0, maxX, maxY))

	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			if robotMap[i][j] == 0 {
				img.Set(j, i, color.White)
				continue
			}
			img.Set(j, i, color.Black)
		}
	}

	f, _ := os.Create(fmt.Sprintf("day14/maps/map%d.png", imageNumber))
	_ = png.Encode(f, img)
	_ = f.Close()
}

func findChristmasTree(robots []Robot, maxX, maxY int) {
	seconds := 0
	for seconds < 10000 {
		seconds++
		robotsMap := moveRobotsInMap(robots, seconds, maxX, maxY)
		fmt.Println()
		fmt.Println("======================================================================================================================")
		fmt.Println("Seconds:", seconds)
		fmt.Println("======================================================================================================================")
		fmt.Println("======================================================================================================================")
		saveMapImage(robotsMap, maxX, maxY, seconds)
	}
}
