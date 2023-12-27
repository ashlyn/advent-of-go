package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/sets"
	"strings"
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
	loop, _ := findLoop(pipes, start)
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

func findLoop(pipes map[string]rune, start grid.Coords) (sets.Set, bool) {
	visited := sets.New()
	startNorth := false
	visited.Add(start.ToString())

	candidateDirections := []string{north, south, east, west}
	current, moving, canMove := start, grid.Origin, false
	for _, direction := range candidateDirections {
		dir := directionMap[direction]
		next := grid.Coords{ X: start.X + dir.X, Y: start.Y + dir.Y }
		current, moving, canMove = getNextCoords(next, dir, pipes[next.ToString()])
		if canMove {
			if direction == north {
				startNorth = true
			}
			visited.Add(next.ToString())
			break
		}
	}

	for pipes[current.ToString()] != 'S' {
		visited.Add(current.ToString())
		current, moving, canMove = getNextCoords(current, moving, pipes[current.ToString()])
	}

	if moving.X == 0 && moving.Y == -1 {
		startNorth = true
	}


	return visited, startNorth
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
	enclosed := sets.New()
	pipes, start := parseInput(input)
	loop, startNorth := findLoop(pipes, start)

	nonNorthTiles := "-7F"
	if !startNorth {
		nonNorthTiles += "S"
	}
	for y := 0; y < len(input) - 1; y++ {
		isEnclosed := false
		for x := 0; x < len(input[y]) - 1; x++ {
			c := grid.Coords{X: x, Y: y}
			key := c.ToString()
			if loop.Has(key) && !strings.Contains(nonNorthTiles, string(pipes[key])) {
				isEnclosed = !isEnclosed
			} else if !loop.Has(key) && isEnclosed {
				enclosed.Add(key)
			}
		}
	}
	return enclosed
}
