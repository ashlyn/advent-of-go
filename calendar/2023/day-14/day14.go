package main

import (
	"advent-of-go/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(14, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	return calculateLoad(tiltNorth(input))
}

func solvePart2(input []string) int {
	return calculateLoad(runCycles(input, 1000000000))
}

// Various tilt functions could definitely be consolidated
// using transposition or variably controlling the loops
// but this version works for now
func tiltNorth(platform []string) []string {
	tilted := make([][]string, len(platform))
	for i := 0; i < len(platform); i++ {
		tilted[i] = strings.Split(platform[i], "")
	}

	for column := 0; column < len(tilted[0]); column++ {
		for row := 1; row < len(tilted); row++ {
			if tilted[row][column] != "O" {
				continue
			}
			tilted[row][column] = "."
			for potentialEmpty := row; potentialEmpty >= 0; potentialEmpty-- {
				if potentialEmpty == 0 {
					tilted[potentialEmpty][column] = "O"
					break
				}
				if tilted[potentialEmpty-1][column] != "." {
					tilted[potentialEmpty][column] = "O"
					break
				}
			}
		}
	}

	result := make([]string, len(tilted))
	for i := 0; i < len(tilted); i++ {
		result[i] = strings.Join(tilted[i], "")
	}
	return result
}

func tiltSouth(platform []string) []string {
	tilted := make([][]string, len(platform))
	for i := 0; i < len(platform); i++ {
		tilted[i] = strings.Split(platform[i], "")
	}

	for column := 0; column < len(tilted[0]); column++ {
		for row := len(tilted) - 2; row >= 0; row-- {
			if tilted[row][column] != "O" {
				continue
			}
			tilted[row][column] = "."
			for potentialEmpty := row; potentialEmpty <= len(tilted) - 1; potentialEmpty++ {
				if potentialEmpty == len(tilted) - 1 {
					tilted[potentialEmpty][column] = "O"
					break
				}
				if tilted[potentialEmpty+1][column] != "." {
					tilted[potentialEmpty][column] = "O"
					break
				}
			}
		}
	}

	result := make([]string, len(platform))
	for i := 0; i < len(tilted); i++ {
		result[i] = strings.Join(tilted[i], "")
	}
	return result
}

func tiltWest(platform []string) []string {
	tilted := make([][]string, len(platform))
	for i := 0; i < len(platform); i++ {
		tilted[i] = strings.Split(platform[i], "")
	}

	for row := 0; row < len(tilted); row++ {
		for column := 1; column < len(tilted[row]); column++ {
			if tilted[row][column] != "O" {
				continue
			}
			tilted[row][column] = "."
			for potentialEmpty := column; potentialEmpty >= 0; potentialEmpty-- {
				if potentialEmpty == 0 {
					tilted[row][potentialEmpty] = "O"
					break
				}
				if tilted[row][potentialEmpty-1] != "." {
					tilted[row][potentialEmpty] = "O"
					break
				}
			}
		}
	}

	result := make([]string, len(tilted))
	for i := 0; i < len(tilted); i++ {
		result[i] = strings.Join(tilted[i], "")
	}
	return result
}

func tiltEast(platform []string) []string {
	tilted := make([][]string, len(platform))
	for i := 0; i < len(platform); i++ {
		tilted[i] = strings.Split(platform[i], "")
	}

	for row := 0; row < len(tilted); row++ {
		for column := len(tilted[row]) - 2; column >= 0; column-- {
			if tilted[row][column] != "O" {
				continue
			}
			tilted[row][column] = "."
			for potentialEmpty := column; potentialEmpty <= len(tilted) - 1; potentialEmpty++ {
				if potentialEmpty == len(tilted[row]) - 1 {
					tilted[row][potentialEmpty] = "O"
					break
				}
				if tilted[row][potentialEmpty+1] != "." {
					tilted[row][potentialEmpty] = "O"
					break
				}
			}
		}
	}

	result := make([]string, len(platform))
	for i := 0; i < len(tilted); i++ {
		result[i] = strings.Join(tilted[i], "")
	}
	return result
}


func calculateLoad(platform []string) int {
	load := 0
	for i := 0; i < len(platform); i++ {
		count := strings.Count(platform[i], "O")
		load += count * (len(platform) - i)
	}
	return load
}


type cachedPlatform struct {
	platform []string
	cachedAt int
}

func runCycles(platform []string, cycles int) []string {
	cycled := make([]string, len(platform))
	copy(cycled, platform)
	cache := map[string]cachedPlatform{}
	directions := []string{"n", "w", "s", "e"}
	directionMap := map[string]func([]string)[]string{
		"n": tiltNorth,
		"s": tiltSouth,
		"w": tiltWest,
		"e": tiltEast,
	}
	cyclesRemaining := cycles
	for i := 0; i < cyclesRemaining; i++ {
		cycleKey := buildCacheKey(cycled, "cycle")
		if cached, ok := cache[cycleKey]; ok {
			cyclesRemaining = i + ((cycles - cached.cachedAt) % (i - cached.cachedAt))
			cycled = cached.platform
			continue
		}

		for _, direction := range directions {
			key := buildCacheKey(cycled, direction)
			if _, ok := cache[key]; ok {
				cycled = cache[key].platform
				continue
			}
			cycled = directionMap[direction](cycled)
			cache[key] = cachedPlatform{ platform: cycled, cachedAt: i }
		}
		cache[cycleKey] = cachedPlatform{ platform: cycled, cachedAt: i }
	}
	return cycled
}

func buildCacheKey(platform []string, direction string) string {
	return strings.Join(platform, "\n") + direction
}