package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"fmt"
	"strings"
	"time"
)

func main() {
	input := files.ReadFile(5, 2025, "\n\n")
	sw := time.Now()
	println(solvePart1(input))
	fmt.Printf("Solved part 1 in %v\n", time.Since(sw))
	sw = time.Now()
	println(solvePart2(input))
	fmt.Printf("Solved part 2 in %v\n", time.Since(sw))
}

func solvePart1(input []string) int {
	result := 0

	ranges, ingredientIds := parseInput(input)
	for i := 0; i < len(ingredientIds); i++ {
		if isFresh(ingredientIds[i], ranges) {
			result++
		}	
	}

	return result
}

func solvePart2(input []string) int {
	result := 0



	return result
}

func parseInput(input []string) ([][2]int, []int) {
	rangesInput, ingredientsInput := strings.Split(input[0], "\n"), strings.Split(input[1], "\n")
	ranges := make([][2]int, len(rangesInput))
	ingredientIds := make([]int, len(ingredientsInput))
	for i, line := range rangesInput {
		ranges[i] = parseRange(line)
	}
	ingredientIds = slices.ParseIntsFromStrings(ingredientsInput)
	return ranges, ingredientIds
}

func parseRange(line string) [2]int {
	parts := strings.Split(line, "-")
	parsed := slices.ParseIntsFromStrings(parts)
	return [2]int{parsed[0], parsed[1]}
}

func isInRange(ingredient int, r [2]int) bool {
	return ingredient >= r[0] && ingredient <= r[1]
}

func isFresh(ingredient int, ranges [][2]int) bool {
	for i := 0; i < len(ranges); i++ {
		if isInRange(ingredient, ranges[i]) {
			return true
		}
	}
	return false
}
