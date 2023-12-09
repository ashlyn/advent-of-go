package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(9, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0
	sequences := parseInput(input)
	for i := 0; i < len(sequences); i++ {
		result += solveSequence(sequences[i])
	}
	return result
}

func solvePart2(input []string) int {
	result := 0
	sequences := parseInput(input)
	for i := 0; i < len(sequences); i++ {
		s := slices.Reverse(sequences[i])
		result += solveSequence(s)
	}
	return result
}

func parseInput(input []string) [][]int {
	sequences := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		values := strings.Fields(input[i])
		sequences[i] = make([]int, len(values))
		for j := 0; j < len(values); j++ {
			value, _ := strconv.Atoi(values[j])
			sequences[i][j] = value
		}
	}
	return sequences
}

func solveSequence(sequence []int) int {
	diffs := [][]int{sequence}
	for !allZeroes(diffs[len(diffs)-1]) {
		diff := make([]int, len(diffs[len(diffs)-1])-1)
		for i := 0; i < len(diffs[len(diffs)-1])-1; i++ {
			diff[i] = diffs[len(diffs)-1][i+1] - diffs[len(diffs)-1][i]
		}
		diffs = append(diffs, diff)
	}

	result := 0
	for i := 0; i < len(diffs); i++ {
		result += diffs[i][len(diffs[i])-1]
	}
	return result
}

func allZeroes(sequence []int) bool {
	for i := 0; i < len(sequence); i++ {
		if sequence[i] != 0 {
			return false
		}
	}
	return true
}
