package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(15, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	steps := parseInput(input)
	for _, s := range steps {
		result += hash(s)
	}

	return result
}

func solvePart2(input []string) int {
	steps := parseInput(input)
	boxes, lenses := buildBoxes(steps)
	return calculateFocusingPower(boxes, lenses)
}

func buildKey(label string, box int) string {
	return fmt.Sprintf("%d-%s", box, label)
}

func parseInput(input []string) []string {
	return strings.Split(input[0], ",")
}

func hash(input string) int {
	currentValue := 0
	for _, char := range input {
		currentValue += int(char)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

func buildBoxes(steps []string) ([][]string, map[string]int) {
	boxes := make([][]string, 256)
	lenseMap := map[string]int{}
	for i := 0; i < len(boxes); i++ {
		boxes[i] = []string{}
	}
	for _, s := range steps {
		if strings.Contains(s, "-") {
			parts := strings.Split(s, "-")
			label := parts[0]
			boxNumber := hash(label)
			lensIndex := slices.IndexOfStr(label, boxes[boxNumber])
			if lensIndex != -1 {
				boxes[boxNumber] = append(boxes[boxNumber][:lensIndex], boxes[boxNumber][lensIndex+1:]...)
			}
		} else if strings.Contains(s, "=") {
			parts := strings.Split(s, "=")
			label := parts[0]
			boxNumber := hash(label)
			focalLength, _ := strconv.Atoi(parts[1])
			lenseMap[buildKey(label, boxNumber)] = focalLength
			if !slices.Contains(boxes[boxNumber], label) {
				boxes[boxNumber] = append(boxes[boxNumber], label)
			}
		}
	}
	return boxes, lenseMap
}

func calculateFocusingPower(boxes [][]string, lensMap map[string]int) int {
	focusingPower := 0
	for boxNumber, boxLenses := range boxes {
		for lensSlot, lensLabel := range boxLenses {
			focalLength := lensMap[buildKey(lensLabel, boxNumber)]
			focusingPower += (1 + boxNumber) * (lensSlot + 1) * focalLength
		}
	}
	return focusingPower
}
