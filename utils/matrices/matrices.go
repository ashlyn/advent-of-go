package matrices

import (
	"advent-of-go/utils/maths"
	"advent-of-go/utils/slices"
	"math"
)

// GaussianElimination performs Gaussian elimination on the given matrix
// Returns the pivot columns and reduced matrix
// https://www.cliffsnotes.com/study-guides/algebra/linear-algebra/linear-systems/gaussian-elimination
// Interchange two rows, Multiply by non-zero constant, or Add multiple of one row to another row
// Goal: reduce the augmented matrix into the triangular form
func GaussianElimination(matrix [][]int) ([]int, [][]int) {
	m := len(matrix)
	if m == 0 {
		return nil, nil
	}
	n := len(matrix[0]) - 1 

	pivotColumns := []int{}
	currentRow := 0

	reducedMatrix := make([][]int, m)
	for i := range matrix {
		reducedMatrix[i] = make([]int, n+1)
		copy(reducedMatrix[i], matrix[i])
	}

	for col := 0; col < n && currentRow < m; col++ {
		pivotRow := -1
		for row := currentRow; row < m; row++ {
			if reducedMatrix[row][col] != 0 {
				pivotRow = row
				break
			}
		}

		if pivotRow == -1 {
			continue
		}

		// swap rows
		reducedMatrix[currentRow], reducedMatrix[pivotRow] = reducedMatrix[pivotRow], reducedMatrix[currentRow]
		pivotColumns = append(pivotColumns, col)

		for row := currentRow + 1; row < m; row++ {
			if reducedMatrix[row][col] != 0 {
				multiplyBy := reducedMatrix[row][col]
				pivotValue := reducedMatrix[currentRow][col]

				for j := col; j <= n; j++ {
					reducedMatrix[row][j] = reducedMatrix[row][j] * pivotValue - reducedMatrix[currentRow][j] * multiplyBy
				}
			}
		}

		currentRow++
	}

	return pivotColumns, reducedMatrix
}

// LimitConfiguration defines configuration options for solving matrices
type LimitConfiguration struct {
	// SearchSpaceExpansionFactor defines how much to expand the search space for a single free variable
	SearchSpaceExpansionFactor int

	// MinimumFreeValue defines the minimum value for free variables
	MinimumFreeValue int

	// MaximumFreeValue defines the maximum value for free variable count over 2
	MaximumFreeValue int
}

// DefaultLimitConfiguration provides default limits for solving matrices
// These values are tuned to work for AoC 2025 Day 10 pt. 2 input and untested w/ negatives
var DefaultLimitConfiguration = LimitConfiguration{
	SearchSpaceExpansionFactor: 2,
	MinimumFreeValue: 0,
	MaximumFreeValue: 200,
}

// Solve modes
// Minimize finds the solution with the minimum sum of variable values
var Minimize string = "minimize"
// Maximize finds the solution with the maximum sum of variable values
var Maximize string = "maximize"

// SolveMatrix solves the given matrix and returns values that solve the system of equations
// Matrix is expected to be in augmented form (i.e. last column is the constants)
func SolveMatrix(matrix [][]int, mode string, limits LimitConfiguration) []int {
	m, n := len(matrix), len(matrix[0]) - 1

	coefficients := make([][]int, m)
	constants := make([]int, m)
	for i := 0; i < m; i++ {
		constants[i] = matrix[i][n]
		coefficients[i] = make([]int, n)
		copy(coefficients[i], matrix[i][:n])
	}

	pivotColumns, reducedMatrix := GaussianElimination(matrix)
	if reducedMatrix == nil {
		return nil
	}

	pivotSet := make(map[int]bool)
	for _, col := range pivotColumns {
		pivotSet[col] = true
	}

	freeVariables := []int{}
	for col := 0; col < n; col++ {
		if !pivotSet[col] {
			freeVariables = append(freeVariables, col)
		}
	}

	bestSolution := make([]int, n)
	bestSum := getDefault(mode)
	if mode == Minimize {
		bestSum = maths.MaxInt()
	}
	var tryRunSolution func(freeValues []int)
	tryRunSolution = func(freeValues []int) {
		solution := make([]int, n)
		for i, row := range freeVariables {
			if i < len(freeValues) {
				solution[row] = freeValues[i]
			}
		}

		for i := len(pivotColumns) - 1; i >= 0; i-- {
			row, col := i, pivotColumns[i]
			constantTerm := reducedMatrix[row][n]

			for c := col + 1; c < n; c++ {
				constantTerm -= reducedMatrix[row][c] * solution[c]
			}

			if reducedMatrix[row][col] == 0 {
				return
			}

			if constantTerm % reducedMatrix[row][col] != 0 {
				return
			}

			value := constantTerm / reducedMatrix[row][col]
			if value < 0 {
				return
			}

			solution[col] = value
		}

		for i := 0; i < m; i++ {
			total := 0
			for j := 0; j < n; j++ {
				if solution[j] > 0 {
					if coefficients[i][j] != 0 {
						total += solution[j]
					}
				}
			}
			if total != constants[i] {
				return
			}
		}

		sum := slices.Sum(solution)
		if bestSum == getDefault(mode) || isBetter(bestSum, sum, mode) {
			bestSum = sum
			copy(bestSolution, solution)
		}
	}

	// Try solutions until best is found satisfying constraints
	// Attempted to simplify the logic here using combinations vs length-based nested loops but it took > 5s to run
	if len(freeVariables) == 0 {
		tryRunSolution([]int{})
	} else if len(freeVariables) == 1 {
		maxConstant := 0
		for i := 0; i < m; i++ {
			if constants[i] > maxConstant {
				maxConstant = constants[i]
			}
		}

		maxConstant *= limits.SearchSpaceExpansionFactor
		for val := limits.MinimumFreeValue; val <= maxConstant; val++ {
			if bestSum != getDefault(mode) && !isBetter(bestSum, val, mode) {
				break
			}
			tryRunSolution([]int{val})
		}
	} else if len(freeVariables) == 2 {
		maxConstant := 0
		for i := limits.MinimumFreeValue; i < m; i++ {
			if constants[i] > maxConstant {
				maxConstant = constants[i]
			}
		}

		maxConstant = maths.Min(limits.MaximumFreeValue, maxConstant)
		for val1 := limits.MinimumFreeValue; val1 <= maxConstant; val1++ {
			for val2 := limits.MinimumFreeValue; val2 <= maxConstant; val2++ {
				currentSum := val1 + val2
				if bestSum != getDefault(mode) && !isBetter(bestSum, currentSum, mode) {
					continue
				}
				tryRunSolution([]int{val1, val2})
			}
		}
	} else if len(freeVariables) == 3 {
		for val1 := limits.MinimumFreeValue; val1 <= limits.MaximumFreeValue; val1++ {
			for val2 := limits.MinimumFreeValue; val2 <= limits.MaximumFreeValue; val2++ {
				for val3 := limits.MinimumFreeValue; val3 <= limits.MaximumFreeValue; val3++ {
					currentSum := val1 + val2 + val3
					if bestSum != getDefault(mode) && !isBetter(bestSum, currentSum, mode) {
						continue
					}
					tryRunSolution([]int{val1, val2, val3})
				}
			}
		}
	} else {
		values := make([]int, len(freeVariables))
		for i := range values {
			values[i] = limits.MinimumFreeValue
		}
		tryRunSolution(values)
	}

	if (bestSum == -1 && mode == Minimize) || (bestSum == -1 && mode == Maximize) {
		return make([]int, n)
	}

	return bestSolution
}

func isBetter(currentBest int, candidate int, mode string) bool {
	switch mode {
		case Minimize:
			return candidate < currentBest
		case Maximize:
			return candidate > currentBest
	}
	return false
}

func getDefault(mode string) int {
	if mode == Minimize {
		return maths.MaxInt()
	}
	return math.MinInt
}
