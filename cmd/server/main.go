package main

import (
	"log"
	"net/http"
	"os"
	"sudoku-app/internal/handlers"
)

func main() {
	// Sudoku-related handlers
	http.HandleFunc("/generate", handlers.GenerateHandler)
	http.HandleFunc("/solve", handlers.SolveHandler)
	http.HandleFunc("/validate", handlers.ValidateHandler)

	// Start the server
	port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port for local development
    }

	log.Printf("Server starting on port %s...\n", port)

	// Serve static files for local development (not for production)
	if port == "8080" { // Local development check (can also use a custom environment variable)
		fs := http.FileServer(http.Dir("./web"))
		http.Handle("/", fs)
		log.Fatal(http.ListenAndServe("localhost:"+port, nil))
	} else {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}
}