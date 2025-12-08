package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/slices"
	"math"
	"sort"
	"strings"
)

func main() {
	input := files.ReadFile(8, 2025, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	coords := parseInput(input)

	rounds := 1000
	circuits, _ := connectCircuits(coords, rounds)
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func solvePart2(input []string) int {
	coords := parseInput(input)
	_, p := connectCircuits(coords, maths.MaxInt())
	return p.a.x * p.b.x
}

func parseInput(input []string) []threeDCoords {
	coords := make([]threeDCoords, len(input))
	for i := 0; i < len(input); i++ {
		parts := strings.Split(input[i], ",")
		asInts := slices.ParseIntsFromStrings(parts)
		coords[i] = threeDCoords{ x: asInts[0], y: asInts[1], z: asInts[2] }
	}
	return coords
}

type threeDCoords struct {
	x, y, z int
}

type pair struct {
	a, b threeDCoords
	distance float64
}

func straightLineDistance(a, b threeDCoords) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2) + math.Pow(float64(a.z-b.z), 2))
}

type circuit map[threeDCoords]bool
func (c circuit) add(coords threeDCoords) {
	c[coords] = true
}

func (c circuit) contains(coords threeDCoords) bool {
	_, exists := c[coords]
	return exists
}

func generateAllPairs(coords []threeDCoords) []pair {
	pairs := []pair{}
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			distance := straightLineDistance(coords[i], coords[j])
			pairs = append(pairs, pair{ a: coords[i], b: coords[j], distance: distance })
		}
	}
	return pairs
}

func connectCircuits(coords []threeDCoords, rounds int) ([]circuit, pair) {
	pairs := generateAllPairs(coords)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})

	circuits := make([]circuit, len(coords))
	for i := 0; i < len(coords); i++ {
		circuits[i] = circuit{}
		circuits[i].add(coords[i])
	}

	var p pair
	for r := 0; r < len(pairs) && len(circuits) > 1 && r < rounds; r++ {
		p = pairs[r]
		var circuitA, circuitB *circuit
		for i := 0; i < len(circuits); i++ {
			if circuits[i].contains(p.a) {
				circuitA = &circuits[i]
			}
			if circuits[i].contains(p.b) {
				circuitB = &circuits[i]
			}
		}

		if (circuitA == circuitB) {
			continue
		}

		for coords := range *circuitB {
				circuitA.add(coords)
		}
		// Remove circuitB
		newCircuits := []circuit{}
		for i := 0; i < len(circuits); i++ {
			if &circuits[i] != circuitB {
				newCircuits = append(newCircuits, circuits[i])
			}
		}
		circuits = newCircuits
	}

	return circuits, p
}
