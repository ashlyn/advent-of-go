package polygon

import (
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
)

// Area calculates the area of a polygon given a list of vertices
// by cross-multiplying vertices per the Shoelace formula
// https://en.wikipedia.org/wiki/Shoelace_formula
func Area(vertices []grid.Coords) int {
	area := 0
	vertexCount := len(vertices)
	for i := 0; i < vertexCount; i++ {
		nextI := (i + 1) % vertexCount
		area += vertices[i].X * vertices[nextI].Y
		area -= vertices[nextI].X * vertices[i].Y
	}

	return maths.Abs(area / 2)
}

// InteriorArea calculates the interior area of a polygon given a list of vertices
// using Pick's theorem
// https://en.wikipedia.org/wiki/Pick%27s_theorem
func InteriorArea(vertices []grid.Coords) int {
	return Area(vertices) - (Perimeter(vertices) / 2) + 1
}

// Perimeter calculates the perimeter of a polygon given a list of vertices
func Perimeter(vertices []grid.Coords) int {
	perimeter := 0
	vertexCount := len(vertices)
	for i := 0; i < vertexCount; i++ {
		next := (i + 1) % vertexCount
		perimeter += vertices[i].ManhattanDistance(vertices[next])
	}
	return perimeter
}