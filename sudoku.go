package sudoku

// Solve Every Sudoku Puzzle
//  See http://norvig.com/sudoku.html
// Throughout this program we have:
//   r is a row,    e.g. "A"
//   c is a column, e.g. "3"
//   s is a square, e.g. "A3"
//   d is a digit,  e.g. "9"
//   u is a unit,   e.g. ["A1","B1","C1","D1","E1","F1","G1","H1","I1"]
//   grid is a grid,e.g. 81 non-blank chars, e.g. starting with ".18...7...
//   values is a dict of possible values, e.g. {"A1":"12349", "A2":"8", ...}

import "fmt"

const (
	Digits = "123456789"
	Rows   = "ABCDEFGHI"
	Cols   = Digits
)

var (
	Squares  []string
	Unitlist [][]string
	Units    map[string][][]string
	Peers    map[string][]string
)

func cross(A, B string) []string {
	ret := make([]string, 0, len(A)*len(B))
	for _, cA := range A {
		for _, cB := range B {
			ret = append(ret, string(cA)+string(cB))
		}
	}
	return ret
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func initUnitlist() {
	Unitlist = make([][]string, 0, 27)
	for _, c := range Cols {
		Unitlist = append(Unitlist, cross(Rows, string(c)))
	}
	for _, r := range Rows {
		Unitlist = append(Unitlist, cross(string(r), Cols))
	}

	for _, rs := range []string{"ABC", "DEF", "GHI"} {
		for _, cs := range []string{"123", "456", "789"} {
			Unitlist = append(Unitlist, cross(rs, cs))
		}
	}
}

func initUnits() {
	Units = make(map[string][][]string)
	for _, s := range Squares {
		Units[s] = make([][]string, 0, 3)
		for _, u := range Unitlist {
			if stringInSlice(s, u) {
				Units[s] = append(Units[s], u)
			}
		}
	}
}

func initPeers() {
	Peers = make(map[string][]string)
	for _, s := range Squares {
		Peers[s] = make([]string, 0, 20)
		for _, u := range Units[s] {
			for _, e := range u {
				if (e != s) && (!stringInSlice(e, Peers[s])) {
					Peers[s] = append(Peers[s], e)
				}
			}
		}
	}
}

func init() {
	fmt.Println("init in sudoku.go")
	Squares = cross(Rows, Cols)
	initUnitlist()
	initUnits()
	initPeers()
}
