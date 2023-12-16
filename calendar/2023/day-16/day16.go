package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/sets"
	"fmt"
)

func main() {
	input := files.ReadFile(16, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	start := beam{position: grid.Coords{X: 0, Y: 0}, velocity: grid.Coords{X: 1, Y: 0}}
	return countEnergized(input, start)
}

func solvePart2(input []string) int {
	maxEnergized := 0

	for x := 0; x < len(input[0]); x++ {
		maxEnergized = maths.Max(maxEnergized, countEnergized(input, beam{position: grid.Coords{X: x, Y: 0}, velocity: grid.Coords{X: 0, Y: 1}}))
		maxEnergized = maths.Max(maxEnergized, countEnergized(input, beam{position: grid.Coords{X: x, Y: len(input) - 1}, velocity: grid.Coords{X: 0, Y: -1}}))
	}

	for y := 1; y < len(input) - 1; y++ {
		maxEnergized = maths.Max(maxEnergized, countEnergized(input, beam{position: grid.Coords{X: 0, Y: y}, velocity: grid.Coords{X: 1, Y: 0}}))
		maxEnergized = maths.Max(maxEnergized, countEnergized(input, beam{position: grid.Coords{X: len(input[y]) - 1, Y: y}, velocity: grid.Coords{X: -1, Y: 0}}))
	}

	return maxEnergized
}

type beam struct {
	position grid.Coords
	velocity grid.Coords
}
func countEnergized(input []string, start beam) int {
	energized, paths := sets.New(), sets.New()
	queue := []beam{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.position.X < 0 || current.position.Y < 0 ||
			current.position.X >= len(input[0]) || current.position.Y >= len(input) {
			continue
		}

		pathKey := fmt.Sprintf("%v %v", current.position.ToString(), current.velocity.ToString())
		if paths.Has(pathKey) {
			continue
		}
		paths.Add(pathKey)
		energized.Add(current.position.ToString())

		space := input[current.position.Y][current.position.X]
		switch space {
		case '.':
			next := beam{
				position: grid.Coords{ X: current.position.X + current.velocity.X, Y: current.position.Y + current.velocity.Y},
				velocity: current.velocity,
			}
			queue = append(queue, next)
		case '/':
			nextVelocity := grid.Coords{X: -current.velocity.Y, Y: -current.velocity.X}
			nextPosition := grid.Coords{X: current.position.X + nextVelocity.X, Y: current.position.Y + nextVelocity.Y}
			next := beam{ position: nextPosition, velocity: nextVelocity }
			queue = append(queue, next)
		case '\\': 
			nextVelocity := grid.Coords{X: current.velocity.Y, Y: current.velocity.X}
			nextPosition := grid.Coords{X: current.position.X + nextVelocity.X, Y: current.position.Y + nextVelocity.Y}
			next := beam{ position: nextPosition, velocity: nextVelocity }
			queue = append(queue, next)
		case '-':
			if current.velocity.Y == 0 {
				next := beam{
					position: grid.Coords{ X: current.position.X + current.velocity.X, Y: current.position.Y + current.velocity.Y},
					velocity: current.velocity,
				}
				queue = append(queue, next)
			} else {
				leftVelocity, rightVelocity  := grid.Coords{X: -1, Y: 0 }, grid.Coords{X: 1, Y: 0 }
				leftPosition, rightPosition := grid.Coords{X: current.position.X + leftVelocity.X, Y: current.position.Y + leftVelocity.Y}, grid.Coords{X: current.position.X + rightVelocity.X, Y: current.position.Y + rightVelocity.Y}
				left, right := beam{ position: leftPosition, velocity: leftVelocity }, beam{ position: rightPosition, velocity: rightVelocity }
				queue = append(queue, left, right)
			}
		case '|':
			if current.velocity.X == 0 {
				next := beam{
					position: grid.Coords{ X: current.position.X + current.velocity.X, Y: current.position.Y + current.velocity.Y},
					velocity: current.velocity,
				}
				queue = append(queue, next)
			} else {
				upVelocity, downVelocity := grid.Coords{X: 0, Y: -1 }, grid.Coords{X: 0, Y: 1 }
				upPosition, downPosition := grid.Coords{X: current.position.X + upVelocity.X, Y: current.position.Y + upVelocity.Y}, grid.Coords{X: current.position.X + downVelocity.X, Y: current.position.Y + downVelocity.Y}
				up, down := beam{ position: upPosition, velocity: upVelocity }, beam{ position: downPosition, velocity: downVelocity }
				queue = append(queue, up, down)
			}
		}
	}

	return energized.Size()
}
