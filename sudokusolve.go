package sudokupackage

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func Sudoku() {
	input := os.Args[1:]
	valid, grid := BuildBase(input)
	if !valid {
		fmt.Println("Error")
		return
	}
	if !WholeTester(grid) {
		fmt.Println("Error")
		return
	}
	if SudokuSolver(grid) {
		PrintSudoku(grid)
	} else {
		fmt.Println("Error")
	}
}

func SudokuSolver(grid [][]int) bool {
	emptyslots, i, j := FindNextEmpty(grid)
	if !emptyslots {
		return true
	}
	for n := 1; n <= 9; n++ {
		grid[i][j] = n
		if WholeTester(grid) {
			if SudokuSolver(grid) {
				return true
			}
			grid[i][j] = 0
		} else {
			grid[i][j] = 0
		}
	}
	return false
}

func BuildBase(input []string) (bool, [][]int) {
	grid := make([][]int, 9)
	for i := 0; i < 9; i++ {
		grid[i] = make([]int, 9)
	}
	if len(input) != 9 {
		return false, grid
	}
	for i := 0; i < 9; i++ {
		sl := []rune(input[i])
		if len(sl) != 9 {
			return false, grid
		} else {
			for j := 0; j < 9; j++ {
				if 49 <= sl[j] && sl[j] <= 57 {
					grid[i][j] = int(sl[j] - 48)
				} else if sl[j] == '.' {
					grid[i][j] = 0
				} else {
					return false, grid
				}
			}
		}
	}
	return true, grid
}

func PrintSudoku(grid [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune(rune(grid[i][j] + 48))
			if j != 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func FindNextEmpty(grid [][]int) (bool, int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return true, i, j
			}
		}
	}
	return false, 0, 0
}

func WholeTester(grid [][]int) bool {
	overall := true
	for i := 0; i < 9; i++ {
		if !RowTest(grid, i) {
			overall = false
		}
		if !ColumnTest(grid, i) {
			overall = false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !SquareTest(grid, i, j) {
				overall = false
			}
		}
	}
	return overall
}

func RowTest(grid [][]int, index int) bool {
	var store []int
	for i := 0; i < 9; i++ {
		if grid[index][i] != 0 {
			store = append(store, grid[index][i])
		}
	}
	if AppearsTwice(store) {
		return false
	} else {
		return true
	}
}

func ColumnTest(grid [][]int, index int) bool {
	var store []int
	for i := 0; i < 9; i++ {
		if grid[i][index] != 0 {
			store = append(store, grid[i][index])
		}
	}
	if AppearsTwice(store) {
		return false
	} else {
		return true
	}
}

func SquareTest(grid [][]int, index1, index2 int) bool {
	var store []int
	for i := index1 * 3; i < (index1+1)*3; i++ {
		for j := index2 * 3; j < (index2+1)*3; j++ {
			if grid[i][j] != 0 {
				store = append(store, grid[i][j])
			}
		}
	}
	if AppearsTwice(store) {
		return false
	} else {
		return true
	}
}

func AppearsTwice(store []int) bool {
	for i := 0; i < len(store); i++ {
		for j := 0; j < len(store) && j != i; j++ {
			if store[i] == store[j] {
				return true
			}
		}
	}
	return false
}
