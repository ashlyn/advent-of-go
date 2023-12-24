package main

import (
	"advent-of-go/utils/conv"
	"advent-of-go/utils/files"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	input := files.ReadFile(24, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	hailstones := parseInput(input)

	// return findIntersectionsInTwoDimensions(hailstones, [2]int{ 200000000000000, 400000000000000 })
	return findIntersectionsInTwoDimensions(hailstones, [2]int{ 7, 27 })
}

func solvePart2(input []string) int {
	hailstones := parseInput(input)
	startPosition := getStartPositionForIntersection(hailstones)
	slopes := make([]float64, 3)
	for i := 0; i < len(slopes); i++ {
		h := hailstones[i]
		label := fmt.Sprintf("%s", string(rune('t' + i)))
		fmt.Printf("x + vx%s = %d%s + %d\n", label, h.velocity.x, label, h.position.x)
		fmt.Printf("y + vy%s = %d%s + %d\n", label, h.velocity.y, label, h.position.y)
		fmt.Printf("z + vz%s = %d%s + %d\n", label, h.velocity.z, label, h.position.z)
		slopes[i] = float64(h.velocity.y) / float64(h.velocity.x)
	}

	fmt.Printf("slopes: %v\n", slopes)
	return startPosition.x + startPosition.y + startPosition.z
}

func findIntersectionsInTwoDimensions(hailstones []hailstone, bounds[2]int) int {
	result := 0

	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			if hailstones[i].willIntersectTwoDimensions(hailstones[j], bounds) {
				result++
			}
		}
	}

	return result
}

type threeDCoord struct {
	x, y, z int
}

type hailstone struct {
	position threeDCoord
	velocity threeDCoord
}

func parseHailstone(input string) hailstone {
	symbolPattern := regexp.MustCompile(`[@,]`)
	parts := conv.ToIntSlice(strings.Fields(symbolPattern.ReplaceAllString(input, "")))
	return hailstone{
		position: threeDCoord{ x: parts[0], y: parts[1], z: parts[2] },
		velocity: threeDCoord{ x: parts[3], y: parts[4], z: parts[5] },
	}
}

func parseInput(input []string) []hailstone {
	result := make([]hailstone, len(input))
	for i, line := range input {
		result[i] = parseHailstone(line)
	}
	return result
}

func (h hailstone) willIntersectTwoDimensions(other hailstone, bounds [2]int) bool {
	a1, b1 := h.coefficients()
	a2, b2 := other.coefficients()

	if a1 * b2 == b1 * a2 {
		return false
	}
	
	x := (b2 - b1) / (a1 - a2)
	y := a1 * x + b1

	// calculate time each hailstone reaches the intersection point
	// to make sure that it's in the future
	time1 := (x - float64(h.position.x)) / float64(h.velocity.x)
	time2 := (x - float64(other.position.x)) / float64(other.velocity.x)

	lower, upper := float64(bounds[0]), float64(bounds[1])
	return time1 >= 0 && time2 >= 0 &&
		lower <= x && x <= upper &&
		lower <= y && y <= upper
}

func (h hailstone) coefficients() (float64, float64) {
	return float64(h.velocity.y) / float64(h.velocity.x), float64(h.position.y) - (float64(h.position.x) * float64(h.velocity.y)) / float64(h.velocity.x)
}

func getStartPositionForIntersection(hailstones []hailstone) threeDCoord {
	return threeDCoord{ 0, 0, 0 }
}
