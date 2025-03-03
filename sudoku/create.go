package sudoku

import (
	"math/rand"
)

func GenerateBoard(dificulty int) (Board, Board) {
	board := createSolvedBoard()

	iterations := 0
	removeNumbers(&board, dificulty, &iterations)

	for rowI, row := range board {
		for colI, cell := range row {
			if cell.Value != 0 {
				board[rowI][colI].Fixed = true
			}
		}
	}

	solution, _ := Solve(board)

	return board, solution
}

func createSolvedBoard() Board {
	board := Board{}

	//fill randomly with numbers 1-9
	for i := 1; i <= 9; i++ {
		// find random empty cell
		row := rand.Intn(9)
		col := rand.Intn(9)
		for board[row][col].Value != 0 {
			row = rand.Intn(9)
			col = rand.Intn(9)
		}
		board[row][col].Value = i
	}

	solved, _ := Solve(board)

	// scramble the numbers, keeping the same structure
	swapNumbersRandomly(&solved)

	return solved
}

func removeNumbers(board *Board, dificulty int, iterations *int) {
	dificulty = max(0, min(6, dificulty))
	clues := 30 - dificulty

	filledCoords := GetFilledCoords(*board)
	filledLen := len(filledCoords)

	if filledLen <= clues {
		return
	}

	// pick a random coord and try to remove it
	coord := filledCoords[rand.Intn(len(filledCoords))]
	if !removeCoordIfPosible(board, coord) {
		*iterations++
		if *iterations > 20 {
			return
		}
		removeNumbers(board, dificulty, iterations)
	} else {
		filledLen--
	}

	if filledLen <= clues {
		return
	}

	diagonalCount := 2
	if filledLen <= 60 {
		diagonalCount = 1
	}

	coordDiagonals := getDiagonals(coord, diagonalCount)

	for _, coord := range coordDiagonals {
		if removeCoordIfPosible(board, coord) {
			filledLen--
			if filledLen <= clues {
				return
			}
		}
	}

	removeNumbers(board, dificulty, iterations)
}

func removeCoordIfPosible(board *Board, coord Coord) bool {
	valBackup := board[coord.row][coord.col].Value
	(*board)[coord.row][coord.col].Value = 0

	_, err := Solve(*board)

	if err != nil {
		(*board)[coord.row][coord.col].Value = valBackup
		return false
	}

	return true
}

func getDiagonals(coord Coord, distance int) []Coord {
	directions := [][]int{
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	randomDirection := directions[rand.Intn(len(directions))]

	diagonals := []Coord{}
	for i := 1; i <= distance; i++ {
		row := coord.row + randomDirection[0]*i
		col := coord.col + randomDirection[1]*i
		if row < 9 && row >= 0 && col < 9 && col >= 0 {
			diagonals = append(diagonals, Coord{row: row, col: col})
		}
		row1 := coord.row + randomDirection[0]*-i
		col1 := coord.col + randomDirection[1]*-i
		if row1 < 9 && row1 >= 0 && col1 < 9 && col1 >= 0 {
			diagonals = append(diagonals, Coord{row: row1, col: col1})
		}
	}

	return diagonals
}
