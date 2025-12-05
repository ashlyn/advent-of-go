package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"fmt"
	"sort"
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
	ranges, _ := parseInput(input)

	// sort ranges by min value so that we can simplify without backtracking
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	// iterate over ranges, simplifying down to non-overlapping ranges
	for i := 0; i < len(ranges)-1; {
		simplified := simplifyTwoRanges(ranges[i], ranges[i+1])
		if len(simplified) == 1 {
			ranges[i] = simplified[0]
			ranges = append(ranges[:i+1], ranges[i+2:]...)
		} else {
			i++
		}
	}

	return countIngredientsInRanges(ranges)
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

func simplifyTwoRanges(r1, r2 [2]int) [][2]int {
	// no overlap
	if r2[1] < r1[0] || r1[1] < r2[0] {
		return [][2]int{r1, r2}
	}
	// complete overlap, r2 in r1
	if r1[0] <= r2[0] && r1[1] >= r2[1] {
		return [][2]int{r1}
	}
	// complete overlap, r1 in r2
	if r2[0] <= r1[0] && r2[1] >= r1[1] {
		return [][2]int{r2}
	}
	// partial overlap
	newRange := [2]int{slices.Min([]int{r1[0], r2[0]}), slices.Max([]int{r1[1], r2[1]})}
	return [][2]int{newRange}
}

func countIngredientsInRange(r [2]int) int {
	return r[1] - r[0] + 1
}

func countIngredientsInRanges(ranges [][2]int) int {
	count := 0
	for i := 0; i < len(ranges); i++ {
		count += countIngredientsInRange(ranges[i])
	}
	return count
}