# Solving Every Sudoku Puzzle 
To learn the Go language, I setup one topic, Sudoku solver, as my first homework. After Googling, I found Dimitri's sudoku github project, https://github.com/dimitri/sudoku, and learned Peter Norvig's Python implementation in http://norvig.com/sudoku.html.


## Sudoku Notation and Preliminary Notions by copying from Dimitri

First we have to agree on some notation. A Sudoku puzzle is a grid of 81 squares; the majority of enthusiasts label the columns 1-9, the rows A-I, and call a collection of nine squares (column, row, or box) a unit and the squares that share a unit the peers. A puzzle leaves some squares blank and fills others with digits, and the whole idea is:

  A puzzle is solved if the squares in each unit are filled with a
  permutation of the digits 1 to 9.
