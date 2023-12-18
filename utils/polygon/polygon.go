package polygon

import (
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
)

// Area calculates the area of a polygon given a list of vertices
// using the Shoelace formula
// https://en.wikipedia.org/wiki/Shoelace_formula
func Area(vertices []grid.Coords) int {
	area := 0
	vertexCount := len(vertices)
	// Cross-multiply verticies per Shoelace formula
	// https://en.wikipedia.org/wiki/Shoelace_formula
	for i := 0; i < vertexCount; i++ {
		area += vertices[i].X * vertices[(i + 1) % vertexCount].Y
		area -= vertices[(i + 1) % vertexCount].X * vertices[i].Y
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
	for i := 0; i < len(vertices); i++ {
		next := (i + 1) % len(vertices)
		perimeter += vertices[i].ManhattanDistance(vertices[next])
	}
	return perimeter
}