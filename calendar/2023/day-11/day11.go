package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/slices"
	"strings"
)

func main() {
	input := files.ReadFile(11, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	return findShortestPaths(input, 2)
}

func solvePart2(input []string) int {
	return findShortestPaths(input, 1000000)
}

func findEmptyRows(input []string) []int {
	emptyRows := []int{}
	for i, row := range input {
		if !strings.Contains(row, "#") {
			emptyRows = append(emptyRows, i)
		}
	}
	return emptyRows
}

func findEmptySpaces(input []string) ([]int, []int) {
	emptyRows := findEmptyRows(input)
	columns := []string{}
	for x := 0; x < len(input[0]); x++ {
		newRow := ""
		for y := len(input) - 1; y >= 0; y-- {
			newRow += string(input[y][x])
		}
		columns = append(columns, newRow)
	}
	emptyColumns := findEmptyRows(columns)
	return emptyRows, emptyColumns
}

func findGalaxiesVirtually(expanded []string, expansionSize int) []grid.Coords {
	galaxies := []grid.Coords{}
	emptyRows, emptyColumns := findEmptySpaces(expanded)
	deltaY := 0
	for y := 0; y < len(expanded); y++ {
		deltaX := 0
		if slices.Contains(emptyRows, y) {
			deltaY += (expansionSize - 1)
			continue
		}
		for x := 0; x < len(expanded[y]); x++ {
			if slices.Contains(emptyColumns, x) {
				deltaX += (expansionSize - 1)
				continue
			}
			if expanded[y][x] == '#' {
				galaxies = append(galaxies, grid.Coords{X: x + deltaX, Y: y + deltaY})
			}
		}
	}
	return galaxies
}

func findShortestPaths(input []string, expansionSize int) int {
	result := 0

	galaxies := findGalaxiesVirtually(input, expansionSize)
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			result += galaxies[i].ManhattanDistance(galaxies[j])
		}
	}

	return result
}
