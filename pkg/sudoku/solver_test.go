package sudoku

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
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
	t.Logf("Testing sudoku init...")

	t.Log(unitlist)
	if len(unitlist) != 27 {
		t.Errorf("len(Unitlist)==%d, want 27", len(unitlist))
	}

	for s := 0; s < 81; s++ {
		if len(units[s]) != 3 {
			t.Errorf("len(Units[%v])==%d, want 3", s, len(units[s]))
		}
		if len(peers[s]) != 20 {
			t.Errorf("len(Peers[%v])==%d, want 20", s, len(peers[s]))
		}
	}

	t.Log("Units[C2(19)] =", units[19])
	unitC2 := [][]int{
		{0, 1, 2, 9, 10, 11, 18, 19, 20},
		{1, 10, 19, 28, 37, 46, 55, 64, 73},
		{18, 19, 20, 21, 22, 23, 24, 25, 26},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 9; j++ {
			if units[19][i][j] != unitC2[i][j] {
				t.Errorf("Units[19][%d][%d]==%d, want %d", i, j, units[19][i][j], unitC2[i][j])
			}
		}
	}
	t.Log("Peers[C2(19)] =", peers[19])
	peerC2 := []int{0, 1, 2, 9, 10, 11, 18, 20, 28, 37, 46, 55, 64, 73, 21, 22, 23, 24, 25, 26}
	for i := 0; i < 20; i++ {
		if peers[19][i] != peerC2[i] {
			t.Errorf("Peers[19][%d]==%d, want %d", i, peers[19][i], peerC2[i])
		}
	}
	t.Logf("Testing sudoku done.")
}

const grid1 = "003020600900305001001806400008102900700000008006708200002609500800203009005010300"
const grid2 = "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"
const grid3 = ".....1...87.9..6..1.....9.3..16.8.3.........6..63.4.......5......4.2..98.59..7.21"
const hard1 = ".....6....59.....82....8....45........3........6..3.54...325..6.................."
const hard2 = ".....5.8....6.1.43..........1.5........1.6...3.......553.....61........4........."

func TestGridValues(t *testing.T) {
	values := GridValues(grid1)
	if values == nil {
		t.Error("Failed to convert grid into the puzzle: Grid 01")
	}
}

func ExampleGridValues() {
	Display(GridValues(grid1))
	// Output:
	// 0 0 3 |0 2 0 |6 0 0
	// 9 0 0 |3 0 5 |0 0 1
	// 0 0 1 |8 0 6 |4 0 0
	// ------+------+------
	// 0 0 8 |1 0 2 |9 0 0
	// 7 0 0 |0 0 0 |0 0 8
	// 0 0 6 |7 0 8 |2 0 0
	// ------+------+------
	// 0 0 2 |6 0 9 |5 0 0
	// 8 0 0 |2 0 3 |0 0 9
	// 0 0 5 |0 1 0 |3 0 0
}

func TestParseGrid(t *testing.T) {
	values := ParseGrid(grid2)
	if values == nil {
		t.Error("Failed to parse the grid into the puzzle: Grid 02")
	}
}

func ExampleParseGrid() {
	Display(ParseGrid(grid2))
	// Output:
	// 4       1679    12679   |139     2369    269     |8       1239    5
	// 26789   3       1256789 |14589   24569   245689  |12679   1249    124679
	// 2689    15689   125689  |7       234569  245689  |12369   12349   123469
	// ------------------------+------------------------+------------------------
	// 3789    2       15789   |3459    34579   4579    |13579   6       13789
	// 3679    15679   15679   |359     8       25679   |4       12359   12379
	// 36789   4       56789   |359     1       25679   |23579   23589   23789
	// ------------------------+------------------------+------------------------
	// 289     89      289     |6       459     3       |1259    7       12489
	// 5       6789    3       |2       479     1       |69      489     4689
	// 1       6789    4       |589     579     5789    |23569   23589   23689
}

func TestSolve(t *testing.T) {
	values := Solve(grid3)
	if !isSolved(values) {
		t.Error("Failed to solve the test puzzle: Grid 03")
	}
}

func ExampleSolve() {
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

func BenchmarkSolveAll(b *testing.B) {
	benchmarks := []struct {
		name     string
		filename string
		grids    []string
	}{
		{"easy", "../../Puzzles/easy50.txt", nil},
		{"hard", "../../Puzzles/top95.txt", nil},
		{"hardest", "../../Puzzles/hardest.txt", nil},
		{"Hard01", "", []string{hard1}},
		// {"Hard02", "", []string{hard2}},
	}

	for _, bm := range benchmarks {
		if bm.grids == nil && len(bm.filename) > 0 {
			var err error
			bm.grids, err = fromFile(bm.filename)
			if err != nil {
				b.Errorf("Failed to readlines from file(%s): %s", bm.filename, err)
				continue
			}
		}

		b.Run(bm.name, func(b *testing.B) {
			for _, grid := range bm.grids {
				values := Solve(grid)
				if !isSolved(values) {
					b.Errorf("Failed to solve puzzle inside %s", bm.name)
					Display(GridValues(grid))
				}
			}
			b.N = len(bm.grids)
		})
	}
}

func randomPuzzle(N int) string {
	values := make(ValuesType, 81)
	for s := 0; s < 81; s++ {
		values[s] = digits
	}

	rand.Seed(time.Now().UnixNano())
	for s := rand.Intn(81); true; s = rand.Intn(81) {
		if len(values[s]) == 1 {
			continue
		}
		d := string(values[s][rand.Intn(len(values[s]))])
		if nil == assign(values, s, d) {
			break
		}

		l := 0
		ret := ""
		for i := 0; i < 81; i++ {
			if len(values[i]) == 1 {
				l++
				ret += values[i]
			} else {
				ret += "."
			}
		}
		if l > N {
			return ret
		}
	}
	return randomPuzzle(N) // Give up and make a new puzzle
}

func TestRandomPuzzle2(t *testing.T) {
	const badGrid = "4..8.6.....9..3..4.............4........6......3..........35.......194373.....8.."
	values := make(ValuesType, 81)
	for s := 0; s < 81; s++ {
		values[s] = digits
	}

	tmp := GridValues(badGrid)
	if tmp == nil {
		t.Error("Failed to create the puzzle by grid!")
		return
	}
	for s, d := range tmp {
		if strings.Contains(digits, d) && (nil == assign(values, s, d)) {
			if (s != 78) || (d != "8") {
				t.Errorf("Error to assign %s at %d square\n", d, s)
			}
			break
		}
	}
}

func TestRandomPuzzle3(t *testing.T) {
	seeds := [16]struct {
		s int
		d string
	}{
		{47, "3"},
		{40, "6"},
		{72, "3"},
		{17, "4"},
		{31, "4"},
		{78, "8"},
		{5, "6"},
		{67, "1"},
		{71, "7"},
		{3, "8"},
		{59, "5"},
		{69, "4"},
		{11, "9"},
		{0, "4"},
		{68, "9"},
		{14, "3"},
	}
	isErrorFound := false
	values := make(ValuesType, 81)
	for s := 0; s < 81; s++ {
		values[s] = digits
	}

	for _, seed := range seeds {
		if len(values[seed.s]) == 1 {
			continue
		}
		if nil == assign(values, seed.s, seed.d) {
			if (seed.s == 14) && (seed.d == "3") {
				isErrorFound = true
			}
			break
		}
		// Display(values)
	}
	if !isErrorFound {
		t.Error("Failed to detect the bad random puzzle!")
	}
}

func BenchmarkRandomPuzzle(b *testing.B) {
	if b.N < 10 {
		b.N = 100
	}
	nFailed := 0
	for i := 0; i < b.N; i++ {
		v := randomPuzzle(17)
		b.StopTimer()
		r := Solve(v)
		b.StartTimer()
		if r == nil {
			nFailed++

			b.Logf("Failed to solve the random puzzle: %s\n", v)
			// Display(gridValues(v))
			r := ParseGrid(v)
			if r == nil {
				b.Errorf("Failed to generate random puzzle, which is not able to be parsed!")
			} else {
				// Display(r)
			}
		}
	}
	if nFailed > 0 {
		b.Logf("Failed to solve %d of total %d random puzzles!", nFailed, b.N)
	}
}
