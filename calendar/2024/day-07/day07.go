package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(7, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	equations := parseInput(input)
	for _, eq := range equations {
		if validateEquation(eq, false) {
			result += eq.testValue
		}
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	equations := parseInput(input)
	for _, eq := range equations {
		if validateEquation(eq, true) {
			result += eq.testValue
		}
	}

	return result
}

type equation struct {
	testValue int
	values []int
}
func parseInput(input []string) []equation {
	count := len(input)
	equations := make([]equation, count)

	for i := 0; i < count; i++ {
		parts := strings.Fields(input[i])
		testValue, _ := strconv.Atoi(strings.Trim(parts[0], ":"))
		values := slices.ParseIntsFromStrings(parts[1:])
		equations[i] = equation{
			testValue: testValue,
			values: values,
		}
	}

	return equations
}

func validateEquation(eq equation, useConcatenation bool) bool {
	return validateEquationRecursive(eq.testValue, eq.values[1:], eq.values[0], useConcatenation)
}

func validateEquationRecursive(testValue int, values []int, total int, useConcatenation bool) bool {
	if len(values) == 0 {
		return testValue == total
	}

	if total > testValue {
		return false
	}

	newTotals := []int{total + values[0], total * values[0]}
	if useConcatenation {
		newTotals = append(newTotals, concatenate(total, values[0]))
	}

	for _, newTotal := range newTotals {
		if validateEquationRecursive(testValue, values[1:], newTotal, useConcatenation) {
			return true
		}
	}
	return false
}

func concatenate(a int, b int) int {
	value, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return value
}
