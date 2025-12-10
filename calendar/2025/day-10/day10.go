package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/str"
	"math"
)

func main() {
	input := files.ReadFile(10, 2025, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	machines := parseInput(input)
	for _, machine := range machines {
		result += mashButtons(machine)
	}

	return result
}

func solvePart2(input []string) int {
	result := 0


	return result
}

type machine struct {
	lightPattern string
	buttons [][]int
	joltageRequirements []int
}

func parseLine(line string) machine {
	lightPattern := str.ParseAllGroupsBetween("[", "]", line)[0]

	buttonMatches := str.ParseAllGroupsBetween("(", ")", line)
	buttons := make([][]int, len(buttonMatches))
	for i := 0; i < len(buttonMatches); i++ {
		buttons[i] = str.ParseDelimetedStringToInts(buttonMatches[i], ",")
	}

	joltageMatches := str.ParseAllGroupsBetween("{", "}", line)
	joltageRequirements := str.ParseDelimetedStringToInts(joltageMatches[0], ",")

	return machine{
		lightPattern: lightPattern,
		buttons: buttons,
		joltageRequirements: joltageRequirements,
	}
}

func parseInput(input []string) []machine {
	machines := make([]machine, len(input))
	for i := 0; i < len(input); i++ {
		machines[i] = parseLine(input[i])
	}
	return machines
}

func pressButton(lights string, button []int) string {
	next := []rune(lights)
	for _, lightIndex := range button {
		if next[lightIndex] == '#' {
			next[lightIndex] = '.'
		} else {
			next[lightIndex] = '#'
		}
	}
	return string(next)
}

func getStartingLightsForPattern(pattern string) string {
	startingLights := ""
	for i := 0; i < len(pattern); i++ {
		startingLights += "."
	}
	return startingLights
}

func mashButtons(machine machine) int {
	currentLights := getStartingLightsForPattern(machine.lightPattern)
	fewestPresses := map[string]int{
		currentLights: 0,
		machine.lightPattern: math.MaxInt,
	}

	queue := []string{ currentLights }
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == machine.lightPattern {
			continue
		}

		for _, button := range machine.buttons {
			next := pressButton(current, button)
			if fewestPresses[next] == 0 || fewestPresses[next] > fewestPresses[current] + 1 {
				fewestPresses[next] = fewestPresses[current] + 1
				queue = append(queue, next)
			}
		}
	}

	return fewestPresses[machine.lightPattern]
}
