package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMutants(t *testing.T) {
	testCases := []struct {
		input      interface{}
		expected   int
		httpMethod string
	}{
		{
			input:      PostData{DNA: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}},
			expected:   http.StatusOK,
			httpMethod: "POST",
		},
		{
			input:      PostData{DNA: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}},
			expected:   http.StatusMethodNotAllowed,
			httpMethod: "GET",
		},
		{
			input:      PostData{DNA: []string{"ATGCGA", "CAGTGC", "TTCTGT", "AGAATG", "CCCATA", "TCACTG"}},
			expected:   http.StatusForbidden,
			httpMethod: "POST",
		},
		{
			input:      PostData{DNA: []string{"ATG", "CAC", "TTC"}},
			expected:   http.StatusInternalServerError,
			httpMethod: "POST",
		},
		{
			input:      1,
			expected:   http.StatusBadRequest,
			httpMethod: "POST",
		},
	}
	for _, testCase := range testCases {
		body, _ := json.Marshal(testCase.input)
		req, err := http.NewRequest(testCase.httpMethod, "/mutants", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		a := app{}
		a.Initialize(nil)
		a.initializeRoutes()
		handler := http.HandlerFunc(a.mutant)
		handler.ServeHTTP(rec, req)

		if status := rec.Code; status != testCase.expected {
			t.Errorf("Expected: %d Received: %d", testCase.expected, status)
		}
	}
}

func TestStats(t *testing.T) {
	testCases := []struct {
		input      interface{}
		expected   int
		httpMethod string
	}{
		{
			expected:   http.StatusMethodNotAllowed,
			httpMethod: "POST",
		},
	}
	for _, testCase := range testCases {
		req, err := http.NewRequest(testCase.httpMethod, "/stats", nil)
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		a := app{}
		a.Initialize(nil)
		a.initializeRoutes()
		handler := http.HandlerFunc(a.stats)
		handler.ServeHTTP(rec, req)

		if status := rec.Code; status != testCase.expected {
			t.Errorf("Expected: %d Received: %d", testCase.expected, status)
		}
	}
}
