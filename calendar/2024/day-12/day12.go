package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/sets"
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

func solvePart2(input []string) int {
	result := 0

	regions := identifyRegions(input)
	for _, r := range regions {
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

			// flood fill current region
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

			r.area = r.points.Size()

			for _, key := range r.points.Iterator() {
				c := grid.ParseCoords(key)
				hasNeighborOrtho, hasNeighborDiagonally := []bool{true, true, true, true}, []bool{true, true, true, true}
				for n := 0; n < len(neighbors); n++ {
					nc := grid.Coords{X: c.X + neighbors[n].X, Y: c.Y + neighbors[n].Y}
					if !r.points.Has(nc.ToString()) {
						r.perimeter++
						hasNeighborOrtho[n] = false
					}
				}

				for n := 0; n < len(diagonals); n++ {
					nc := grid.Coords{X: c.X + diagonals[n].X, Y: c.Y + diagonals[n].Y}
					if !r.points.Has(nc.ToString()) {
						hasNeighborDiagonally[n] = false
					}
				}

				// convex corners
				if !hasNeighborOrtho[0] && !hasNeighborOrtho[2] {
					r.sides++
				}
				if !hasNeighborOrtho[0] && !hasNeighborOrtho[3] {
					r.sides++
				}
				if !hasNeighborOrtho[1] && !hasNeighborOrtho[2] {
					r.sides++
				}
				if !hasNeighborOrtho[1] && !hasNeighborOrtho[3] {
					r.sides++
				}

				// concave corners
				if hasNeighborOrtho[0] && hasNeighborOrtho[2] && !hasNeighborDiagonally[0] {
					r.sides++
				}
				if hasNeighborOrtho[0] && hasNeighborOrtho[3] && !hasNeighborDiagonally[2] {
					r.sides++
				}
				if hasNeighborOrtho[1] && hasNeighborOrtho[2] && !hasNeighborDiagonally[1] {
					r.sides++
				}
				if hasNeighborOrtho[1] && hasNeighborOrtho[3] && !hasNeighborDiagonally[3] {
					r.sides++
				}
			}

			regions = append(regions, r)
		}
	}

	return regions
}

func calculateFencePricing(r region) (int, int) {
	return r.area * r.perimeter, r.area * r.sides
}
