package main

import (
	"advent-of-go/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(19, 2024, "\n\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	towels, patterns := parseInput(input)
	cache := make(map[string]int)
	for i := 0; i < len(patterns); i++ {
		if buildPattern(patterns[i], towels, cache) > 0 {
			result++
		}
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	towels, patterns := parseInput(input)
	cache := make(map[string]int)
	for i := 0; i < len(patterns); i++ {
		waysToBuild := buildPattern(patterns[i], towels, cache)
		if waysToBuild > 0 {
			result += waysToBuild
		}
	}

	return result
}

func parseInput(input []string) ([]string, []string) {
	availableTowels := strings.Split(input[0], ", ")
	patterns := strings.Split(input[1], "\n")
	return availableTowels, patterns
}

func buildPattern(pattern string, towels []string, cache map[string]int) int {
	if len(pattern) == 0 {
		return 1
	}

	if result, ok := cache[pattern]; ok {
		return result
	}

	for i := 0; i < len(towels); i++ {
		if len(towels[i]) > len(pattern) {
			continue
		}
		rest, startsWith := strings.CutPrefix(pattern, towels[i])
		if startsWith {
			cache[pattern] += buildPattern(rest, towels, cache)
		}
	}
	return cache[pattern]
}
