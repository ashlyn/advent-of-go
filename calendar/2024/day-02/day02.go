package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/slices"
	"strings"
)

func main() {
	input := files.ReadFile(2, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	for _, line := range input {
		if isSafe(slices.ParseIntsFromStrings(strings.Fields(line))) {
			result++
		}
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	for _, line := range input {
		numbers := slices.ParseIntsFromStrings(strings.Fields(line))
		if isSafe(numbers) {
			result++
		} else {
			for i := 0; i < len(numbers); i++ {
				amended := make([]int, i)
				copy(amended, numbers[:i])
				amended = append(amended, numbers[i+1:]...)
				if isSafe(amended) {
					result++
					break
				}
			}
		}
	}

	return result
}

func isSafe(numbers []int) bool {
	directions := map[int]int{}
	for i := 0; i < len(numbers) - 1; i++ {
		diff := numbers[i] - numbers[i + 1]
		if (diff > 0) {
			directions[1]++
		} else if (diff < 0) {
			directions[-1]++
		}
		diffAbs := maths.Abs(diff)
		if diffAbs > 3 || diffAbs < 1 || len(directions) > 1 {
			return false
		}
	}
	return true
}
