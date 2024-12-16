package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/sets"
	"sort"
	"strings"
)

func main() {
	input := files.ReadFile(13, 2018, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) string {
	tracks, carts, cartPositions := prepareTracks(input)
	crash := simulateUntilCrash(tracks, carts, cartPositions)
	return crash.ToString()
}

func solvePart2(input []string) int {
	result := 0



	return result
}

type turn int
const (
	left turn = iota
	straight turn = iota
	right turn = iota
)
type cart struct {
	position grid.Coords
	direction grid.Coords
	nextTurn turn
}
func prepareTracks(input []string) ([]string, []cart, *sets.Set) {
	tracks := make([]string, len(input))
	carts := []cart{}
	cartPositions := sets.New()

	for y := 0; y < len(input); y++ {
		sb := strings.Builder{}
		for x := 0; x < len(input[y]); x++ {
			current := input[y][x]
			position := grid.Coords{X: x, Y: y}
			switch current {
			case '^':
				sb.WriteString("|")
				carts = append(carts, cart{position, grid.Coords{X: 0, Y: -1}, left})
				cartPositions.Add(position.ToString())
			case 'v':
				sb.WriteString("|")
				carts = append(carts, cart{position, grid.Coords{X: 0, Y: 1}, left})
				cartPositions.Add(position.ToString())
			case '<':
				sb.WriteString("-")
				carts = append(carts, cart{position, grid.Coords{X: -1, Y: 0}, left})
				cartPositions.Add(position.ToString())
			case '>':
				sb.WriteString("-")
				carts = append(carts, cart{position, grid.Coords{X: 1, Y: 0}, left})
				cartPositions.Add(position.ToString())
			default:
				sb.WriteByte(current)
			}
		}
		tracks[y] = sb.String()
	}
	return tracks, carts, &cartPositions
}

func (c *cart) move(tracks []string) {
	c.position = grid.Coords{X: c.position.X + c.direction.X, Y: c.position.Y + c.direction.Y}
	currentTrack := tracks[c.position.Y][c.position.X]
	if currentTrack == '+' {
		c.direction = makeTurn(c.direction, c.nextTurn)
		c.nextTurn = nextTurn(c.nextTurn)
	} else if currentTrack == '/' {
		if c.direction.Y == 0 {
			c.direction = makeTurn(c.direction, left)
		} else if c.direction.X == 0 {
			c.direction = makeTurn(c.direction, right)
		}
	} else if currentTrack == '\\' {
		if c.direction.Y == 0 {
			c.direction = makeTurn(c.direction, right)
		} else if c.direction.X == 0 {
			c.direction = makeTurn(c.direction, left)
		}
	}
}

func nextTurn(current turn) turn {
	if current == left {
		return straight
	}

	if current == straight {
		return right
	}

	return left;
}

func makeTurn(currentDirection grid.Coords, turnDirection turn) grid.Coords {
	if turnDirection == left {
		if currentDirection.X == 1 {
			return grid.Coords{X: 0, Y: -1}
		}
		if currentDirection.X == -1 {
			return grid.Coords{X: 0, Y: 1}
		}
		if currentDirection.Y == 1 {
			return grid.Coords{X: 1, Y: 0}
		}
		if currentDirection.Y == -1 {
			return grid.Coords{X: -1, Y: 0}
		}
	}

	if turnDirection == right {
		if currentDirection.X == 1 {
			return grid.Coords{X: 0, Y: 1}
		}
		if currentDirection.X == -1 {
			return grid.Coords{X: 0, Y: -1}
		}
		if currentDirection.Y == 1 {
			return grid.Coords{X: -1, Y: 0}
		}
		if currentDirection.Y == -1 {
			return grid.Coords{X: 1, Y: 0}
		}
	}

	return currentDirection
}

func simulateUntilCrash(tracks []string, carts []cart, cartPositions *sets.Set) grid.Coords {
	for {
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].position.Y == carts[j].position.Y {
				return carts[i].position.X < carts[j].position.X
			}
			return carts[i].position.Y < carts[j].position.Y
		})
		for i := 0; i < len(carts); i++ {
			c := &carts[i]
			cartPositions.Remove(c.position.ToString())
			c.move(tracks)
			if cartPositions.Has(c.position.ToString()) {
				return c.position
			}
			cartPositions.Add(c.position.ToString())
		}
	}
}
