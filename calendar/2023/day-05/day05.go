package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/ranges"
	"advent-of-go/utils/slices"
	"math"
	"strings"
)

type almanacMap struct {
	destinationRange ranges.Range
	sourceRange      ranges.Range
}

func main() {
	input := files.ReadFile(5, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	seeds, almanac := parseInput(input, false)
	return getMinimumLocations(almanac, seeds)
}

func solvePart2(input []string) int {
	seeds, almanac := parseInput(input, true)

	return getMinimumLocations(almanac, seeds)
}

func getMinimumLocations(almanac [][]*almanacMap, seeds []ranges.Range) int {
	// keep seeds/current values (e.g. turn to "soil" as ranges)
	// split ranges based on current transformation map

	currentValues := make([]ranges.Range, len(seeds))
	for i := range seeds {
		currentValues[i] = seeds[i]
	}

	for round := range almanac {
		nextValues := []ranges.Range{}

		for i := 0; i < len(currentValues); i++ {
			current := currentValues[i]
			hasOverlap := false
			for _, a := range almanac[round] {
				if !current.Overlaps(a.sourceRange) {
					continue
				}
				hasOverlap = true
				splits := current.SplitOn(a.sourceRange)

				for _, split := range splits {
					if (a.sourceRange.ContainsRange(split)) {
						delta := a.sourceRange.Start - a.destinationRange.Start
						transformed := ranges.New(split.Start - delta, split.End - delta)
						nextValues = append(nextValues, transformed)
					} else {
						currentValues = append(currentValues, split)
					}
				}
			}
			if !hasOverlap {
				nextValues = append(nextValues, current)
			}
		}
		currentValues = make([]ranges.Range, len(nextValues))
		copy(currentValues, nextValues)
	}

	min := math.MaxInt
	for _, v := range currentValues {
		min = maths.Min(min, v.Start)
	}

	return min
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
		destinationRange: ranges.NewWithLength(values[0], values[2]),
		sourceRange:      ranges.NewWithLength(values[1], values[2]),
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
		if m.sourceRange.Contains(source) {
			return m.destinationRange.Start + (source - m.sourceRange.Start)
		}
	}

	return source
}

