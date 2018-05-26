package models

import (
	"errors"
	"github.com/viniciusfeitosa/matrix_mutante/db"
	"testing"
)

func TestCollectStats(t *testing.T) {
	testCases := []struct {
		input           []string
		humansExpected  int
		mutantsExpected int
		errExpected     error
		mockErr         []error
	}{
		{
			input:           []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			humansExpected:  6,
			mutantsExpected: 3,
			errExpected:     nil,
			mockErr:         []error{nil, nil},
		},
		{
			input:           []string{"ATGCAA", "CTGTGC", "TTATGT", "AGAAGG", "CCGCTA", "TCACTG"},
			humansExpected:  9,
			mutantsExpected: 0,
			errExpected:     nil,
			mockErr:         []error{nil, nil},
		},
		{
			input:           []string{"AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA"},
			humansExpected:  0,
			mutantsExpected: 18,
			errExpected:     nil,
			mockErr:         []error{nil, nil},
		},
		{
			input:           []string{"AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA"},
			humansExpected:  0,
			mutantsExpected: 18,
			errExpected:     errors.New("test set error"),
			mockErr:         []error{errors.New("test set error"), nil},
		},
		{
			input:           []string{"AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA"},
			humansExpected:  0,
			mutantsExpected: 18,
			errExpected:     errors.New("test enqueue error"),
			mockErr:         []error{nil, errors.New("test enqueue error")},
		},
	}

	for _, testCase := range testCases {
		matrix, _ := CreateMatrix(testCase.input)
		stats, err := NewStats(&db.MockDB{Err: testCase.mockErr}).CreateStats(matrix)
		if testCase.humansExpected != stats.CountHumanDna {
			t.Errorf("Humans Expected: %v, Received: %v", testCase.humansExpected, stats.CountHumanDna)
		}
		if testCase.mutantsExpected != stats.CountMutantDna {
			t.Errorf("Mutants Expected: %v, Received: %v", testCase.mutantsExpected, stats.CountMutantDna)
		}
		if err != nil && testCase.errExpected.Error() != err.Error() {
			t.Errorf("Err Expected: %v, Received: %v", testCase.errExpected, err)
		}
	}
}

func TestGetStats(t *testing.T) {
	testCases := []struct {
		valueExpected []byte
		errExpected   error
		mockValue     string
		mockErr       []error
	}{
		{
			valueExpected: []byte(`{"count_mutant_dna": 27,"count_human_dna": 18,"ratio": 0.6}`),
			errExpected:   nil,
			mockValue:     `{"count_mutant_dna": 27,"count_human_dna": 18,"ratio": 0.6}`,
			mockErr:       []error{nil},
		},
		{
			valueExpected: []byte{},
			errExpected:   errors.New("test get error"),
			mockValue:     `{"count_mutant_dna": 27,"count_human_dna": 18,"ratio": 0.6}`,
			mockErr:       []error{errors.New("test get error")},
		},
	}

	for _, testCase := range testCases {
		value, err := NewStats(&db.MockDB{Value: testCase.mockValue, Err: testCase.mockErr}).GetStats()
		if string(testCase.valueExpected) != string(value) {
			t.Errorf("Value Expected: %v, Received: %v", testCase.valueExpected, value)
		}
		if err != nil && testCase.errExpected.Error() != err.Error() {
			t.Errorf("Err Expected: %v, Received: %v", testCase.errExpected, err)
		}
	}
}

func TestTruncate(t *testing.T) {
	testCases := []struct {
		input    float32
		expected float32
	}{
		{
			input:    0.55555,
			expected: 0.55,
		},
		{
			input:    1.0,
			expected: 1,
		},
	}

	for _, testCase := range testCases {
		stats := NewStats(&db.MockDB{})
		value := stats.truncate(testCase.input)
		if testCase.expected != value {
			t.Errorf("Expected: %v, Received: %v", testCase.expected, value)
		}
	}
}
