package main

import (
	"advent-of-go/utils/files"
	"regexp"
	"strconv"
)

func main() {
	input := files.ReadFile(3, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	return runInstructions(input, false)
}

func solvePart2(input []string) int {
	return runInstructions(input, true)
}

func runInstructions(input []string, useDoDont bool) int {
	result := 0

	instructionsPattern := *regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	enabled := true
	for _, line := range input {
		matches := instructionsPattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
			} else if useDoDont && match[0] == "don't()" {
				enabled = false
			} else if enabled {
				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])
				result += a * b
			}
		}
	}

	return result
}
