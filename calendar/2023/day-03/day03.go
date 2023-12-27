package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"fmt"
	"regexp"
	"strconv"
)

type number struct {
	value int
	left *grid.Coords
	right *grid.Coords
}

func main() {
	input := files.ReadFile(3, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	numbers := parseNumbers(input)
	gearMap := make(map[string][]int)
	for _, number := range numbers {
		if checkForSymbols(number, input, &gearMap) {
			result += number.value
		}
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	numbers := parseNumbers(input)
	gearMap := make(map[string][]int)
	for _, number := range numbers {
		checkForSymbols(number, input, &gearMap)
	}

	for _, partNumbers := range gearMap {
		if len(partNumbers) == 2 {
			result += partNumbers[0] * partNumbers[1]
		}
	}

	return result
}

func parseNumbers(input []string) []number {
	allNumbers := []number{}
	numberPattern := regexp.MustCompile(`\d+`)
	for y, line := range input {
		matches := numberPattern.FindAllString(line, -1)
		matchIndexes := numberPattern.FindAllStringIndex(line, -1)
		for i, match := range matches {
			num, _ := strconv.Atoi(match)
			number := number{
				value: num,
				left: &grid.Coords{ X: matchIndexes[i][0], Y: y },
				right: &grid.Coords{ X: matchIndexes[i][1] -1, Y: y },
			}

			allNumbers = append(allNumbers, number)
		}
	}
	return allNumbers
}

func checkForSymbols(number number, grid []string, gearMap *map[string][]int) bool {
	spaceLeft, spaceRight := 0, 0

	if number.left.X > 0 {
		spaceLeft = 1
		leftSymbol := grid[number.left.Y][number.left.X - 1]
		if isSymbol(rune(leftSymbol)) {
			if leftSymbol == '*' {
				key := fmt.Sprintf("%d,%d", number.left.X - 1, number.left.Y)
				(*gearMap)[key] = append((*gearMap)[key], number.value)
			}
			return true
		}
	}

	if number.right.X < len(grid[number.left.Y]) - 1 {
		spaceRight = 1
		rightSymbol := grid[number.right.Y][number.right.X + 1]
		if isSymbol(rune(rightSymbol)) {
			if rightSymbol == '*' {
				key := fmt.Sprintf("%d,%d", number.right.X + 1, number.right.Y)
				(*gearMap)[key] = append((*gearMap)[key], number.value)
			}
			return true
		}
	}

	if number.left.Y > 0 {
		topRange := grid[number.left.Y - 1][number.left.X - spaceLeft:number.right.X + spaceRight + 1]
		for i, character := range topRange {
			if isSymbol(rune(character)) {
				if character == '*' {
					key := fmt.Sprintf("%d,%d", number.left.X - spaceLeft + i, number.left.Y - 1)
					(*gearMap)[key] = append((*gearMap)[key], number.value)
				}
				return true
			}
		}
	}

	if number.left.Y < len(grid) - 1 {
		bottomRange := grid[number.left.Y + 1][number.left.X - spaceLeft:number.right.X + spaceRight + 1]
		for i, character := range bottomRange {
			if isSymbol(rune(character)) {
				if character == '*' {
					key := fmt.Sprintf("%d,%d", number.left.X - spaceLeft + i, number.left.Y + 1)
					(*gearMap)[key] = append((*gearMap)[key], number.value)
				}
				return true
			}
		}
	}

	return false
}

func isSymbol(character rune) bool {
	if character == '.' {
		return false
	}
	return character < '0' || character > '9';
}
