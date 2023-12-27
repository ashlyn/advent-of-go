package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/polygon"
	"advent-of-go/utils/sets"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(18, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	instructions := parseInstructions(input, parseInstruction)
	perimeter, minCoords, maxCoords := digPerimeter(instructions)
	area := floodfillCalculateArea(perimeter, minCoords, maxCoords)
	return area.Size()
}

func solvePart2(input []string) int {
	instructions := parseInstructions(input, parseHexInstruction)
	vertices := findLagoonVertices(instructions)
	return polygon.Perimeter(vertices) + polygon.InteriorArea(vertices) 
}

func findLagoonVertices(instructions []instruction) []grid.Coords {
	x, y := 0, 0
	vertices := []grid.Coords{}
	for _, inst := range instructions {
		endX, endY := x + (inst.direction.X * inst.spaces), y + (inst.direction.Y * inst.spaces)
		vertices = append(vertices, grid.Coords{ X: endX, Y: endY })
		x, y = endX, endY
	}
	return vertices
}

func digPerimeter(instructions []instruction) (sets.Set, grid.Coords, grid.Coords) {
	maxX, maxY := 0, 0
	minX, minY := math.MaxInt, math.MaxInt
	perimeter := sets.New()
	current := grid.Origin
	for _, inst := range instructions {
		for i := 0; i < inst.spaces; i++ {
			current.X, current.Y = current.X + inst.direction.X, current.Y + inst.direction.Y
			if current.X > maxX {
				maxX = current.X
			}
			if current.Y > maxY {
				maxY = current.Y
			}
			if current.X < minX {
				minX = current.X
			}
			if current.Y < minY {
				minY = current.Y
			}
			perimeter.Add(current.ToString())
		}
	}
	return perimeter, grid.Coords{ X: minX, Y: minY }, grid.Coords{ X: maxX, Y: maxY }
}

func floodfillCalculateArea(perimeter sets.Set, minCoords grid.Coords, maxCoords grid.Coords) sets.Set {
	area := perimeter.Copy()
	queue := []grid.Coords{ { X: 1, Y: 1 } }
	directions := []grid.Coords{ { X: 1, Y: 0 }, { X: -1, Y: 0 }, { X: 0, Y: 1 }, { X: 0, Y: -1 } }
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, direction := range directions {
			next := grid.Coords{ X: current.X + direction.X, Y: current.Y + direction.Y }
			if !area.Has(next.ToString()) && !perimeter.Has(next.ToString()) &&
				next.X >= minCoords.X && next.X <= maxCoords.X &&
				next.Y >= minCoords.Y && next.Y <= maxCoords.Y {
				area.Add(next.ToString())
				queue = append(queue, next)
			}
		}
	}
	return area
}

type instruction struct {
	direction grid.Coords
	spaces int
}
func parseInstructions(input []string, parse func(input string) instruction) []instruction {
	instructions := make([]instruction, len(input))
	for i := 0; i < len(input); i++ {
		instructions[i] = parse(input[i])
	}
	return instructions
}

func parseInstruction(input string) instruction {
	parts := strings.Fields(input)
	direction := grid.Origin
	switch parts[0] {
	case "R":
		direction = grid.Coords{ X: 1, Y: 0 }
	case "L":
		direction = grid.Coords{ X: -1, Y: 0 }
	case "U":
		direction = grid.Coords{ X: 0, Y: -1 }
	case "D":
		direction = grid.Coords{ X: 0, Y: 1 }
	}
	spaces, _ := strconv.Atoi(parts[1])
	return instruction{
		direction: direction,
		spaces: spaces,
	}
}

func parseHexInstruction(input string) instruction {
	parts := strings.Fields(input)
	hex := parts[2][2:len(parts[2]) - 1]
	direction := grid.Origin
	switch hex[5] {
	case '0':
		direction = grid.Coords{ X: 1, Y: 0 }
	case '2':
		direction = grid.Coords{ X: -1, Y: 0 }
	case '3':
		direction = grid.Coords{ X: 0, Y: -1 }
	case '1':
		direction = grid.Coords{ X: 0, Y: 1 }
	}
	spaces, _ := strconv.ParseInt(hex[:5], 16, 64)
	return instruction{
		direction: direction,
		spaces: int(spaces),
	}
}
