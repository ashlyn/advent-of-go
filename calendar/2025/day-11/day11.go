package main

import (
	"advent-of-go/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(11, 2025, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	adjacencyList := parseInput(input)
	validPaths := bfs(adjacencyList, "you", "out")
	return len(validPaths)
}

func solvePart2(input []string) int {
	result := 0



	return result
}

func parseInput(input []string) map[string][]string {
	adjacencyList := map[string][]string{}
	for _, line := range input {
		parts := strings.Fields(line)
		device := parts[0][0:len(parts[0])-1]
		outputs := parts[1:]
		adjacencyList[device] = outputs
	}
	return adjacencyList
}

func bfs(adjacencyList map[string][]string, start, end string) [][]string {
	validPaths := [][]string{}
	visited := map[string]bool{}

	queue := [][]string{{start}}

	for len(queue) > 0 {
		currentPath := queue[0]
		queue = queue[1:]
		device := currentPath[len(currentPath)-1]

		if device == end {
			validPaths = append(validPaths, currentPath)
			continue
		}

		pathKey := strings.Join(currentPath, "-")
		if !visited[pathKey] {
			visited[pathKey] = true
			for _, output := range adjacencyList[device] {
				if !strings.Contains(pathKey, output) {
					newPath := make([]string, len(currentPath))
					copy(newPath, currentPath)
					newPath = append(newPath, output)
					queue = append(queue, newPath)
				}
			}
		}
	}

	return validPaths
}
