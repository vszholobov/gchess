package board

import "math"

type MoveValidator interface {
	Validate(chessboard *Board, move Move) bool
}

func initValidators() map[FigureType][]MoveValidator {
	validators := make(map[FigureType][]MoveValidator, 6)
	validators[King] = []MoveValidator{
		BordersBreachValidator{},
		DepartureEqualsDestinationValidator{},
		NotAllyChessmanValidator{},
		KingMoveValidator{},
		CastlingMoveValidator{},
	}
	validators[Pawn] = []MoveValidator{
		BordersBreachValidator{},
		DepartureEqualsDestinationValidator{},
		NotAllyChessmanValidator{},
		LinePathValidator{},
		PawnMoveValidator{},
	}
	validators[Rook] = []MoveValidator{
		BordersBreachValidator{},
		DepartureEqualsDestinationValidator{},
		NotAllyChessmanValidator{},
		LinePathValidator{},
		RookMoveValidator{},
	}
	validators[Knight] = []MoveValidator{
		BordersBreachValidator{},
		DepartureEqualsDestinationValidator{},
		NotAllyChessmanValidator{},
		KingMoveValidator{},
	}
	validators[Bishop] = []MoveValidator{
		BordersBreachValidator{},
		DepartureEqualsDestinationValidator{},
		NotAllyChessmanValidator{},
		DiagonalPathValidator{},
		BishopMoveValidator{},
	}
	validators[Queen] = []MoveValidator{
		BordersBreachValidator{},
		DepartureEqualsDestinationValidator{},
		NotAllyChessmanValidator{},
		LinePathValidator{},
		DiagonalPathValidator{},
		QueenMoveValidator{},
	}
	return validators
}

type BordersBreachValidator struct{}

func (BordersBreachValidator) Validate(_ *Board, move Move) bool {
	destinationCords := move.Destination.Cords
	valid := destinationCords.Col >= 0 && destinationCords.Col < ChessboardSize &&
		destinationCords.Row >= 0 && destinationCords.Row < ChessboardSize
	return valid
}

type DepartureEqualsDestinationValidator struct{}

func (DepartureEqualsDestinationValidator) Validate(_ *Board, move Move) bool {
	return !move.Departure.Cords.Equal(move.Destination.Cords)
}

type NotAllyChessmanValidator struct{}

func (NotAllyChessmanValidator) Validate(_ *Board, move Move) bool {
	if !move.Destination.Filled {
		return true
	}
	return move.Departure.Figure.FigureSide != move.Destination.Figure.FigureSide
}

type LinePathValidator struct{}

func (LinePathValidator) Validate(chessboard *Board, move Move) bool {
	startCol := move.Departure.Cords.Col
	startRow := move.Departure.Cords.Row

	destCol := move.Destination.Cords.Col
	destRow := move.Destination.Cords.Row

	if destRow != startRow && destCol != startCol {
		return true
	}

	temp := startCol
	startCol = min(startCol, destCol)
	destCol = max(temp, destCol)

	temp = startRow
	startRow = min(startRow, destRow)
	destRow = max(temp, destRow)

	if destRow == startRow {
		for col := startCol + 1; col < destCol; col++ {
			if chessboard.GetField(Cords{Col: col, Row: startRow}).Filled {
				return false
			}
		}
	} else {
		for row := startRow + 1; row < destRow; row++ {
			if chessboard.GetField(Cords{Col: startCol, Row: row}).Filled {
				return false
			}
		}
	}
	return true
}

type DiagonalPathValidator struct{}

func (DiagonalPathValidator) Validate(chessboard *Board, move Move) bool {
	startCol := move.Departure.Cords.Col
	startRow := move.Departure.Cords.Row

	destCol := move.Destination.Cords.Col
	destRow := move.Destination.Cords.Row

	if math.Abs(float64(destRow-startRow)) != math.Abs(float64(destCol-startCol)) {
		return true
	}

	if (destCol > startCol && destRow < startRow) || (destCol < startCol && destRow > startRow) {
		for col := startCol; col < destCol; col++ {
			if chessboard.GetField(Cords{Col: col, Row: ChessboardSize - col - 1}).Filled {
				return false
			}
		}
	} else {
		for i := 0; i < destCol; i++ {
			if chessboard.GetField(Cords{Col: i, Row: i}).Filled {
				return false
			}
		}
	}
	return true
}

type KnightMoveValidator struct{}

func (KnightMoveValidator) Validate(_ *Board, move Move) bool {
	startCol := move.Departure.Cords.Col
	startRow := move.Departure.Cords.Row

	destCol := move.Destination.Cords.Col
	destRow := move.Destination.Cords.Row

	isValid := (math.Abs(float64(destCol-startCol)) == 2 && math.Abs(float64(destRow-startRow)) == 1) ||
		(math.Abs(float64(destCol-startCol)) == 1 && math.Abs(float64(destRow-startRow)) == 2)
	return isValid
}

type QueenMoveValidator struct{}

func (QueenMoveValidator) Validate(_ *Board, move Move) bool {
	startCol := move.Departure.Cords.Col
	startRow := move.Departure.Cords.Row

	destCol := move.Destination.Cords.Col
	destRow := move.Destination.Cords.Row

	if destRow == startRow || destCol == startCol {
		return true
	}

	if math.Abs(float64(destRow-startRow)) == math.Abs(float64(destCol-startCol)) {
		return true
	}

	return false
}

type RookMoveValidator struct{}

func (RookMoveValidator) Validate(_ *Board, move Move) bool {
	startCol := move.Departure.Cords.Col
	startRow := move.Departure.Cords.Row

	destCol := move.Destination.Cords.Col
	destRow := move.Destination.Cords.Row

	isValid := destRow == startRow || destCol == startCol
	return isValid
}

type BishopMoveValidator struct{}

func (BishopMoveValidator) Validate(_ *Board, move Move) bool {
	startCol := move.Departure.Cords.Col
	startRow := move.Departure.Cords.Row

	destCol := move.Destination.Cords.Col
	destRow := move.Destination.Cords.Row

	isValid := math.Abs(float64(destRow-startRow)) == math.Abs(float64(destCol-startCol))
	return isValid
}

type PawnMoveValidator struct{}

func (PawnMoveValidator) Validate(_ *Board, move Move) bool {
	startCol := move.Departure.Cords.Col
	destCol := move.Destination.Cords.Col

	if startCol != destCol {
		// TODO: взятие на проходе (амперсанд)
		// TODO: в валидатор амперсанд можно прокинуть указатель на последний ход. Сделать указатель на переменную "последний ход"
		isValid := move.Destination.Filled &&
			move.Departure.Figure.FigureSide != move.Destination.Figure.FigureSide &&
			math.Abs(float64(startCol-destCol)) == 1.0
		return isValid
	} else {
		var diff int
		if move.Departure.Figure.FigureSide == White {
			diff = move.Destination.Cords.Row - move.Departure.Cords.Row
		} else {
			diff = move.Departure.Cords.Row - move.Destination.Cords.Row
		}
		return diff == 1 || !move.Departure.Figure.Moved && diff == 2
	}
}

type KingMoveValidator struct{}

func (KingMoveValidator) Validate(_ *Board, move Move) bool {
	colDiff := math.Abs(float64(move.Departure.Cords.Col - move.Destination.Cords.Col))
	rowDiff := math.Abs(float64(move.Departure.Cords.Row - move.Destination.Cords.Row))
	return colDiff <= 1 && rowDiff <= 1
}

type CastlingMoveValidator struct{}

func (CastlingMoveValidator) Validate(board *Board, move Move) bool {
	king := move.Departure.Figure
	if king.Moved {
		return false
	}
	var row int
	if king.FigureSide == White {
		row = 0
	} else {
		row = 7
	}
	var rookCol int
	if longCastleCords := (Cords{Col: 2, Row: row}); longCastleCords == move.Destination.Cords {
		rookCol = 0
	} else if shortCastleCords := (Cords{Col: 6, Row: row}); shortCastleCords == move.Destination.Cords {
		rookCol = 7
	} else {
		return false
	}

	// if castle side rook isn't a rook or moved before, then can't castle
	c := Cords{Col: rookCol, Row: row}
	if rook := board.GetField(c).Figure; rook.FigureType != Rook || rook.Moved || rook.FigureSide != king.FigureSide {
		return false
	}

	// if any field between the king and the rook are filled then can't castle
	kingCol := move.Departure.Cords.Col
	for col := min(kingCol, rookCol) + 1; col < max(kingCol, rookCol); col++ {
		if board.GetField(Cords{Col: col, Row: row}).Filled {
			return false
		}
	}

	// if any field between the king and the destination are attacked then can't castle
	for col := min(move.Destination.Cords.Col, kingCol); col <= max(move.Destination.Cords.Col, kingCol); col++ {
		if board.IsFieldAttackedByOpposedSide(Cords{Col: col, Row: row}, king.FigureSide) {
			return false
		}
	}

	return true
}

// TODO: валидатор атаки короля после хода фигуры
// TODO: поле в валидаторы можно прокинуть указателем в сами структуры, чтобы не передавать в метод. Сделать указатель на переменную "актуальное поле"
