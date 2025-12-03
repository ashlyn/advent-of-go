package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/slices"
)

func main() {
	input := files.ReadFile(3, 2025, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	banks := parseInput(input)
	
	for i := range banks {
		result += maxJoltage(banks[i], 2)
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	banks := parseInput(input)
	
	for i := range banks {
		result += maxJoltage(banks[i], 12)
	}

	return result
}

func parseInput(input []string) [][]int {
	parsed := make([][]int, len(input))
	for i, line := range input {
		nums := make([]int, len(line))
		for j, char := range line {
			nums[j] = int(char - '0')
		}
		parsed[i] = nums
	}
	return parsed
}

func maxJoltage(bank []int, n int) int {
	if n == 1 {
		return slices.Max(bank)
	}
	front := bank[:len(bank)-(n-1)]
	max := slices.Max(front)
	maxIndex := slices.IndexOfInt(max, front)
	return (max * maths.Pow(10, n-1)) + maxJoltage(bank[maxIndex+1:], n-1)
}
