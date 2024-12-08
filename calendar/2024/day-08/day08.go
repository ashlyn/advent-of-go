package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/sets"
)

func main() {
	input := files.ReadFile(8, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	antennas := parseInput(input)
	allAntinodes := sets.New()
	for _, coords := range antennas {
		pairs := generatePairs(coords)
		for _, pair := range pairs {
			antinodes := getAntinodesForPair(pair[0], pair[1])
			if (isInbounds(antinodes[0].X, antinodes[0].Y, input)) {
				allAntinodes.Add(antinodes[0].ToString())
			}
			if isInbounds(antinodes[1].X, antinodes[1].Y, input) {
				allAntinodes.Add(antinodes[1].ToString())
			}
		}
	}

	return allAntinodes.Size()
}

func solvePart2(input []string) int {
	antennas := parseInput(input)
	allAntinodes := sets.New()
	for _, coords := range antennas {
		pairs := generatePairs(coords)
		for _, pair := range pairs {
			antinodes := getAllAntinodesForPair(pair[0], pair[1], input)
			for _, antinode := range antinodes {
				if isInbounds(antinode.X, antinode.Y, input) {
					allAntinodes.Add(antinode.ToString())
				}
			}
		}
	}

	return allAntinodes.Size()
}

func parseInput(input []string) map[byte][]*grid.Coords {
	antennas := map[byte][]*grid.Coords{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			value := input[y][x]
			if value == '.' || value == '#' {
				continue
			}
			antennas[value] = append(antennas[value], &grid.Coords{X: x, Y: y})
		}
	}
	return antennas
}

func getAntinodesForPair(a *grid.Coords, b *grid.Coords) [2]*grid.Coords {
	dx, dy := b.X-a.X, b.Y-a.Y
	return [2]*grid.Coords{
		{X: a.X - dx, Y: a.Y - dy},
		{X: b.X + dx, Y: b.Y + dy},
	}
}

func getAllAntinodesForPair(a *grid.Coords, b *grid.Coords, input []string) []*grid.Coords {
	dx, dy := b.X-a.X, b.Y-a.Y
	antinodes := []*grid.Coords{}

	x, y := a.X, a.Y
	for isInbounds(x, y, input) {
		antinodes = append(antinodes, &grid.Coords{X: x, Y: y})
		x -= dx
		y -= dy
	}

	x, y = b.X, b.Y
	for isInbounds(x, y, input) {
		antinodes = append(antinodes, &grid.Coords{X: x, Y: y})
		x += dx
		y += dy
	}

	return antinodes
}

func generatePairs(antennas []*grid.Coords) [][2]*grid.Coords {
	pairs := make([][2]*grid.Coords, 0)
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			pairs = append(pairs, [2]*grid.Coords{antennas[i], antennas[j]})
		}
	}
	return pairs
}

func isInbounds(x int, y int, input []string) bool {
	return x >= 0 && y >= 0 && y < len(input) && x < len(input[y])
}

func printGrid(input []string, antinodes sets.Set) {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if antinodes.Has(grid.Coords{X: x, Y: y}.ToString()) {
				print("#")
			} else {
				print(string(input[y][x]))
			}
		}
		println()
	}
}