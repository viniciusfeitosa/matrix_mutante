package main

import "testing"

func checkSlices(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestColsChecker(t *testing.T) {
	testCases := []struct {
		input            []string
		expected         []string
		quantityExpected int
	}{
		{
			input:            []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			expected:         []string{"GGGG"},
			quantityExpected: 1,
		},
		{
			input:            []string{"ATGCAA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			expected:         []string{},
			quantityExpected: 0,
		},
		{
			input:            []string{"AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA"},
			quantityExpected: 12,
		},
		{
			input:            []string{"AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA"},
			quantityExpected: 18,
		},
		{
			input:            []string{"AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA"},
			quantityExpected: 6,
		},
		{
			input:            []string{"AAAAA", "AAAAA", "AAAAA", "AAAAA", "AAAAA", "AAAAA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "AAAA", "AAAA"},
			quantityExpected: 5,
		},
		{
			input:            []string{"AAAAA", "AAAAA", "AAAAA"},
			expected:         []string{},
			quantityExpected: 0,
		},
	}

	for _, testCase := range testCases {
		matrix, _ := createMatrix(testCase.input)
		result := matrix.colsChecker()
		if !checkSlices(testCase.expected, result) {
			t.Errorf("Expected: %v, Received: %v", testCase.expected, result)
		}
		if testCase.quantityExpected != len(result) {
			t.Errorf("Expected: %d, Received: %d", testCase.quantityExpected, len(result))
		}
	}
}

func TestRowsChecker(t *testing.T) {
	testCases := []struct {
		input            []string
		expected         []string
		quantityExpected int
	}{
		{
			input:            []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			expected:         []string{"CCCC"},
			quantityExpected: 1,
		},
		{
			input:            []string{"ATGCAA", "CAGTGC", "TTATGT", "AGAAGG", "ACCCTA", "TCACTG"},
			expected:         []string{},
			quantityExpected: 0,
		},
		{
			input:            []string{"AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA", "AAAAAAAAAAAA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA"},
			quantityExpected: 18,
		},
		{
			input:            []string{"AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA"},
			quantityExpected: 12,
		},
		{
			input:            []string{"AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "AAAA", "AAAA", "AAAA"},
			quantityExpected: 6,
		},
		{
			input:            []string{"AAAAA", "AAAAA", "AAAAA", "AAAAA", "AAAAA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "AAAA", "AAAA"},
			quantityExpected: 5,
		},
		{
			input:            []string{"AAA", "AAA", "AAA", "AAA"},
			expected:         []string{},
			quantityExpected: 0,
		},
	}

	for _, testCase := range testCases {
		matrix, _ := createMatrix(testCase.input)
		result := matrix.rowsChecker()
		if !checkSlices(testCase.expected, result) {
			t.Errorf("Expected: %v, Received: %v", testCase.expected, result)
		}
		if testCase.quantityExpected != len(result) {
			t.Errorf("Expected: %d, Received: %d", testCase.quantityExpected, len(result))
		}
	}
}

func TestDiagonalLeftToRightChecker(t *testing.T) {
	testCases := []struct {
		input            []string
		expected         []string
		quantityExpected int
	}{
		{
			input:            []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			expected:         []string{"AAAA"},
			quantityExpected: 1,
		},
		{
			input:            []string{"ATGCGAATGCGA", "CAGTGCATGCGA", "TCATGTATGCGA", "ATCAGGATGCGA", "CCCCAAATGCGA", "TCACCAATGCGA", "TCACTCATGCGA", "TCACTACAGCGA", "TCACTAACACGA", "TTACTAATCAGA", "TCTCTAATGCAA", "TCATTAATGCCA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "CCCC", "CCCC", "TTTT"},
			quantityExpected: 6,
		},
		{
			input:            []string{"CTGCAA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			expected:         []string{},
			quantityExpected: 0,
		},
		{
			input:            []string{"AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "AAAA", "AAAA"},
			quantityExpected: 5,
		},
		{
			input:            []string{"AAAAA", "AAAAA", "AAAAA", "AAAAA", "AAAAA", "AAAAA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "AAAA"},
			quantityExpected: 4,
		},
		{
			input:            []string{"AAAAAAA", "AAAAAAA", "AAAAAAA", "AAAAAAA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "AAAA"},
			quantityExpected: 4,
		},
		{
			input:            []string{"AAA", "AAA", "AAA"},
			expected:         []string{},
			quantityExpected: 0,
		},
	}

	for _, testCase := range testCases {
		matrix, _ := createMatrix(testCase.input)
		result := matrix.diagonalLeftToRightChecker()
		if !checkSlices(testCase.expected, result) {
			t.Errorf("Expected: %v, Received: %v", testCase.expected, result)
		}
		if testCase.quantityExpected != len(result) {
			t.Errorf("Expected: %d, Received: %d", testCase.quantityExpected, len(result))
		}
	}
}

func TestDiagonalRightToLeftChecker(t *testing.T) {
	testCases := []struct {
		input            []string
		expected         []string
		quantityExpected int
	}{
		{
			input:            []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			expected:         []string{},
			quantityExpected: 0,
		},
		{
			input:            []string{"ATGCGAATGCGA", "CAGTGCATGCAA", "TTATGTATGAGA", "AGAAGGATACGA", "CCCCTAAAGCGA", "TCACTGATGCGA", "ATGCGAATGCGA", "CAGTACATGCGA", "TTAAGTATGCGA", "AGAAGGATGCGA", "CACCTAATGCGA", "ACACTGATGCGA"},
			expected:         []string{"AAAA", "AAAA", "AAAA"},
			quantityExpected: 3,
		},
		{
			input:            []string{"AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA", "AAAAAA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "AAAA", "AAAA"},
			quantityExpected: 5,
		},
		{
			input:            []string{"AAAAA", "AAAAA", "AAAAA", "AAAAA", "AAAAA", "AAAAA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "AAAA"},
			quantityExpected: 4,
		},
		{
			input:            []string{"AAAAAAA", "AAAAAAA", "AAAAAAA", "AAAAAAA"},
			expected:         []string{"AAAA", "AAAA", "AAAA", "AAAA"},
			quantityExpected: 4,
		},
		{
			input:            []string{"AAA", "AAA", "AAA"},
			expected:         []string{},
			quantityExpected: 0,
		},
	}

	for _, testCase := range testCases {
		matrix, _ := createMatrix(testCase.input)
		result := matrix.diagonalRightToLeftChecker()
		if !checkSlices(testCase.expected, result) {
			t.Errorf("Expected: %v, Received: %v", testCase.expected, result)
		}
		if testCase.quantityExpected != len(result) {
			t.Errorf("Expected: %d, Received: %d", testCase.quantityExpected, len(result))
		}
	}
}

func TestValidateMatrix(t *testing.T) {
	testCases := []struct {
		input    []string
		expected error
	}{
		{
			input:    []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			expected: nil,
		},
		{
			input:    []string{"AAAAAA", "AAAAAA", "AAAAAA", "AAAAA", "AAAAAA", "AAAAAA"},
			expected: errInvalidDNASequence,
		},
		{
			input:    []string{"AAA", "AAA", "AAA"},
			expected: errSmallerMatrix,
		},
		{
			input:    []string{"AAA"},
			expected: errSmallerMatrix,
		},
		{
			input:    []string{},
			expected: errEmptyMatrix,
		},
	}

	for _, testCase := range testCases {
		_, err := createMatrix(testCase.input)
		if testCase.expected != err {
			t.Errorf("Expected: %v, Received: %v", testCase.expected, err)
		}
	}
}
