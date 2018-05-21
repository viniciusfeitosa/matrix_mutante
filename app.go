package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type app struct {
	DB     *sql.DB
	Router *http.ServeMux
}

// Initialize create the DB connection and prepare all the routes
func (a *app) Initialize(db *sql.DB) {
	a.DB = db
	a.Router = http.NewServeMux()
}

func (a *app) initializeRoutes() {
	a.Router.HandleFunc("/mutant", a.mutant)
	a.Router.HandleFunc("/stats", a.stats)
}

// Run initialize the server
func (a *app) Run(addr string) {
	log.Println("Aplication running on", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *app) mutant(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "This endpoint works just with method POST", http.StatusMethodNotAllowed)
		return
	}

	var postData PostData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&postData); err != nil {
		http.Error(w, "Invalid resquest payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	matrix, err := createMatrix(postData.DNA)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var hasMutant int
	hasMutant += len(matrix.rowsChecker())
	hasMutant += len(matrix.colsChecker())
	hasMutant += len(matrix.diagonalLeftToRightChecker())
	hasMutant += len(matrix.diagonalRightToLeftChecker())

	w.Header().Set("Content-Type", "text/plain")
	if hasMutant > 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}

func (a *app) stats(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "This endpoint works just with method GET", http.StatusMethodNotAllowed)
		return
	}
}
