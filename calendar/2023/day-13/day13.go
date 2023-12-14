package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/slices"
	"advent-of-go/utils/str"
	"strings"
)

func main() {
	input := files.ReadFile(13, 2023, "\n\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	columns, rows := 0, 0
	for i := 0; i < len(input); i++ {
		c, r := findReflection(strings.Split(input[i], "\n"))
		if len(c) > 0 {
			columns += c[0]
		} else if len(r) > 0 {
			rows += r[0]
		}
	}

	return columns + (100 * rows)
}

func solvePart2(input []string) int {
	columns, rows := 0, 0
	for i := 0; i < len(input); i++ {
		frame := strings.Split(input[i], "\n")
		c, r := findReflection(frame)
		originalColumn, originalRow := 0, 0
		if len(c) > 0 {
			originalColumn = c[0]
		} else if len(r) > 0 {
			originalRow = r[0]
		}
		findSmudge:
			for y := 0; y < len(frame); y++ {
				for x := 0; x < len(frame[y]); x++ {
					frameCopy := make([]string, len(frame))
					copy(frameCopy, frame)
					newSymbol := "#"
					if frameCopy[y][x] == '#' {
						newSymbol = "."
					}
					frameCopy[y] = str.ReplaceCharAt(frameCopy[y], newSymbol, x)
					columnValues, rowValues := findReflection(frameCopy)
					if len(columnValues) > 0 {
						column := findNewValue(columnValues, originalColumn)
						if column != originalColumn {
							columns += column
							break findSmudge
						}
					}
					if len(rowValues) > 0 {
						row := findNewValue(rowValues, originalRow)
						if row != originalRow {
							rows += row
							break findSmudge
						}
					}
				}
			}
	}

	return columns + (100 * rows)
}

func isReflectedAt(input string, index int) bool {
	if len(input) < 2 || index == 0 || index > len(input) - 1 {
		return false
	}
	charactersReflected := maths.Min(index, len(input) - index)
	left, right := input[index - charactersReflected:index], input[index:index + charactersReflected]
	return str.Reverse(left) == right
}

func reflectionColumns(frame []string) []int {
	reflections := []int{}
	for i := 1; i < len(frame[0]); i++ {
		reflectedAtI := true
		for row := 0; row < len(frame) && reflectedAtI; row++ {
			reflectedAtI = isReflectedAt(frame[row], i)
		}
		if reflectedAtI {
			reflections = append(reflections, i)
		}
	}
	return reflections
}

func reflectionRows(frame []string) []int {
	reflections := []int{}
	if len(frame) < 2 {
		return reflections
	}

	for i := 1; i < len(frame); i++ {
		rowsReflected := maths.Min(i, len(frame) - i)
		reflectedAtI := true
		above, below := frame[i - rowsReflected:i], slices.Reverse(frame[i:i + rowsReflected])
		for row := 0; row < len(above) && reflectedAtI; row++ {
			reflectedAtI = above[row]	== below[row]
		}
		if reflectedAtI {
			reflections = append(reflections, i)
		}
	}
	return reflections
}

func findReflection(frame []string) ([]int, []int) {
	return reflectionColumns(frame), reflectionRows(frame)
}

func findNewValue(values []int, originalValue int) int {
	if len(values) == 1 {
		return values[0]
	}
	for _, v := range values {
		if v != originalValue {
			return v
		}
	}
	return originalValue
}
