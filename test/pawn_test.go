package test

import (
	"chess/board"
	"chess/session"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPawnMoveForwardPassEnemyPawn(t *testing.T) {
	chessBoard := board.MakeBoard()
	whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
	whitePawnCords := board.Cords{Col: 0, Row: 5}
	whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
	chessBoard.SetField(whitePawnField)

	blackPawn := board.Figure{FigureType: board.Pawn, FigureSide: board.Black, Moved: false}
	blackPawnCords := board.Cords{Col: 0, Row: 6}
	blackPawnField := board.Field{Figure: blackPawn, Cords: blackPawnCords, Filled: true}
	chessBoard.SetField(blackPawnField)

	destinationField := chessBoard.GetField(board.Cords{Col: 0, Row: 7})

	move := board.MakeMove(whitePawnField, destinationField, board.EmptyType)
	validator := board.PawnMoveValidator{ActualBoard: &chessBoard}

	assert.False(t, validator.Validate(move))
}

func TestPawnMoveForwardToEnemyPawn(t *testing.T) {
	chessBoard := board.MakeBoard()
	whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
	whitePawnCords := board.Cords{Col: 0, Row: 5}
	whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
	chessBoard.SetField(whitePawnField)

	blackPawn := board.Figure{FigureType: board.Pawn, FigureSide: board.Black, Moved: false}
	blackPawnCords := board.Cords{Col: 0, Row: 6}
	blackPawnField := board.Field{Figure: blackPawn, Cords: blackPawnCords, Filled: true}
	chessBoard.SetField(blackPawnField)

	move := board.MakeMove(whitePawnField, blackPawnField, board.EmptyType)
	validator := board.PawnMoveValidator{ActualBoard: &chessBoard}

	assert.False(t, validator.Validate(move))
}

func TestPawnEnPassant_Fail_KillDestinationTooFarRow(t *testing.T) {
	chessBoard := board.MakeBoard()

	whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
	whitePawnCords := board.Cords{Col: 0, Row: 4}
	whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
	chessBoard.SetField(whitePawnField)
	killDestination := chessBoard.GetField(board.Cords{Col: 1, Row: 6})
	move := board.MakeMove(whitePawnField, killDestination, board.EmptyType)
	validator := board.PawnMoveValidator{ActualBoard: &chessBoard}

	assert.False(t, validator.Validate(move))
}

func TestPawnEnPassant_Fail_KillDestinationTooFarCol(t *testing.T) {
	chessBoard := board.MakeBoard()

	whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
	whitePawnCords := board.Cords{Col: 0, Row: 4}
	whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
	chessBoard.SetField(whitePawnField)
	killDestination := chessBoard.GetField(board.Cords{Col: 2, Row: 5})
	move := board.MakeMove(whitePawnField, killDestination, board.EmptyType)
	validator := board.PawnMoveValidator{ActualBoard: &chessBoard}

	assert.False(t, validator.Validate(move))
}

func TestPawnEnPassant_Fail_ClosePawnMovedShort(t *testing.T) {
	chessBoard := board.MakeBoard()

	whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
	whitePawnCords := board.Cords{Col: 0, Row: 4}
	whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
	chessBoard.SetField(whitePawnField)
	blackPawn := board.Figure{FigureType: board.Pawn, FigureSide: board.Black, Moved: true}
	blackPawnCords := board.Cords{Col: 1, Row: 5}
	blackPawnField := board.Field{Figure: blackPawn, Cords: blackPawnCords, Filled: true}
	chessBoard.SetField(blackPawnField)
	blackPawnShortMoveDestinationField := chessBoard.GetField(board.Cords{Col: 1, Row: 4})
	chessBoard = chessBoard.Move(board.MakeMove(blackPawnField, blackPawnShortMoveDestinationField, board.EmptyType))
	whitePawnEnPassantMoveDestinationField := chessBoard.GetField(board.Cords{Col: 1, Row: 5})
	move := board.MakeMove(whitePawnField, whitePawnEnPassantMoveDestinationField, board.EmptyType)
	validator := board.PawnMoveValidator{ActualBoard: &chessBoard}

	assert.False(t, validator.Validate(move))
}

func TestPawnEnPassant_Success(t *testing.T) {
	chessBoard := board.MakeBoard()

	whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
	whitePawnCords := board.Cords{Col: 0, Row: 4}
	whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
	chessBoard.SetField(whitePawnField)
	blackPawn := board.Figure{FigureType: board.Pawn, FigureSide: board.Black, Moved: true}
	blackPawnCords := board.Cords{Col: 1, Row: 6}
	blackPawnField := board.Field{Figure: blackPawn, Cords: blackPawnCords, Filled: true}
	chessBoard.SetField(blackPawnField)
	blackPawnLongMoveDestinationField := chessBoard.GetField(board.Cords{Col: 1, Row: 4})
	chessBoard = chessBoard.Move(board.MakeMove(blackPawnField, blackPawnLongMoveDestinationField, board.EmptyType))
	whitePawnEnPassantMoveDestinationField := chessBoard.GetField(board.Cords{Col: 1, Row: 5})
	move := board.MakeMove(whitePawnField, whitePawnEnPassantMoveDestinationField, board.EmptyType)
	validator := board.PawnMoveValidator{ActualBoard: &chessBoard}

	assert.True(t, validator.Validate(move))
}

func TestPawnKillFigureValidation_Success(t *testing.T) {
	for col := 0; col < board.ChessboardSize; col++ {
		if col != 0 {
			assert.True(t, testPawnKill(col, col-1))
		}
		if col != board.ChessboardSize-1 {
			assert.True(t, testPawnKill(col, col+1))
		}
	}
}

func testPawnKill(col int, destCol int) bool {
	chessBoard := board.MakeBoard()

	whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
	whitePawnCords := board.Cords{Col: col, Row: 1}
	whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
	chessBoard.SetField(whitePawnField)
	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black, Moved: true}
	blackRookCords := board.Cords{Col: destCol, Row: 2}
	blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
	chessBoard.SetField(blackRookField)
	move := board.MakeMove(whitePawnField, blackRookField, board.EmptyType)
	validator := board.PawnMoveValidator{ActualBoard: &chessBoard}

	return validator.Validate(move)
}

func TestPromotionMoveAllowedTypes(t *testing.T) {
	allowedFigureTypes := []board.FigureType{board.Queen, board.Bishop, board.Knight, board.Rook}
	for _, figureType := range allowedFigureTypes {
		chessBoard := board.MakeBoard()

		whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
		whitePawnCords := board.Cords{Col: 0, Row: 6}
		whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
		chessBoard.SetField(whitePawnField)
		destinationCords := board.Cords{Col: 0, Row: 7}
		gameSession := session.MakeSession(&chessBoard)
		moveRequest := session.MoveRequest{
			DepartureCords:   whitePawnCords,
			DestinationCords: destinationCords,
			PromoteToType:    figureType,
		}

		isMoved := gameSession.Move(moveRequest)
		actualBoard := gameSession.ActualBoard

		assert.True(t, isMoved)
		assert.Equal(t, figureType, actualBoard.GetField(destinationCords).Figure.FigureType)
	}
}

func TestPromotionMoveNotAllowedTypes(t *testing.T) {
	notAllowedFigureTypes := []board.FigureType{board.King, board.Pawn}
	for _, figureType := range notAllowedFigureTypes {
		chessBoard := board.MakeBoard()

		whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
		whitePawnCords := board.Cords{Col: 0, Row: 6}
		whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
		chessBoard.SetField(whitePawnField)
		destinationCords := board.Cords{Col: 0, Row: 7}

		gameSession := session.MakeSession(&chessBoard)
		moveRequest := session.MoveRequest{
			DepartureCords:   whitePawnCords,
			DestinationCords: destinationCords,
			PromoteToType:    figureType,
		}

		isMoved := gameSession.Move(moveRequest)

		assert.False(t, isMoved)
	}
}

func TestPromotionKillMoveAllowedTypes(t *testing.T) {
	allowedFigureTypes := []board.FigureType{board.Queen, board.Bishop, board.Knight, board.Rook}
	for _, figureType := range allowedFigureTypes {
		chessBoard := board.MakeBoard()

		whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
		whitePawnCords := board.Cords{Col: 0, Row: 6}
		whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
		chessBoard.SetField(whitePawnField)

		blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black, Moved: true}
		blackRookCords := board.Cords{Col: 1, Row: 7}
		blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
		chessBoard.SetField(blackRookField)

		gameSession := session.MakeSession(&chessBoard)
		moveRequest := session.MoveRequest{
			DepartureCords:   whitePawnCords,
			DestinationCords: blackRookCords,
			PromoteToType:    figureType,
		}

		isMoved := gameSession.Move(moveRequest)
		actualBoard := gameSession.ActualBoard

		assert.True(t, isMoved)
		assert.Equal(t, figureType, actualBoard.GetField(blackRookCords).Figure.FigureType)
	}
}

func TestPromotionKillMoveNotAllowedTypes(t *testing.T) {
	allowedFigureTypes := []board.FigureType{board.King, board.Pawn}
	for _, figureType := range allowedFigureTypes {
		chessBoard := board.MakeBoard()

		whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
		whitePawnCords := board.Cords{Col: 0, Row: 6}
		whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
		chessBoard.SetField(whitePawnField)

		blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black, Moved: true}
		blackRookCords := board.Cords{Col: 1, Row: 7}
		blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
		chessBoard.SetField(blackRookField)

		gameSession := session.MakeSession(&chessBoard)
		moveRequest := session.MoveRequest{
			DepartureCords:   whitePawnCords,
			DestinationCords: blackRookCords,
			PromoteToType:    figureType,
		}

		isMoved := gameSession.Move(moveRequest)
		assert.False(t, isMoved)
	}
}
