package models

import "testing"

func TestCollectStats(t *testing.T) {
	testCases := []struct {
		input           []string
		humansExpected  int
		mutantsExpected int
		ratioExpected   float64
	}{
		{
			input:           []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			humansExpected:  6,
			mutantsExpected: 3,
			ratioExpected:   0.3333333333333333,
		},
		{
			input:           []string{"ATGCAA", "CTGTGC", "TTATGT", "AGAAGG", "CCGCTA", "TCACTG"},
			humansExpected:  9,
			mutantsExpected: 0,
			ratioExpected:   0.0,
		},
		{
			input:           []string{"AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA"},
			humansExpected:  0,
			mutantsExpected: 18,
			ratioExpected:   1.0,
		},
	}

	for _, testCase := range testCases {
		matrix, _ := CreateMatrix(testCase.input)
		stats := NewStats(nil).CollectStats(matrix)
		if testCase.humansExpected != stats.CountHumanDna {
			t.Errorf("Humans Expected: %v, Received: %v", testCase.humansExpected, stats.CountHumanDna)
		}
		if testCase.mutantsExpected != stats.CountMutantDna {
			t.Errorf("Mutants Expected: %v, Received: %v", testCase.mutantsExpected, stats.CountMutantDna)
		}
		if testCase.ratioExpected != stats.Ratio {
			t.Errorf("Ratio Expected: %v, Received: %v", testCase.ratioExpected, stats.Ratio)
		}
	}
}
