package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/sets"
	"fmt"
)

func main() {
	input := files.ReadFile(10, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	trailheads := findTrailheads(input)
	for _, trailhead := range trailheads {
		score, _ := walkFromTrailhead(input, trailhead)
		result += score
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	trailheads := findTrailheads(input)
	for _, trailhead := range trailheads {
		_, rating := walkFromTrailhead(input, trailhead)
		result += rating
	}

	return result
}

func findTrailheads(input []string) []*grid.Coords {
	trailheads := []*grid.Coords{}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == '0' {
				trailheads = append(trailheads, &grid.Coords{X: x, Y: y})
			}
		}
	}
	
	return trailheads
}

func walkFromTrailhead(input []string, trailhead *grid.Coords) (int, int) {
	visitedPaths := map[string][]grid.Coords{}

	distinctPaths, distinctEnds := sets.New(), sets.New()

	inititalPath := []grid.Coords{*trailhead}
	initialKey := toPathKey(inititalPath)

	visitedPaths[initialKey] = inititalPath

	queue := sets.New()
	queue.Add(initialKey)

	for queue.Size() > 0 {
		currentPathKey := queue.Random()
		queue.Remove(currentPathKey)
		currentPath := visitedPaths[currentPathKey]

		current := currentPath[len(currentPath) - 1]
		value := input[current.Y][current.X]
		
		if value == '9' {
			distinctPaths.Add(currentPathKey)
			distinctEnds.Add(current.ToString())
		} else {
			neighbors := []grid.Coords{
				{ X: current.X + 1, Y: current.Y },
				{ X: current.X - 1, Y: current.Y },
				{ X: current.X, Y: current.Y + 1 },
				{ X: current.X, Y: current.Y - 1 },
			}
			for _, n := range neighbors {
				if n.Y >= 0 && n.Y < len(input) && n.X >= 0 && n.X < len(input[n.Y]) {
					newValue := input[n.Y][n.X]
					if newValue == value + 1 {
						if !pathHasCoords(currentPath, n) {
							newPath := make([]grid.Coords, len(currentPath) + 1)
							copy(newPath, currentPath)
							newPath[len(newPath) - 1] = n
							newPathKey := toPathKey(newPath)
							if _, visited := visitedPaths[newPathKey]; !visited {
								visitedPaths[newPathKey] = newPath
								queue.Add(newPathKey)
							}
						}
					}
				}
			}
		}
	}

	return distinctEnds.Size(), distinctPaths.Size()
}

func toPathKey(path []grid.Coords) string {
	result := ""
	for _, p := range path {
		result += fmt.Sprintf(" %s", p.ToString())
	}
	return result
}

func pathHasCoords(path []grid.Coords, coords grid.Coords) bool {
	for _, p := range path {
		if p.X == coords.X && p.Y == coords.Y {
			return true
		}
	}
	return false
}
