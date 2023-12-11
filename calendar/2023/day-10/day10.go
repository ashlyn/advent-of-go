package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/sets"
	"advent-of-go/utils/slices"
	"fmt"
)

const (
	north string = "north"
	south = "south"
	east = "east"
	west = "west"
)
var directionNames = map[string]string {
	grid.Coords{X: 0, Y: 1}.ToString(): south,
	grid.Coords{X: 0, Y: -1}.ToString(): north,
	grid.Coords{X: 1, Y: 0}.ToString(): east,
	grid.Coords{X: -1, Y: 0}.ToString(): west,
}
var directionMap = map[string]grid.Coords {
	north: {X: 0, Y: -1},
	south: {X: 0, Y: 1},
	east: {X: 1, Y: 0},
	west: {X: -1, Y: 0},
}

func main() {
	input := files.ReadFile(10, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	pipes, start := parseInput(input)
	loop := findLoop(pipes, start)
	return loop.Size() / 2 
}

func solvePart2(input []string) int {
	enclosed := findEnclosed(input)
	return enclosed.Size()
}

func parseInput(input []string) (map[string]rune, grid.Coords) {
	var start grid.Coords
	pipes, notPipes := map[string]rune{}, sets.New()
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			c := grid.Coords{X: x, Y: y}
			current := rune(input[y][x])
			if current == 'S' {
				start = c
			}
			if current != '.' {
				pipes[c.ToString()] = rune(input[y][x])
			} else {
				notPipes.Add(c.ToString())
			}
		}
	}

	return pipes, start
}

func findLoop(pipes map[string]rune, start grid.Coords) sets.Set {
	visited := sets.New()
	visited.Add(start.ToString())

	candidateDirections := []string{north, south, east, west}
	current, moving, canMove := start, grid.Origin, false
	for _, direction := range candidateDirections {
		dir := directionMap[direction]
		next := grid.Coords{ X: start.X + dir.X, Y: start.Y + dir.Y }
		current, moving, canMove = getNextCoords(next, dir, pipes[next.ToString()])
		if canMove {
			visited.Add(next.ToString())
			break
		}
	}

	visited.Add(current.ToString())
	for pipes[current.ToString()] != 'S' {
		current, moving, canMove = getNextCoords(current, moving, pipes[current.ToString()])
		visited.Add(current.ToString())
	}
	return visited
}

func getNextCoords(currentPosition grid.Coords, currentDirection grid.Coords, piece rune) (grid.Coords, grid.Coords, bool) {
	newDirection := grid.Origin
	canMove := false
	direction := directionNames[currentDirection.ToString()]
	switch piece {
	case '|':
		if direction == north || direction == south {
			newDirection = currentDirection
			canMove = true
		}
	case '-':
		if direction == east || direction == west {
			newDirection = currentDirection
			canMove = true
		}
	case 'L':
		if direction == west {
			newDirection = directionMap[north]
			canMove = true
		} else if direction == south {
			newDirection = directionMap[east]
			canMove = true
		}
	case 'J':
		if direction == south {
			newDirection = directionMap[west]
			canMove = true
		} else if direction == east {
			newDirection = directionMap[north]
			canMove = true
		}
	case '7':
		if direction == north {
			newDirection = directionMap[west]
			canMove = true
		} else if direction == east {
			newDirection = directionMap[south]
			canMove = true
		}
	case 'F':
		if direction == north {
			newDirection = directionMap[east]
			canMove = true
		} else if direction == west {
			newDirection = directionMap[south]
			canMove = true
		}
	}
	return grid.Coords{ X: currentPosition.X + newDirection.X, Y: currentPosition.Y + newDirection.Y }, newDirection, canMove
}

func findEnclosed(input []string) sets.Set {
	enclosed, visited := sets.New(), sets.New()
	pipes, start := parseInput(input)
	loop := findLoop(pipes, start)

	// go through each tile
	// if in loop, do nothing
	// already enclosed, do nothing
	// otherwise
	// add to current group
	// if current group presumed enclosed
	// if edge, current group not enclosed
	// if loop and not enclosed, current group not enclosed
	// if not in loop, add neighbors to queue
	// if queue empty, add current group to appropriate set
	// empty current group

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			currentGroup, isEnclosed := sets.New(), true
			queue := sets.New()
			queue.Add((grid.Coords{ X: x, Y: y }).ToString())
			for queue.Size() > 0 {
				key := queue.Iterator()[0]
				queue.Remove(key)
				c := grid.ParseCoords(key)

				if enclosed.Has(key) || visited.Has(key) || currentGroup.Has(key) {
					continue
				}
				visited.Add(key)
				if loop.Has(key) {
					if hasGap(c, pipes) {
						println("gap")
						// todo handle gaps better
						isEnclosed = false
					}
					continue
				}
				if c.X == 0 || c.Y == 0 || c.X == len(input[c.Y]) - 1 || c.Y == len(input) - 1 {
					isEnclosed = false
				}
				currentGroup.Add(key)
				if c.X > 0 {
					queue.Add((grid.Coords{ X: c.X - 1, Y: c.Y }).ToString())
				}
				if c.X < len(input[c.Y]) - 1 {
					queue.Add((grid.Coords{ X: c.X + 1, Y: c.Y }).ToString())
				}
				if c.Y > 0 {
					queue.Add((grid.Coords{ X: c.X, Y: c.Y - 1 }).ToString())
				}
				if c.Y < len(input) - 1 {
					queue.Add((grid.Coords{ X: c.X, Y: c.Y + 1 }).ToString())
				}
			}
			if isEnclosed && currentGroup.Size() > 0 {
				fmt.Printf("Enclosed: %v\n", currentGroup)
				enclosed = enclosed.Union(currentGroup)
			}
		}
	}
	println(visited.Size(), loop.Size(), enclosed.Size())

	return enclosed
}

func getConnections(pipe rune)	[]string {
	var directionsToCheck []string
	
	switch pipe {
	case '|':
		directionsToCheck = []string{north, south}
	case '-':
		directionsToCheck = []string{east, west}
	case 'L':
		directionsToCheck = []string{north, east}
	case 'J':
		directionsToCheck = []string{north, west}
	case '7':
		directionsToCheck = []string{south, west}
	case 'F':
		directionsToCheck = []string{south, east}
	case 'S':
		directionsToCheck = []string{north, south, east, west}
	}
	return directionsToCheck
}

func hasGap(c grid.Coords, pipes map[string]rune) bool {
	// todo: fix gap logic
	current := pipes[c.ToString()]
	if current == 'S' {
		return false
	}
	directionsToCheck := getConnections(current)

	for _, direction := range directionsToCheck {
		dir := directionMap[direction]
		next := grid.Coords{ X: c.X + dir.X, Y: c.Y + dir.Y }
		possible := getConnections(pipes[next.ToString()])
		if current == '7' {
			fmt.Printf("%v%v%v\n", possible, next, string(pipes[next.ToString()]))
		}
		if direction == north && !slices.Contains(possible, south) {
			return true
		}
		if direction == south && !slices.Contains(possible, north) {
			return true
		}
		if direction == east && !slices.Contains(possible, west) {
			return true
		}
		if direction == west && !slices.Contains(possible, east) {
			return true
		}
	}

	return false
}
