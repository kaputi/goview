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
	solution Board
	placed   Board
	notes1   Board
	notes2   Board
}

func NewSudoku() *Sudoku {
	// TODO: dificulty
	return &Sudoku{}
}

func (s *Sudoku) getBoard(layer string) Board {
	var board Board
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

func (s *Sudoku) GetCell(row, col int, layer string) SudokuCell {
	if s.placed[row][col].value != 0 {
		return s.placed[row][col]
	}

	board := s.getBoard(layer)
	return board[row][col]
}

func (s *Sudoku) SetCell(row, col int, layer string, value int) {
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

func (s *Sudoku) updatePlacedErrors() {
	for rowI, row := range s.placed {
		for colI, cell := range row {
			cell.isPlacedError = !isValidCell(cell.value, rowI, colI, s.placed)
		}
	}
}
