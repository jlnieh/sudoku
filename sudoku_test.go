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
}
