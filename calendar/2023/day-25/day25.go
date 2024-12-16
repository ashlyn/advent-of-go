package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/priorityqueue"
	"advent-of-go/utils/slices"
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

func main() {
	input := files.ReadFile(25, 2023, "\n")
	sw := time.Now()
	println(solvePart1(input))
	fmt.Printf("Part 1 took %s\n", time.Since(sw))
	println(solvePart2(input))
}

type edgeCount struct {
	edge []string
	count int
}
func solvePart1(input []string) int {
	graph, _ := parseInput(input)
	// Able to reliably get the correct answer by using the top 8 edges after 25 traversals in ~5s
	// lower sample sizes = faster execution times but higher likelihood of missing the correct combination
	return trySplittingGraph(graph, 3, 25, 10)
}

func trySplittingGraph(graph map[string][]string, edgesToRemove, traversalAttempts, useTopNEdges int) int {
	traversalCount := map[string]int{}

	// Run dijkstra n times to get a sample of which edges are traversed the most
	for i := 0; i < traversalAttempts; i++ {
		node := ""
		for key := range graph {
			node = key
			break
		}
		findGroupSizes(graph, node, traversalCount)
	}

	// Take the top n most traversed edges and try removing any random three
	edgeCounts := make([]edgeCount, len(traversalCount))
	i := 0
	for key, count := range traversalCount {
		edgeCounts[i] = edgeCount{edge: parseKey(key), count: count}
		i++
	}
	sort.Slice(edgeCounts, func(i, j int) bool {
		return edgeCounts[i].count < edgeCounts[j].count
	})
	candidateEdges := []string{}
	for _, edgeCount := range edgeCounts[len(edgeCounts) - useTopNEdges:] {
		candidateEdges = append(candidateEdges, getKey(edgeCount.edge[0], edgeCount.edge[1]))
	}
	combinations := slices.GenerateCombinationsLengthNGeneric(candidateEdges, edgesToRemove)
	for _, combination := range combinations {
		graphCopy := copyGraph(graph)
		for _, c := range combination {
			graphCopy = removeLink(graphCopy, parseKey(c))
		}
		group1, group2 := findGroupSizes(graphCopy, parseKey(combination[0])[0], traversalCount)
		if group1 != 0 && group2 != 0 {
			fmt.Printf("Combination: %v\n", combination)
			fmt.Printf("Group 1: %d, Group 2: %d\n", group1, group2)
			return group1 * group2
		}
	}
	return -1
}

func getTwoRandomNodes(graph map[string][]string) (string, string){
	node1, node2 := "", ""
	for key := range graph {
		node1 = key
		break
	}	
	for key := range graph {
		if key != node1 {
			node2 = key
			break
		}
	}
	return node1, node2
}

func solvePart2(input []string) string {
	return "Merry Christmas! 50*"
}

func removeLink(graph map[string][]string, link []string) map[string][]string {
	component1, component2 := link[0], link[1]
	graphCopy := copyGraph(graph)
	i, j := slices.IndexOfStr(component1, graphCopy[component2]), slices.IndexOfStr(component2, graphCopy[component1])
	graphCopy[component2] = append(graphCopy[component2][:i], graphCopy[component2][i+1:]...)
	graphCopy[component1] = append(graphCopy[component1][:j], graphCopy[component1][j+1:]...)
	return graphCopy
}

const separator = "/"
func parseKey(key string) []string {
	parts := strings.Split(key, separator)
	return []string{ parts[0], parts[1]}
}

func getKey(component1, component2 string) string {
	if component1 < component2 {
		return fmt.Sprintf("%s%s%s", component1, separator, component2)
	}
	return fmt.Sprintf("%s%s%s", component2, separator, component1)
}

func copyGraph(graph map[string][]string) map[string][]string {
	newGraph := map[string][]string{}
	for key, values := range graph {
		newGraph[key] = append([]string{}, values...)
	}
	return newGraph
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

func findGroupSizes(graph map[string][]string, source string, traversalCount map[string]int) (int, int) {
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
					key := getKey(u.Value, v)
					traversalCount[key]++
					dist[v] = alt
					queue.Update(v, alt)
				}
			}
		}
	}

	group1, group2 := 0, 0
	for _, v := range dist {
		if v == math.MaxInt || v < 0 {
			group2++
		} else {
			group1++
		}
	}

	return group1, group2
}
