package main

import (
	"github.com/jlnieh/sudoku/pkg/sudoku"
)

func main() {
	const grid = ".....1...87.9..6..1.....9.3..16.8.3.........6..63.4.......5......4.2..98.59..7.21"
	sudoku.Display(sudoku.Solve(grid))
}
