package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"fmt"
	"strconv"
)

func main() {
	input := files.ReadFile(9, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) uint {
	blocks := parseBlocks(input[0])
	reallocated := reallocateBlocks(blocks)
	return calculateChecksum(reallocated)
}

func solvePart2(input []string) uint {
	blocks := parseBlocks(input[0])
	reallocated := reallocateFiles(blocks)
	return calculateChecksum(reallocated)
}

func parseBlocks(input string) []int {
	blocks := []int{}
	for i := 0; i < len(input); i++ {
		length, _ := strconv.Atoi(input[i:i+1])
		isFile := i % 2 == 0
		value := -1
		if isFile {
			value = i / 2
		}
		for j := 0; j < length; j++ {
			blocks = append(blocks, value)
		}
	}
	return blocks
}

func reallocateBlocks(blocks []int) []int {
	newBlocks := make([]int, len(blocks))
	copy(newBlocks, blocks)
	for slices.Contains(newBlocks, -1) {
		newBlocks = slices.TrimRight(newBlocks, -1)
		left := slices.IndexOfInt(-1, newBlocks)
		if left > -1 {
			newBlocks[left] = newBlocks[len(newBlocks)-1]
			newBlocks = newBlocks[0:len(newBlocks)-1]
		}
	}

	return slices.TrimRight(newBlocks, -1)
}

func reallocateFiles(blocks []int) []int {
	newBlocks := make([]int, len(blocks))
	copy(newBlocks, blocks)
	
	id, length := newBlocks[len(newBlocks)-1], 0
	for left := len(newBlocks) - 1; left > 0; {
		if newBlocks[left] == -1 || newBlocks[left] > id {
			left--
			continue
		}

		id, length = newBlocks[left], 0
		for left >= 0 && blocks[left] == id {
			left--
			length++
		}

		if length == 0 {
			left--
			continue
		}

		fileStart := left + 1

		emptyLength, emptyStart := 0, 0

		for right := 0; right < fileStart; {
			if newBlocks[right] == -1 {
				emptyLength++
			} else {
				emptyLength = 0
				emptyStart = right + 1
			}

			if emptyLength >= length {
				for i := 0; i < length; i++ {
					newBlocks[emptyStart+i] = newBlocks[fileStart+i]
					newBlocks[fileStart+i] = -1
				}
 				break
			}
			right++
		}
	}

	return newBlocks
}

func calculateChecksum(blocks []int) uint {
	var checksum uint = 0

	for i := 0; i < len(blocks); i++ {
		if (blocks[i] != -1) {
			checksum += uint(i * blocks[i])
		}
	}

	return checksum
}

func printAsExample(blocks []int) {
	for i := 0; i < len(blocks); i++ {
		if (blocks[i] == -1) {
			fmt.Print(".")
		} else {
			fmt.Print(blocks[i])
		}
	}
	fmt.Println()
}
