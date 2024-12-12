package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/sets"
	"fmt"
)

func main() {
	input := files.ReadFile(12, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	regions := identifyRegions(input)
	for _, r := range regions {
		price, _ := calculateFencePricing(r)
		result += price
	}

	return result
}

// 801312 too low
// 804574 too low
func solvePart2(input []string) int {
	result := 0

	regions := identifyRegions(input)
	for _, r := range regions {
		fmt.Println(r)
		_, discount := calculateFencePricing(r)
		result += discount
	}

	return result
}

type region struct {
	crop string
	area int
	perimeter int
	sides int
	points *sets.Set
}

func identifyRegions(input []string) []region {
	// left, right, up, down
	neighbors := []grid.Coords{{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: 1}}
	// top-left, top-right, bottom-left, bottom-right
	diagonals := []grid.Coords{{X: -1, Y: -1}, {X: 1, Y: -1}, {X: -1, Y: 1}, {X: 1, Y: 1}}

	regions := []region{}
	visited := sets.New()

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			c := grid.Coords{X: x, Y: y}
			if visited.Has(c.ToString()) {
				continue
			}

			inRegion := sets.New()
			r := region{crop: input[y][x:x+1], area: 0, perimeter: 0, points: &inRegion}
			queue := sets.New()
			queue.Add(c.ToString())

			for queue.Size() > 0 {
				currentKey := queue.Random()
				queue.Remove(currentKey)
				currentC := grid.ParseCoords(currentKey)

				if currentC.Y < 0 || currentC.Y >= len(input) || currentC.X < 0 || currentC.X >= len(input[currentC.Y]) || visited.Has(currentC.ToString()) {
					continue
				}

				currentCrop := input[currentC.Y][currentC.X:currentC.X+1]
				if currentCrop == r.crop && !r.points.Has(currentC.ToString()) {
					visited.Add(currentC.ToString())
					r.points.Add(currentC.ToString())
					for n := 0; n < len(neighbors); n++ {
						queue.Add(grid.Coords{X: currentC.X + neighbors[n].X, Y: currentC.Y + neighbors[n].Y}.ToString())
					}
				}
			}

			borders := sets.New()
			for _, key := range r.points.Iterator() {
				c := grid.ParseCoords(key)
				for n := 0; n < len(neighbors); n++ {
					nc := grid.Coords{X: c.X + neighbors[n].X, Y: c.Y + neighbors[n].Y}
					if !r.points.Has(nc.ToString()) {
						r.perimeter++
						borders.Add(nc.ToString())
					}
				}
				for n := 0; n < len(diagonals); n++ {
					nc := grid.Coords{X: c.X + diagonals[n].X, Y: c.Y + diagonals[n].Y}
					if !r.points.Has(nc.ToString()) {
						borders.Add(nc.ToString())
					}
				}
			}

			println(r.crop, " ")
			for _, key := range borders.Iterator() {
				c := grid.ParseCoords(key)
				touchPointsOrtho, touchPointsDiag := []bool{false, false, false, false}, []bool{false, false, false, false}
				touchPointsOrthoN, touchPointsDiagN := 0, 0
				for n := 0; n < len(neighbors); n++ {
					nc := grid.Coords{X: c.X + neighbors[n].X, Y: c.Y + neighbors[n].Y}
					if r.points.Has(nc.ToString()) {
						touchPointsOrtho[n] = true
						touchPointsOrthoN++
					}
				}
				for n := 0; n < len(diagonals); n++ {
					nc := grid.Coords{X: c.X + diagonals[n].X, Y: c.Y + diagonals[n].Y}
					if r.points.Has(nc.ToString()) {
						touchPointsDiag[n] = true
						touchPointsDiagN++
					}
				}
				fmt.Printf("%s: %d %d -> ", c.ToString(), touchPointsOrthoN, touchPointsDiagN)
				if touchPointsOrthoN == 0 {
					// fmt.Printf("%s corner at %v\n", c.ToString(), touchPointsDiagN)
					fmt.Println(touchPointsDiagN)
					r.sides += touchPointsDiagN
				} else if touchPointsOrthoN == 1 &&
					((touchPointsOrtho[0] && (touchPointsDiag[1] || touchPointsDiag[3])) ||
					(touchPointsOrtho[1] && (touchPointsDiag[0] || touchPointsDiag[2])) ||
					(touchPointsOrtho[2] && (touchPointsDiag[2] || touchPointsDiag[3])) ||
					(touchPointsOrtho[3] && (touchPointsDiag[0] || touchPointsDiag[1]))) {
					// adjacent 0 and 0 or 2, 1 and 1 or 3, 2 and 0 or 1, 3 and 2 or 3
					// TODO: this is the issue--where there's an ortho touch point and unrelated diag touch point
					// ortho and non-adjacent diag
					// fmt.Printf("%s weird corner at %v\n", c.ToString(), touchPointsDiagN / 2)
					// r.sides += touchPointsDiagN / 2
					r.sides++
					println("???")
				} else if touchPointsOrthoN == 2 && !(touchPointsOrtho[0] && touchPointsOrtho[1]) && !(touchPointsOrtho[2] && touchPointsOrtho[3]) {
					// fmt.Print(c.ToString())
					// fmt.Println(touchPointsOrtho)
					println("1")
					r.sides++
				}	else if touchPointsOrthoN == 3 {
					// fmt.Printf("%s interior corner (3)\n", c.ToString())
					println("2")
					r.sides += 2
				} else if touchPointsOrthoN == 4 {
					// fmt.Printf("%s surrounded\n", c.ToString())
					println("4")
					r.sides += 4
				} else {
					// fmt.Printf("%s normal %d %d\n", c.ToString(), touchPointsOrthoN, touchPointsDiagN)
					println("0")
					fmt.Println(touchPointsOrtho, touchPointsDiag)
				}
				// if touchPointsOrtho == 0 {
				// 	for n := 0; n < len(diagonals); n++ {
				// 		nc := grid.Coords{X: c.X + diagonals[n].X, Y: c.Y + diagonals[n].Y}
				// 		if r.points.Has(nc.ToString()) {
				// 			r.sides++
				// 		}
				// 	}
				// } else if touchPointsOrtho == 2 && !(
				// 	(r.points.Has(grid.Coords{X: c.X - 1, Y: c.Y}.ToString()) && r.points.Has(grid.Coords{X: c.X + 1, Y: c.Y}.ToString())) ||
				// 	(r.points.Has(grid.Coords{X: c.X, Y: c.Y - 1}.ToString()) && r.points.Has(grid.Coords{X: c.X, Y: c.Y + 1}.ToString()))) {
				// 	r.sides++
				// } else if touchPointsOrtho == 3 {
				// 	r.sides += 2
				// } else if touchPointsOrtho == 4 {
				// 	r.sides += 4
				// }
			}
			println(r.sides)
			println()

			r.area = r.points.Size()
			regions = append(regions, r)
		}
	}

	return regions
}

func identifyRegionsNaive(input []string) map[byte][]int {
	neighbors := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	regions := make(map[byte][]int)

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			crop := input[y][x]
			if _, ok := regions[crop]; !ok {
				regions[crop] = []int{0, 0}
			}
			regions[crop][0]++

			for n := 0; n < len(neighbors); n++ {
				nx := x + neighbors[n][0]
				ny := y + neighbors[n][1]

				if nx >= 0 && nx < len(input[y]) && ny >= 0 && ny < len(input) {
					if input[ny][nx] != crop {
						regions[crop][1]++
					}
				} else {
					regions[crop][1]++
				}
			}
		}
	}

	return regions
}

func calculateFencePricing(r region) (int, int) {
	return r.area * r.perimeter, r.area * r.sides
}
