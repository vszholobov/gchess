package board

import "fmt"

const SIZE = 8

type Board struct {
	board          [][]Field
	moveValidators map[FigureType][]MoveValidator
}

// Copy returns deep Copy of given Board
func (board *Board) Copy() Board {
	duplicate := make([][]Field, SIZE)
	for i := range board.board {
		duplicate[i] = make([]Field, SIZE)
		copy(duplicate[i], board.board[i])
	}
	return Board{board: duplicate, moveValidators: board.moveValidators}
}

// GetField returns Field at given Cords
func (board *Board) GetField(cords Cords) Field {
	return board.board[cords.Row][cords.Col]
}

// SetField puts given Field to given Cords
func (board *Board) SetField(field Field) {
	board.board[field.Cords.Row][field.Cords.Col] = field
}

func (board *Board) Move(
	departureCords Cords,
	destinationCords Cords,
	moveSide FigureSide,
) (bool, *Board) {
	departure := board.GetField(departureCords)
	if departure.Figure.FigureSide != moveSide {
		return false, nil
	}
	destination := board.GetField(destinationCords)
	move := MakeMove(departure, destination)

	for _, validator := range board.moveValidators[departure.Figure.FigureType] {
		validMove := validator.Validate(board, move)
		if !validMove {
			return false, nil
		}
	}

	departure.Figure.Moved = true

	// Перемещение. TODO: учесть рокировку и повышение фигуры
	newDeparture := Field{Cords: departure.Cords, Filled: false}
	newDestination := Field{Figure: departure.Figure, Cords: destination.Cords, Filled: true}

	actualBoard := board.Copy()
	actualBoard.SetField(newDeparture)
	actualBoard.SetField(newDestination)

	return true, &actualBoard
}

// isFieldAttackedByOpposedSide checks whether field at given cords is attacked by any figure of opposed side
func (board *Board) isFieldAttackedByOpposedSide(cords Cords, side FigureSide) bool {
	// TODO: implement
	var opposedSide FigureSide
	if side == White {
		opposedSide = Black
	} else {
		opposedSide = White
	}
	fmt.Println(opposedSide)
	return false
}

// MakeBoard returns initialized board
func MakeBoard() Board {
	board := Board{make([][]Field, SIZE), initValidators()}
	for row := range board.board {
		board.board[row] = make([]Field, SIZE)
		for col := 0; col < SIZE; col++ {
			board.board[row][col] = Field{Cords: Cords{Col: col, Row: row}, Filled: false}
		}
	}
	return board
}

func InitDefaultBoard() *Board {
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

	// empty
	for row := 2; row < 6; row++ {
		for col := 0; col < SIZE; col++ {
			chessboard.SetField(Field{Cords: Cords{Col: col, Row: row}})
		}
	}

	return &chessboard
}
