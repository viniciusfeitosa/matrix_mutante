package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/viniciusfeitosa/matrix_mutante/db"
	"log"
	"sync"
)

const (
	statsKey   = "STATS_KEY"
	statsQueue = "STATS_QUEUE"
)

// Stats the struct that contains the data status about all the matrix
type Stats struct {
	db             db.DB
	CountMutantDna int    `json:"count_mutant_dna"`
	CountHumanDna  int    `json:"count_human_dna"`
	Ratio          string `json:"ratio"`
}

// NewStats returns a stats instance
func NewStats(db db.DB) *Stats {
	return &Stats{db: db}
}

// CreateStats get the information from the matrix and populate the stats
func (s *Stats) CreateStats(matrix Matrix) (*Stats, error) {
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
	} else {
		s.CountMutantDna = mutants
		s.CountHumanDna = (gridSize / 4) - mutants
	}

	data, _ := json.Marshal(s)
	uid := uuid.New()
	if err := s.db.SetValue(uid.ID(), string(data)); err != nil {
		return s, err
	}

	if err := s.db.EnqueueValue(statsQueue, uid.ID()); err != nil {
		return s, err
	}
	return s, nil
}

// GetStats is a method that return the stats cosolidated
func (s *Stats) GetStats() ([]byte, error) {
	value, err := s.db.GetValue(statsKey)
	if err != nil {
		return []byte{}, err
	}
	statsBytes, err := json.Marshal(value)
	if err != nil {
		return []byte{}, err
	}
	return statsBytes, err
}

// JoinStatsWorker is the responsible to sumarise the status
func (s *Stats) JoinStatsWorker(numWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			s.workerProcess(id)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}

func (s *Stats) workerProcess(id int) {
	log.Println("Worker:", id)
	for {
		uuid, values, err := s.db.PopQueue(statsQueue, id)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("values from queue:", values)

		stats := Stats{}
		if err := json.Unmarshal([]byte(values), &stats); err != nil {
			s.db.EnqueueValue(statsQueue, uuid)
			continue
		}
		log.Println("stats unmarshaled:", stats)

		value, err := s.db.GetValue(statsKey)
		if err != nil {
			s.db.EnqueueValue(statsQueue, uuid)
			continue
		}

		if len(value) > 0 {
			statsToConsolidate := Stats{}
			if err := json.Unmarshal([]byte(value), &statsToConsolidate); err != nil {
				s.db.EnqueueValue(statsQueue, uuid)
				continue
			}
			statsToConsolidate.CountHumanDna += stats.CountHumanDna
			statsToConsolidate.CountMutantDna += stats.CountMutantDna
			ratio := float64(statsToConsolidate.CountMutantDna*4) / float64((statsToConsolidate.CountMutantDna+statsToConsolidate.CountHumanDna)*4)
			statsToConsolidate.Ratio = fmt.Sprintf("%.2f", ratio)
			data, _ := json.Marshal(statsToConsolidate)
			if err := s.db.SetValue(statsKey, string(data)); err != nil {
				s.db.EnqueueValue(statsQueue, uuid)
				continue
			}
		} else {
			ratio := float64(stats.CountMutantDna*4) / float64((stats.CountMutantDna+stats.CountHumanDna)*4)
			stats.Ratio = fmt.Sprintf("%.2f", ratio)
			data, _ := json.Marshal(stats)
			if err := s.db.SetValue(statsKey, string(data)); err != nil {
				s.db.EnqueueValue(statsQueue, uuid)
				continue
			}
		}

		s.db.DelValue(uuid)
	}
}
