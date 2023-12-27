package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/sets"
	"advent-of-go/utils/slices"
	"fmt"
	"sort"
)

func main() {
	input := files.ReadFile(21, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	stepGoal := 64
	reachable := calculateExactSteps(input, []int{ stepGoal })
	return reachable[stepGoal]
}

// Took a lot of help/inspiration from the subreddit discussions
// While I identified the odd/even oscillation and that there was
// generally a polynomial (quadratic specifically) relationship
// between steps and the number of reachable plots, I didn't understand
// the pattern or math enough to figure out the exact polynomial and
// needed to learn about LaGrange polynomials and the step/offset patterns
func solvePart2(input []string) int {
	stepGoal := 26501365
	length := len(input)
	offset := stepGoal % length
	x0, x1, x2 := offset, offset + length, offset + (length * 2)
	plotsMap := calculateExactSteps(input, []int{ x0, x1, x2 })
	
	coefficients := maths.FindLagrangeCoefficients(float64(plotsMap[x0]), float64(plotsMap[x1]), float64(plotsMap[x2]))

	s := float64((stepGoal - offset) / length)
	return int(maths.CalculatePolynomial(coefficients, s))

}

// Not optimized for larger step counts and has long runtime for part 2
// Started with a Dijkstra approach but realized that it wasn't shortest paths and had backtracking
var directions = []grid.Coords{ {X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0} }
func calculateExactSteps(input []string, steps []int) map[int]int {
	plots := sets.New()
	start := ""
	gridSize := grid.Coords{ X: len(input[0]), Y: len(input) }

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			currentCharacter, key := input[y][x], fmt.Sprintf("%d,%d", x, y)
			if currentCharacter != '#' {
				plots.Add(key)
			}
			if currentCharacter == 'S' {
				start = key
			}
		}
	}

	reachable := sets.New()
	reachable.Add(start)

	reachableAtSteps := map[int]int{}
	sort.Ints(steps)
	maxSteps := steps[len(steps) - 1]

	for i := 1; i <= maxSteps; i++ {
		newReachable := sets.New()
		for reachable.Size() > 0 {
			key := reachable.Random()
			reachable.Remove(key)
			coords := grid.ParseCoords(key)
			for _, d := range directions {
				next := grid.Coords{X: coords.X + d.X, Y: coords.Y + d.Y}
				nextKey, translatedKey := next.ToString(), translateRepeatedCoords(next, gridSize).ToString()
				if plots.Has(translatedKey) {
					newReachable.Add(nextKey)
				}
			}
		}
		if slices.Contains(steps, i) {
			reachableAtSteps[i] = newReachable.Size()
		}
		reachable = newReachable
	}

	return reachableAtSteps
}

func translateRepeatedCoords(repeated grid.Coords, gridSize grid.Coords) grid.Coords {
	return grid.Coords{ X: (repeated.X % gridSize.X + gridSize.X) % gridSize.X, Y: (repeated.Y % gridSize.Y + gridSize.Y) % gridSize.Y }
}
