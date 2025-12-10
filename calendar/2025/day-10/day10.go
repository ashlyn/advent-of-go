package main

import (
	"advent-of-go/utils/files"
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
		presses := pressJoltageButtons(machine)
		result += presses
		fmt.Println(presses)
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

func pressButtonForJoltage(joltage []int, button []int) []int {
	next := make([]int, len(joltage))
	copy(next, joltage)
	for _, index := range button {
		next[index]++
	}
	return next
}

func getStartingLightsForPattern(pattern string) string {
	startingLights := ""
	for i := 0; i < len(pattern); i++ {
		startingLights += "."
	}
	return startingLights
}

func getStartingJoltageForRequirements(requirements []int) []int {
	startingJoltage := make([]int, len(requirements))
	for i := 0; i < len(requirements); i++ {
		startingJoltage[i] = 0
	}
	return startingJoltage
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

func key(s []int) string {
	return fmt.Sprint(s)
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func exceedsJoltageRequirements(joltage []int, requirements []int) bool {
	for i := 0; i < len(joltage); i++ {
		if joltage[i] > requirements[i] {
			return true
		}
	}
	return false
}

func pressJoltageButtons(machine machine) int {
	currentJoltage := getStartingJoltageForRequirements(machine.joltageRequirements)

	fewestPresses := map[string]int{
		key(currentJoltage): 0,
		key(machine.joltageRequirements): math.MaxInt,
	}

	queue := [][]int{currentJoltage}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if equals(current, machine.joltageRequirements) {
			continue
		}

		if exceedsJoltageRequirements(current, machine.joltageRequirements) {
			continue
		}

		for _, button := range machine.buttons {
			nextJoltage := pressButtonForJoltage(current, button)
			nextKey := key(nextJoltage)
			currentKey := key(current)
			if fewestPresses[nextKey] == 0 || fewestPresses[nextKey] > fewestPresses[currentKey] + 1 {
				fewestPresses[nextKey] = fewestPresses[currentKey] + 1
				queue = append(queue, nextJoltage)
			}
		}
	}

	return fewestPresses[key(machine.joltageRequirements)]
}

