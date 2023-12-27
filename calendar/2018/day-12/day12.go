package main

import (
	"advent-of-go/utils/files"
	"regexp"
	"strings"
)

func main() {
	input := files.ReadFile(12, 2018, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	initial, transformations := parseInput(input)
	return transformGenerations(initial, transformations, 20)
}

func solvePart2(input []string) int {
	initial, transformations := parseInput(input)
	return transformGenerations(initial, transformations, 50000000000)
}

var plant, empty, separator = "#", ".", " => "
func parseInput(input []string) (string, map[string]string) {
	initial := strings.Fields(input[0])[2]
	transformations := map[string]string{}
	for i := 2; i < len(input); i++ {
		parts := strings.Split(input[i], separator)
		transformations[parts[0]] = parts[1]
	}
	return initial, transformations
}

var windowSizeLeft, windowSizeRight = 2, 2
var windowSize = windowSizeLeft + windowSizeRight + 1
func transform(initial string, transformations map[string]string) (string, int) {
	transformed := ""
	first, last := strings.Index(initial, plant), strings.LastIndex(initial, plant)
	leftPadding, rightPadding := (windowSize - 1) - first, (windowSize - 1) - (len(initial) - last - 1)
	padded := initial
	if leftPadding > 0 {
		padded = strings.Repeat(empty, leftPadding) + initial
	}
	if rightPadding > 0 {
		padded = padded + strings.Repeat(empty, rightPadding)
	}
	for i := windowSizeLeft; i < len(padded) - windowSizeRight; i++ {
		window := padded[i-windowSizeLeft:i+windowSizeRight+1]
		result, hasTransformation := transformations[window]
		if hasTransformation {
			transformed += result
		} else {
			transformed += empty
		}
	}
	extraEmpty := strings.Index(transformed, plant)
	return strings.Trim(transformed, empty), leftPadding - windowSizeLeft - extraEmpty
}

func transformGenerations(initial string, transformations map[string]string, generations int) int {
	pots, leftOffset, potNumberSum := initial, 0, 0
	for i := 1; i <= generations; i++ {
		newPots, additionalLeftOffset := transform(pots, transformations)
		if pots == newPots {
			newPotNumberSum := sumPotNumbers(newPots, leftOffset - additionalLeftOffset)
			growthRate := newPotNumberSum - potNumberSum
			return newPotNumberSum + (generations - i) * growthRate
		}
		leftOffset -= additionalLeftOffset
		pots = newPots
		potNumberSum = sumPotNumbers(pots, leftOffset)
	}

	return potNumberSum
}

func sumPotNumbers(pots string, left int) int {
	potPattern := regexp.MustCompile(plant)
	potSum := 0
	for _, p := range potPattern.FindAllStringIndex(pots, -1) {
		potSum += p[0] + left
	}
	return potSum
}
