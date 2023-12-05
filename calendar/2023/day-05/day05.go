package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"math"
	"strings"
)

type almanacMap struct {
	destinationStart int
	sourceStart int
	rangeLength int
}

func main() {
	input := files.ReadFile(5, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	seeds, almanac := parseInput(input)

	allDestinations := getMinimumLocation(almanac, seeds)

	min := math.MaxInt
	for location := range allDestinations[len(allDestinations)-1] {
		if location < min {
			min = location
		}
	}
	return min
}

func solvePart2(input []string) int {
	seedRange, almanac := parseInput(input)

	min := math.MaxInt
	for i := 0; i < len(seedRange); i += 2 {
		println(i, seedRange[i], seedRange[i+1])
		seeds := make([]int, seedRange[i+1])
		for j := 0; j < seedRange[i+1]; j++ {
			seeds[j] = seedRange[i] + j
		}
		allDestinations := getMinimumLocation(almanac, seeds)
		for location := range allDestinations[len(allDestinations)-1] {
			if location < min {
				min = location
			}
		}
	}
	return min
}

func getMinimumLocation(almanac [][]*almanacMap, seeds []int) (allDestinations []map[int]int) {
	for i := 0; i < len(almanac)+1; i++ {
		allDestinations = append(allDestinations, make(map[int]int))
	}
	for _, s := range seeds {
		allDestinations[0][s] = -1
	}

	for round, a := range almanac {
		for sourceValue, destinationValue := range allDestinations[round] {
			if destinationValue == -1 {
				dest := getDestination(sourceValue, a)
				allDestinations[round][sourceValue] = dest
				allDestinations[round+1][dest] = -1
			}
		}
	}

	return
}

func parseSeeds(seedsLine string) []int {
	fields := strings.Fields(seedsLine)
	return slices.ParseIntsFromStrings(fields[1:])
}

func parseMap(mapLine string) almanacMap {
	values := slices.ParseIntsFromStrings(strings.Fields(mapLine))
	return almanacMap{
		destinationStart: values[0],
		sourceStart: values[1],
		rangeLength: values[2],
	}
}

func parseInput(input []string) (seeds []int, maps [][]*almanacMap) {
	seeds = parseSeeds(input[0])
	currentMapSet := []*almanacMap{}
	for i := 3; i < len(input); i++ {
		if input[i] == "" {
			maps = append(maps, currentMapSet)
			currentMapSet = []*almanacMap{}
			i++
			continue
		}

		nextMap := parseMap(input[i])
		currentMapSet = append(currentMapSet, &nextMap)
	}

	maps = append(maps, currentMapSet)
	return
}

func getDestination(source int, maps []*almanacMap) int {
	for _, m := range maps {
		if source >= m.sourceStart && source < m.sourceStart + m.rangeLength {
			return m.destinationStart + (source - m.sourceStart)
		}
	}

	return source
}

