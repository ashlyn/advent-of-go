package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/priorityqueue"
	"advent-of-go/utils/slices"
	"container/heap"
	"math"
)

type cheatingRules struct {
	maxTimeCheating int
	minTimeSaved int
}

func main() {
	input := files.ReadFile(20, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	walls, start, end, size := parseRacetrack(input)
	bestTimes := dijkstraPriorityQueue(walls, start, size)
	return findBetterPaths(walls, start, end, size, bestTimes, cheatingRules{maxTimeCheating: 2, minTimeSaved: 100})
}

func solvePart2(input []string) int {
	walls, start, end, size := parseRacetrack(input)
	bestTimes := dijkstraPriorityQueue(walls, start, size)
	return findBetterPaths(walls, start, end, size, bestTimes, cheatingRules{maxTimeCheating: 20, minTimeSaved: 100})
}

func dijkstraPriorityQueue(walls []grid.Coords, start, size grid.Coords) map[grid.Coords]int {
	times := map[grid.Coords]int{}

	times[start] = 0
	queue := make(priorityqueue.PriorityQueue, 0)

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			v := grid.Coords{X: x, Y: y}
			if v != start {
				times[v] = math.MaxInt
			}
			heap.Push(&queue, &priorityqueue.Item{Value: v.ToString(), Priority: times[v]})
		}
	}

	for queue.Len() > 0 {
		u := heap.Pop(&queue).(*priorityqueue.Item)
		point := grid.ParseCoords(u.Value)
		neighbors := getNeighbors(point, walls, size)
		for i := 0; i < len(neighbors); i++ {
			v := neighbors[i]
			key := v.ToString()
			alt := times[point] + 1
			if alt < times[v] && alt > 0 {
				times[v] = alt
				queue.Update(key, alt)
			}
		}
	}

	return times
}

func parseRacetrack(input []string) ([]grid.Coords, grid.Coords, grid.Coords, grid.Coords) {
	var start, end, size grid.Coords
	walls := make([]grid.Coords, 0)

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			switch input[y][x] {
			case '#':
				walls = append(walls, grid.Coords{X: x, Y: y})
			case 'S':
				start = grid.Coords{X: x, Y: y}
			case 'E':
				end = grid.Coords{X: x, Y: y}
			}
		}
	}

	size = grid.Coords{X: len(input[0]), Y: len(input)}
	return walls, start, end, size
}

func findBetterPaths(walls []grid.Coords, start, end, size grid.Coords, bestTimes map[grid.Coords]int, cheatingInfo cheatingRules) int {
	previous, current := grid.Coords{X: -1, Y: -1}, start
	time := 0

	for current != end {
		time += findTimeSavingPathsFromPoint(walls, current, end, size, bestTimes, cheatingInfo)
		neighbors := getNeighbors(current, walls, size)
		for i := 0; i < len(neighbors); i++ {
			next := neighbors[i]
			if next != previous {
				previous = grid.Coords { X: current.X, Y: current.Y }
				current = grid.Coords { X: next.X, Y: next.Y }
				break
			}
		}
	}
	return time
}

func findTimeSavingPathsFromPoint(walls []grid.Coords, point, end, size grid.Coords, bestTimes map[grid.Coords]int, cheatingInfo cheatingRules) int {
	if point == end {
		return 0
	}

	time := 0
	neighbors := getNeighborsWithCheating(point, walls, size, cheatingInfo.maxTimeCheating)
	for i := 0; i < len(neighbors); i++ {
		next := neighbors[i]
		isSavingTime := bestTimes[next] > bestTimes[point]
		isSavingTargetTime := bestTimes[next] - bestTimes[point] - next.ManhattanDistance(point) >= cheatingInfo.minTimeSaved
		if isSavingTime && isSavingTargetTime {
			time++
		}
	}
	return time
}

var directions = []grid.Coords{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: -1, Y: 0}}

func getNeighbors(c grid.Coords, walls []grid.Coords, size grid.Coords) []grid.Coords {
	neighbors := []grid.Coords{}
	for _, direction := range directions {
		next := grid.Coords{X: c.X + direction.X, Y: c.Y + direction.Y}
		if isValidPoint(next, walls, size) {
			neighbors = append(neighbors, next)
		}
	}
	return neighbors
}

func getNeighborsWithCheating(c grid.Coords, walls []grid.Coords, size grid.Coords, maxDistance int) []grid.Coords {
	neighbors := []grid.Coords{}
	if !slices.Contains(walls, c) {
		for x := 0; x < size.X; x++ {
			for y := 0; y < size.Y; y++ {
				destination := grid.Coords{X: x, Y: y}
				if destination.ManhattanDistance(c) <= maxDistance && isValidPoint(destination, walls, size) {
					neighbors = append(neighbors, destination)
				}
			}
		}
	}
	return neighbors
}

func isValidPoint(c grid.Coords, walls []grid.Coords, size grid.Coords) bool {
	return c.X >= 0 && c.X < size.X && c.Y >= 0 && c.Y < size.Y && !slices.Contains(walls, c)
}