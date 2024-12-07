package sudoku

import "sudoku-app/internal/sudoku"

func Generate(difficulty int) [][]int {
    return sudoku.GenerateSudoku(difficulty)
}

func Solve(grid [][]int) ([][]int, bool) {
    return sudoku.SolveSudoku(grid)
}