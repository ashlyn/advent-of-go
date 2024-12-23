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

func solvePart2(input []string) string {
	matrix := parseAsMatrix(input)
	adjacencyList := parseAsAdjacencyList(input)
	degrees := map[int][]string{}
	maxCliqueSize := 0
	for n, connections := range adjacencyList {
		n, degree := n, len(connections)
		if degree > maxCliqueSize {
			maxCliqueSize = degree
		}
		degrees[degree] = append(degrees[degree], n)
	}
	nodes := []string{}
	for node := range matrix {
		nodes = append(nodes, node)
	}
	for cliqueSize := maxCliqueSize; cliqueSize > 0; cliqueSize-- {
		for _, parent := range degrees[maxCliqueSize] {
			neighbors := adjacencyList[parent]
			nodeSet := make([]string, cliqueSize)
			copy(nodeSet, neighbors)
			nodeSet = append(nodeSet, parent)
			combinations := slices.GenerateCombinationsLengthN(nodeSet, cliqueSize + 1)
			for _, c := range combinations {
				if isClique(matrix, c) {
					sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
					return strings.Join(c, ",")
				}
			}	
		}
	}
	return ""
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

func parseAsMatrix(input []string) map[string]map[string]int {
	matrix := make(map[string]map[string]int)

	for _, line := range input {
		computers := strings.Split(line, "-")
		if _, ok := matrix[computers[0]]; !ok {
			matrix[computers[0]] = make(map[string]int)
		}
		matrix[computers[0]][computers[1]] = 1
		if _, ok := matrix[computers[1]]; !ok {
			matrix[computers[1]] = make(map[string]int)
		}
		matrix[computers[1]][computers[0]] = 1
	}

	return matrix
}

func isClique(matrix map[string]map[string]int, nodes []string) bool {
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if _, ok := matrix[nodes[i]][nodes[j]]; !ok {
				return false
			}
		}
	}

	return true
}
