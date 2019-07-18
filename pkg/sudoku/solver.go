package sudoku

import (
	"fmt"
	"log"
	"strings"
)

// Solve Every Sudoku Puzzle
//  See http://norvig.com/sudoku.html
// Throughout this program we have:
//   r is a row
//   c is a column
//   s is a square, e.g. 0~80
//   d is a digit,  e.g. "9"
//   u is a unit,   e.g. [0,9,18,27,36,45,54,63,72]
//   grid is a grid,e.g. 81 non-blank chars, e.g. starting with ".18...7...
//   values is a array of possible values, e.g. {"12349", "8", ...}

// ValuesType is the data type of current status of the grid of sudoku puzzle
type ValuesType []string

const (
	digits = "123456789"
)

var (
	// squares  []int	0~80
	unitlist [][]int
	units    [][][]int
	peers    [][]int
)

func elementInSlice(a int, unit []int) bool {
	for _, b := range unit {
		if b == a {
			return true
		}
	}
	return false
}

func initUnitlist() {
	unitlist = make([][]int, 0, 27)
	for i := 0; i < 9; i++ {
		oneRow := make([]int, 9)
		oneCol := make([]int, 9)
		oneBox := make([]int, 9)
		for j := 0; j < 9; j++ {
			oneRow[j] = i*9 + j
			oneCol[j] = i + 9*j
			oneBox[j] = (i/3)*27 + (i%3)*3 + (j/3)*9 + (j % 3)
		}
		unitlist = append(unitlist, oneRow, oneCol, oneBox)
	}
}

func initUnits() {
	units = make([][][]int, 81)
	for s := 0; s < 81; s++ {
		units[s] = make([][]int, 0, 3)
		for _, u := range unitlist {
			if elementInSlice(s, u) {
				units[s] = append(units[s], u)
			}
		}
	}
}

func initPeers() {
	peers = make([][]int, 81)
	for s := 0; s < 81; s++ {
		peers[s] = make([]int, 0, 20)
		for _, u := range units[s] {
			for _, e := range u {
				if (e != s) && (!elementInSlice(e, peers[s])) {
					peers[s] = append(peers[s], e)
				}
			}
		}
	}
}

func init() {
	// fmt.Println("init in sudoku.go")
	initUnitlist()
	initUnits()
	initPeers()
}

// ################ Parse a Grid ################
func parseGrid(grid string) ValuesType {
	values := make(ValuesType, 81)
	for s := 0; s < 81; s++ {
		values[s] = digits
	}

	tmp := gridValues(grid)
	if tmp == nil {
		return nil
	}
	for s, d := range tmp {
		if strings.Contains(digits, d) && (nil == assign(values, s, d)) {
			return nil
		}
	}
	return values
}

func gridValues(grid string) ValuesType {
	chars := make([]string, 0, 81)
	for _, c := range grid {
		if (c >= '0' && c <= '9') || (c == '.') {
			chars = append(chars, string(c))
		}
	}
	if len(chars) != 81 {
		log.Printf("Length of the input grid(%d) is not correct!\n", len(chars))
		return nil
	}

	ret := make(ValuesType, 81)
	for i, c := range chars {
		ret[i] = string(c)
	}
	return ret
}

func assign(values ValuesType, s int, d string) ValuesType {
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

func eliminate(values ValuesType, s int, d string) ValuesType {
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
	}

	// (2) If a unit u is reduced to only one place for a value d, then put it there.
	for _, u := range units[s] {
		var dplaces []int
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
	newValues := make(ValuesType, 81)
	copy(newValues, values)
	return newValues
}
func search(values ValuesType) ValuesType {
	if nil == values {
		return nil
	}

	minLen := 10
	minS := -1
	isAllDone := true
	for s := 0; s < 81; s++ {
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
	if nil == values {
		fmt.Println("Not a valid puzzle!")
		return
	}

	width := 1
	for s := 0; s < 81; s++ {
		if (1 + len(values[s])) > width {
			width = 1 + len(values[s])
		}
	}
	line := strings.Join([]string{strings.Repeat("-", width*3), strings.Repeat("-", width*3), strings.Repeat("-", width*3)}, "+")
	for s := 0; s < 81; s++ {
		fmt.Print(values[s])
		if ((s + 1) % 9) == 0 {
			fmt.Println()
			if (s == 26) || (s == 53) {
				fmt.Println(line)
			}
		} else {
			fmt.Print(strings.Repeat(" ", width-len(values[s])))
			if ((s + 1) % 3) == 0 {
				fmt.Print("|")
			}
		}
	}
}
