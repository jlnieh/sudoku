package sudoku

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func fromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func isSolved(values ValuesType) bool {
	if values == nil {
		return false
	}

	for _, unit := range unitlist {
		sum := 0
		for _, s := range unit {
			v, err := strconv.Atoi(values[s])
			if err != nil {
				return false
			}
			sum += v
		}
		if sum != 45 {
			return false
		}
	}
	return true
}

func TestInit(t *testing.T) {
	fmt.Println("Testing sudoku init...")
	// fmt.Println(Digits)
	// fmt.Println(Squares)
	if len(squares) != 81 {
		t.Errorf("len(Squares)==%d, want 81", len(squares))
	}
	// fmt.Println(Unitlist)
	if len(unitlist) != 27 {
		t.Errorf("len(Unitlist)==%d, want 27", len(unitlist))
	}

	for _, s := range squares {
		if len(units[s]) != 3 {
			t.Errorf("len(Units[%v])==%d, want 3", s, len(units[s]))
		}
		if len(peers[s]) != 20 {
			t.Errorf("len(Peers[%v])==%d, want 20", s, len(peers[s]))
		}
	}
	fmt.Println("Units[C2] =", units["C2"])
	fmt.Println("Peers[C2] =", peers["C2"])
	fmt.Println("Testing sudoku done.")
}

const grid1 = "003020600900305001001806400008102900700000008006708200002609500800203009005010300"
const grid2 = "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"
const grid3 = ".....1...87.9..6..1.....9.3..16.8.3.........6..63.4.......5......4.2..98.59..7.21"
const hard1 = ".....6....59.....82....8....45........3........6..3.54...325..6.................."
const hard2 = ".....5.8....6.1.43..........1.5........1.6...3.......553.....61........4........."

func TestParseGrid(t *testing.T) {
	fmt.Println("Grid 01: input")
	Display(gridValues(grid1))
	fmt.Println("Grid 01: parsed")
	Display(parseGrid(grid1))

	fmt.Println("Grid 02: input")
	Display(gridValues(grid2))
	fmt.Println("Grid 02: parsed")
	Display(parseGrid(grid2))
}

func TestSolve(t *testing.T) {
	values := Solve(grid2)
	if !isSolved(values) {
		t.Error("Failed to solve the test puzzle: Grid 02")
		Display(values)
	}
	// Display(Solve(grid2))
	// Output:
	// Grid 02: solved
	// 4 1 7 |3 6 9 |8 2 5
	// 6 3 2 |1 5 8 |9 4 7
	// 9 5 8 |7 2 4 |3 1 6
	// ------+------+------
	// 8 2 5 |4 3 7 |1 6 9
	// 7 9 1 |5 8 6 |4 3 2
	// 3 4 6 |9 1 2 |7 5 8
	// ------+------+------
	// 2 8 9 |6 4 3 |5 7 1
	// 5 7 3 |2 9 1 |6 8 4
	// 1 6 4 |8 7 5 |2 9 3
}

func ExampleDisplay() {
	Display(Solve(grid3))
	// Output:
	// 9 6 3 |2 8 1 |7 5 4
	// 8 7 5 |9 4 3 |6 1 2
	// 1 4 2 |7 6 5 |9 8 3
	// ------+------+------
	// 4 9 1 |6 7 8 |2 3 5
	// 3 8 7 |5 9 2 |1 4 6
	// 5 2 6 |3 1 4 |8 7 9
	// ------+------+------
	// 2 1 8 |4 5 9 |3 6 7
	// 7 3 4 |1 2 6 |5 9 8
	// 6 5 9 |8 3 7 |4 2 1
}

func BenchmarkSolveHardest(b *testing.B) {
	benchmarks := []struct {
		name string
		grid string
	}{
		{"Hard01", hard1},
		{"Hard02", hard2},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			values := Solve(bm.grid)
			if !isSolved(values) {
				b.Errorf("Failed to solve puzzle inside %s", bm.name)
			}
			b.N = 1
		})
	}
}

func BenchmarkSolveAll(b *testing.B) {
	benchmarks := []struct {
		name     string
		filename string
		grids    []string
	}{
		{"easy", "Puzzles/easy50.txt", nil},
		{"hard", "Puzzles/top95.txt", nil},
		{"hardest", "Puzzles/hardest.txt", nil},
	}

	for _, bm := range benchmarks {
		if bm.grids == nil && len(bm.filename) > 0 {
			var err error
			bm.grids, err = fromFile(bm.filename)
			if err != nil {
				b.Errorf("Failed to readlines from file(%s): %s", bm.filename, err)
			}
		}

		b.Run(bm.name, func(b *testing.B) {
			for _, grid := range bm.grids {
				values := Solve(grid)
				if !isSolved(values) {
					b.Errorf("Failed to solve puzzle inside %s", bm.filename)
				}
			}
			b.N = len(bm.grids)

		})
	}
}
