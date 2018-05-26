package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/viniciusfeitosa/matrix_mutante/db"
	"github.com/viniciusfeitosa/matrix_mutante/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMutants(t *testing.T) {
	testCases := []struct {
		input      interface{}
		expected   int
		httpMethod string
	}{
		{
			input:      models.PostData{DNA: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}},
			expected:   http.StatusOK,
			httpMethod: "POST",
		},
		{
			input:      models.PostData{DNA: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}},
			expected:   http.StatusMethodNotAllowed,
			httpMethod: "GET",
		},
		{
			input:      models.PostData{DNA: []string{"ATGCGA", "CAGTGC", "TTCTGT", "AGAATG", "CCCATA", "TCACTG"}},
			expected:   http.StatusForbidden,
			httpMethod: "POST",
		},
		{
			input:      models.PostData{DNA: []string{"ATG", "CAC", "TTC"}},
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
		a.Initialize(&db.MockDB{})
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
		input        db.DataBase
		bodyExpected string
		codeExpected int
		httpMethod   string
	}{
		{
			input:        &db.MockDB{},
			bodyExpected: `This endpoint works just with method GET`,
			codeExpected: http.StatusMethodNotAllowed,
			httpMethod:   "POST",
		},
		{
			input:        &db.MockDB{Value: `{"count_mutant_dna": 27,"count_human_dna": 18,"ratio": 0.6}`},
			bodyExpected: `{"count_mutant_dna": 27,"count_human_dna": 18,"ratio": 0.6}`,
			codeExpected: http.StatusOK,
			httpMethod:   "GET",
		},
		{
			input:        &db.MockDB{Err: []error{errors.New("test get error")}},
			bodyExpected: `test get error`,
			codeExpected: http.StatusInternalServerError,
			httpMethod:   "GET",
		},
	}
	for _, testCase := range testCases {
		req, err := http.NewRequest(testCase.httpMethod, "/stats", nil)
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		a := app{}
		a.Initialize(testCase.input)
		a.initializeRoutes()
		handler := http.HandlerFunc(a.stats)
		handler.ServeHTTP(rec, req)

		if status := rec.Code; status != testCase.codeExpected {
			t.Errorf("Header Expected: %d Received: %d", testCase.codeExpected, status)
		}

		if testCase.bodyExpected != strings.TrimSpace(rec.Body.String()) {
			t.Errorf("Body Expected: %s Received: %s", testCase.bodyExpected, rec.Body.String())
		}
	}
}
