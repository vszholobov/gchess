package board

const SIZE = 8

type Board struct {
	board [][]Field
}

// MakeBoard returns initialized isFilled board
func MakeBoard() Board {
	board := Board{make([][]Field, SIZE)}
	for i := range board.board {
		board.board[i] = make([]Field, SIZE)
	}
	return board
}

// Copy returns deep Copy of given Board
func (board Board) Copy() Board {
	duplicate := make([][]Field, SIZE)
	for i := range board.board {
		duplicate[i] = make([]Field, SIZE)
		copy(duplicate[i], board.board[i])
	}
	return Board{board: duplicate}
}

// GetField returns Field at given Cords
func (board Board) GetField(cords Cords) Field {
	return board.board[cords.Col][cords.Row]
}

// SetField puts given Field to given Cords
func (board Board) SetField(field Field) {
	board.board[field.Cords.Row][field.Cords.Col] = field
}
