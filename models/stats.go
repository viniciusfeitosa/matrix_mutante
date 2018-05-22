package models

import (
	"database/sql"
)

// Stats the struct that contains the data status about all the matrix
type Stats struct {
	db             *sql.DB
	CountMutantDna int     `json:"count_mutant_dna" db:"count_mutant_dna"`
	CountHumanDna  int     `json:"count_human_dna" db:"count_human_dna"`
	Ratio          float64 `json:"ratio" db:"ratio"`
}

// NewStats returns a stats instance
func NewStats(db *sql.DB) *Stats {
	return &Stats{db: db}
}

// CollectStats get the information from the matrix and populate the stats
func (s *Stats) CollectStats(matrix Matrix) *Stats {
	var mutants int
	gridSize := matrix.rows * matrix.cols
	c := make(chan []string, 4)
	go matrix.colsChecker(c)
	go matrix.rowsChecker(c)
	go matrix.diagonalLeftToRightChecker(c)
	go matrix.diagonalRightToLeftChecker(c)
	mutants = len(<-c) + len(<-c) + len(<-c) + len(<-c)
	if (mutants * 4) >= gridSize {
		s.CountMutantDna = gridSize / 4
		s.CountHumanDna = 0
		s.Ratio = 1.0
	} else {
		s.CountMutantDna = mutants
		s.CountHumanDna = (gridSize / 4) - mutants
		s.Ratio = float64(mutants) / float64(gridSize/4)
	}
	return s
}
