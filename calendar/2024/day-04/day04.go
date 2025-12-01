package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/str"
	"strings"
)

func main() {
	input := files.ReadFile(4, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	rotated := rotateGrid(input)

	result += findHorizontalOccurances(input, "XMAS")
	result += findHorizontalOccurances(rotated, "XMAS")

	result += findHorizontalOccurances(buildDiagonals(input), "XMAS")
	result += findHorizontalOccurances(buildDiagonals(rotated), "XMAS")

	return result
}

func solvePart2(input []string) int {
	result := 0

	for y := 1; y < len(input) - 1; y++ {
		for x := 1; x < len(input[y]) - 1; x++ {
			if (input[y][x] == 'A') {
				topLeft, topRight, bottomLeft, bottomRight := input[y - 1][x - 1], input[y - 1][x + 1], input[y + 1][x - 1], input[y + 1][x + 1]
				if ((topLeft == 'M' && bottomRight == 'S') || (topLeft == 'S' && bottomRight == 'M')) &&
					((topRight == 'M' && bottomLeft == 'S') || (topRight == 'S' && bottomLeft == 'M')) {
					result++
				}
			}
		}
	}

	return result
}

func findHorizontalOccurances(input []string, searchString string) int {
	result := 0

	for _, line := range input {
		result += strings.Count(line, searchString)
		result += strings.Count(str.Reverse(line), searchString)
	}

	return result
}

func rotateGrid(input []string) []string {
	rotated := []string{}
	split := [][]string{}

	for _, line := range input {
		split = append(split, strings.Split(line, ""))
	}

	rotatedChars := grid.Rotate90(split)
	for _, line := range rotatedChars {
		rotated = append(rotated, strings.Join(line, ""))
	}
	
	return rotated
}

func buildDiagonals(input []string) []string {
	diagonals := []string{}

	rows, columns := len(input), len(input[0])
	for line := 1; line <= rows + columns - 1; line++ {
		start := maths.Max(0, line - rows)
		items := maths.Min(line, maths.Min(columns - start, rows))

		d := ""
		for i := 0; i < items; i++ {
			d += input[maths.Min(rows, line) - i - 1][start + i:start + i + 1]
		}
		diagonals = append(diagonals, d)
	}

	return diagonals
}
