package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"fmt"
	"strings"
	"time"
)

type row struct {
	contents string
	pattern []int
}

func main() {
	input := files.ReadFile(12, 2023, "\n")
	sw := time.Now()
	println(solvePart1(input))
	fmt.Printf("Solved part 1 in %v\n", time.Since(sw))
	sw = time.Now()
	println(solvePart2(input))
	fmt.Printf("Solved part 2 in %v\n", time.Since(sw))
}

func solvePart1(input []string) int {
	result := 0

	rows := parseInput(input)
	for i := 0; i < len(rows); i++ {
		result += rows[i].findValidArrangementsBruteForce()
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	rows := parseInput(input)
	cache := map[string]int{}
	for i := 0; i < len(rows); i++ {
		unfolded := rows[i].unfold()
		result += findValidArrangementsRecursiveCached(unfolded.contents, unfolded.pattern, cache)
	}

	return result
}

func (r row) findValidArrangementsBruteForce() int {
	totalDamagedTarget, totalDamaged := slices.Sum(r.pattern), strings.Count(r.contents, "#")
	damagedToAdd := totalDamagedTarget - totalDamaged
	if damagedToAdd == 0 {
		return 1
	}

	// generate all combinations of indexes where n damaged springs could be placed
	unknownIndexes := getUnknownIndexes(r.contents)
	possibleArrangements := slices.GenerateCombinationsLengthN(unknownIndexes, damagedToAdd)
	arrangements := 0

	for i := 0; i < len(possibleArrangements); i++ {
		arrangementSlice := strings.Split(r.contents, "")
		for j := 0; j < len(possibleArrangements[i]); j++ {
			arrangementSlice[possibleArrangements[i][j]] = "#"
		}
		arrangement := strings.ReplaceAll(strings.Join(arrangementSlice, ""), "?", ".")
		if matchesPattern(arrangement, r.pattern) {
			arrangements++
		}
	}

	return arrangements
}

func findValidArrangementsRecursiveCached(line string, pattern []int, cache map[string]int) int {
	key := buildMemoKey(line, pattern)
	previous, inCache := cache[key]
	if inCache {
		return previous
	}

	// no pattern groups left to match
	if len(pattern) == 0 {
		// if there are still damaged springs, this is not a valid arrangement
		if strings.Contains(line, "#") {
			cache[key] = 0
			return 0
		}
		cache[key] = 1
		return 1
	}

	arrangements, currentGroup, remaining := 0, pattern[0], slices.Sum(pattern[1:])
	// leave enough space for remaining springs at the end
	for i := 0; i < len(line) - remaining; i++ {
		for j := i; j < i + currentGroup; j++ {
			if j == i && j > 0 && line[j - 1] == '#' {
				cache[key] = arrangements
				return arrangements
			}
			if j >= len(line) {
				cache[key] = arrangements
				return arrangements
			}
			// does not match current pattern group
			if line[j] == '.' {
				break
			}
			if j == i + currentGroup - 1 {
				// groups must be separated, bad ending
				if j + 1 < len(line) && line[j + 1] == '#' {
					break
				}

				// process next group
				index := j + 2
				newLine := ""
				if index < len(line) {
					newLine = line[index:]
				}
				arrangements += findValidArrangementsRecursiveCached(newLine, pattern[1:], cache)
			}
		}
	}

	cache[key] = arrangements
	return arrangements
}

func parseInput(input []string) []row {
	result := []row{}
	for i := 0; i < len(input); i++ {
		parts := strings.Fields(input[i])
		result = append(result, row{contents: parts[0], pattern: parsePattern(parts[1])})
	}
	return result
}

func parsePattern(patternStr string) []int {
	parts := strings.Split(patternStr, ",")
	return slices.ParseIntsFromStrings(parts)
}

func matchesPattern(line string, pattern []int) bool {
	parts := strings.Fields(strings.ReplaceAll(line, ".", " "))
	if len(parts) != len(pattern) {
		return false
	}
	for i := 0; i < len(parts); i++ {
		if len(parts[i]) != pattern[i] {
			return false
		}
	}
	return true
}

func getUnknownIndexes(line string) []int {
	result := []int{}
	for i := 0; i < len(line); i++ {
		if line[i] == '?' {
			result = append(result, i)
		}
	}
	return result
}

func (r row) unfold() row {
	newContents, newPattern := make([]string, 5), make([]int, len(r.pattern) * 5)
	for i := 0; i < 5; i++ {
		newContents[i] = r.contents
		for j := 0; j < len(r.pattern); j++ {
			newPattern[i * len(r.pattern) + j] = r.pattern[j]
		}
	}
	return row{contents: strings.Join(newContents, "?"), pattern: newPattern}
}

func buildMemoKey(line string, pattern []int) string {
	return fmt.Sprintf("%s-%v", line, pattern)
}
