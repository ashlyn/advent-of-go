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
	adjacencyList := parseInput(input)

	// Valid paths take one of two forms:
	//         / ① --> dac --> ② --> fft --> ③ \
	//  svr -->                                   --> out
	//         \ ① --> fft --> ② --> dac --> ③ /

	// first segments
	svrToDacPaths := countPathsDfs(adjacencyList, "svr", "dac")
	svrToFftPaths := countPathsDfs(adjacencyList, "svr", "fft")

	// middle segments
	dacToFftPaths := countPathsDfs(adjacencyList, "dac", "fft")
	fftToDacPaths := countPathsDfs(adjacencyList, "fft", "dac")

	// end segments
	dacToOutPaths := countPathsDfs(adjacencyList, "dac", "out")
	fftToOutPaths := countPathsDfs(adjacencyList, "fft", "out")

	dacThenFftPaths := svrToDacPaths * dacToFftPaths * fftToOutPaths
	fftThenDacPaths := svrToFftPaths * fftToDacPaths * dacToOutPaths

	return dacThenFftPaths + fftThenDacPaths
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

func countPathsDfs(adjacencyList map[string][]string, start, end string) int {
	pathCountFromKeyToEnd := map[string]int{}
	for device, outputs := range adjacencyList {
		pathCountFromKeyToEnd[device] = -1
		for _, output := range outputs {
			pathCountFromKeyToEnd[output] = -1
		}
	}
	pathCountFromKeyToEnd[end] = 1

	for pathsFromStart, _ := pathCountFromKeyToEnd[start]; pathsFromStart == -1; pathsFromStart, _ = pathCountFromKeyToEnd[start] {
		viablePathFound := false
		for device, pathCount := range pathCountFromKeyToEnd {
			if pathCount == -1 {
				pathsFromAllOutputsExplored := true
				sum := 0
				for _, output := range adjacencyList[device] {
					if pathCountFromKeyToEnd[output] != -1 {
						sum += pathCountFromKeyToEnd[output]
					} else {
						pathsFromAllOutputsExplored = false
					}
				}
				if pathsFromAllOutputsExplored {
					pathCountFromKeyToEnd[device] = sum
					viablePathFound = true
				}
			}
		}
		// If no viable path was found in this iteration, break as there is no path from start to end
		if !viablePathFound {
			return 0
		}
	}

	return pathCountFromKeyToEnd[start]
}
