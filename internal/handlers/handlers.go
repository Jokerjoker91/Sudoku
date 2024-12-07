package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sudoku-app/internal/sudoku"
)

// ValidateHandler checks if the inserted number is correct in the specified position
func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Grid   [][]int `json:"grid"`
		Row    int     `json:"row"`
		Col    int     `json:"col"`
		Number int     `json:"number"`
	}
	
	// Parse the request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Check if the number is valid at the specified position
	isValid := sudoku.IsValid(req.Grid, req.Row, req.Col, req.Number)

	// Return the validation result
	json.NewEncoder(w).Encode(map[string]bool{"valid": isValid})
}

// GenerateHandler handles Sudoku puzzle generation
func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	difficulty := 9 // Default difficulty level
	if d, ok := r.URL.Query()["difficulty"]; ok && len(d) > 0 {
		_, err := fmt.Sscanf(d[0], "%d", &difficulty)
		if err != nil || difficulty < 1 || difficulty > 9 {
			http.Error(w, "Invalid difficulty level", http.StatusBadRequest)
			return
		}
	}

	grid := sudoku.GenerateSudoku(difficulty)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grid)
}

// SolveHandler handles requests to solve a Sudoku puzzle
func SolveHandler(w http.ResponseWriter, r *http.Request) {
	// Define a struct to parse the JSON request
	var req struct {
		Grid [][]int `json:"grid"`
	}

	// Decode the JSON request body into the struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Solve the Sudoku puzzle using the Sudoku library
	solution, solvable := sudoku.SolveSudoku(req.Grid)

	// Encode the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"solution": solution,
		"solvable": solvable,
	})
}