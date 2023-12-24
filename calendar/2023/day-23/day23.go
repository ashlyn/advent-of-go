package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/sets"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	input := files.ReadFile(23, 2023, "\n")
	sw := time.Now()
	println(solvePart1(input))
	fmt.Printf("Solved part 1 in %v\n", time.Since(sw))
	sw = time.Now()
	println(solvePart2(input))
	fmt.Printf("Solved part 2 in %v\n", time.Since(sw))
}

func solvePart1(input []string) int {
	return findLongestPath(input)
}

func solvePart2(input []string) int {
	withoutSlopes := make([]string, len(input))
	slopePattern := regexp.MustCompile(`[<>^v]`)
	for i, line := range input {
		withoutSlopes[i] = slopePattern.ReplaceAllString(line, ".")
	}
	graph := buildGraph(withoutSlopes)
	return traverseGraph(graph)
}

type queueItem struct {
	current grid.Coords
	visited *sets.Set
}
var directions = []grid.Coords{ { X: 0, Y: -1 }, { X: 1, Y: 0 }, { X: 0, Y: 1 }, { X: -1, Y: 0 } }
func findLongestPath(input []string) int {
	start, target := findStartAndTargetCoords(input)
	visited, walkable := sets.New(), findWalkable(input)
	visited.Add(start.ToString())

	longestPath := -1

	queue := []queueItem{ { current: start, visited: &visited } }
	for len(queue) > 0 {
		currentItem := queue[0]
		queue = queue[1:]
		if currentItem.current == target && currentItem.visited.Size() - 1 > longestPath {
			longestPath = currentItem.visited.Size() - 1
			continue
		}
		for _, d := range directions {
			nextCoords := grid.Coords{ X: currentItem.current.X + d.X, Y: currentItem.current.Y + d.Y }
			nextKey := nextCoords.ToString()
			if walkable.Has(nextKey) && !currentItem.visited.Has(nextKey) {
				nextCharacter := input[nextCoords.Y][nextCoords.X]
				if (nextCharacter == '.' ||
					(nextCharacter == '>' && d.X == 1) ||
					(nextCharacter == '<' && d.X == -1) ||
					(nextCharacter == '^' && d.Y == -1) ||
					(nextCharacter == 'v' && d.Y == 1)) {
					newVisited := currentItem.visited.Copy()
					newVisited.Add(nextKey)
					queue = append(queue, queueItem{ current: nextCoords, visited: &newVisited })
				}
			}
		}

	}

	return longestPath
}

func findWalkable(input []string) *sets.Set {
	walkable := sets.New()

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			currentChar := input[y][x]
			currentCoords := grid.Coords{ X: x, Y: y}
			key := currentCoords.ToString()
			if strings.ContainsRune("<>^v.", rune(currentChar)) {
				walkable.Add(key)
			}
		}
	}
	return &walkable
}

func findStartAndTargetCoords(input []string) (grid.Coords, grid.Coords) {
	start := grid.Coords{ X: strings.Index(input[0], "."), Y: 0 }
	target := grid.Coords{ X: strings.Index(input[len(input) - 1], "."), Y: len(input) - 1 }
	return start, target
}

func buildGraph(input []string) [][]int {
	start, target := findStartAndTargetCoords(input)
	walkable := findWalkable(input)

	adj := map[grid.Coords][]grid.Coords{}
	nodeIndex := 0
	intersectionNodeIndexMap := map[grid.Coords]int{
		start: nodeIndex,
	}
	nodeIndex++
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			currentCoords := grid.Coords{ X: x, Y: y}
			if !walkable.Has(currentCoords.ToString()) {
				continue
			}
			adj[currentCoords] = []grid.Coords{}
			for _, d := range directions {
				nextCoords := grid.Coords{ X: x + d.X, Y: y + d.Y }
				if walkable.Has(nextCoords.ToString()) {
					adj[currentCoords] = append(adj[currentCoords], nextCoords)
				}
			}

			if len(adj[currentCoords]) > 2 {
				intersectionNodeIndexMap[currentCoords] = nodeIndex
				nodeIndex++
			}
		}
	}
	nodeIndex++
	intersectionNodeIndexMap[target] = nodeIndex

	graph := make([][]int, nodeIndex + 1)
	for i := range graph {
		graph[i] = make([]int, nodeIndex + 1)
	}
	for c := range intersectionNodeIndexMap {
		buildEdges(c, intersectionNodeIndexMap, adj, graph)
	}
	return graph
}

type queueEdge struct {
	current grid.Coords
	distance int
}
func buildEdges(coords grid.Coords, nodes map[grid.Coords]int, adj map[grid.Coords][]grid.Coords, graph [][]int) {
	currentIndex, queue, visited := nodes[coords], []queueEdge{{ current: coords, distance: 1 }}, sets.New()
	visited.Add(coords.ToString())

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		for _, neighbor := range adj[current.current] {
			if !visited.Has(neighbor.ToString()) {
				neighborIndex, ok := nodes[neighbor]
				if ok {
					graph[currentIndex][neighborIndex] = current.distance
					graph[neighborIndex][currentIndex] = current.distance
				} else {
					queue = append(queue, queueEdge{ current: neighbor, distance: current.distance + 1 })
				}
				visited.Add(neighbor.ToString())
			}
		}
	}
}

func traverseGraph(graph [][]int) int {
	end := len(graph) - 1
	v := sets.New()
	return traverseGraphRecursively(0, 0, end, &v, graph)
}

func traverseGraphRecursively(currentNodeIndex int, previousNodeIndex int, endNodeIndex int, visisted *sets.Set, graph [][]int) int {
	if currentNodeIndex == endNodeIndex {
		return 0
	}
	maxPathLength := -1
	for i := 0; i <= endNodeIndex; i++ {
		key := fmt.Sprintf("%d", i)
		if graph[currentNodeIndex][i] != 0 && i != previousNodeIndex && !visisted.Has(key) {
			visistedCopy := visisted.Copy()
			visistedCopy.Add(key)
			d := traverseGraphRecursively(i, currentNodeIndex, endNodeIndex, &visistedCopy, graph)
			if d != -1 {
				maxPathLength = maths.Max(maxPathLength, d + graph[currentNodeIndex][i])
			}
		}
	}
	return maxPathLength
}
