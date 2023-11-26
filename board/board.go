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

func InitDefaultBoard() Board {
	chessboard := MakeBoard()

	// pawn
	whitePawn := Figure{FigureType: Pawn, FigureSide: White}
	blackPawn := Figure{FigureType: Pawn, FigureSide: Black}
	for col := 0; col < SIZE; col++ {
		chessboard.SetField(Field{Figure: whitePawn, Cords: Cords{Col: col, Row: 1}, Filled: true})
		chessboard.SetField(Field{Figure: blackPawn, Cords: Cords{Col: col, Row: 6}, Filled: true})
	}

	// rook
	whiteRook := Figure{FigureType: Rook, FigureSide: White}
	chessboard.SetField(Field{Figure: whiteRook, Cords: Cords{Col: 0, Row: 0}, Filled: true})
	chessboard.SetField(Field{Figure: whiteRook, Cords: Cords{Col: 7, Row: 0}, Filled: true})
	blackRook := Figure{FigureType: Rook, FigureSide: Black}
	chessboard.SetField(Field{Figure: blackRook, Cords: Cords{Col: 0, Row: 7}, Filled: true})
	chessboard.SetField(Field{Figure: blackRook, Cords: Cords{Col: 7, Row: 7}, Filled: true})

	// knight
	whiteKnight := Figure{FigureType: Knight, FigureSide: White}
	chessboard.SetField(Field{Figure: whiteKnight, Cords: Cords{Col: 1, Row: 0}, Filled: true})
	chessboard.SetField(Field{Figure: whiteKnight, Cords: Cords{Col: 6, Row: 0}, Filled: true})
	blackKnight := Figure{FigureType: Knight, FigureSide: Black}
	chessboard.SetField(Field{Figure: blackKnight, Cords: Cords{Col: 1, Row: 7}, Filled: true})
	chessboard.SetField(Field{Figure: blackKnight, Cords: Cords{Col: 6, Row: 7}, Filled: true})

	// bishop
	whiteBishop := Figure{FigureType: Bishop, FigureSide: White}
	chessboard.SetField(Field{Figure: whiteBishop, Cords: Cords{Col: 2, Row: 0}, Filled: true})
	chessboard.SetField(Field{Figure: whiteBishop, Cords: Cords{Col: 5, Row: 0}, Filled: true})
	blackBishop := Figure{FigureType: Bishop, FigureSide: Black}
	chessboard.SetField(Field{Figure: blackBishop, Cords: Cords{Col: 2, Row: 7}, Filled: true})
	chessboard.SetField(Field{Figure: blackBishop, Cords: Cords{Col: 5, Row: 7}, Filled: true})

	// queen
	whiteQueen := Figure{FigureType: Queen, FigureSide: White}
	chessboard.SetField(Field{Figure: whiteQueen, Cords: Cords{Col: 3, Row: 0}, Filled: true})
	blackQueen := Figure{FigureType: Queen, FigureSide: Black}
	chessboard.SetField(Field{Figure: blackQueen, Cords: Cords{Col: 3, Row: 7}, Filled: true})

	// king
	whiteKing := Figure{FigureType: King, FigureSide: White}
	chessboard.SetField(Field{Figure: whiteKing, Cords: Cords{Col: 4, Row: 0}, Filled: true})
	blackKing := Figure{FigureType: King, FigureSide: Black}
	chessboard.SetField(Field{Figure: blackKing, Cords: Cords{Col: 4, Row: 7}, Filled: true})

	return chessboard
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
	return board.board[cords.Row][cords.Col]
}

// SetField puts given Field to given Cords
func (board Board) SetField(field Field) {
	board.board[field.Cords.Row][field.Cords.Col] = field
}
