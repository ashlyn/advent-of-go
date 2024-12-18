package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/priorityqueue"
	"advent-of-go/utils/slices"
	"container/heap"
	"math"
)

func main() {
	input := files.ReadFile(18, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	coords := parseInput(input)
	return dijkstraPriorityQueue(coords[0:1024], grid.Coords{X: 71, Y: 71})
}

func solvePart2(input []string) string {
	coords := parseInput(input)
	start, end := 1024, len(input)
	// binary search until we find the first byte with no paths
	for start != end {
		mid := (start + end) / 2
		score := dijkstraPriorityQueue(coords[0:mid], grid.Coords{X: 71, Y: 71})
		if score == math.MaxInt {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return coords[start-1].ToString()
}

func parseInput(input []string) []grid.Coords {
	coords := make([]grid.Coords, len(input))
	for i := 0; i < len(input); i++ {
		coords[i] = grid.ParseCoords(input[i])
	}
	return coords
}

func dijkstraPriorityQueue(corrupted []grid.Coords, size grid.Coords) int {
	dist := map[string]int{}

	start, end := grid.Coords{X: 0, Y: 0}, grid.Coords{X: size.X - 1, Y: size.Y - 1}

	dist[start.ToString()] = 0
	queue := make(priorityqueue.PriorityQueue, 0)

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			v := grid.Coords{X: x, Y: y}
			key := v.ToString()
			if v != start {
				dist[key] = math.MaxInt
			}
			heap.Push(&queue, &priorityqueue.Item{Value: key, Priority: dist[key]})
		}
	}

	for queue.Len() > 0 {
		u := heap.Pop(&queue).(*priorityqueue.Item)
		point := grid.ParseCoords(u.Value)
		for _, d := range directions {
			v := grid.Coords{X: point.X + d.X, Y: point.Y + d.Y}
			key := v.ToString()
			if v.X < 0 || v.X >= size.X || v.Y < 0 || v.Y >= size.Y {
				continue
			}
			if slices.Contains(corrupted, v) {
				continue
			}
			if !queue.Has(key) {
				continue
			}

			alt := dist[u.Value] + 1
			if alt < dist[key] && alt > 0 {
				dist[v.ToString()] = alt
				queue.Update(key, alt)
			}
		}
	}

	return dist[end.ToString()]
}

var directions = []grid.Coords{
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: -1, Y: 0},
}
