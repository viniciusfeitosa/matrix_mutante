package main

// PostData is the struct responsible by translate the data that comes in the mutant endpoint
type PostData struct {
	DNA []string `json:"dna"`
}
