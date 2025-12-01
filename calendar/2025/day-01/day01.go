package main

import (
	"advent-of-go/utils/circulararray"
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"strconv"
)

func main() {
	input := files.ReadFile(1, 2025, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	min, max, current := 0, 99, 50
	lockSize := max - min + 1

	instructions := parseInput(input)

	for _, steps := range instructions {
		current = circulararray.CircularIndex(current, steps, lockSize)
		if current == 0 {
			result++
		}
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	min, max, current := 0, 99, 50

	instructions := parseInput(input)

	for _, steps := range instructions {
		for i := 0; i < maths.Abs(steps); i++ {
			if steps < 0 {
				current--
				if current < min {
					current = max
				}
			} else {
				current++
				if current > max {
					current = min
				}
			}
			if current == 0 {
				result++
			}
		}
	}

	return result
}

func parseInput(input []string) []int {
	instructions := make([]int, len(input))
	for i, line := range input {
		direction := line[0:1]
		count, _ := strconv.Atoi(line[1:])
		if direction == "L" {
			count = -count
		}
		instructions[i] = count
	}
	return instructions
}
