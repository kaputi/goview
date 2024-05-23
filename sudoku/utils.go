package sudoku

import (
	"math/rand"
)

func swapNumbersRandomly(board *Board) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	scrambleSlice(&numbers)

	for rowIdx, row := range *board {
		for colIdx, cell := range row {
			(*board)[rowIdx][colIdx].value = numbers[cell.value-1]
		}
	}
}

func scrambleSlice[T any](slice *[]T) {
	for i := range *slice {
		j := rand.Intn(i + 1)
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
}

func getNextEmpty(board Board) (int, int) {
	for rowIdx, row := range board {
		for colIdx, cell := range row {
			if cell.value == 0 {
				return rowIdx, colIdx
			}
		}
	}

	return -1, -1
}

func boardIsFull(board Board) bool {
	for _, rowIdx := range board {
		for _, cell := range rowIdx {
			if cell.value == 0 {
				return false
			}
		}
	}

	return true
}

func getCoords(board Board, filled bool) []Coord {
	coords := []Coord{}
	for rowI, row := range board {
		for colI, cell := range row {
			coord := Coord{row: rowI, col: colI}
			if filled && cell.value != 0 {
				coords = append(coords, coord)
			}
			if !filled && cell.value == 0 {
				coords = append(coords, coord)
			}
		}
	}

	return coords
}

func GetEmptyCoords(board Board) []Coord {
	return getCoords(board, false)
}

func GetFilledCoords(board Board) []Coord {
	return getCoords(board, true)
}
