package main

import (
	"advent-of-go/utils/files"
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
	return solve(instructions, nodes)
}

func solvePart2(input []string) int {
	result := 0



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
	nodePattern := regexp.MustCompile(`[A-Z]+`)
	parts := nodePattern.FindAllString(line, -1)
	return parts[0], node{ left: parts[1], right: parts[2]}
}

func solve(instructions string, nodes map[string]node) int {
	current, steps := "AAA", 0
	for current != "ZZZ" {
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
