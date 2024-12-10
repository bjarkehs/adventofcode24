package day9

import (
	"adventofcode2024-go/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Block struct {
	size   int
	fileId int
}

func Puzzle2() {
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

	blocks := make([]Block, len(numbers))
	fileId := 0
	for i, number := range numbers {
		if i%2 == 0 {
			blocks[i] = Block{size: number, fileId: fileId}
			fileId++
		} else {
			blocks[i] = Block{size: number, fileId: -1}
		}
	}

	movedBlocks := moveBlocks2(blocks)
	fmt.Println(calculateChecksum(movedBlocks))
}

func moveBlocks2(blocks []Block) []Block {
	for i := len(blocks) - 1; i > 0; i-- {
		block := blocks[i]
		if block.fileId == -1 {
			continue
		}

		for j := 0; j < len(blocks); j++ {
			if j >= i {
				break
			}
			checkedBlock := blocks[j]
			if checkedBlock.fileId != -1 {
				continue
			}

			if checkedBlock.size >= block.size {
				returnedBlocks := make([]Block, 0)
				returnedBlocks = append(returnedBlocks, blocks[:j]...)
				returnedBlocks = append(returnedBlocks, block)
				if checkedBlock.size-block.size > 0 {
					returnedBlocks = append(returnedBlocks, Block{size: checkedBlock.size - block.size, fileId: -1})
				}
				returnedBlocks = append(returnedBlocks, blocks[j+1:i]...)
				if i+1 < len(blocks) {
					returnedBlocks = append(returnedBlocks, Block{size: block.size, fileId: -1})
					returnedBlocks = append(returnedBlocks, blocks[i+1:]...)
				}

				return moveBlocks2(returnedBlocks)
			}
		}
	}

	return blocks
}

func checksumForNumber(block Block, offset int) int {
	if block.fileId == -1 {
		return 0
	}
	checksum := 0
	for i := 0; i < block.size; i++ {
		checksum += (i + offset) * block.fileId
	}
	return checksum
}

func calculateChecksum(blocks []Block) int {
	checksum := 0
	offset := 0
	for i := 0; i < len(blocks); i++ {
		block := blocks[i]
		checksum += checksumForNumber(block, offset)
		offset += block.size
	}

	return checksum
}
