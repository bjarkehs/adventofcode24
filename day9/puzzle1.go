package day9

import (
	"adventofcode2024-go/helpers"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int
	for scanner.Scan() {
		lineText := scanner.Text()
		characters := strings.Split(lineText, "")
		numbers = helpers.MapToInts(characters)
	}

	diskSize := 0
	for _, number := range numbers {
		diskSize += number
	}

	disk := make([]string, diskSize)
	diskLocation := 0
	fileId := 0
	totalFileSize := 0
	for index, number := range numbers {
		var blockId string
		if index%2 == 0 {
			blockId = strconv.Itoa(fileId)
			totalFileSize += number
			fileId++
		} else {
			blockId = "."
		}
		for i := 0; i < number; i++ {
			disk[diskLocation+i] = blockId
		}
		diskLocation += number
	}

	moveBlocks1(disk, totalFileSize)
}

func moveBlocks1(disk []string, totalFileSize int) {
	for i := 0; i < len(disk); i++ {
		block := disk[i]
		if block != "." {
			continue
		}

		for j := len(disk) - 1; j > 0; j-- {
			if j == i {
				break
			}

			if disk[j] == "." {
				continue
			}

			disk[i] = disk[j]
			disk[j] = "."
			break
		}
	}

	checksum := 0
	for i := 0; i < totalFileSize; i++ {
		number, _ := strconv.Atoi(disk[i])
		checksum += number * i
	}

	fmt.Println(checksum)
}
