package main

import (
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
	input := files.ReadFile(16, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	score, _ := dijkstraPiroityQueue(input)
	return score
}

func solvePart2(input []string) int {
	_, placesToSit := dijkstraPiroityQueue(input)
	return placesToSit
}

var directions = []grid.Coords{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: -1, Y: 0}}
type reindeerState struct {
	position grid.Coords
	direction grid.Coords
}
func dijkstraPiroityQueue(maze []string) (int, int) {
	scores, previous := map[reindeerState]int{}, map[reindeerState][]reindeerState{}

	start, end := parseMaze(maze)

	initialState := reindeerState{position: start, direction: grid.Coords{X: 1, Y: 0}}
	scores[initialState] = 0
	initialPath := sets.New()
	initialPath.Add(initialState.position.ToString())
	
	q := make(priorityqueue.PriorityQueue, 0)
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			if maze[y][x] != '#' {
				position := grid.Coords{X: x, Y: y}
				for _, direction := range directions {
					if position == start && direction != initialState.direction {
						continue
					}
					state := reindeerState{position: position, direction: direction}
					if position != start {
						scores[state] = math.MaxInt
					}
					heap.Push(&q, &priorityqueue.Item{Priority: scores[state], Value: state.toString()})
				}
			}
		}
	}

	for q.Len() > 0 {
		item := heap.Pop(&q).(*priorityqueue.Item)
		state := parseState(item.Value)
		neighbors := getNeighbors(maze, state)
		for _, neighbor := range neighbors {
			altScore := scores[state] + getScoreForMove(state, neighbor)
			if altScore <= scores[neighbor] && altScore > 0 {
				scores[neighbor] = altScore
				q.Update(neighbor.toString(), altScore)
				if _, ok := previous[neighbor]; !ok {
					previous[neighbor] = []reindeerState{}
				}
				previous[neighbor] = append(previous[neighbor], state)
			}
		}
	}

	minScore := math.MaxInt
	for _, direction := range directions {
		state := reindeerState{position: end, direction: direction}
		if scores[state] < minScore {
			minScore = scores[state]
		}
	}
	
	placesToSit := sets.New()
	for _, direction := range directions {
		state := reindeerState{position: end, direction: direction}
		if scores[state] == minScore {
			queue := sets.New()
			queue.Add(state.toString())
			for queue.Size() > 0 {
				current := parseState(queue.Random())
				queue.Remove(current.toString())
				placesToSit.Add(current.position.ToString())
				for _, p := range previous[current] {
					if !queue.Has(p.toString()) {
						queue.Add(p.toString())
					}
				}
			}
		}
	}
	return minScore, placesToSit.Size()
}

func parseMaze(input []string) (grid.Coords, grid.Coords) {
	start, end := grid.Coords{}, grid.Coords{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == 'S' {
				start = grid.Coords{X: x, Y: y}
			}
			if input[y][x] == 'E' {
				end = grid.Coords{X: x, Y: y}
			}
		}
	}
	return start, end
}

func (r *reindeerState) toString() string {
	return fmt.Sprintf("%s %s", r.position.ToString(), r.direction.ToString())
}

func parseState(input string) reindeerState {
	parts := strings.Split(input, " ")
	position := grid.ParseCoords(parts[0])
	direction := grid.ParseCoords(parts[1])
	return reindeerState{position: position, direction: direction}
}

func getScoreForMove(from reindeerState, to reindeerState) int {
	if from.direction == to.direction {
		return 1
	}
	turns := maths.Abs(slices.IndexOf(from.direction, directions) - slices.IndexOf(to.direction, directions))
	if turns == 3 {
		turns = 1
	}
	return (turns * 1000) + 1
}

func getNeighbors(maze []string, state reindeerState) []reindeerState {
	neighbors := []reindeerState{}
	for _, direction := range directions {
		newPosition := grid.Coords{X: state.position.X + direction.X, Y: state.position.Y + direction.Y}
		if newPosition.X < 0 || newPosition.X >= len(maze[0]) || newPosition.Y < 0 || newPosition.Y >= len(maze) {
			continue
		}
		if maze[newPosition.Y][newPosition.X] == '#' {
			continue
		}
		neighbors = append(neighbors, reindeerState{position: newPosition, direction: direction})
	}
	return neighbors
}
