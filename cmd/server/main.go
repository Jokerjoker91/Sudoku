package main

import (
	"net/http"
	"sudoku-app/internal/handlers"
)

func main() {
	// Sudoku-related handlers
	http.HandleFunc("/generate", handlers.GenerateHandler)
	http.HandleFunc("/solve", handlers.SolveHandler)
	http.HandleFunc("/validate", handlers.ValidateHandler)

	// Serve static files for the frontend
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	// Start the server
	http.ListenAndServe("localhost:8080", nil)
}