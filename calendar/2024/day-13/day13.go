package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"regexp"
	"strconv"
)

type machine struct {
	a grid.Coords
	b grid.Coords
	prize grid.Coords
}

func main() {
	input := files.ReadFile(13, 2024, "\n\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	machines := parseInput(input, 0)
	for _, m := range machines {
		aPresses, bPresses := calculateButtonPressesToPrize(m)
		result += calculateTotalTokenCost(aPresses, bPresses)
	}

	return result
}

func solvePart2(input []string) int {
	result := 0
	
	machines := parseInput(input, 10000000000000)
	for _, m := range machines {
		aPresses, bPresses := calculateButtonPressesToPrize(m)
		result += calculateTotalTokenCost(aPresses, bPresses)
	}

	return result
}

func parseInput(input []string, extraDistance int) []machine {
	machines := make([]machine, len(input))
	coordPattern := regexp.MustCompile(`[XY][+=](\d+)`)
	for i := 0; i < len(input); i++ {
		matches := coordPattern.FindAllStringSubmatch(input[i], -1)
		ax, _ := strconv.Atoi(matches[0][1])
		ay, _ := strconv.Atoi(matches[1][1])
		bx, _ := strconv.Atoi(matches[2][1])
		by, _ := strconv.Atoi(matches[3][1])
		prizex, _ := strconv.Atoi(matches[4][1])
		prizey, _ := strconv.Atoi(matches[5][1])
		machines[i] = machine{
			a: grid.Coords{X: ax, Y: ay},
			b: grid.Coords{X: bx, Y: by},
			prize: grid.Coords{X: prizex + extraDistance, Y: prizey + extraDistance},
		}
	}
	return machines
}

// Using Cramer's Rule
// https://www.chilimath.com/lessons/advanced-algebra/cramers-rule-with-two-variables/
func calculateButtonPressesToPrize(m machine) (int, int) {
	aPresses, bPresses := 0, 0

	coeffientsDeterminant := (m.a.X * m.b.Y) - (m.a.Y * m.b.X)
	if coeffientsDeterminant != 0 {
		aDeterminant := (m.prize.X * m.b.Y) - (m.prize.Y * m.b.X)
		aPresses = aDeterminant / coeffientsDeterminant
	}

	if m.b.Y != 0 {
		bPresses = (m.prize.Y - (aPresses * m.a.Y)) / m.b.Y
	}

	if ((m.a.X * aPresses) + (m.b.X * bPresses)) == m.prize.X && ((m.a.Y * aPresses) + (m.b.Y * bPresses)) == m.prize.Y {
		return aPresses, bPresses
	}

	return 0, 0
}

func calculateTotalTokenCost(aPresses, bPresses int) int {
	return (3 * aPresses) + bPresses
}
