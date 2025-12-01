package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"go/types"
	"strings"
)

func main() {
	input := files.ReadFile(15, 2024, "\n\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	wh, moves := parseInput(input)
	wh.moveRobot(moves)
	return wh.sumGpsCoordinates()
}

// Deserves some cleaning up and consolidation if time/interest allows
func solvePart2(input []string) int {
	wh, moves := parseInputWithBigBoxes(input)
	wh.moveRobot(moves)
	return wh.sumGpsCoordinates()
}

func (w *warehouse) moveRobot(moves []grid.Coords) {
	for _, move := range moves {
		newPositions := []grid.Coords{{X: w.robot.X + move.X, Y: w.robot.Y + move.Y}}
		foundEmpty := false
		for !foundEmpty && len(newPositions) > 0 {
			if _, ok := w.walls[newPositions[len(newPositions) - 1]]; ok {
				newPositions = []grid.Coords{}
				continue
			}
			if _, ok := w.boxes[newPositions[len(newPositions) - 1]]; ok {
				newPositions = append(newPositions, grid.Coords{X: newPositions[len(newPositions) - 1].X + move.X, Y: newPositions[len(newPositions) - 1].Y + move.Y})
			} else {
				foundEmpty = true
			}
		}
		for i := 0; i < len(newPositions); i++ {
			if i == 0 {
				w.robot = newPositions[i]
				delete(w.boxes, newPositions[i])
			} else {
				w.boxes[newPositions[i]] = types.Nil{}
			}
		}
		// fmt.Printf("Move %c\n", movesCharMap[move])
		// fmt.Println(w.String())
	}
}

type warehouse struct {
	size grid.Coords
	walls map[grid.Coords]types.Nil
	boxes map[grid.Coords]types.Nil
	robot grid.Coords
}

func parseInputWithBigBoxes(input []string) (bigBoxWarehouse, []grid.Coords) {
	w := parseBigBoxWarehouse(builtBigBoxWarehouseInput(input[0]))
	moves := parseMoves(strings.ReplaceAll(input[1], "\n", ""))
	return w, moves
}

func parseInput(input []string) (warehouse, []grid.Coords) {
	w := parseGrid(strings.Split(input[0], "\n"))
	moves := parseMoves(strings.ReplaceAll(input[1], "\n", ""))
	return w, moves
}

func parseGrid(input []string) warehouse {
	warehouse := warehouse{
		walls: make(map[grid.Coords]types.Nil),
		boxes: make(map[grid.Coords]types.Nil),
		robot: grid.Origin,
		size: grid.Coords{X: len(input[0]), Y: len(input)},
	}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == '#' {
				warehouse.walls[grid.Coords{X: x, Y: y}] = types.Nil{}
			} else if input[y][x] == 'O' {
				warehouse.boxes[grid.Coords{X: x, Y: y}] = types.Nil{}
			} else if input[y][x] == '@' {
				warehouse.robot = grid.Coords{X: x, Y: y}
			}
		}
	}
	return warehouse
}

var up, down, left, right = grid.Coords{X: 0, Y: -1}, grid.Coords{X: 0, Y: 1}, grid.Coords{X: -1, Y: 0}, grid.Coords{X: 1, Y: 0}
var movesMap = map[rune]grid.Coords {
	'^': up,
	'v': down,
	'<': left,
	'>': right,
}
var movesCharMap = map[grid.Coords]rune {
	up: '^',
	down: 'v',
	left: '<',
	right: '>',
}
func parseMoves(input string) []grid.Coords {
	moves := make([]grid.Coords, len(input))
	for i, r := range input {
		moves[i] = movesMap[r]
	}
	return moves
}

func (w warehouse) String() string {
	var sb strings.Builder
	for y := 0; y < w.size.Y; y++ {
		for x := 0; x < w.size.X; x++ {
			if _, ok := w.walls[grid.Coords{X: x, Y: y}]; ok {
				sb.WriteRune('#')
			} else if _, ok := w.boxes[grid.Coords{X: x, Y: y}]; ok {
				sb.WriteRune('O')
			} else if w.robot.X == x && w.robot.Y == y {
				sb.WriteRune('@')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func movesString(moves []grid.Coords) string {
	var sb strings.Builder
	for i := 0; i < len(moves); i++ {
		sb.WriteRune(movesCharMap[moves[i]])
	}
	return sb.String()
}

func gpsCoordinate(c grid.Coords) int {
	return (c.Y * 100) + c.X
}

func (w warehouse) sumGpsCoordinates() int {
	result := 0
	for k := range w.boxes {
		result += gpsCoordinate(k)
	}
	return result
}

func builtBigBoxWarehouseInput(input string) []string {
	newInput := input
	replacements := map[string]string {
		"#": "##",
		"O": "[]",
		".": "..",
	}
	for old, new := range replacements {
		newInput = strings.ReplaceAll(newInput, old, new)
	}

	return strings.Split(strings.ReplaceAll(newInput, "@", "@."), "\n")
}

type bigBoxWarehouse struct {
	size grid.Coords
	walls map[grid.Coords]types.Nil
	boxes map[grid.Coords]types.Nil
	robot grid.Coords
}

func parseBigBoxWarehouse(input []string) bigBoxWarehouse {
	warehouse := bigBoxWarehouse{
		walls: make(map[grid.Coords]types.Nil),
		boxes: make(map[grid.Coords]types.Nil),
		robot: grid.Origin,
		size: grid.Coords{X: len(input[0]), Y: len(input)},
	}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == '#' {
				warehouse.walls[grid.Coords{X: x, Y: y}] = types.Nil{}
			} else if input[y][x] == '[' {
				warehouse.boxes[grid.Coords{X: x, Y: y}] = types.Nil{}
			} else if input[y][x] == '@' {
				warehouse.robot = grid.Coords{X: x, Y: y}
			}
		}
	}
	return warehouse
}

func (w bigBoxWarehouse) String() string {
	var sb strings.Builder
	for y := 0; y < w.size.Y; y++ {
		for x := 0; x < w.size.X; x++ {
			if _, ok := w.walls[grid.Coords{X: x, Y: y}]; ok {
				sb.WriteRune('#')
			} else if _, ok := w.boxes[grid.Coords{X: x, Y: y}]; ok {
				sb.WriteString("[]")
				x++
			} else if w.robot.X == x && w.robot.Y == y {
				sb.WriteRune('@')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (w bigBoxWarehouse) robotCanMove(direction grid.Coords) bool {
	c := grid.Coords{X: w.robot.X + direction.X, Y: w.robot.Y + direction.Y}
	if _, ok := w.walls[c]; ok {
		return false
	}
	occupied := w.occupiedByBox(c)
	if !occupied {
		return true
	}
	leftSide := w.getLeft(c)
	// fmt.Printf("Must move box at %v\n", leftSide)
	return w.boxCanMove(leftSide, direction)
}

func (w bigBoxWarehouse) boxCanMove(leftSide grid.Coords, direction grid.Coords) bool {
	newLeft := grid.Coords{X: leftSide.X + direction.X, Y: leftSide.Y + direction.Y}
	newRight := grid.Coords{X: leftSide.X + 1 + direction.X, Y: leftSide.Y + direction.Y}

	_, leftWall := w.walls[newLeft]
	_, rightWall := w.walls[newRight]

	if leftWall || rightWall {
		return false
	}

	leftBoxExists, rightBoxExists := false, false
	if direction != left {
		rightBoxExists = w.occupiedByBox(newRight)
	}
	if direction != right {
		leftBoxExists = w.occupiedByBox(newLeft)
	}

	boxAtLeft := w.getLeft(newLeft)
	boxAtRight := w.getLeft(newRight)

	if !leftBoxExists && !rightBoxExists {
		return true
	} else if leftBoxExists && rightBoxExists {
		// fmt.Printf("Must move box at %v\n", boxAtLeft)
		// fmt.Printf("Must move box at %v\n", boxAtRight)
		return w.boxCanMove(boxAtLeft, direction) && w.boxCanMove(boxAtRight, direction)
	} else if leftBoxExists {
		// fmt.Printf("Must move box at %v\n", boxAtLeft)
		return w.boxCanMove(boxAtLeft, direction)
	}
	// fmt.Printf("Must move box at %v\n", boxAtRight)
	return w.boxCanMove(boxAtRight, direction)
}

func (w bigBoxWarehouse) occupiedByBox(c grid.Coords) bool {
	_, leftSide := w.boxes[c]
	_, rightSide := w.boxes[w.getLeft(c)]
	return leftSide || rightSide
}

func (w bigBoxWarehouse) getLeft(c grid.Coords) grid.Coords {
	if _, ok := w.boxes[c]; ok {
		return c
	}
	return grid.Coords{X: c.X - 1, Y: c.Y}
}

func (w bigBoxWarehouse) moveBox(leftSide grid.Coords, direction grid.Coords) {
	newLeft := grid.Coords{X: leftSide.X + direction.X, Y: leftSide.Y + direction.Y}
	newRight := grid.Coords{X: leftSide.X + 1 + direction.X, Y: leftSide.Y + direction.Y}

	leftBoxExists, rightBoxExists := false, false
	if direction != left {
		rightBoxExists = w.occupiedByBox(newRight)
	}
	if direction != right {
		leftBoxExists = w.occupiedByBox(newLeft)
	}

	boxAtLeft, boxAtRight := w.getLeft(newLeft), w.getLeft(newRight)
	if leftBoxExists {
		w.moveBox(w.getLeft(newLeft), direction)
	}
	if rightBoxExists && boxAtRight != boxAtLeft {
		w.moveBox(w.getLeft(newRight), direction)
	}

	w.boxes[newLeft] = types.Nil{}
	delete(w.boxes, leftSide)
}

func (w bigBoxWarehouse) moveRobot(moves []grid.Coords) {
	for _, move := range moves {
		newPosition := grid.Coords{X: w.robot.X + move.X, Y: w.robot.Y + move.Y}
		// fmt.Printf("Move %c\n", movesCharMap[move])
		canMove := w.robotCanMove(move)
		mustMoveBoxes := w.occupiedByBox(newPosition)
		// println(canMove, mustMoveBoxes)
		if canMove {
			w.robot = newPosition
			if mustMoveBoxes {
				w.moveBox(w.getLeft(w.robot), move)
			}
		}
		// fmt.Println(w.String())
	}
}

func (w bigBoxWarehouse) sumGpsCoordinates() int {
	result := 0
	for k := range w.boxes {
		result += gpsCoordinate(k)
	}
	return result
}
