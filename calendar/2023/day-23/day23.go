package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/priorityqueue"
	"advent-of-go/utils/sets"
	"container/heap"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func main() {
	input := files.ReadFile(23, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	return -1
	return findLongest(input)
}

func solvePart2(input []string) int {
	withoutSlopes := make([]string, len(input))
	slopePattern := regexp.MustCompile(`[<>^v]`)
	for i, line := range input {
		withoutSlopes[i] = slopePattern.ReplaceAllString(line, ".")
	}
	start, end, graph := buildAdjacencyList(withoutSlopes)
	fmt.Printf("%v %v\n%v\n", start, end, graph)

	return bruteForceDag(start, end, graph)
	findLongestPathInDag(start, end, graph)
	return -1

	return findLongest(withoutSlopes)
}

var directions = []grid.Coords{ { X: 0, Y: -1 }, { X: 1, Y: 0 }, { X: 0, Y: 1 }, { X: -1, Y: 0 } }
func longestPath(input []string) int {
	dist := map[string]int{}
	queue := make(priorityqueue.PriorityQueue, 0)
	var target grid.Coords

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			currentChar := input[y][x]
			currentCoords := grid.Coords{ X: x, Y: y}
			key := currentCoords.ToString()
			if strings.ContainsRune("<>^v.", rune(currentChar)) {
				if y != 0 {
					dist[key] = math.MaxInt
				} else {
					dist[key] = 0
				}
				if y == len(input) - 1 {
					target = currentCoords
				}
				heap.Push(&queue, &priorityqueue.Item{ Value: key, Priority: dist[key] })
			}
		}
	}

	for queue.Len() > 0 {
		currentItem := heap.Pop(&queue).(*priorityqueue.Item)
		currentCoords := grid.ParseCoords(currentItem.Value)
		currentKey := currentCoords.ToString()
		validNeighbors := []grid.Coords{}
		for _, d := range directions {
			nextCoords := grid.Coords{ X: currentCoords.X + d.X, Y: currentCoords.Y + d.Y }
			nextKey := nextCoords.ToString()
			if queue.Has(nextKey) {
				nextCharacter := input[nextCoords.Y][nextCoords.X]
				if nextCharacter == '.' ||
					(nextCharacter == '>' && d.X == 1) ||
					(nextCharacter == '<' && d.X == -1) ||
					(nextCharacter == '^' && d.Y == -1) ||
					(nextCharacter == 'v' && d.Y == 1) {
						alt := dist[currentKey] + 1
						if alt < dist[nextKey] {
							dist[nextKey] = alt
							queue.Update(nextKey, alt)
						}
						validNeighbors = append(validNeighbors, nextCoords)
					}
				if nextCoords == target {
					println("found target", dist[nextKey])
				}
			}
		}
		if len(validNeighbors) > 1 {
			println("branching path", len(validNeighbors))
		}
	}
	
	return  dist[target.ToString()]
}

func buildAdjacencyList(input []string) (grid.Coords, grid.Coords, map[grid.Coords]map[grid.Coords]int) {
	start := grid.Coords{ X: strings.Index(input[0], "."), Y: 0 }
	target := grid.Coords{ X: strings.Index(input[len(input) - 1], "."), Y: len(input) - 1 }
	walkable := sets.New()
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			currentChar := input[y][x]
			currentCoords := grid.Coords{ X: x, Y: y}
			key := currentCoords.ToString()
			if strings.ContainsRune("<>^v.", rune(currentChar)) {
				walkable.Add(key)
				// queue = append(queue, currentCoords)
			}
		}
	}

	nodes := map[grid.Coords][]grid.Coords{
		start: { { X: start.X, Y: start.Y + 1 } },
		target: { { X: target.X, Y: target.Y - 1 } },
	}
	v := sets.New()
	v.Add(start.ToString())
	queue := []*queueItem{ { current: start, visited: &v } }
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentCharacter := input[current.current.Y][current.current.X]

		validNeighbors := []grid.Coords{}
		for _, d := range directions {
			nextCoords := grid.Coords{ X: current.current.X + d.X, Y: current.current.Y + d.Y }
			nextKey := nextCoords.ToString()
			if walkable.Has(nextKey) && !current.visited.Has(nextKey) && isValidMove(currentCharacter, d) {
				validNeighbors = append(validNeighbors, nextCoords)
			}
		}
		if len(validNeighbors) != 0 {
			current.visited.Add(validNeighbors[0].ToString())
			queue = append(queue, &queueItem{ current: validNeighbors[0], visited: current.visited })
		}
		if len(validNeighbors) > 1 {
			nodes[current.current] = validNeighbors
			for i := 1; i < len(validNeighbors); i++ {
				newVisited := current.visited.Copy()
				newVisited.Add(validNeighbors[i].ToString())
				queue = append(queue, &queueItem{ current: validNeighbors[i], visited: &newVisited })
			}
		}
	}

	graph := map[grid.Coords]map[grid.Coords]int{}
	for node, neighbors := range nodes {
		graph[node] = map[grid.Coords]int{}
		for _, n := range neighbors {
			current := n
			dist := 0
			visited := sets.New()
			visited.Add(current.ToString())
			visited.Add(node.ToString())
			_, isNode := nodes[current]
			currentCharacter := input[current.Y][current.X]
			for !isNode {
				for _, d := range directions {
					nextCoords := grid.Coords{ X: current.X + d.X, Y: current.Y + d.Y }
					if walkable.Has(nextCoords.ToString()) && !visited.Has(nextCoords.ToString()) && isValidMove(currentCharacter, d) {
						dist++
						current = nextCoords
						visited.Add(nextCoords.ToString())
						_, isNode = nodes[current]
						break
					}
				}
			}
			graph[node][current] = dist
		}
	}
	return start, target, graph
}

func isValidMove(currentCharacter byte, direction grid.Coords) bool {
	return currentCharacter == '.' ||
		(currentCharacter == '>' && direction.X == 1) ||
		(currentCharacter == '<' && direction.X == -1) ||
		(currentCharacter == '^' && direction.Y == -1) ||
		(currentCharacter == 'v' && direction.Y == 1)
}

type queueItem2 struct {
	current grid.Coords
	visited *sets.Set
	pathLength int
}
func bruteForceDag(start grid.Coords, end grid.Coords, graph map[grid.Coords]map[grid.Coords]int) int {
	longestPath := math.MinInt
	v := sets.New()
	v.Add(start.ToString())
	queue := []queueItem2{ { current: start, visited: &v, pathLength: 0 } }

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.current == end {
			if current.pathLength > longestPath {
				fmt.Printf("%v %v %v\n", current.current, current.pathLength, current.visited)
				longestPath = current.pathLength
			}
			continue
		}
		for neighbor, weight := range graph[current.current] {
			if !current.visited.Has(neighbor.ToString()) {
				newVisited := current.visited.Copy()
				newVisited.Add(neighbor.ToString())
				queue = append(queue, queueItem2{ current: neighbor, visited: &newVisited, pathLength: current.pathLength + weight })
			}
		}
	}

	return longestPath
}

func findLongestPathInDag(start grid.Coords, end grid.Coords, graph map[grid.Coords]map[grid.Coords]int) int {
	visited := sets.New()
	stack := []grid.Coords{ }
	dist := map[grid.Coords]int{}

	for v := range graph {
		if !visited.Has(v.ToString()) {
			topologicalSort(v, graph, &visited, &stack)
		}
		if v == start {
			dist[v] = 0
		} else {
			dist[v] = math.MinInt
		}
	}

	for len(stack) > 0 {
		current := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if dist[current] != math.MinInt {
			for neighbor, weight := range graph[current] {
				if dist[neighbor] < dist[current] + weight {
					dist[neighbor] = dist[current] + weight
				}
			}
		}
	}

	fmt.Printf("%v\n", dist)
	println(dist[end])

	return -1
}

func topologicalSort(current grid.Coords, graph map[grid.Coords]map[grid.Coords]int, visited *sets.Set, stack *[]grid.Coords) {
	visited.Add(current.ToString())
	for neighbor := range graph[current] {
		if !visited.Has(neighbor.ToString()) {
			topologicalSort(neighbor, graph, visited, stack)
		}
	}
	*stack = append(*stack, current)
}

type queueItem struct {
	current grid.Coords
	visited *sets.Set
}
func findLongest(input []string) int {
	start := grid.Coords{ X: strings.Index(input[0], "."), Y: 0 }
	target := grid.Coords{ X: strings.Index(input[len(input) - 1], "."), Y: len(input) - 1 }
	visited, walkable := sets.New(), sets.New()
	visited.Add(start.ToString())

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

	longestPath := -1

	queue := []queueItem{ { current: start, visited: &visited } }
	for len(queue) > 0 {
		currentItem := queue[0]
		queue = queue[1:]
		if currentItem.current == target && currentItem.visited.Size() - 1 > longestPath {
			println("found target", currentItem.visited.Size() - 1, len(queue))
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
