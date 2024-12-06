package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/sets"
	"fmt"
)

func main() {
	input := files.ReadFile(6, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	visited, _ := traverseGrid(input)
	return visited
}

// This brute-force solution is very slow
// May optimize later if time/interest allows
func solvePart2(input []string) int {
	loops := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			if input[y][x] == '.' {
				newInput := make([]string, len(input))
				copy(newInput, input)
				newInput[y] = newInput[y][:x] + "#" + newInput[y][x+1:]
				_, loop := traverseGrid(newInput)
				if loop {
					loops++
				}
			}
		}
	}
	return loops
}

func getStartingCoords(input []string) (c *grid.Coords) {
	for y, line := range input {
		for x, char := range line {
			if char == '^' {
				return &grid.Coords{X: x, Y: y}
			}
		}
	}
	return nil
}

func traverseGrid(input []string) (int, bool) {
	start := getStartingCoords(input)
	visited := sets.New()
	loopTracker := sets.New()

	directions := []*grid.Coords{
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
	}

	currentDirecton := 0
	direction := directions[currentDirecton]

	current := start

	for {
		visited.Add(current.ToString())

		nextKey := fmt.Sprintf("%s %d", current.ToString(), currentDirecton)
		if loopTracker.Has(nextKey) {
			return visited.Size(), true
		}
		loopTracker.Add(nextKey)

		next := &grid.Coords{X: current.X + direction.X, Y: current.Y + direction.Y}
		if !(next.X >= 0 && next.X < len(input[0]) && next.Y >= 0 && next.Y < len(input)) {
			return visited.Size(), false
		}

		if input[next.Y][next.X] == '#' {
			currentDirecton = (currentDirecton + 1) % 4
			direction = directions[currentDirecton]
		} else {
			current = next
		}
	}
}
