package main

import (
	"log"
	"net/http"
	"os"
	"sudoku-app/internal/handlers"

	"github.com/rs/cors"
)

func main() {

	// Enable CORS for the frontend hosted on GitHub Pages
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"https://jokerjoker91.github.io"}, // Allow requests from GitHub Pages
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Allow specific methods
        AllowedHeaders: []string{"Content-Type"}, // Allow specific headers
    })

	mux := http.NewServeMux()

    // Sudoku-related handlers
    mux.HandleFunc("/generate", handlers.GenerateHandler)
    mux.HandleFunc("/solve", handlers.SolveHandler)
    mux.HandleFunc("/validate", handlers.ValidateHandler)

	// Start the server
	port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port for local development
    }

	log.Printf("Server starting on port %s...\n", port)

	// Serve static files for local development (not for production)
	if port == "8080" { // Local development check (can also use a custom environment variable)
		fs := http.FileServer(http.Dir("./web"))
		mux.Handle("/", fs)
		log.Fatal(http.ListenAndServe("localhost:"+port, c.Handler(mux)))
	} else {
		log.Fatal(http.ListenAndServe(":"+port, c.Handler(mux)))
	}
}