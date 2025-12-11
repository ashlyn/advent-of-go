package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/matrices"
	"advent-of-go/utils/slices"
	"advent-of-go/utils/str"
	"fmt"
	"math"
	"sort"
	"time"
)

func main() {
	input := files.ReadFile(10, 2025, "\n")
	sw := time.Now()
	println(solvePart1(input))
	fmt.Printf("Solved part 1 in %v\n", time.Since(sw))
	sw = time.Now()
	println(solvePart2(input))
	fmt.Printf("Solved part 2 in %v\n", time.Since(sw))
}

func solvePart1(input []string) int {
	result := 0

	machines := parseInput(input)
	for _, machine := range machines {
		result += pressLightButtons(machine)
	}

	return result
}

func solvePart2(input []string) int {
	result := 0
	
	machines := parseInput(input)
	for _, machine := range machines {
		m := buildMatrix(machine)
		solution := matrices.SolveMatrix(m, matrices.Minimize, matrices.DefaultLimitConfiguration)
		result += slices.Sum(solution)
	}

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
		buttons[i] = str.ParseDelimitedStringToInts(buttonMatches[i], ",")
	}
	sort.Slice(buttons, func(i, j int) bool {
		return len(buttons[i]) >= len(buttons[j])
	})

	joltageMatches := str.ParseAllGroupsBetween("{", "}", line)
	joltageRequirements := str.ParseDelimitedStringToInts(joltageMatches[0], ",")

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

func pressButtonForLights(lights string, button []int) string {
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

func pressLightButtons(machine machine) int {
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
			next := pressButtonForLights(current, button)
			if fewestPresses[next] == 0 || fewestPresses[next] > fewestPresses[current] + 1 {
				fewestPresses[next] = fewestPresses[current] + 1
				queue = append(queue, next)
			}
		}
	}

	return fewestPresses[machine.lightPattern]
}

func buildMatrix(ma machine) [][]int {
	m, n := len(ma.joltageRequirements), len(ma.buttons)
	matrix := make([][]int, m)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n + 1)
		for j := 0; j < n; j++ {
			incrementsButtonAtI := false
			for _, pos := range ma.buttons[j] {
				if pos == i {
					incrementsButtonAtI = true
					break
				}
			}
			if incrementsButtonAtI {
				matrix[i][j] = 1
			}
		}
		matrix[i][n] = ma.joltageRequirements[i]
	}

	return matrix
}
