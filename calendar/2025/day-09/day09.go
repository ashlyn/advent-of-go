package main

import (
	"advent-of-go/utils/colors"
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
	"fmt"
	"sort"
)

func main() {
	input := files.ReadFile(9, 2025, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	coords, _, _ := parseInputAsRedTilesOnly(input)
	pairs := generateAllPairs(coords)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].area > pairs[j].area
	})
	return pairs[0].area
}

// 1302141939 too low
// 2875248828 too high
func solvePart2(input []string) int {
	return checkRectangles(input)
}

type pair struct {
	a, b grid.Coords
	area int
}

func generateAllPairs(coords []grid.Coords) []pair {
	pairs := []pair{}
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			area := getAreaInRectangle(coords[i], coords[j])
			pairs = append(pairs, pair{ a: coords[i], b: coords[j], area: area })
		}
	}
	return pairs
}

func parseInputAsRedTilesOnly(input []string) ([]grid.Coords, grid.Coords, grid.Coords) {
	coords := make([]grid.Coords, len(input))
	minX, maxX, minY, maxY := maths.MaxInt(), 0, maths.MaxInt(), 0
	for i := 0; i < len(input); i++ {
		c := grid.ParseCoords(input[i])
		if c.X < minX {
			minX = c.X
		}
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y < minY {
			minY = c.Y
		}
		if c.Y > maxY {
			maxY = c.Y
		}
		coords[i] = c
	}
	return coords, grid.Coords{ X: minX, Y: minY }, grid.Coords{ X: maxX, Y: maxY }
}

func getAreaInRectangle(a, b grid.Coords) int {
	return (maths.Abs(a.X - b.X) + 1) * (maths.Abs(a.Y - b.Y) + 1)
}

var red, green byte = 'r', 'g'
func parseInputAsRedGreenTiles(input []string) (redTiles []grid.Coords, redGreenTiles map[grid.Coords]byte, min, max grid.Coords) {
	redTiles, min, max = parseInputAsRedTilesOnly(input)
	redGreenTiles = make(map[grid.Coords]byte)
	for r := 0; r < len(redTiles); r++ {
		redGreenTiles[redTiles[r]] = red
	}

	for r := 0; r < len(redTiles); r++ {
		between := []grid.Coords{}
		if r < len(redTiles) - 1 {
			between = getTilesBetween(redTiles[r], redTiles[r+1])
		} else {
			between = getTilesBetween(redTiles[r], redTiles[0])
		}
		for b := 0; b < len(between); b++ {
			redGreenTiles[between[b]] = green
		}
	}

	fmt.Println("built edges, starting raycasting")
	insidesFound := 0

	for y := min.Y; y <= max.Y; y++ {
		onEdge := false
		edgeCount := 0
		for x := min.X; x <= max.X; x++ {
			c := grid.Coords{ X: x, Y: y }
			_, exists := redGreenTiles[c]
			if onEdge && !exists {
				onEdge = false
				edgeCount++
			}
			if exists {
				onEdge = true
			}
			if edgeCount % 2 == 1 && !exists {
				insidesFound++
				if insidesFound % 1000 == 0 {
					fmt.Println(insidesFound, min, max, x, y)
				}
				redGreenTiles[c] = green
			}
		}
	}

	fmt.Println("finished parsing grid")
	return redTiles, redGreenTiles, min, max
}

func getTilesBetween(a, b grid.Coords) []grid.Coords {
	tiles := []grid.Coords{}
	minX, maxX, minY, maxY := maths.Min(a.X, b.X), maths.Max(a.X, b.X), maths.Min(a.Y, b.Y), maths.Max(a.Y, b.Y)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			c := grid.Coords{ X: x, Y: y }
			if c != a && c != b {
				tiles = append(tiles, c)
			}
		}
	}
	return tiles
}

func printGrid(tiles map[grid.Coords]byte, min, max grid.Coords) {
	for y := min.Y - 1; y <= max.Y + 1; y++ {
		line := ""
		for x := min.X - 1; x <= max.X + 1; x++ {
			c := grid.Coords{ X: x, Y: y }
			if color, exists := tiles[c]; exists {
				if color == red {
					line += colors.RedString("#")
				} else if color == green {
					line += colors.GreenString("X")
				}
			} else {
				line += colors.WhiteString(".")
			}
		}
		fmt.Println(line)
	}
}

func checkRectangles(input []string) int {
	redTiles, min, _ := parseInputAsRedTilesOnly(input)
	edges := make(map[grid.Coords]byte)
	for r := 0; r < len(redTiles); r++ {
		edges[redTiles[r]] = red
	}

	for r := 0; r < len(redTiles); r++ {
		between := []grid.Coords{}
		if r < len(redTiles) - 1 {
			between = getTilesBetween(redTiles[r], redTiles[r+1])
		} else {
			between = getTilesBetween(redTiles[r], redTiles[0])
		}
		for b := 0; b < len(between); b++ {
			edges[between[b]] = green
		}
	}

	pairs := generateAllPairs(redTiles)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].area > pairs[j].area
	})

	for i := 0; i < len(pairs); i++ {
		topLeft := grid.Coords{ X: maths.Min(pairs[i].a.X, pairs[i].b.X), Y: maths.Min(pairs[i].a.Y, pairs[i].b.Y) }
		topRight := grid.Coords{ X: maths.Max(pairs[i].a.X, pairs[i].b.X), Y: maths.Min(pairs[i].a.Y, pairs[i].b.Y) }
		bottomLeft := grid.Coords{ X: maths.Min(pairs[i].a.X, pairs[i].b.X), Y: maths.Max(pairs[i].a.Y, pairs[i].b.Y) }
		bottomRight := grid.Coords{ X: maths.Max(pairs[i].a.X, pairs[i].b.X), Y: maths.Max(pairs[i].a.Y, pairs[i].b.Y) }
		corners := [4]grid.Coords{ topLeft, topRight, bottomLeft, bottomRight }
		cornersOnEdges := 0
		for _, corner := range corners {
			if _, exists := edges[corner]; !exists {
				cornersOnEdges++
			}
		}
		if cornersOnEdges > 1 {
			// rectangles with more than one corner not on an edge can't be valid
			continue
		}
		area := getAreaInRectangle(pairs[i].a, pairs[i].b)

		// raycast across the edges to see if this is a valid rectangle
		valid := true
		topEdgeCount, bottomEdgeCount := 0, 0
		topOnEdge, bottomOnEdge := false, false
		for x := min.X; x <= topRight.X; x++ {
			topY, bottomY := topLeft.Y, bottomLeft.Y
			topC := grid.Coords{ X: x, Y: topY }
			bottomC := grid.Coords{ X: x, Y: bottomY }
			_, topExists := edges[topC]
			_, bottomExists := edges[bottomC]
			if topOnEdge && !topExists {
				topOnEdge = false
				topEdgeCount++
			}
			if topExists {
				topOnEdge = true
			}
			if bottomOnEdge && !bottomExists {
				bottomOnEdge = false
				bottomEdgeCount++
			}
			if bottomExists {
				bottomOnEdge = true
			}
			if x >= topLeft.X && x <= topRight.X && !topExists && topEdgeCount % 2 == 0 {
				valid = false
				break
			}
			if x >= bottomLeft.X && x <= bottomRight.X && !bottomExists && bottomEdgeCount % 2 == 0 {
				valid = false
				break
			}
		}

		leftEdgeCount, rightEdgeCount := 0, 0
		leftOnEdge, rightOnEdge := false, false
		for y := min.Y; y <= bottomLeft.Y; y++ {
			leftX, rightX := topLeft.X, topRight.X
			leftC := grid.Coords{ X: leftX, Y: y }
			rightC := grid.Coords{ X: rightX, Y: y }
			_, leftExists := edges[leftC]
			_, rightExists := edges[rightC]
			if leftOnEdge && !leftExists {
				leftOnEdge = false
				leftEdgeCount++
			}
			if leftExists {
				leftOnEdge = true
			}
			if rightOnEdge && !rightExists {
				rightOnEdge = false
				rightEdgeCount++
			}
			if rightExists {
				rightOnEdge = true
			}
			if y >= topLeft.Y && y <= bottomLeft.Y && !leftExists && leftEdgeCount % 2 == 0 {
				valid = false
				break
			}
			if y >= topRight.Y && y <= bottomRight.Y && !rightExists && rightEdgeCount % 2 == 0 {
				valid = false
				break
			}
		}

		if valid {
			fmt.Println(i)
			return area
		}
	}

	return -1
}