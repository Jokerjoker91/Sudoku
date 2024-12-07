package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Shuffle an array of numbers
func shuffle(numbers []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
}

// Generate a complete Sudoku grid with randomization
func generateFullGrid(grid [][]int) bool {
	row, col := findEmptyCell(grid)
	if row == -1 {
		return true // Grid is full
	}

	// Create a shuffled array of numbers from 1 to 9
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	shuffle(numbers)

	// Attempt to place each number
	for _, num := range numbers {
		if isValid(grid, row, col, num) {
			grid[row][col] = num
			if generateFullGrid(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
}

// Check if a number is valid in the given position
func isValid(grid [][]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		// Check row and column
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}

	// Check the 3x3 subgrid
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

// Find an empty cell in the grid
func findEmptyCell(grid [][]int) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

// Count the number of solutions for a given Sudoku puzzle
func countSolutions(grid [][]int) int {
	row, col := findEmptyCell(grid)
	if row == -1 {
		return 1 // A valid solution is found
	}

	solutions := 0
	for num := 1; num <= 9; num++ {
		if isValid(grid, row, col, num) {
			grid[row][col] = num
			solutions += countSolutions(grid)
			grid[row][col] = 0

			// Early exit if more than one solution is found
			if solutions > 1 {
				break
			}
		}
	}
	return solutions
}

func removeNumbers(grid [][]int, difficulty int) {
	rand.Seed(time.Now().UnixNano())

	// Determine the target number of cells to leave filled, based on difficulty
	minFilledCells := 81 - difficulty*8 // Higher difficulty leaves fewer cells filled
	if minFilledCells < 17 {
		minFilledCells = 17 // Sudoku must have at least 17 clues to be solvable
	}

	for {
		remainingFilledCells := countFilledCells(grid)
		if remainingFilledCells <= minFilledCells {
			break
		}

		row := rand.Intn(9)
		col := rand.Intn(9)
		if grid[row][col] != 0 {
			// Temporarily remove the number
			temp := grid[row][col]
			grid[row][col] = 0

			// Check if the puzzle remains uniquely solvable
			if countSolutions(grid) != 1 {
				grid[row][col] = temp // Restore the number if multiple solutions exist
			}
		}
	}
}

// Helper function to count the number of filled cells in the grid
func countFilledCells(grid [][]int) int {
	count := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] != 0 {
				count++
			}
		}
	}
	return count
}

// Create a Sudoku puzzle based on difficulty
func generateSudoku(difficulty int) [][]int {
	if difficulty < 1 || difficulty > 9 {
		panic("Difficulty level must be between 1 and 9")
	}

	// Initialize a 9x9 grid with zeros
	grid := make([][]int, 9)
	for i := range grid {
		grid[i] = make([]int, 9)
	}

	// Generate a full grid
	generateFullGrid(grid)

	// Remove numbers to create a puzzle
	removeNumbers(grid, difficulty)

	return grid
}

// Print the Sudoku grid
func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, num := range row {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}

func solveSudoku(grid [][]int) ([][]int, bool) {
	solution := make([][]int, 9)
	for i := range grid {
		solution[i] = make([]int, 9)
		copy(solution[i], grid[i])
	}

	if solve(solution) {
		return solution, true
	}
	return nil, false
}

// Backtracking solver for Sudoku
func solve(grid [][]int) bool {
	row, col := findEmptyCell(grid)
	if row == -1 {
		return true // Grid is full
	}

	for num := 1; num <= 9; num++ {
		if isValid(grid, row, col, num) {
			grid[row][col] = num
			if solve(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
}

// Estimate Sudoku difficulty
func perceiveDifficulty(grid [][]int) int {
	emptyCells := 0
	branchingFactor := 0

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				emptyCells++
				branchingFactor += countValidNumbers(grid, i, j)
			}
		}
	}

	// Calculate a score based on empty cells and branching factor
	score := emptyCells + branchingFactor/9
	if score <= 30 {
		return 1 // Easiest
	} else if score <= 40 {
		return 3
	} else if score <= 50 {
		return 5
	} else if score <= 60 {
		return 7
	} else {
		return 9 // Hardest
	}
}

// Count the number of valid numbers for a given cell
func countValidNumbers(grid [][]int, row, col int) int {
	count := 0
	for num := 1; num <= 9; num++ {
		if isValid(grid, row, col, num) {
			count++
		}
	}
	return count
}

func main() {
	start := time.Now() // Start timing
	difficulty := 7 // Choose a difficulty level between 1 (easiest) and 9 (hardest)
	puzzle := generateSudoku(difficulty)
	fmt.Println("Generated Sudoku Puzzle:")
	printGrid(puzzle)

	fmt.Println("\nNumber of solutions for this puzzle:")
	solutions := countSolutions(puzzle)
	fmt.Println(solutions)

	fmt.Println("\nSolved Sudoku Puzzle:")
	solvedGrid,_:=(solveSudoku(puzzle)) 
	printGrid(solvedGrid)

	fmt.Printf("Percieved difficulty: %d", perceiveDifficulty(puzzle))

	elapsed := time.Since(start) // Measure elapsed time
	fmt.Printf("\nExecution Time: %s\n", elapsed)
}
