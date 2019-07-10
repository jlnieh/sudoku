# Solving Every Sudoku Puzzle 
To learn the Go language, I setup one topic, Sudoku solver, as my first homework. After Googling, I found Dimitri's sudoku github project, https://github.com/dimitri/sudoku, and learned Peter Norvig's Python implementation in http://norvig.com/sudoku.html.


## Sudoku Notation and Preliminary Notions by copying from Dimitri

First we have to agree on some notation. A Sudoku puzzle is a grid of 81 squares; the majority of enthusiasts label the columns 1-9, the rows A-I, and call a collection of nine squares (column, row, or box) a unit and the squares that share a unit the peers. A puzzle leaves some squares blank and fills others with digits, and the whole idea is:

  A puzzle is solved if the squares in each unit are filled with a
  permutation of the digits 1 to 9.

## Performances

### Original Python version

With Dimitri's modification and port to Python 3 on my PC,

    C:\> sudoku.dim.py
    All tests pass.
    Solved 50 of 50 easy puzzles (avg 0.01 secs (199 Hz), max 0.01 secs).
    Solved 95 of 95 hard puzzles (avg 0.02 secs (51 Hz), max 0.10 secs).
    Solved 11 of 11 hardest puzzles (avg 0.01 secs (143 Hz), max 0.01 secs).
    Solved 99 of 99 random puzzles (avg 0.01 secs (177 Hz), max 0.01 secs).

### Golang version
