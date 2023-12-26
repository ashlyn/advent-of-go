package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/priorityqueue"
	"advent-of-go/utils/sets"
	"advent-of-go/utils/slices"
	"container/heap"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := files.ReadFile(25, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	graph, source := parseInput(input)
	findGroupSizes(graph, source)

	// naive solution, could optimize for "weakest links" in the graph
	pairSet := sets.New()
	for key, connected := range graph {
		for _, v := range connected {
			pairSet.Add(getKey(key, v))
		}
	}
	combinations := slices.GenerateCombinationsLengthNGeneric(pairSet.Iterator(), 3)
	for _, combination := range combinations {
		graphCopy := copyGraph(graph)
		for _, c := range combination {
			graphCopy = removeLink(graphCopy, parseKey(c))
		}
		group1, group2 := findGroupSizes(graphCopy, "bvb")
		if group1 != 0 && group2 != 0 {
			fmt.Printf("\nGraph: %v\n\n", graph)
			fmt.Printf("Graph: %v\n", graphCopy)
			fmt.Printf("Group 1: %d, Group 2: %d\n", group1, group2)
			fmt.Printf("Combination: %v\n", combination)
			return group1 * group2
		}
	}

	return 0
}

func removeLink(graph map[string][]string, link []string) map[string][]string {
	component1, component2 := link[0], link[1]
	graphCopy := copyGraph(graph)
	i, j := slices.IndexOf(component1, graphCopy[component2]), slices.IndexOf(component2, graphCopy[component1])
	graphCopy[component2] = append(graphCopy[component2][:i], graphCopy[component2][i+1:]...)
	graphCopy[component1] = append(graphCopy[component1][:j], graphCopy[component1][j+1:]...)
	return graphCopy
}

func parseKey(key string) []string {
	parts := strings.Fields(key)
	return []string{ parts[0], parts[1]}
}

func getKey(component1, component2 string) string {
	if component1 < component2 {
		return fmt.Sprintf("%s %s", component1, component2)
	}
	return fmt.Sprintf("%s %s", component2, component1)
}

func copyGraph(graph map[string][]string) map[string][]string {
	newGraph := map[string][]string{}
	for key, values := range graph {
		newGraph[key] = append([]string{}, values...)
	}
	return newGraph
}

func solvePart2(input []string) int {
	result := 0



	return result
}

func parseInput(input []string) (map[string][]string, string) {
	componentGraph := map[string][]string{}
	source := ""
	for _, line := range input {
		c, connected := parseComponent(line)
		if source == "" {
			source = c
		}
		if _, ok := componentGraph[c]; !ok {
			componentGraph[c] = []string{}
		}
		for _, v := range connected {
			if !slices.Contains(componentGraph[c], v) {
				componentGraph[c] = append(componentGraph[c], v)
			}
			if _, ok := componentGraph[v]; !ok {
				componentGraph[v] = []string{}
			}
			if !slices.Contains(componentGraph[v], c) {
				componentGraph[v] = append(componentGraph[v], c)
			}
		}
	}
	return componentGraph, source
}

func parseComponent(line string) (string, []string) {
	parts := strings.Fields(line)
	return parts[0][:len(parts[0])-1], parts[1:]
}

func dijkstra(graph map[string][]string, source string) map[string]int {
	dist := map[string]int{}
	queue := make(priorityqueue.PriorityQueue, 0)
	for key, values := range graph {
		if key == source {
			dist[key] = 0
		} else {
			dist[key] = math.MaxInt
		}
		if !queue.Has(key) {
			heap.Push(&queue, &priorityqueue.Item{Value: key, Priority: dist[key]})
		}

		for _, v := range values {
			if v == source {
				dist[v] = 0
			} else {
				dist[v] = math.MaxInt
			}
			if !queue.Has(v) {
				heap.Push(&queue, &priorityqueue.Item{Value: v, Priority: dist[v]})
			}
		}
	}

	for queue.Len() > 0 {
		u := heap.Pop(&queue).(*priorityqueue.Item)
		for _, v := range graph[u.Value] {
			if queue.Has(v) {
				alt := dist[u.Value] + 1
				if alt < dist[v] {
					dist[v] = alt
					queue.Update(v, alt)
				}
			}
		}
	}

	return dist
}

func findGroupSizes(graph map[string][]string, source string) (int, int) {
	dist := dijkstra(graph, source)

	group1, group2 := 0, 0
	for _, v := range dist {
		if v == math.MaxInt || v == math.MinInt {
			group2++
		} else {
			group1++
		}
	}

	return group1, group2
}

func findShortestPath(graph map[string][]string, source, dest string) int {
	dist := dijkstra(graph, source)
	return dist[dest]
}	
