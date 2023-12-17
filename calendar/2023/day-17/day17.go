package main

import (
	"advent-of-go/utils/colors"
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/priorityqueue"
	"advent-of-go/utils/sets"
	"advent-of-go/utils/slices"
	"container/heap"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := files.ReadFile(17, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

// 822 too high
// not 807, 810
func solvePart1(input []string) int {
	blocks := parseInput(input)
	/* minHeatLoss := math.MaxInt
	for i := 0; i < 10; i++ {
		minHeatLoss = maths.Min(minHeatLoss, findShortestPath(blocks))
	}
	return minHeatLoss */
	return dijkstraPriorityQueue(blocks)
}

func solvePart2(input []string) int {
	result := 0



	return result
}

func parseInput(input []string) [][]int {
	result := make([][]int, len(input))
	for i, line := range input {
		strs := strings.Split(line, "")
		result[i] = slices.ParseIntsFromStrings(strs)
	}
	return result
}

type path struct {
	visited *sets.Set
	heatLoss int
	direction grid.Coords
	stepsInDirection int
}
func findShortestPath(blocks [][]int) int {
	shortestPaths := map[string]*path{}
	queue := sets.New()
	winningPaths := []*path{}
	killswitch := 0

	startVisited := sets.New()
	start := grid.Origin
	startVisited.Add(start.ToString())
	startPath := path{
		visited: &startVisited,
		heatLoss: 0,
		direction: grid.Origin,
		stepsInDirection: 0,
	}
	shortestPaths[start.ToString()] = &startPath
	queue.Add(start.ToString())
	for queue.Size() > 0 && killswitch < 99999999 {
		currentStr := queue.Random()
		queue.Remove(currentStr)
		current := grid.ParseCoords(currentStr)

		if current.X == len(blocks[0]) - 1 && current.Y == len(blocks) - 1 {
			path := shortestPaths[currentStr]
			path.visited.Add(currentStr)
			winningPaths = append(winningPaths, path)
			continue
		}

		neighbors := []grid.Coords{
			{ X: current.X + 1, Y: current.Y },
			{ X: current.X - 1, Y: current.Y },
			{ X: current.X, Y: current.Y + 1 },
			{ X: current.X, Y: current.Y - 1 },
		}
		currentPath := shortestPaths[currentStr]
		currentHeatLoss := blocks[current.Y][current.X]
		for _, n := range neighbors {
			nextKey := n.ToString()
			nextPath, hasVisited := shortestPaths[nextKey]
			dir := grid.Coords{X: n.X - current.X, Y: n.Y - current.Y}
			if grid.IsInGrid(n, blocks) &&
				!currentPath.visited.Has(nextKey) &&
				(!hasVisited || currentPath.heatLoss + currentHeatLoss < nextPath.heatLoss) {
				if dir.X == currentPath.direction.X && dir.Y == currentPath.direction.Y {
					if currentPath.stepsInDirection < 3 {
						queue.Add(nextKey)
						nextVisited := currentPath.visited.Copy()
						nextVisited.Add(nextKey)
						newPath := path{
							visited: &nextVisited,
							heatLoss: currentPath.heatLoss + currentHeatLoss,
							direction: dir,
							stepsInDirection: currentPath.stepsInDirection + 1,
						}
						shortestPaths[nextKey] = &newPath
					}
				} else {
					queue.Add(nextKey)
					nextVisited := currentPath.visited.Copy()
					nextVisited.Add(nextKey)
					newPath := path{
						visited: &nextVisited,
						heatLoss: currentPath.heatLoss + currentHeatLoss,
						direction: dir,
						stepsInDirection: 1,
					}
					shortestPaths[nextKey] = &newPath
				}
			}
		}
	}

	// printPath(blocks, shortestPaths[grid.Coords{X: len(blocks[0]) - 1, Y: len(blocks) - 1}.ToString()])

	startHeathLoss, endHeatLoss := blocks[0][0], blocks[len(blocks) - 1][len(blocks[0]) - 1]
	minHeatLoss := maths.MaxInt()
	for _, path := range winningPaths {
		// printPath(blocks, path)
		minHeatLoss = maths.Min(minHeatLoss, path.heatLoss + endHeatLoss - startHeathLoss)
	}
	return minHeatLoss
}

func printPath(blocks [][]int, path *path) {
	total := -1 * blocks[0][0]
	for y, line := range blocks {
		for x, value := range line {
			if path.visited.Has(grid.Coords{X: x, Y: y}.ToString()) {
				total += value
				print(colors.Green + fmt.Sprintf("%v", value) + colors.Reset)
			} else {
				print(value)
			}
		}
		println()
	}
	println()
	println(colors.Purple + fmt.Sprintf("Summed %v", total) + colors.Reset)
	println(colors.Red + fmt.Sprintf("Heat loss %v", path.heatLoss - blocks[0][0] + blocks[len(blocks) - 1][len(blocks[0]) - 1]) + colors.Reset)
}

type pqItem struct {
	coords grid.Coords
	direction grid.Coords
	stepsInDirection int
}
func buildItemKey(coords grid.Coords, direction grid.Coords, stepsInDirection int) string {
	return fmt.Sprintf("%v,%v,%v,%v,%v", coords.X, coords.Y, direction.X, direction.Y, stepsInDirection)
}
func parseItem(value string) pqItem {
	values := slices.ParseIntsFromStrings(strings.Split(value, ","))
	return pqItem{
		coords: grid.Coords{X: values[0], Y: values[1]},
		direction: grid.Coords{X: values[2], Y: values[3]},
		stepsInDirection: values[4],
	}
}

func dijkstraPriorityQueue(blocks [][]int) int {
	dist := map[string]int{}
	lenX, lenY := len(blocks[0]), len(blocks)
	target, source := grid.Coords{ X: len(blocks[0]) - 1, Y: len(blocks) - 1 }, grid.Origin
	q := make(priorityqueue.PriorityQueue, 0)
	directions := []grid.Coords{
		{ X: 1, Y: 0 },
		{ X: -1, Y: 0 },
		{ X: 0, Y: 1 },
		{ X: 0, Y: -1 },
	}
	targetKeys := []string{}
	for _, direction := range directions {
		for i := 1; i <= 3; i++ {
			key := buildItemKey(target, direction, i)
			targetKeys = append(targetKeys, key)
		}
	}

	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			next := grid.Coords{X: x, Y: y}
			for _, direction := range directions {
				value, start := math.MaxInt, 1
				if next.ToString() == source.ToString() {
					value = 0
					start = 0
					if direction.X < 0 || direction.Y < 0 {
						continue
					}
				}
				for i := start; i <= 3; i++ {
					key := buildItemKey(next, direction, i)
					dist[key] = value
					heap.Push(&q, &priorityqueue.Item{Priority: dist[key], Value: key})
				}
			}
		}
	}

	for q.Len() > 0 {
		u := heap.Pop(&q).(*priorityqueue.Item)
		currentValue := parseItem(u.Value)
		curentCoords := currentValue.coords
		for _, direction := range directions {
			steps := 1
			if direction.X == currentValue.direction.X && direction.Y == currentValue.direction.Y {
				steps = currentValue.stepsInDirection + 1
			}
			next := grid.Coords{X: curentCoords.X + direction.X, Y: curentCoords.Y + direction.Y}
			minAlt := math.MaxInt
			if grid.IsInGrid(next, blocks) {
				for i := steps; i <= 3; i++ {
					minAlt = maths.Min(minAlt, dist[u.Value] + blocks[next.Y][next.X])
				}
			}
			for i := steps; i <= 3; i++ {
				key := buildItemKey(next, direction, i)
				if q.Has(key) && minAlt < dist[key] {
					nextKey := buildItemKey(next, direction, steps)
					dist[nextKey] = minAlt
					q.Update(nextKey, minAlt)
				}
			}
		}
	}

	minHeatLoss := maths.MaxInt()
	for _, key := range targetKeys {
		println(key, dist[key])
		minHeatLoss = maths.Min(minHeatLoss, dist[key])
	}
	return minHeatLoss
}
