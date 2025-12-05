package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	input := files.ReadFile(4, 2025, "\n")
	sw := time.Now()
	println(solvePart1(input))
	fmt.Printf("Solved part 1 in %v\n", time.Since(sw))
	sw = time.Now()
	println(solvePart2(input))
	fmt.Printf("Solved part 2 in %v\n", time.Since(sw))
}

func solvePart1(input []string) int {
	warehouse := parseInput(input)
	return len(getAccessibleRolls(warehouse))
}

func solvePart2(input []string) int {
	totalRemoved := 0

	warehouse := parseInput(input)
	for accessibleCount := math.MaxInt; accessibleCount > 0; {
		accessible := getAccessibleRolls(warehouse)
		accessibleCount = len(accessible)
		totalRemoved += accessibleCount

		for _, coords := range accessible {
			warehouse[coords.Y][coords.X] = "."
		}
	}

	return totalRemoved
}

func parseInput(input []string) [][]string {
	parsed := make([][]string, len(input))
	for i, line := range input {
		parsed[i] = strings.Split(line, "")
	}
	return parsed
}

func isAccessible(coords grid.Coords, warehouse [][]string) bool {
	if !grid.IsInGrid(coords, warehouse) || warehouse[coords.Y][coords.X] != "@" {
		return false
	}
	x, y := coords.X, coords.Y
	if warehouse[y][x] == "@" {
		neighborCount := 0
		for neighborX := x - 1; neighborX <= x+1; neighborX++ {
			for neighborY := y - 1; neighborY <= y+1; neighborY++ {
				neighbor := grid.Coords{X: neighborX, Y: neighborY}
				if (neighborX != x || neighborY != y) && grid.IsInGrid(neighbor, warehouse) && warehouse[neighborY][neighborX] == "@" {
					neighborCount++
				}
			}
		}
		if neighborCount < 4 {
			return true
		}
	}
	return false
}

func getAccessibleRolls(warehouse [][]string) []grid.Coords {
	accessible := []grid.Coords{}
	maxX, maxY := grid.Size(warehouse)
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if warehouse[y][x] == "@" {
				if isAccessible(grid.Coords{X: x, Y: y}, warehouse) {
					accessible = append(accessible, grid.Coords{X: x, Y: y})
				}
			}
		}
	}
	return accessible
}
