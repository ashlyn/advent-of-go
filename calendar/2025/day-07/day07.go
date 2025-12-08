package main

import (
	"advent-of-go/utils/colors"
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/slices"
	"advent-of-go/utils/str"
	"fmt"
	"strings"
	"time"
)

func main() {
	input := files.ReadFile(7, 2025, "\n")
	sw := time.Now()
	println(solvePart1(input))
	fmt.Printf("%s Solved part 1 in %v\n", colors.GreenString("✓"), time.Since(sw))
	sw = time.Now()
	println(solvePart2(input))
	fmt.Printf("%s Solved part 2 in %v\n", colors.GreenString("✓"), time.Since(sw))
}

func solvePart1(input []string) int {
	result := 0

	start, splitters := parseInput(input)
	current := map[grid.Coords]bool{ start: true }
	for y := start.Y + 1; y >= 0 && y < len(input); y++ {
		next := make(map[grid.Coords]bool)
		for c := range current {
			toCheck := grid.Coords{ X: c.X, Y: y }
			if _, exists := splitters[toCheck]; exists {
				result++
				next[grid.Coords{ X: c.X - 1, Y: y }] = true
				next[grid.Coords{ X: c.X + 1, Y: y }] = true
			} else {
				next[toCheck] = true
			}
		}
		current = next
	}

	return result
}

func solvePart2(input []string) int {
	start, _ := parseInput(input)

	splitsAtColumns := make([]int, len(input[0]))
	splitsAtColumns[start.X] = 1
	for y := start.Y; y < len(input); y++ {
		line := input[y]
		for x := 0; x < len(line); x++ {
			if line[x] == '^' {
				splitsAtColumns[x-1] += splitsAtColumns[x]
				splitsAtColumns[x+1] += splitsAtColumns[x]
				splitsAtColumns[x] = 0
			}
		}
	}

	return slices.Sum(splitsAtColumns)
}


// BROKEN -- either doesn't terminate correctly or this solution set is too large for this recursive approach (needs memoized)
func countTimelines(input []string, current grid.Coords, splitters map[grid.Coords]bool) int {
	if current.X < 0 || current.X >= len(input[0]) {
		return 0
	}
	if current.Y < 0 || current.Y >= len(input) {
		return 1
	}
	if _, exists := splitters[current]; exists {
		l, r := current.X-1, current.X+1
		if l < 0 {
			return countTimelines(input, grid.Coords{ X: r, Y: current.Y + 1 }, splitters)
		}
		if r >= len(input[0]) {
			return countTimelines(input, grid.Coords{ X: l, Y: current.Y + 1 }, splitters)
		}
		
		left :=countTimelines(input, grid.Coords{ X: current.X - 1, Y: current.Y + 1 }, splitters)
		right := countTimelines(input, grid.Coords{ X: current.X + 1, Y: current.Y + 1 }, splitters)
		return left + right
	}
	return countTimelines(input, grid.Coords{ X: current.X, Y: current.Y + 1 }, splitters)
}

func parseInput(input []string) (start grid.Coords, splitters map[grid.Coords]bool) {
	splitters = make(map[grid.Coords]bool)
	for i := 0; i < len(input); i++ {
		row := input[i]
		sIndex := strings.Index(row, "S")
		if sIndex > -1 {
			start = grid.Coords{ X: sIndex, Y: i }
		}
		splitterIndexes := str.IndexesAny(row, "^")
		for s := 0; s < len(splitterIndexes); s++ {
			splitters[grid.Coords{ X: splitterIndexes[s], Y: i }] = true
		}
	}
	return start, splitters
}
