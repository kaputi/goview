package sudoku

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createBoardFromInts(intBoard [][]int) Board {
	board := Board{}
	for rowI, row := range intBoard {
		for colI, val := range row {
			board[rowI][colI].value = val
		}
	}
	return board
}

func printBoard(board Board, what string) {
	fmt.Println("-----------------------------------------")
	ints := [][]string{}
	for _, row := range board {
		rowInt := []string{}
		for _, cell := range row {
			val := "-"
			switch what {
			case "solutionError":
				val = fmt.Sprintf("%t", cell.isSolutionError)
			case "placedError":
				val = fmt.Sprintf("%t", cell.isPlacedError)
			case "placed":
				if cell.value != 0 {
					val = fmt.Sprintf("%d", cell.value)
				}
			}
			rowInt = append(rowInt, val)
		}
		ints = append(ints, rowInt)
	}

	for _, row := range ints {
		fmt.Println(row)
	}

	fmt.Println("-----------------------------------------")
}

var board1 = [][]int{
	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	// -------------------
	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	// -------------------
	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	{0, 0, 0, 0, 8, 0, 0, 7, 9},
}

var solution1 = [][]int{
	{5, 3, 4, 6, 7, 8, 9, 1, 2},
	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	{1, 9, 8, 3, 4, 2, 5, 6, 7},
	{8, 5, 9, 7, 6, 1, 4, 2, 3},
	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 4, 5, 2, 8, 6, 1, 7, 9},
}

var multipleSolutions = [][]int{
	{2, 9, 5, 7, 4, 3, 8, 6, 1},
	{4, 3, 1, 8, 6, 5, 9, 0, 0},
	{8, 7, 6, 1, 9, 2, 5, 4, 3},
	{3, 8, 7, 4, 5, 9, 2, 1, 6},
	{6, 1, 2, 3, 8, 7, 4, 9, 5},
	{5, 4, 9, 2, 1, 6, 7, 3, 8},
	{7, 6, 3, 5, 2, 4, 1, 8, 9},
	{9, 2, 8, 6, 7, 1, 3, 5, 4},
	{1, 5, 4, 9, 3, 8, 6, 0, 0},
}

var unsovable = [][]int{
	{2, 0, 0, 9, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 6, 0},
	{0, 0, 0, 0, 0, 1, 0, 0, 0},
	{5, 0, 2, 6, 0, 0, 4, 0, 7},
	{0, 0, 0, 0, 0, 4, 1, 0, 0},
	{0, 0, 0, 0, 9, 8, 0, 2, 3},
	{0, 0, 0, 0, 0, 3, 0, 8, 0},
	{0, 0, 5, 0, 1, 0, 0, 0, 0},
	{0, 0, 7, 0, 0, 0, 0, 0, 0},
}

// func TestSolve(t *testing.T) {
// 	board := createBoardFromInts(board1)
// 	expect := createBoardFromInts(solution1)
// 	soluition, err := Solve(board)
// 	assert.Nil(t, err)
// 	assert.Equal(t, expect, soluition)
// }

// func TestMultipleSolutions(t *testing.T) {
// 	board := createBoardFromInts(multipleSolutions)
// 	_, err := Solve(board)
// 	assert.EqualError(t, err, "invalid, multiple solutions")
// }

// func TestCountSolutions(t *testing.T) {
// 	board := createBoardFromInts(multipleSolutions)
// 	count := CountSolutions(board)
// 	assert.Equal(t, 2, count)
// }

// func TestUnsolvable(t *testing.T) {
// 	board := createBoardFromInts(unsovable)
// 	_, err := Solve(board)

// 	assert.EqualError(t, err, "unsolvable")
// }

// func TestCreateSolvedBoard(t *testing.T) {
// 	board := createSolvedBoard()
// 	assert.True(t, ValidatePlacements(board))
// }

// func TestCreate(t *testing.T) {
// 	board, expect := GenerateBoard(6)
// 	// printBoard(board)
// 	// printBoard(solution)

// 	solution, err := Solve(board)
// 	assert.Nil(t, err)

// 	assert.Equal(t, expect, solution)
// }

func TestSudoku(t *testing.T) {
	dificulty := 5
	s := NewSudoku(dificulty)

	filled := GetFilledCoords(*s.placed)
	fixedCell := s.GetPlacedCell(filled[0].row, filled[0].col)
	assert.True(t, fixedCell.fixed)

	empty := GetEmptyCoords(*s.placed)
	emptyCell := s.GetPlacedCell(empty[0].row, empty[0].col)
	assert.False(t, emptyCell.fixed)
	assert.Equal(t, 0, emptyCell.value)

	solutionCell := s.GetSolutionCell(filled[0].row, filled[0].col)
	assert.Equal(t, fixedCell.value, solutionCell.value)

	solutionEmpty := s.GetSolutionCell(empty[0].row, empty[0].col)
	solutionError := solutionEmpty.value + 1
	if solutionError == 9 {
		solutionError = 1
	} else if solutionError == 1 {
		solutionError = 9
	}

	s.SetPlacedCell(empty[0].row, empty[0].col, solutionError)
	assert.Equal(t, solutionError, s.GetPlacedCell(empty[0].row, empty[0].col).value)

	assert.True(t, s.GetPlacedCell(empty[0].row, empty[0].col).isSolutionError)

	for i := 0; i < 9; i++ {
		if s.GetPlacedCell(i, empty[0].col).value == 0 {
			s.SetPlacedCell(i, empty[0].col, solutionError)
			assert.True(t, s.GetPlacedCell(i, empty[0].col).isPlacedError)
			break
		}
	}
}
