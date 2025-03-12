package main

import (
	"encoding/json"
	"net/http"
)

// Request represents a JSON request containing the number of items to order.
type Request struct {
	Items int `json:"items"`
}

// Response represents the JSON response containing the calculated pack distribution.
type Response struct {
	Packs map[int]int `json:"packs"`
}

// handler processes API requests for pack calculation.
// It handles CORS and parses incoming JSON requests.
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	packs := calculatePacks(req.Items)
	response := Response{Packs: packs}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
