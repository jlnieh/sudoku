package sudoku

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	fmt.Println("Testing sudoku init...")
	// fmt.Println(Digits)
	// fmt.Println(Squares)
	if len(Squares) != 81 {
		t.Errorf("len(Squares)==%d, want 81", len(Squares))
	}
	// fmt.Println(Unitlist)
	if len(Unitlist) != 27 {
		t.Errorf("len(Unitlist)==%d, want 27", len(Unitlist))
	}

	for _, s := range Squares {
		if len(Units[s]) != 3 {
			t.Errorf("len(Units[%v])==%d, want 3", s, len(Units[s]))
		}
		if len(Peers[s]) != 20 {
			t.Errorf("len(Peers[%v])==%d, want 20", s, len(Peers[s]))
		}
	}
	fmt.Println("Units[C2] =", Units["C2"])
	fmt.Println("Peers[C2] =", Peers["C2"])
	fmt.Println("Testing sudoku done.")
}

const grid1 = "003020600900305001001806400008102900700000008006708200002609500800203009005010300"
const grid2 = "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"
const hard1 = ".....6....59.....82....8....45........3........6..3.54...325..6.................."
const hard2 = ".....1...87.9..6..1.....9.3..16.8.3.........6..63.4.......5......4.2..98.59..7.21"

func TestParseGrid(t *testing.T) {
	fmt.Println("Testing parse grid...")

	fmt.Println("Grid 01: input")
	display(gridValues(grid1))
	fmt.Println("Grid 01: parsed")
	display(parseGrid(grid1))

	fmt.Println("Grid 02: input")
	display(gridValues(grid2))
	fmt.Println("Grid 02: parsed")
	display(parseGrid(grid2))
	fmt.Println("Grid 02: solved")
	display(solve(grid2))

	fmt.Println("Hard 01: input")
	display(gridValues(hard1))
	fmt.Println("Hard 01: solved")
	display(solve(hard1))

	fmt.Println("Hard 02: input")
	display(gridValues(hard2))
	fmt.Println("Hard 02: solved")
	display(solve(hard2))

	fmt.Println("Testing parse grid done.")
}
