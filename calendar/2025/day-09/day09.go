package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
	"fmt"
	"go/types"
	"sort"
	"time"
)

func main() {
	input := files.ReadFile(9, 2025, "\n")
	sw := time.Now()
	println(solvePart1(input))
	fmt.Printf("Solved part 1 in %v\n", time.Since(sw))
	sw = time.Now()
	println(solvePart2(input))
	fmt.Printf("Solved part 2 in %v\n", time.Since(sw))
}

func solvePart1(input []string) int {
	coords := parseCorners(input)
	pairs := generateAllPairs(coords)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].area > pairs[j].area
	})
	return pairs[0].area
}

func solvePart2(input []string) int {
	coords := parseCorners(input)
	pairs := generateAllPairs(coords)
	compressedCoords, compressedPairs, _, _ := compressData(coords, pairs)
	compressedPair := checkRectangles(compressedCoords, compressedPairs)
	return compressedPair.area
}

type pair struct {
	a, b grid.Coords
	area int
}

type compressedPair struct {
	ua, ub grid.Coords
	ca, cb grid.Coords
	area   int
}
func compressData(coords []grid.Coords,pairs []pair) ([]grid.Coords, []compressedPair, []int, []int) {
	uniqueX, uniqueY := map[int]types.Nil{}, map[int]types.Nil{}
	for _, p := range pairs {
		uniqueX[p.a.X] = types.Nil{}
		uniqueY[p.a.Y] = types.Nil{}
		uniqueX[p.b.X] = types.Nil{}
		uniqueY[p.b.Y] = types.Nil{}
	}

	sortedX, sortedY := []int{}, []int{}
	for x := range uniqueX {
		sortedX = append(sortedX, x)
	}
	for y := range uniqueY {
		sortedY = append(sortedY, y)
	}
	sort.Ints(sortedX)
	sort.Ints(sortedY)

	xMapping, yMapping := map[int]int{}, map[int]int{}
	reverseXMapping, reverseYMapping := map[int]int{}, map[int]int{}
	for i, x := range sortedX {
		xMapping[i] = x
		reverseXMapping[x] = i
	}
	for i, y := range sortedY {
		yMapping[i] = y
		reverseYMapping[y] = i
	}

	compressedPairs := make([]compressedPair, len(pairs))
	for i, c := range pairs {
		compressedPairs[i] = compressedPair{
			ua: c.a,
			ub: c.b,
			ca: grid.Coords{ X: reverseXMapping[c.a.X], Y: reverseYMapping[c.a.Y] },
			cb: grid.Coords{ X: reverseXMapping[c.b.X], Y: reverseYMapping[c.b.Y] },
			area: c.area,
		}
	}

	compressedCoords := make([]grid.Coords, len(coords))
	for i, c := range coords {
		compressedCoords[i] = grid.Coords{
			X: reverseXMapping[c.X],
			Y: reverseYMapping[c.Y],
		}
	}

	return compressedCoords, compressedPairs, sortedX, sortedY
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


func parseCorners(input []string) []grid.Coords {
	coords := make([]grid.Coords, len(input))
	for i := 0; i < len(input); i++ {
		c := grid.ParseCoords(input[i])
		coords[i] = c
	}
	return coords
}

func getAreaInRectangle(a, b grid.Coords) int {
	return (maths.Abs(a.X - b.X) + 1) * (maths.Abs(a.Y - b.Y) + 1)
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

var red, green byte = 'r', 'g'
func checkRectangles(corners []grid.Coords, pairs []compressedPair) compressedPair {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].area >= pairs[j].area
	})

	min, max := grid.Coords{}, grid.Coords{}

	tileMap := map[grid.Coords]byte{}
	cache := make(map[grid.Coords]bool)
	for r := 0; r < len(corners); r++ {
		c := corners[r]
		next := (r + 1) % len(corners)
		if c.X < min.X {
			min.X = c.X
		}
		if c.X > max.X {
			max.X = c.X + 1
		}
		if c.Y < min.Y {
			min.Y = c.Y
		}
		if c.Y > max.Y {
			max.Y = c.Y + 1
		}

		between := getTilesBetween(c, corners[next])
		for b := range between {
			tileMap[between[b]] = green
			cache[between[b]] = true
		}
		tileMap[c] = red
		cache[c] = true
	}

	for i := 0; i < len(pairs); i++ {
		p := pairs[i]
		min := grid.Coords{ X: maths.Min(p.ca.X, p.cb.X), Y: maths.Min(p.ca.Y, p.cb.Y) }
		max := grid.Coords{ X: maths.Max(p.ca.X, p.cb.X), Y: maths.Max(p.ca.Y, p.cb.Y) }
		isValid := true
		// check if any of the inner tiles are on the border
		// will not work for pieces that are fully outside (e.g. 1-2 sides on perimeter) as in smaple
		for x := min.X + 1; x < max.X; x++ {
			for y := min.Y + 1; y < max.Y; y++ {
				if _, isBorder := tileMap[grid.Coords{ X: x, Y: y }]; isBorder {
					isValid = false
					continue
				}
			}
		}

		if !isValid {
			continue
		}

		topBorderCount, bottomBorderCount, leftBorderCount, rightBorderCount := 0, 0, 0, 0
		for x := min.X; x <= max.X; x++ {
			if _, exists := tileMap[grid.Coords{ X: x, Y: min.Y }]; exists {
				topBorderCount++
			}
			if _, exists := tileMap[grid.Coords{ X: x, Y: max.Y }]; exists {
				bottomBorderCount++
			}
		}
		for y := min.Y; y <= max.Y; y++ {
			if _, exists := tileMap[grid.Coords{ X: min.X, Y: y }]; exists {
				leftBorderCount++
			}
			if _, exists := tileMap[grid.Coords{ X: max.X, Y: y }]; exists {
				rightBorderCount++
			}
		}
		horizontalLength := max.X - min.X + 1
		verticalLength := max.Y - min.Y + 1
		diffs := [4]int { horizontalLength - topBorderCount, horizontalLength - bottomBorderCount,
			verticalLength - leftBorderCount, verticalLength - rightBorderCount }
		zeroCount := 0
		for _, d := range diffs {
			if d == 0 {
				zeroCount++
			}
		}
		if zeroCount == 2 {
			isValid = false
		}
		if isValid {
			fmt.Println(p)
			return p
		}
	}

	return compressedPair{ area: -1 }
}
