package board

import (
	mapset "github.com/deckarep/golang-set/v2"
	"math"
)

type MoveValidator interface {
	Validate(move Move) bool
}

func InitValidators(actualBoard *Board) map[FigureType][]MoveValidator {
	validators := make(map[FigureType][]MoveValidator, 6)
	bordersBreachValidator := BordersBreachValidator{}
	departureEqualsDestinationValidator := DepartureEqualsDestinationValidator{}
	notAllyChessmanValidator := NotAllyChessmanValidator{}
	kingIsNotAttackedAfterMoveValidator := KingIsNotAttackedAfterMoveValidator{actualBoard}
	linePathValidator := LinePathValidator{actualBoard}
	diagonalPathValidator := DiagonalPathValidator{actualBoard}
	validators[King] = []MoveValidator{
		bordersBreachValidator,
		departureEqualsDestinationValidator,
		notAllyChessmanValidator,
		kingIsNotAttackedAfterMoveValidator,
		KingMoveValidator{},
		CastlingMoveValidator{actualBoard},
	}
	validators[Pawn] = []MoveValidator{
		bordersBreachValidator,
		departureEqualsDestinationValidator,
		notAllyChessmanValidator,
		kingIsNotAttackedAfterMoveValidator,
		linePathValidator,
		PawnMoveValidator{actualBoard},
		PromotionMoveValidator{},
	}
	validators[Rook] = []MoveValidator{
		bordersBreachValidator,
		departureEqualsDestinationValidator,
		notAllyChessmanValidator,
		kingIsNotAttackedAfterMoveValidator,
		linePathValidator,
		RookMoveValidator{},
	}
	validators[Knight] = []MoveValidator{
		bordersBreachValidator,
		departureEqualsDestinationValidator,
		notAllyChessmanValidator,
		kingIsNotAttackedAfterMoveValidator,
		KnightMoveValidator{},
	}
	validators[Bishop] = []MoveValidator{
		bordersBreachValidator,
		departureEqualsDestinationValidator,
		notAllyChessmanValidator,
		kingIsNotAttackedAfterMoveValidator,
		diagonalPathValidator,
		BishopMoveValidator{},
	}
	validators[Queen] = []MoveValidator{
		bordersBreachValidator,
		departureEqualsDestinationValidator,
		notAllyChessmanValidator,
		kingIsNotAttackedAfterMoveValidator,
		linePathValidator,
		diagonalPathValidator,
		QueenMoveValidator{},
	}
	return validators
}

type BordersBreachValidator struct{}

func (BordersBreachValidator) Validate(move Move) bool {
	destinationCords := move.Destination().Cords
	valid := destinationCords.Col >= 0 && destinationCords.Col < ChessboardSize &&
		destinationCords.Row >= 0 && destinationCords.Row < ChessboardSize
	return valid
}

type DepartureEqualsDestinationValidator struct{}

func (DepartureEqualsDestinationValidator) Validate(move Move) bool {
	return !move.Departure().Cords.Equal(move.Destination().Cords)
}

type NotAllyChessmanValidator struct{}

func (NotAllyChessmanValidator) Validate(move Move) bool {
	if !move.Destination().Filled {
		return true
	}
	return move.Departure().Figure.FigureSide != move.Destination().Figure.FigureSide
}

type LinePathValidator struct {
	ActualBoard *Board
}

func (moveValidator LinePathValidator) Validate(move Move) bool {
	startCol := move.Departure().Cords.Col
	startRow := move.Departure().Cords.Row

	destCol := move.Destination().Cords.Col
	destRow := move.Destination().Cords.Row

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
			if moveValidator.ActualBoard.GetField(Cords{Col: col, Row: startRow}).Filled {
				return false
			}
		}
	} else {
		for row := startRow + 1; row < destRow; row++ {
			if moveValidator.ActualBoard.GetField(Cords{Col: startCol, Row: row}).Filled {
				return false
			}
		}
	}
	return true
}

type DiagonalPathValidator struct {
	ActualBoard *Board
}

func (moveValidator DiagonalPathValidator) Validate(move Move) bool {
	startCol := move.Departure().Cords.Col
	startRow := move.Departure().Cords.Row

	destCol := move.Destination().Cords.Col
	destRow := move.Destination().Cords.Row

	if math.Abs(float64(destRow-startRow)) != math.Abs(float64(destCol-startCol)) {
		return true
	}

	if (destCol > startCol && destRow < startRow) || (destCol < startCol && destRow > startRow) {
		for col := startCol; col < destCol; col++ {
			if moveValidator.ActualBoard.GetField(Cords{Col: col, Row: ChessboardSize - col - 1}).Filled {
				return false
			}
		}
	} else {
		for i := 0; i < destCol; i++ {
			if moveValidator.ActualBoard.GetField(Cords{Col: i, Row: i}).Filled {
				return false
			}
		}
	}
	return true
}

type KnightMoveValidator struct{}

func (KnightMoveValidator) Validate(move Move) bool {
	startCol := move.Departure().Cords.Col
	startRow := move.Departure().Cords.Row

	destCol := move.Destination().Cords.Col
	destRow := move.Destination().Cords.Row

	isValid := (math.Abs(float64(destCol-startCol)) == 2 && math.Abs(float64(destRow-startRow)) == 1) ||
		(math.Abs(float64(destCol-startCol)) == 1 && math.Abs(float64(destRow-startRow)) == 2)
	return isValid
}

type QueenMoveValidator struct{}

func (QueenMoveValidator) Validate(move Move) bool {
	startCol := move.Departure().Cords.Col
	startRow := move.Departure().Cords.Row

	destCol := move.Destination().Cords.Col
	destRow := move.Destination().Cords.Row

	if destRow == startRow || destCol == startCol {
		return true
	}

	if math.Abs(float64(destRow-startRow)) == math.Abs(float64(destCol-startCol)) {
		return true
	}

	return false
}

type RookMoveValidator struct{}

func (RookMoveValidator) Validate(move Move) bool {
	startCol := move.Departure().Cords.Col
	startRow := move.Departure().Cords.Row

	destCol := move.Destination().Cords.Col
	destRow := move.Destination().Cords.Row

	isValid := destRow == startRow || destCol == startCol
	return isValid
}

type BishopMoveValidator struct{}

func (BishopMoveValidator) Validate(move Move) bool {
	startCol := move.Departure().Cords.Col
	startRow := move.Departure().Cords.Row

	destCol := move.Destination().Cords.Col
	destRow := move.Destination().Cords.Row

	isValid := math.Abs(float64(destRow-startRow)) == math.Abs(float64(destCol-startCol))
	return isValid
}

type PawnMoveValidator struct {
	ActualBoard *Board
}

func (moveValidator PawnMoveValidator) Validate(move Move) bool {
	startCol := move.Departure().Cords.Col
	destCol := move.Destination().Cords.Col
	movingPawn := move.Departure().Figure

	var rowDistance int
	if movingPawn.FigureSide == White {
		rowDistance = move.Destination().Cords.Row - move.Departure().Cords.Row
	} else {
		rowDistance = move.Departure().Cords.Row - move.Destination().Cords.Row
	}

	if startCol == destCol {
		return rowDistance == 1 || !movingPawn.Moved && rowDistance == 2
	} else {
		colDistance := math.Abs(float64(startCol - destCol))
		if colDistance != 1 || rowDistance != 1 {
			return false
		}
		if move.Destination().Filled {
			return movingPawn.FigureSide != move.Destination().Figure.FigureSide
		} else {
			lastMove := moveValidator.ActualBoard.GetLastMove()
			distance := math.Abs(float64(lastMove.Destination().Cords.Row - lastMove.Departure().Cords.Row))
			return lastMove.Departure().Figure.FigureType == Pawn && distance == 2 &&
				lastMove.Destination().Cords.Col == destCol && lastMove.Destination().Cords.Row == move.Departure().Cords.Row
		}
	}
}

var promotionAllowedTypes = mapset.NewSet(Queen, Rook, Bishop, Knight)

type PromotionMoveValidator struct{}

func (PromotionMoveValidator) Validate(move Move) bool {
	if promotionMove, isPromotionMove := move.(PromotionMove); isPromotionMove {
		return promotionAllowedTypes.Contains(promotionMove.promoteToType)
	} else {
		return true
	}
}

type KingMoveValidator struct{}

func (KingMoveValidator) Validate(move Move) bool {
	colDiff := math.Abs(float64(move.Departure().Cords.Col - move.Destination().Cords.Col))
	rowDiff := math.Abs(float64(move.Departure().Cords.Row - move.Destination().Cords.Row))
	return colDiff <= 1 && rowDiff <= 1 || colDiff == 2 && rowDiff == 0
}

type CastlingMoveValidator struct {
	ActualBoard *Board
}

func (moveValidator CastlingMoveValidator) Validate(move Move) bool {
	king := move.Departure().Figure
	board := moveValidator.ActualBoard
	if king.Moved {
		return false
	}
	row := GetDefaultRowBySide(king.FigureSide)
	var rookCol int
	if longCastleCords := (Cords{Col: 2, Row: row}); longCastleCords == move.Destination().Cords {
		rookCol = 0
	} else if shortCastleCords := (Cords{Col: 6, Row: row}); shortCastleCords == move.Destination().Cords {
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
	kingCol := move.Departure().Cords.Col
	for col := min(kingCol, rookCol) + 1; col < max(kingCol, rookCol); col++ {
		if board.GetField(Cords{Col: col, Row: row}).Filled {
			return false
		}
	}

	// if any field between the king and the destination are attacked then can't castle
	for col := min(move.Destination().Cords.Col, kingCol); col <= max(move.Destination().Cords.Col, kingCol); col++ {
		if board.IsFieldAttackedByOpposedSide(Cords{Col: col, Row: row}, king.FigureSide) {
			return false
		}
	}

	return true
}

type KingIsNotAttackedAfterMoveValidator struct {
	ActualBoard *Board
}

func (moveValidator KingIsNotAttackedAfterMoveValidator) Validate(move Move) bool {
	departure := move.Departure()
	movingFigure := departure.Figure
	if moveValidator.ActualBoard.GetKingCords(movingFigure.FigureSide) == nil {
		return true
	}
	validationBoard := moveValidator.ActualBoard.Copy()
	departure.Figure = Figure{}
	departure.Filled = false
	destination := move.Destination()
	destination.Figure = movingFigure
	destination.Filled = true

	validationBoard.SetField(departure)
	validationBoard.SetField(destination)

	movingFigureSide := movingFigure.FigureSide
	kingCords := validationBoard.GetKingCords(movingFigureSide)

	return !validationBoard.IsFieldAttackedByOpposedSide(*kingCords, movingFigureSide)
}
