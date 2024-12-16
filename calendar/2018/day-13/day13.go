package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"sort"
	"strings"
)

func main() {
	input := files.ReadFile(13, 2018, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) string {
	tracks, carts := prepareTracks(input)
	crash := simulate(tracks, carts, false)
	return crash.ToString()
}

func solvePart2(input []string) string {
	tracks, carts := prepareTracks(input)
	crash := simulate(tracks, carts, true)
	return crash.ToString()
}

type turn int
const (
	left turn = iota
	straight turn = iota
	right turn = iota
)
type cart struct {
	id int
	position grid.Coords
	direction grid.Coords
	nextTurn turn
	active bool
}
func prepareTracks(input []string) ([]string, []cart) {
	tracks := make([]string, len(input))
	carts := []cart{}

	for y := 0; y < len(input); y++ {
		sb := strings.Builder{}
		for x := 0; x < len(input[y]); x++ {
			current := input[y][x]
			position := grid.Coords{X: x, Y: y}
			switch current {
			case '^':
				sb.WriteString("|")
				carts = append(carts, cart{position: position, direction: grid.Coords{X: 0, Y: -1}, nextTurn: left, id: len(carts), active: true})
			case 'v':
				sb.WriteString("|")
				carts = append(carts, cart{position: position, direction: grid.Coords{X: 0, Y: 1}, nextTurn: left, id: len(carts), active: true})
			case '<':
				sb.WriteString("-")
				carts = append(carts, cart{position: position, direction: grid.Coords{X: -1, Y: 0}, nextTurn: left, id: len(carts), active: true})
			case '>':
				sb.WriteString("-")
				carts = append(carts, cart{position: position, direction: grid.Coords{X: 1, Y: 0}, nextTurn: left, id: len(carts), active: true})
			default:
				sb.WriteByte(current)
			}
		}
		tracks[y] = sb.String()
	}
	return tracks, carts
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

func simulate(tracks []string, carts []cart, removeCrashes bool) grid.Coords {
	for {
		carts = filterInactiveCards(carts)
		if len(carts) == 1 {
			return carts[0].position
		}
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].position.Y == carts[j].position.Y {
				return carts[i].position.X < carts[j].position.X
			}
			return carts[i].position.Y < carts[j].position.Y
		})
		for i := 0; i < len(carts); i++ {
			if !carts[i].active {
				continue
			}
			c := &carts[i]
			c.move(tracks)
			collisionIndex := c.collisionIndex(carts)
			if collisionIndex != -1 && !removeCrashes {
				return c.position
			}
			if collisionIndex != -1 && removeCrashes {
				c.active = false
				collidedWith := &carts[collisionIndex]
				collidedWith.active = false
			}
		}
	}
}

func filterInactiveCards(carts []cart) []cart {
	activeCarts := []cart{}
	for _, c := range carts {
		if c.active {
			activeCarts = append(activeCarts, c)
		}
	}
	return activeCarts
}

func (c *cart) collisionIndex(carts []cart) int {
	for i, cart := range carts {
		if cart.position == c.position && cart.id != c.id {
			return i
		}
	}
	return -1
}
