package sudoku

import (
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
	row, col := FindEmptyCell(grid)
	if row == -1 {
		return true // Grid is full
	}

	// Create a shuffled array of numbers from 1 to 9
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	shuffle(numbers)

	// Attempt to place each number
	for _, num := range numbers {
		if IsValid(grid, row, col, num) {
			grid[row][col] = num
			if generateFullGrid(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
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
			if CountSolutions(grid) != 1 {
				grid[row][col] = temp // Restore the number if multiple solutions exist
			}
		}
	}
}

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

// GenerateSudoku creates a Sudoku puzzle based on difficulty
func GenerateSudoku(difficulty int) [][]int {
    grid := make([][]int, 9)
    for i := range grid {
        grid[i] = make([]int, 9)
    }

    rand.Seed(time.Now().UnixNano())
    generateFullGrid(grid)
    removeNumbers(grid, difficulty)
    return grid
}