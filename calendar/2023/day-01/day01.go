package main

import (
	"advent-of-go/utils/files"
	"regexp"
	"strconv"
)

func main() {
	input := files.ReadFile(1, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	pattern := regexp.MustCompile(`[0-9]`)
	for _, line := range input {
		result += getCalibrationValue(line, pattern)
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	pattern := regexp.MustCompile(`([0-9]|one|two|three|four|five|six|seven|eight|nine)`)
	for _, line := range input {
		result += getCalibrationValue(line, pattern)
	}

	return result
}

func getCalibrationValue(input string, pattern *regexp.Regexp) int {
	lastMatch := pattern.FindString(input)
	calStr := wordToDigit(lastMatch)
	for i := 1; i < len(input); i++ {
		match := pattern.FindString(input[i:])
		if match == "" {
			break
		} else {
			lastMatch = match
		}
	}
	calStr += wordToDigit(lastMatch)
	calVal, _ := strconv.Atoi(calStr)
	return calVal
}

func wordToDigit(word string) string {
	digits := map[string]string{
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}
	
	digit, ok := digits[word]
	if !ok {
		return word
	}
	return digit
}
