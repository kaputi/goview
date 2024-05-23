package sudoku

const (
	SOLUTION = "solution"
	PLACED   = "placed"
	NOTES1   = "notes1"
	NOTES2   = "notes2"
)

type Coord struct {
	row int
	col int
}

type SudokuCell struct {
	value           int
	fixed           bool
	isPlacedError   bool
	isSolutionError bool
}

type Board [9][9]SudokuCell

type Sudoku struct {
	solution *Board
	placed   *Board
	notes1   *Board
	notes2   *Board
}

func NewSudoku(dificulty int) *Sudoku {
	board, solution := GenerateBoard(dificulty)
	return &Sudoku{
		solution: &solution,
		placed:   &board,
	}
}

func (s *Sudoku) GetPlacedCell(row, col int) SudokuCell {
	return s.getCell(row, col, PLACED)
}

func (s *Sudoku) GetSolutionCell(row, col int) SudokuCell {
	return s.getCell(row, col, SOLUTION)
}

func (s *Sudoku) GetNotes1Cell(row, col int) SudokuCell {
	return s.getCell(row, col, NOTES1)
}

func (s *Sudoku) GetNotes2Cell(row, col int) SudokuCell {
	return s.getCell(row, col, NOTES2)
}

func (s *Sudoku) SetPlacedCell(row, col, value int) {
	s.setCell(row, col, PLACED, value)
}

func (s *Sudoku) SetNotes1Cell(row, col, value int) {
	s.setCell(row, col, NOTES1, value)
}

func (s *Sudoku) SetNotes2Cell(row, col, value int) {
	s.setCell(row, col, NOTES2, value)
}

func (s *Sudoku) getCell(row, col int, layer string) SudokuCell {
	// solution layer has preference, if layer == solution
	if layer == SOLUTION {
		return s.solution[row][col]
	}
	// all other layers give preference to the placed layer
	if s.placed[row][col].value != 0 {
		return s.placed[row][col]
	}

	board := s.getBoard(layer)
	return board[row][col]
}

func (s *Sudoku) setCell(row, col int, layer string, value int) {
	board := s.getBoard(layer)
	if board[row][col].fixed {
		return
	}
	board[row][col].value = value

	if layer != PLACED {
		return
	}

	if s.solution[row][col].value != value {
		board[row][col].isSolutionError = true
	} else {
		board[row][col].isSolutionError = false
	}

	s.updatePlacedErrors()
}

func (s *Sudoku) getBoard(layer string) *Board {
	var board *Board
	switch layer {
	case SOLUTION:
		board = s.solution
	case NOTES1:
		board = s.notes1
	case NOTES2:
		board = s.notes2
	default:
		board = s.placed
	}
	return board
}

func (s *Sudoku) updatePlacedErrors() {
	for rowI, row := range s.placed {
		for colI, cell := range row {
			s.placed[rowI][colI].isPlacedError = !isValidCell(cell.value, rowI, colI, *s.placed)
		}
	}
}
