package sudoku

import "errors"

func CountSolutions(board Board) int {
	// create a copy and init the counter
	solved := board
	counter := 0

	solveHelper(&board, &solved, &counter, true)

	return counter
}

func Solve(board Board) (Board, error) {
	// create a copy
	solved := board
	counter := 0

	solveHelper(&board, &solved, &counter, false)

	if counter == 0 {
		return solved, errors.New("unsolvable")
	}

	if counter != 1 {
		return solved, errors.New("invalid, multiple solutions")
	}

	return solved, nil
}

func ValidatePlacements(board Board) bool {
	return len(GetInvalidCoords(board)) == 0
}

func GetInvalidCoords(board Board) []Coord {
	coords := []Coord{}
	for rowIdx, row := range board {
		for colIdx, cell := range row {
			if cell.value != 0 && !isValidCell(cell.value, rowIdx, colIdx, board) {
				coords = append(coords, Coord{row: rowIdx, col: colIdx})
			}
		}
	}

	return coords
}

func solveHelper(board, solved *Board, counter *int, countAll bool) bool {
	// get the next empty cell
	row, col := getNextEmpty(*board)

	for value := 1; value <= 9; value++ {
		if !isValidCell(value, row, col, *board) {
			continue
		}

		(*board)[row][col].value = value

		if boardIsFull(*board) {
			*counter++
			if *counter > 1 && !countAll {
				return false
			}
			*solved = *board
			break
		} else if solveHelper(board, solved, counter, countAll) {
			return true
		}
	}

	(*board)[row][col].value = 0
	return false
}

func isValidCell(number, row, col int, board Board) bool {
	if number == 0 {
		return true
	}

	// check that number doesnt repeat in row or column
	for i := 0; i < 9; i++ {
		if board[row][i].value == number && i != col {
			return false
		}

		if board[i][col].value == number && i != row {
			return false
		}
	}

	// check that number doesnt repeat in the 3x3 grid
	firstRowIdxInGrid := row - row%3
	firstColIdxInGrid := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			checRow := firstRowIdxInGrid + i
			checkCol := firstColIdxInGrid + j
			if checRow == row && checkCol == col {
				continue
			}
			if board[checRow][checkCol].value == number {
				return false
			}

		}
	}

	return true
}
