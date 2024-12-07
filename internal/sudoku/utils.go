package sudoku

// IsValid checks if a number is valid in the given position
func IsValid(grid [][]int, row, col, num int) bool {
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
func FindEmptyCell(grid [][]int) (int, int) {
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
func CountSolutions(grid [][]int) int {
	row, col := FindEmptyCell(grid)
	if row == -1 {
		return 1 // A valid solution is found
	}

	solutions := 0
	for num := 1; num <= 9; num++ {
		if IsValid(grid, row, col, num) {
			grid[row][col] = num
			solutions += CountSolutions(grid)
			grid[row][col] = 0

			// Early exit if more than one solution is found
			if solutions > 1 {
				break
			}
		}
	}
	return solutions
}