package sudoku

// SolveSudoku returns the solved Sudoku grid
func SolveSudoku(grid [][]int) ([][]int, bool) {
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
	row, col := FindEmptyCell(grid)
	if row == -1 {
		return true // Grid is full
	}

	for num := 1; num <= 9; num++ {
		if IsValid(grid, row, col, num) {
			grid[row][col] = num
			if solve(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
}