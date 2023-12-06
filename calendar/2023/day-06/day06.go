package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"strconv"
	"strings"
)

type race struct {
	timeAllowed int
	distanceToBeat int
}

func main() {
	input := files.ReadFile(6, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	races := parseRaces(input)
	return solveBoatRaces(races)
}

func solvePart2(input []string) int {
	races := []race { parseAsSingleRace(input) }
	return solveBoatRaces(races)
}

func parseRaces(input []string) []race {
	times := slices.ParseIntsFromStrings(strings.Fields(input[0])[1:])
	distances := slices.ParseIntsFromStrings(strings.Fields(input[1])[1:])

	races := make([]race, len(times))
	for i := 0; i < len(times); i++ {
		races[i] = race{ timeAllowed: times[i], distanceToBeat: distances[i] }
	}
	return races
}

func parseAsSingleRace(input []string) race {
	time, _ := strconv.Atoi(strings.Join(strings.Fields(input[0])[1:], ""))
	distance, _ := strconv.Atoi(strings.Join(strings.Fields(input[1])[1:], ""))
	return race{ timeAllowed: time, distanceToBeat: distance }
}

func (r race) runSimulation() []int {
	results := make([]int, r.timeAllowed)
	for timePressed := 0; timePressed < r.timeAllowed; timePressed++ {
		distance := timePressed * (r.timeAllowed - timePressed)
		results[timePressed] = distance
	}
	return results
}

func solveBoatRaces(races []race) int {
	solution := 1

	for _, r := range races {
		results := r.runSimulation()
		wins := 0
		for _, result := range results {
			if result > r.distanceToBeat {
				wins++
			}
		}
		solution *= wins
	}

	return solution
}
