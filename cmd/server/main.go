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

	// Serve static files for the frontend
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	// Start the server
	port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port for local development
    }

    log.Printf("Server starting on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))

	servUrl:= "localhost:" + port

	http.ListenAndServe(servUrl, nil)


}