package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"sort"
	"strings"
)

func main() {
	input := files.ReadFile(23, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	adjacencyList := parseAsAdjacencyList(input)
	networks := [][]string{}
	
	for node, connections := range adjacencyList {
		combinations := slices.GenerateCombinationsLengthN(connections, 2)
		for i := 0; i < len(combinations); i++ {
			pair := combinations[i]
			list1, list2 := adjacencyList[pair[0]], adjacencyList[pair[1]]
			if slices.Contains(list1, pair[1]) && slices.Contains(list1, node) &&
				slices.Contains(list2, pair[0]) && slices.Contains(list2, node) {
					network := []string{node, pair[0], pair[1]}
					sort.Slice(network, func(i, j int) bool { return network[i] < network[j] })
					inList := false
					for _, n := range networks {
						if slices.Equals(network, n) {
							inList = true
							break
						}
					}
					if !inList {
						networks = append(networks, network)
					}
			}
		}
	}

	tNetworks := [][]string{}
	for _, network := range networks {
		for _, node := range network {
			if node[0] == 't' {
				tNetworks = append(tNetworks, network)
				break
			}
		}
	}

	return len(tNetworks)
}

func solvePart2(input []string) int {
	result := 0



	return result
}

func parseAsAdjacencyList(input []string) map[string][]string {
	adjacencyList := make(map[string][]string)

	for _, line := range input {
		computers := strings.Split(line, "-")
		if _, ok := adjacencyList[computers[0]]; !ok {
			adjacencyList[computers[0]] = []string{}
		}
		adjacencyList[computers[0]] = append(adjacencyList[computers[0]], computers[1])
		if _, ok := adjacencyList[computers[1]]; !ok {
			adjacencyList[computers[1]] = []string{}
		}
		adjacencyList[computers[1]] = append(adjacencyList[computers[1]], computers[0])
	}

	return adjacencyList
}
