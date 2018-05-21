package main

import (
	"errors"
	"fmt"
	"strings"
)

var errEmptyMatrix = errors.New("Matrix empty")
var errSmallerMatrix = errors.New("Matrix is smaller than required (4x3 or 3x4)")
var errInvalidDNASequence = errors.New("Matrix invalid to get a DNA sequence")

const (
	identificatorNumber = 4
	possibleDNA         = "ACGT"
)

var example = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

// Matrix is a data structure about DNA
type Matrix struct {
	grid            [][]string
	cols            int
	rows            int
	biggestDiagonal int
}

func (matrix Matrix) colsChecker() []string {
	var mutants []string

	if matrix.rows < identificatorNumber {
		return nil
	}

	posBlockMoviments := matrix.rows - identificatorNumber
	for i := 0; i < matrix.cols; i++ {
		rowIndex := 0
		for rowIndex <= posBlockMoviments {
			v1 := matrix.grid[rowIndex][i]
			v2 := matrix.grid[rowIndex+1][i]
			v3 := matrix.grid[rowIndex+2][i]
			v4 := matrix.grid[rowIndex+3][i]

			if strings.Count(fmt.Sprint(v1, v2, v3, v4), v1) == 4 {
				mutants = append(mutants, fmt.Sprint(v1, v2, v3, v4))
				rowIndex += identificatorNumber
			} else {
				rowIndex++
			}
		}
	}
	return mutants
}

func (matrix Matrix) rowsChecker() []string {
	var mutants []string

	if matrix.cols < identificatorNumber {
		return nil
	}

	posBlockMoviments := matrix.cols - identificatorNumber
	for i := 0; i < matrix.rows; i++ {
		colIndex := 0
		for colIndex <= posBlockMoviments {
			v1 := matrix.grid[i][colIndex]
			v2 := matrix.grid[i][colIndex+1]
			v3 := matrix.grid[i][colIndex+2]
			v4 := matrix.grid[i][colIndex+3]

			if strings.Count(fmt.Sprint(v1, v2, v3, v4), v1) == 4 {
				mutants = append(mutants, fmt.Sprint(v1, v2, v3, v4))
				colIndex += identificatorNumber
			} else {
				colIndex++
			}
		}
	}
	return mutants
}

func (matrix Matrix) diagonalLeftToRightChecker() []string {
	var mutants []string

	if matrix.biggestDiagonal < identificatorNumber {
		return nil
	}

	// Diagonal per line
	diagonals := matrix.biggestDiagonal - identificatorNumber
	for i := 0; i <= diagonals; i++ {
		colIndex := 0
		rowIndex := i
		for rowIndex <= diagonals {
			if rowIndex+3 >= matrix.rows || colIndex+3 >= matrix.cols {
				break
			}
			v1 := matrix.grid[rowIndex][colIndex]
			v2 := matrix.grid[rowIndex+1][colIndex+1]
			v3 := matrix.grid[rowIndex+2][colIndex+2]
			v4 := matrix.grid[rowIndex+3][colIndex+3]

			if strings.Count(fmt.Sprint(v1, v2, v3, v4), v1) == 4 {
				mutants = append(mutants, fmt.Sprint(v1, v2, v3, v4))
				colIndex += identificatorNumber
				rowIndex += identificatorNumber
			} else {
				colIndex++
				rowIndex++
			}
		}
	}

	// Diagonal per column
	for x, y := 0, 1; x < diagonals; x, y = x+1, y+1 {
		rowIndex := 0
		colIndex := y
		for rowIndex < diagonals {
			if rowIndex+3 >= matrix.rows || colIndex+3 >= matrix.cols {
				break
			}
			v1 := matrix.grid[rowIndex][colIndex]
			v2 := matrix.grid[rowIndex+1][colIndex+1]
			v3 := matrix.grid[rowIndex+2][colIndex+2]
			v4 := matrix.grid[rowIndex+3][colIndex+3]

			if strings.Count(fmt.Sprint(v1, v2, v3, v4), v1) == 4 {
				mutants = append(mutants, fmt.Sprint(v1, v2, v3, v4))
				colIndex += identificatorNumber
				rowIndex += identificatorNumber
			} else {
				colIndex++
				rowIndex++
			}
		}
	}

	return mutants
}

func (matrix Matrix) diagonalRightToLeftChecker() []string {
	var mutants []string

	if matrix.biggestDiagonal < identificatorNumber {
		return nil
	}

	// Diagonal per line
	diagonals := matrix.biggestDiagonal - identificatorNumber
	for i := 0; i <= diagonals; i++ {
		colIndex := matrix.cols - 1
		rowIndex := i
		for rowIndex <= diagonals {
			if rowIndex+3 >= matrix.rows || colIndex-3 < 0 {
				break
			}
			v1 := matrix.grid[rowIndex][colIndex]
			v2 := matrix.grid[rowIndex+1][colIndex-1]
			v3 := matrix.grid[rowIndex+2][colIndex-2]
			v4 := matrix.grid[rowIndex+3][colIndex-3]

			if strings.Count(fmt.Sprint(v1, v2, v3, v4), v1) == 4 {
				mutants = append(mutants, fmt.Sprint(v1, v2, v3, v4))
				colIndex -= identificatorNumber
				rowIndex += identificatorNumber
			} else {
				colIndex--
				rowIndex++
			}
		}
	}

	// Diagonal per column
	for x, y := 0, (matrix.cols - 2); x < diagonals; x, y = x+1, y-1 {
		rowIndex := 0
		colIndex := y
		for rowIndex < diagonals {
			if rowIndex+3 >= matrix.rows || colIndex-3 < 0 {
				break
			}
			v1 := matrix.grid[rowIndex][colIndex]
			v2 := matrix.grid[rowIndex+1][colIndex-1]
			v3 := matrix.grid[rowIndex+2][colIndex-2]
			v4 := matrix.grid[rowIndex+3][colIndex-3]

			if strings.Count(fmt.Sprint(v1, v2, v3, v4), v1) == 4 {
				mutants = append(mutants, fmt.Sprint(v1, v2, v3, v4))
				colIndex -= identificatorNumber
				rowIndex += identificatorNumber
			} else {
				colIndex--
				rowIndex++
			}
		}
	}

	return mutants
}

func validateMatrix(mat []string) error {

	if len(mat) <= 0 {
		return errEmptyMatrix
	}

	if len(mat) < identificatorNumber && len(mat[0]) < identificatorNumber {
		return errSmallerMatrix
	}

	baseRow := len(mat[0])
	for _, m := range mat {
		if len(m) != baseRow {
			return errInvalidDNASequence
		}
	}
	return nil
}

func createMatrix(mat []string) (Matrix, error) {

	if err := validateMatrix(mat); err != nil {
		return Matrix{}, err
	}

	var biggestDiagonal int
	rows := len(mat)
	cols := len(mat[0])
	if cols > rows {
		biggestDiagonal = cols
	} else {
		biggestDiagonal = rows
	}
	grid := [][]string{}
	for _, value := range mat {
		grid = append(grid, strings.Split(value, ""))
	}
	return Matrix{grid: grid, cols: cols, rows: rows, biggestDiagonal: biggestDiagonal}, nil
}
