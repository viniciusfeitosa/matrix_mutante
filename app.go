package main

import (
	"encoding/json"
	"github.com/viniciusfeitosa/matrix_mutante/db"
	"github.com/viniciusfeitosa/matrix_mutante/models"
	"log"
	"net/http"
)

type app struct {
	DB     db.DB
	Router *http.ServeMux
}

// Initialize create the DB connection and prepare all the routes
func (a *app) Initialize(db db.DB) {
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

	var postData models.PostData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&postData); err != nil {
		http.Error(w, "Invalid resquest payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	matrix, err := models.CreateMatrix(postData.DNA)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stats, err := models.NewStats(a.DB).CreateStats(matrix)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	if stats.CountMutantDna > 0 {
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

	stats := models.NewStats(a.DB)
	response, err := stats.GetStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
