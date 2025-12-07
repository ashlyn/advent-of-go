package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(6, 2025, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	problems := parseInputPart1(input)
	for _, problem := range problems {
		result += problem.solve()
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	problems := parseInputPart2(input)
	for _, problem := range problems {
		result += problem.solve()
	}

	return result
}

type problem struct {
	numbers []int
	operation string
}

func parseInputPart1(input []string) []problem {
	numbers := invert(parseNumbers(input[:len(input) - 1]))
	problems := make([]problem, len(numbers))
	operations := strings.Fields(input[len(input) - 1])
	for i := 0; i < len(numbers); i++ {
		problems[i] = problem{
			numbers: numbers[i],
			operation: operations[i],
		}
	}
	return problems
}

func parseInputPart2(input []string) []problem {
	operatorsLine := input[len(input) - 1]
	ranges := [][2]int{}
	problemCounter, start := 0, 0
	for end := 1; end < len(operatorsLine); end++ {
		if operatorsLine[end] == '+' || operatorsLine[end] == '*' {
			ranges = append(ranges, [2]int{start, end - 2})
			start = end
		} else {
			problemCounter++
		}
	}
	ranges = append(ranges, [2]int{start, len(operatorsLine) - 1})

	problems := make([]problem, len(ranges))
	for i := 0; i < len(ranges); i++ {
		asStrings := make([]string, len(input) - 1)
		for j := 0; j < len(input) - 1; j++ {
			asStrings[j] = input[j][ranges[i][0]:ranges[i][1] + 1]
		}
		numbers := parseRtlColumnNumbers(asStrings)
		operation := string(operatorsLine[ranges[i][0]])
		problems[i] = problem{
			numbers: numbers,
			operation: operation,
		}
	}
	return problems
}

func parseNumbers(input []string) [][]int {
	numbers := make([][]int, len(input))
	for i, line := range input {
		fields := strings.Fields(line)
		nums := make([]int, len(fields))
		for j, field := range fields {
			value, _ := strconv.Atoi(field)
			nums[j] = value
		}
		numbers[i] = nums
	}
	return numbers
}

func invert[T comparable](input [][]T) [][]T {
	inverted := make([][]T, len(input[0]))
	for x := 0; x < len(input[0]); x++ {
		newRow := make([]T, len(input))
		for y := 0; y < len(input); y++ {
			newRow[y] = input[y][x]
		}
		inverted[x] = newRow
	}
	return inverted
}

func (p *problem) solve() int {
	if p.operation == "+" {
		return slices.Sum(p.numbers)
	}
	if p.operation == "*" {
		return slices.Product(p.numbers)
	}
	return 0
}

func parseRtlColumnNumbers(asStrings []string) []int {
	numbers := make([]int, len(asStrings[0]))
	for i := 0; i < len(asStrings[0]); i++ {
		columnDigits := ""
		for j := 0; j < len(asStrings); j++ {
			digit := asStrings[j][len(asStrings[j]) - 1 - i]
			if digit != ' ' {
				columnDigits += string(digit)
				continue
			}
		}
		columnValue, _ := strconv.Atoi(columnDigits)
		numbers[i] = columnValue
	}

	return numbers
}
