package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/priorityqueue"
	"advent-of-go/utils/slices"
	"container/heap"
	"fmt"
	"strings"
)

var (
	left  = grid.Coords{ X: -1, Y: 0 }
	right = grid.Coords{ X: 1, Y: 0 }
	up    = grid.Coords{ X: 0, Y: -1 }
	down  = grid.Coords{ X: 0, Y: 1 }
)
var directions = []grid.Coords{left, right, up, down}

var validTurningDirections = map[grid.Coords][]grid.Coords{
	left:  { up, down },
	right: { up, down },
	up:    { left, right },
	down:  { left, right },
}
var backtrackDirection = map[grid.Coords]grid.Coords{
	left:  right,
	right: left,
	up:    down,
	down:  up,
}

func main() {
	input := files.ReadFile(17, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	blocks := parseInput(input)
	return findLowestHeatloss(blocks, 3, 0)
}

func solvePart2(input []string) int {
	blocks := parseInput(input)
	return findLowestHeatloss(blocks, 10, 4)
}

// Roughly based on 2021/day-15/day15.go
func findLowestHeatloss(blocks [][]int, maxStepsInDirection int, minStepsInDirection int) int {
	target := grid.Coords{X: len(blocks[0]) - 1, Y: len(blocks) - 1}
	startRight := pqItem{coords: grid.Origin, direction: grid.Coords{X: 1, Y: 0}, stepsInDirection: 0}
	startDown := pqItem{coords: grid.Origin, direction: grid.Coords{X: 0, Y: 1}, stepsInDirection: 0}
	rightKey, downKey := startRight.buildKey(), startDown.buildKey()
	q := make(priorityqueue.PriorityQueue, 0)
	heap.Push(&q, &priorityqueue.Item{Priority: 0, Value: rightKey })
	heap.Push(&q, &priorityqueue.Item{Priority: 0, Value: downKey })

	dist := map[string]int{ rightKey: 0, downKey: 0}

	for len(q) > 0 {
		current := heap.Pop(&q).(*priorityqueue.Item)
		currentItem := parseItem(current.Value)

		if dist[current.Value] < current.Priority {
			// shorter path already explored
			continue
		}

		if currentItem.coords.X == target.X && currentItem.coords.Y == target.Y &&
			 currentItem.stepsInDirection >= minStepsInDirection {
			return current.Priority
		}

		for _, nextDirection := range directions {
			// don't explore directions that are in excess of allowed moves, invalid rotations, or backtracking
			if currentItem.stepsInDirection == maxStepsInDirection &&
				 !slices.Contains(validTurningDirections[currentItem.direction], nextDirection) ||
				 nextDirection == backtrackDirection[currentItem.direction] {
				continue
			}

			nextCoords := grid.Coords{ X: currentItem.coords.X + nextDirection.X, Y: currentItem.coords.Y + nextDirection.Y }
			if !grid.IsInGrid(nextCoords, blocks) {
				continue
			}
			steps := currentItem.stepsInDirection

			if currentItem.stepsInDirection < minStepsInDirection {
				if nextDirection != currentItem.direction {
					// turning when not allowed
					continue
				}
				steps++
			} else if  nextDirection != currentItem.direction {
				// turning
				steps = 1
			} else {
				steps = steps % maxStepsInDirection + 1
			}

			nextState := pqItem{coords: nextCoords, direction: nextDirection, stepsInDirection: steps}
			nextKey := nextState.buildKey()
			nextHeatLoss := blocks[nextCoords.Y][nextCoords.X]
			previousDist, hasVisited := dist[nextKey]
			if hasVisited && previousDist <= current.Priority+nextHeatLoss {
				continue
			}

			dist[nextKey] = current.Priority + nextHeatLoss
			heap.Push(&q, &priorityqueue.Item{Priority: current.Priority + nextHeatLoss, Value: nextKey})
		}
	}

	return -1
}

func parseInput(input []string) [][]int {
	result := make([][]int, len(input))
	for i, line := range input {
		strs := strings.Split(line, "")
		result[i] = slices.ParseIntsFromStrings(strs)
	}
	return result
}

type pqItem struct {
	coords grid.Coords
	direction grid.Coords
	stepsInDirection int
}
func (item *pqItem) buildKey() string {
	return fmt.Sprintf("%v,%v,%v,%v,%v", item.coords.X, item.coords.Y, item.direction.X, item.direction.Y, item.stepsInDirection)
}
func parseItem(value string) pqItem {
	values := slices.ParseIntsFromStrings(strings.Split(value, ","))
	return pqItem{
		coords: grid.Coords{X: values[0], Y: values[1]},
		direction: grid.Coords{X: values[2], Y: values[3]},
		stepsInDirection: values[4],
	}
}
