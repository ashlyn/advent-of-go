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

func pointsOnLineSegment(a, b grid.Coords) []grid.Coords {
	points := []grid.Coords{}
	if a.X == b.X {
		minY, maxY := maths.Min(a.Y, b.Y), maths.Max(a.Y, b.Y)
		for y := minY; y <= maxY; y++ {
			points = append(points, grid.Coords{ X: a.X, Y: y })
		}
	} else if a.Y == b.Y {
		minX, maxX := maths.Min(a.X, b.X), maths.Max(a.X, b.X)
		for x := minX; x <= maxX; x++ {
			points = append(points, grid.Coords{ X: x, Y: a.Y })
		}
	}
	return points
}

// LineSegmentsOverlap determines if two line segments A and B intersect
func LineSegmentsOverlap(a, b, c, d grid.Coords) bool {
	orientations := []bool{
		areCounterClockwise(a, b, c),
		areCounterClockwise(b, c, d),
		areCounterClockwise(a, b, c),
		areCounterClockwise(a, b, d),
	}

	return orientations[0] != orientations[1] && orientations[2] != orientations[3]
}

func areCounterClockwise(a, b, c grid.Coords) bool {
	return (c.Y - a.Y) * (b.X - a.X) > (b.Y - a.Y) * (c.X - a.X)
}
