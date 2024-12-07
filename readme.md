# Sudoku Generator and Solver

This is a web-based Sudoku game that allows users to generate new puzzles, solve existing ones, and validate the correctness of their inputs. The app uses a backend written in Go (Golang) to handle the logic for generating, solving, and validating the Sudoku puzzles.

## Features

- **Generate Sudoku**: Generates a new Sudoku puzzle with a configurable difficulty level.
- **Solve Sudoku**: Solves the currently displayed Sudoku puzzle.
- **Validate Sudoku**: Checks if the current Sudoku grid is valid.
- **Interactive Grid**: Users can edit the puzzle and see immediate validation feedback (blue for correct entries, red for incorrect ones).

## Tech Stack

- **Frontend**: HTML, CSS, JavaScript
- **Backend**: Go (Golang)
- **Deployment**: GitHub Pages for the frontend, local server for the Go backend.

## Getting Started

### Prerequisites

- **Go**: You need Go installed on your machine to run the backend locally.

### Installing

1. **Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/Sudoku.git
   cd Sudoku

   ```

2. **Backend Setup (Go)**:

- Navigate to the cmd/server directory and run the Go server:

  ```bash
  cd cmd/server
  go run main.go

  ```

- The Go server will be available at http://localhost:8080.
