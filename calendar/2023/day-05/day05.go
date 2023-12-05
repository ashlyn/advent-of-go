package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/ranges"
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
	seeds, almanac := parseInput(input, false)

	allDestinations := computeLocations(almanac, seeds)

	min := math.MaxInt
	for location := range allDestinations[len(allDestinations)-1] {
		if location < min {
			min = location
		}
	}
	return min
}

func solvePart2(input []string) int {
	seeds, almanac := parseInput(input, true)

	min := math.MaxInt
	for i := 0; i < len(seeds); i++ {
		allDestinations := computeLocations(almanac, seeds)
		for location := range allDestinations[len(allDestinations)-1] {
			if location < min {
				min = location
			}
		}
	}
	return min
}

func computeLocations(almanac [][]*almanacMap, seeds []ranges.Range) (allDestinations []map[int]int) {
	for i := 0; i < len(almanac)+1; i++ {
		allDestinations = append(allDestinations, make(map[int]int))
	}
	for _, s := range seeds {
		allDestinations[0][s.Start] = -1
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

func parseSeeds(seedsLine string, seedsAsRanges bool) []ranges.Range {
	fields := strings.Fields(seedsLine)
	seedValues := slices.ParseIntsFromStrings(fields[1:])
	seeds := []ranges.Range{}
	for i := 0; i < len(seedValues); {
		if seedsAsRanges {
			seeds = append(seeds, ranges.NewWithLength(seedValues[i], seedValues[i+1]))
			i += 2
		} else {
			seeds = append(seeds, ranges.NewWithLength(seedValues[i], 1))
			i++
		}
	}
	return seeds;
}

func parseMap(mapLine string) almanacMap {
	values := slices.ParseIntsFromStrings(strings.Fields(mapLine))
	return almanacMap{
		destinationStart: values[0],
		sourceStart: values[1],
		rangeLength: values[2],
	}
}

func parseInput(input []string, seedsAsRanges bool) (seeds []ranges.Range, maps [][]*almanacMap) {
	seeds = parseSeeds(input[0], seedsAsRanges)
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

