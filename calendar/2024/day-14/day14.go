package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(14, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	size := grid.Coords{ X: 11, Y: 7}
	if len(input) > 12 {
		size = grid.Coords{ X: 101, Y: 103 }
	}
	robots := parseInput(input)
	newRobots := map[grid.Coords][]grid.Coords{}
	for p, velocities := range robots {
		for _, v := range velocities {
			newPosition := move(p, v, size, 100)
			if _, ok := newRobots[newPosition]; !ok {
				newRobots[newPosition] = []grid.Coords{}
			}
			newRobots[newPosition] = append(newRobots[newPosition], v)
		}
	}

	quadrants := []int{0, 0, 0, 0}

	for p, v := range newRobots {
		if p.X < size.X / 2 && p.Y < size.Y / 2 {
			quadrants[0] += len(v)
		} else if p.X > size.X / 2 && p.Y < size.Y / 2 {
			quadrants[1] += len(v)
		} else if p.X < size.X / 2 && p.Y > size.Y / 2 {
			quadrants[2] += len(v)
		} else if p.X > size.X / 2 && p.Y > size.Y / 2 {
			quadrants[3] += len(v)
		}
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func solvePart2(input []string) int {
	size := grid.Coords{ X: 11, Y: 7}
	if len(input) > 12 {
		size = grid.Coords{ X: 101, Y: 103 }
	}

	robots := parseInput(input)
	for seconds := 1; seconds < size.X * size.Y; seconds++ {
		newRobots := map[grid.Coords][]grid.Coords{}
		for p, velocities := range robots {
			for _, v := range velocities {
				newPosition := move(p, v, size, 1)
				if _, ok := newRobots[newPosition]; !ok {
					newRobots[newPosition] = []grid.Coords{}
				}
				newRobots[newPosition] = append(newRobots[newPosition], v)
			}
		}
		image := robotsToString(newRobots, size, false, false)
		if strings.Contains(image, strings.Repeat("#", 10)) {
			println(image)
			return seconds
		}
		
		robots = newRobots
	}

	return -1
}

func parseInput(input []string) map[grid.Coords][]grid.Coords {
	robots := map[grid.Coords][]grid.Coords{}
	coordPattern := regexp.MustCompile(`-?\d+`)
	for i := 0; i < len(input); i++ {
		matches := coordPattern.FindAllString(input[i], -1)
		px, _ := strconv.Atoi(matches[0])
		py, _ := strconv.Atoi(matches[1])
		vx, _ := strconv.Atoi(matches[2])
		vy, _ := strconv.Atoi(matches[3])
		p := grid.Coords{X: px, Y: py}
		if _, ok := robots[p]; !ok {
			robots[p] = []grid.Coords{}
		}
		robots[p] = append(robots[p], grid.Coords{X: vx, Y: vy})
	}
	return robots
}

func move(p grid.Coords, v grid.Coords, bounds grid.Coords, seconds int) grid.Coords {
	x, y := circularIndex(p.X, v.X * seconds, bounds.X), circularIndex(p.Y, v.Y * seconds, bounds.Y)
	return grid.Coords{X: x, Y: y}
}

func circularIndex(index int, increment int, length int) int {
	if increment < 0 {
		return circularBackwards(index, increment, length)
	} else if increment > 0 {
		return circularForwards(index, increment, length)
	}
	return index
}

func circularBackwards(index int, distance int, length int) int {
	return (index + (length + (distance % length))) % length
}

func circularForwards(index int, distance int, length int) int {
	return (index + distance) % length
}

func robotsToString(robots map[grid.Coords][]grid.Coords, size grid.Coords, withGap bool, withCount bool) string {
	output := ""
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			if withGap && (x == size.X / 2 || y == size.Y / 2) {
				output += " "
			} else if _, ok := robots[grid.Coords{X: x, Y: y}]; ok {
				if withCount {
					output += fmt.Sprint(len(robots[grid.Coords{X: x, Y: y}]))
				} else {
					output += "#"
				}
			} else {
				output += "."
			}
		}
		output += "\n"
	}
	return output
}
