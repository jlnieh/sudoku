package sudoku

import (
	"fmt"
	"strings"
)

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

// ValuesType is the data type of current status of the grid of sudoku puzzle
type ValuesType map[string]string

const (
	digits = "123456789"
	rows   = "ABCDEFGHI"
	cols   = digits
)

var (
	squares  []string
	unitlist [][]string
	units    map[string][][]string
	peers    map[string][]string
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
	unitlist = make([][]string, 0, 27)
	for _, c := range cols {
		unitlist = append(unitlist, cross(rows, string(c)))
	}
	for _, r := range rows {
		unitlist = append(unitlist, cross(string(r), cols))
	}

	for _, rs := range []string{"ABC", "DEF", "GHI"} {
		for _, cs := range []string{"123", "456", "789"} {
			unitlist = append(unitlist, cross(rs, cs))
		}
	}
}

func initUnits() {
	units = make(map[string][][]string)
	for _, s := range squares {
		units[s] = make([][]string, 0, 3)
		for _, u := range unitlist {
			if stringInSlice(s, u) {
				units[s] = append(units[s], u)
			}
		}
	}
}

func initPeers() {
	peers = make(map[string][]string)
	for _, s := range squares {
		peers[s] = make([]string, 0, 20)
		for _, u := range units[s] {
			for _, e := range u {
				if (e != s) && (!stringInSlice(e, peers[s])) {
					peers[s] = append(peers[s], e)
				}
			}
		}
	}
}

func init() {
	// fmt.Println("init in sudoku.go")
	squares = cross(rows, cols)
	initUnitlist()
	initUnits()
	initPeers()
}

// ################ Parse a Grid ################
func parseGrid(grid string) ValuesType {
	values := make(ValuesType)
	for _, s := range squares {
		values[s] = digits
	}

	for s, d := range gridValues(grid) {
		if strings.Contains(digits, d) && (nil == assign(values, s, d)) {
			return nil
		}
	}
	return values
}

func gridValues(grid string) ValuesType {
	ret := make(ValuesType)
	chars := make([]string, 0, 81)
	for _, c := range grid {
		if (c >= '0' && c <= '9') || (c == '.') {
			chars = append(chars, string(c))
		}
	}
	if len(chars) != 81 {
		panic("Length of the input grid is not correct!")
	}
	for i, c := range chars {
		ret[squares[i]] = string(c)
	}
	return ret
}

// Eliminate all the other values (except d) from values[s] and propagate.
// Return values, except return False if a contradiction is detected.
func assign(values ValuesType, s string, d string) ValuesType {
	// fmt.Println("AI: ", s, d, values[s])
	otherValues := strings.ReplaceAll(values[s], d, "")
	for _, d2 := range otherValues {
		if nil == eliminate(values, s, string(d2)) {
			return nil
		}
	}
	// fmt.Println("AO: ", s, d, values[s])
	return values
}

// Eliminate d from values[s]; propagate when values or places <= 2.
// Return values, except return False if a contradiction is detected.
func eliminate(values ValuesType, s string, d string) ValuesType {
	// fmt.Println("EI: ", s, d, values[s])
	if !strings.Contains(values[s], d) {
		// fmt.Println("Ea: ", s, d, values[s])
		return values // Already eliminated
	}
	values[s] = strings.ReplaceAll(values[s], d, "")

	// (1) If a square s is reduced to one value d2, then eliminate d2 from the peers.
	if len(values[s]) == 0 {
		return nil // Contradiction: removed last value
	} else if len(values[s]) == 1 {
		d2 := values[s]
		for _, s2 := range peers[s] {
			if nil == eliminate(values, s2, d2) {
				return nil
			}
		}
		return values
	}

	// (2) If a unit u is reduced to only one place for a value d, then put it there.
	for _, u := range units[s] {
		var dplaces []string
		for _, s2 := range u {
			if strings.Contains(values[s2], d) {
				dplaces = append(dplaces, s2)
			}
		}
		if len(dplaces) == 0 {
			return nil
		} else if (len(dplaces) == 1) && (len(values[dplaces[0]]) > 1) {
			// fmt.Println("Eb: ", dplaces[0], d, values[dplaces[0]])
			if nil == assign(values, dplaces[0], d) {
				return nil
			}
		}
	}

	// fmt.Println("EO: ", s, d, values[s])
	return values
}

func copyValues(values ValuesType) ValuesType {
	newValues := make(ValuesType)
	for k, v := range values {
		newValues[k] = v
	}
	return newValues
}
func search(values ValuesType) ValuesType {
	if nil == values {
		return nil
	}

	minLen := 10
	minS := ""
	isAllDone := true
	for _, s := range squares {
		l := len(values[s])
		if l > 1 {
			isAllDone = false
			if l < minLen {
				minLen = l
				minS = s
			}
		}
	}
	if isAllDone {
		return values
	}

	for _, d := range values[minS] {
		v := search(assign(copyValues(values), minS, string(d)))
		if v != nil {
			return v
		}
	}
	return nil
}

// Solve the sudoku puzzle
func Solve(grid string) ValuesType {
	return search(parseGrid(grid))
}

// Display the 2-D sudoku puzzle
func Display(values ValuesType) {
	width := 1
	for _, s := range squares {
		if (1 + len(values[s])) > width {
			width = 1 + len(values[s])
		}
	}
	line := strings.Join([]string{strings.Repeat("-", width*3), strings.Repeat("-", width*3), strings.Repeat("-", width*3)}, "+")
	for i, s := range squares {
		fmt.Print(values[s])
		if ((i + 1) % 9) == 0 {
			fmt.Println()
			if (i == 26) || (i == 53) {
				fmt.Println(line)
			}
		} else {
			fmt.Print(strings.Repeat(" ", width-len(values[s])))
			if ((i + 1) % 3) == 0 {
				fmt.Print("|")
			}
		}
	}
}
