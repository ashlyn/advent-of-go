package main

import (
	"advent-of-go/utils/files"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(2, 2025, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	idRanges := parseInput(input)

	for r := range idRanges {
		min, max := idRanges[r][0], idRanges[r][1]
		for id := min; id <= max; id++ {
			idStr := strconv.Itoa(id)
			if isDoubledString(idStr) {
				result += id
			}
		}
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	idRanges := parseInput(input)

	for r := range idRanges {
		min, max := idRanges[r][0], idRanges[r][1]
		for id := min; id <= max; id++ {
			idStr := strconv.Itoa(id)
			if isRepeatedString(idStr) {
				result += id
			}
		}
	}

	return result
}

func parseInput(input []string) [][2]int {
	ranges := strings.Split(strings.Join(input, ""), ",")
	idRanges := make([][2]int, len(ranges))
	for i, r := range ranges {
		minMax := strings.Split(r, "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		idRanges[i][0] = min
		idRanges[i][1] = max
	}
	return idRanges
}

func isDoubledString(s string) bool {
	middle := len(s) / 2
	firstHalf, secondHalf := s[0:middle], s[middle:]
	return firstHalf == secondHalf
}


// regex such as `(\d+?)\1+` would work but go's regex library does not support backreferences
func isRepeatedString(s string) bool {
	middle := len(s) / 2
	for i := 1; i <= middle; i++ {
		if strings.ReplaceAll(s, s[0:i], "") == "" {
			return true
		}
	}
	return false
}

