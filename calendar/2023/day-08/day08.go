package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"regexp"
	"strings"
)

type node struct {
	right string
	left string
}

func main() {
	input := files.ReadFile(8, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	instructions, nodes, _ := parseMap(input)
	return solve(instructions, nodes, "AAA")
}

func solvePart2(input []string) int {
	result := 1
	instructions, nodes, endsWithA := parseMap(input)
	for _, start := range endsWithA {
		result = maths.Lcm(result, solve(instructions, nodes, start))
	}
	return result
}

func parseMap(input []string) (string, map[string]node, []string) {
	nodes := map[string]node{}
	endsWithA := []string{}
	for _, line := range input[2:] {
		name, node := parseLine(line)
		if strings.HasSuffix(name, "A") {
			endsWithA = append(endsWithA, name)
		}
		nodes[name] = node
	}
	return input[0], nodes, endsWithA
}

func parseLine(line string) (string, node) {
	nodePattern := regexp.MustCompile(`[A-Z0-9]+`)
	parts := nodePattern.FindAllString(line, -1)
	return parts[0], node{ left: parts[1], right: parts[2]}
}

func solve(instructions string, nodes map[string]node, start string) int {
	current, steps := start, 0
	for !strings.HasSuffix(current, "Z") {
		i := steps % len(instructions)
		if instructions[i] == 'L' {
			current = nodes[current].left
		} else if instructions[i] == 'R' {
			current = nodes[current].right
		}
		steps++
	}

	return steps
}
