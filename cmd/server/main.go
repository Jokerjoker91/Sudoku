package main

import (
	"log"
	"net/http"
	"os"
	"sudoku-app/internal/handlers"

	"time"

	"github.com/rs/cors"
)

// Middleware to log client analytics (IP, user agent, and request time)
func analyticsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get client IP address (might need to adjust for proxies in production)
		clientIP := r.RemoteAddr

		// Get the User-Agent (browser and device information)
		userAgent := r.UserAgent()

		// Get the current time of the request
		requestTime := time.Now().Format(time.RFC1123)

		// Log the analytics data
		log.Printf("Client IP: %s, User-Agent: %s, Request Time: %s, Requested Path: %s\n", clientIP, userAgent, requestTime, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {

	// Enable CORS for the frontend hosted on GitHub Pages
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"https://jokerjoker91.github.io"}, // Allow requests from GitHub Pages
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Allow specific methods
        AllowedHeaders: []string{"Content-Type"}, // Allow specific headers
    })

	mux := http.NewServeMux()

    // Sudoku-related handlers
    mux.Handle("/generate", analyticsMiddleware(http.HandlerFunc(handlers.GenerateHandler)))
    mux.Handle("/solve", analyticsMiddleware(http.HandlerFunc(handlers.SolveHandler)))
    mux.Handle("/validate", analyticsMiddleware(http.HandlerFunc(handlers.ValidateHandler)))

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