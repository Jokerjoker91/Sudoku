package sudoku

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
		if IsValid(grid, row, col, num) {
			count++
		}
	}
	return count
}