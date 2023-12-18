package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
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
	return findLagoonArea(parseInstructions(input, parseHexInstruction))
}

func findPolygonPerimeter(instructions []instruction) ([][2]int, int) {
	x, y := 0, 0
	perimeter := 0
	vertices := [][2]int{}
	for _, inst := range instructions {
		endX, endY := x + (inst.direction.X * inst.spaces), y + (inst.direction.Y * inst.spaces)
		perimeter += inst.spaces
		vertices = append(vertices, [2]int{ endX, endY })
		x, y = endX, endY
	}
	return vertices, perimeter
}

func calculatePolygonArea(vertices [][2]int) int {
	area := 0
	vertexCount := len(vertices)
	// Cross-multiply verticies per Shoelace formula
	// https://en.wikipedia.org/wiki/Shoelace_formula
	for i := 0; i < vertexCount; i++ {
		area += vertices[i][0] * vertices[(i + 1) % vertexCount][1]
		area -= vertices[(i + 1) % vertexCount][0] * vertices[i][1]
	}
	return area / 2
}

func findLagoonArea(instructions []instruction) int {
	vertices, perimeter := findPolygonPerimeter(instructions)
	area := calculatePolygonArea(vertices)
	// find interior/enclosed area using Pick's theorem
	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	interiorArea := area - (perimeter / 2) + 1
	return interiorArea + perimeter
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
