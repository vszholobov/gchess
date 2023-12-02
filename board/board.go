package board

import (
	mapset "github.com/deckarep/golang-set/v2"
	"math"
)

const ChessboardSize = 8

var lineFiguresToSearch = mapset.NewSet(Queen, Rook)
var diagonalFiguresToSearch = mapset.NewSet(Queen, Bishop)

type Board struct {
	board          [][]Field
	moveValidators map[FigureType][]MoveValidator
}

// Copy returns deep Copy of given Board
func (board *Board) Copy() Board {
	duplicate := make([][]Field, ChessboardSize)
	for i := range board.board {
		duplicate[i] = make([]Field, ChessboardSize)
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

	newDeparture := Field{Cords: departure.Cords, Filled: false}
	newDestination := Field{Figure: departure.Figure, Cords: destination.Cords, Filled: true}
	actualBoard := board.Copy()
	actualBoard.SetField(newDeparture)
	actualBoard.SetField(newDestination)

	// TODO: учесть повышение фигуры
	if castleMove, isCastleMove := move.(CastleMove); isCastleMove {
		rook := board.GetField(castleMove.RookDepartureCords()).Figure
		rook.Moved = true
		newRookDeparture := Field{Cords: castleMove.RookDepartureCords(), Filled: false}
		newRookDestination := Field{Figure: rook, Cords: castleMove.RookDestinationCords(), Filled: true}
		actualBoard.SetField(newRookDeparture)
		actualBoard.SetField(newRookDestination)
	}

	return true, &actualBoard
}

// IsFieldAttackedByOpposedSide checks whether field at given cords is attacked by any figure of opposed side
func (board *Board) IsFieldAttackedByOpposedSide(cords Cords, side FigureSide) bool {
	isAttacked := false
	for colDelta := -1; colDelta <= 1; colDelta++ {
		for rowDelta := -1; rowDelta <= 1; rowDelta++ {
			if colDelta == 0 && rowDelta == 0 {
				continue
			}
			var figuresToSearch mapset.Set[FigureType]
			if math.Abs(float64(colDelta)) == math.Abs(float64(rowDelta)) {
				figuresToSearch = diagonalFiguresToSearch
			} else {
				figuresToSearch = lineFiguresToSearch
			}
			isAttacked = isAttacked || isLineAttacked(board, cords, colDelta, rowDelta, side, figuresToSearch)
		}
	}
	isAttacked = isAttacked || isAttackedByPawn(board, cords, side)
	isAttacked = isAttacked || isAttackedByKing(board, cords, side)
	return isAttacked
}

func isAttackedByKing(board *Board, cords Cords, side FigureSide) bool {
	for row := cords.Row - 1; row <= cords.Row+1; row++ {
		for col := cords.Col - 1; col <= cords.Col+1; col++ {
			curCords := Cords{Col: col, Row: row}
			if curCords == cords || ChessboardSize <= row || row < 0 || ChessboardSize <= col || col < 0 {
				continue
			}
			field := board.GetField(curCords)
			figure := field.Figure
			if field.Filled && figure.FigureSide != side && figure.FigureType == King {
				return true
			}
		}
	}
	return false
}

func isLineAttacked(
	board *Board,
	cords Cords,
	colDelta int,
	rowDelta int,
	side FigureSide,
	figuresToSearch mapset.Set[FigureType],
) bool {
	col, row := cords.Col+colDelta, cords.Row+rowDelta
	for ; 0 <= row && row < ChessboardSize && 0 <= col && col < ChessboardSize; col, row = col+colDelta, row+rowDelta {
		field := board.GetField(Cords{col, row})
		if !field.Filled {
			continue
		} else if side == field.Figure.FigureSide {
			return false
		} else if figuresToSearch.Contains(field.Figure.FigureType) {
			return true
		}
	}
	return false
}

func isAttackedByPawn(
	board *Board,
	cords Cords,
	side FigureSide,
) bool {
	var rowDelta int
	if side == White {
		rowDelta = 1
	} else {
		rowDelta = -1
	}

	if row := cords.Row + rowDelta; 0 <= row && row < ChessboardSize {
		isAttackedByPawn := false
		if col := cords.Col - 1; col >= 0 {
			isAttackedByPawn = isAttackedByPawn || isFieldAttackedByPawn(board, Cords{Col: col, Row: row}, side)
		}
		if col := cords.Col + 1; col < ChessboardSize {
			isAttackedByPawn = isAttackedByPawn || isFieldAttackedByPawn(board, Cords{Col: col, Row: row}, side)
		}
		return isAttackedByPawn
	} else {
		return false
	}
}

func isFieldAttackedByPawn(board *Board, cords Cords, side FigureSide) bool {
	field := board.GetField(cords)
	figure := field.Figure
	return field.Filled && figure.FigureSide != side && figure.FigureType == Pawn
}

// MakeBoard returns initialized board
func MakeBoard() Board {
	board := Board{make([][]Field, ChessboardSize), initValidators()}
	for row := range board.board {
		board.board[row] = make([]Field, ChessboardSize)
		for col := 0; col < ChessboardSize; col++ {
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
	for col := 0; col < ChessboardSize; col++ {
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
		for col := 0; col < ChessboardSize; col++ {
			chessboard.SetField(Field{Cords: Cords{Col: col, Row: row}})
		}
	}

	return &chessboard
}

func GetDefaultRowBySide(side FigureSide) int {
	if side == White {
		return 0
	} else {
		return 7
	}
}
