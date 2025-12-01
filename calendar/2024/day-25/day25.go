package main

import (
	"advent-of-go/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(25, 2024, "\n\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	locks, keys := parseInput(input)

	for _, lock := range locks {
		for _, key := range keys {
			if fits(lock, key) {
				result++
			}
		}
	}

	return result
}

func solvePart2(input []string) string {
	return "Merry Christmas!"
}

func parseInput(input []string) ([][]int, [][]int) {
	locks, keys := [][]int{}, [][]int{}

	for d := 0; d < len(input); d++ {
		diagram := strings.Split(input[d], "\n")
		parsed := parseDiagram(diagram)
		if diagram[0][0] == '#' {
			locks = append(locks, parsed)
		} else if diagram[0][0] == '.' {
			keys = append(keys, parsed)
		}
	}

	return locks, keys
}

func parseDiagram(input []string) []int {
	lock := make([]int, len(input[0]))
	for column := 0; column < len(input[0]); column++ {
		lock[column] = -1
		for row := 0; row < len(input); row++ {
			if input[row][column] == '#' {
				lock[column]++
			}
		}
	}
	return lock
}

func fits(lock, key []int) bool {
	for i := 0; i < len(lock); i++ {
		if key[i] + lock[i] > 5 {
			return false
		}
	}
	return true
}
