package test

import (
	"chess/board"
	"chess/session"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeMove_LongCastleMove(t *testing.T) {
	chessBoard := board.MakeBoard()
	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 4, Row: 0}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)
	castleCords := board.Cords{Col: 2, Row: 0}
	assert.IsType(t, board.CastleMove{}, board.MakeMove(
		whiteKingField,
		board.Field{Cords: castleCords},
		board.EmptyType,
	))
}

func TestMakeMove_CastleMove(t *testing.T) {
	chessBoard := board.MakeBoard()
	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 4, Row: 0}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)
	castleCords := board.Cords{Col: 6, Row: 0}
	assert.IsType(t, board.CastleMove{}, board.MakeMove(
		whiteKingField,
		board.Field{Cords: castleCords},
		board.EmptyType,
	))
}

func TestLongCastleMove(t *testing.T) {
	chessBoard := board.MakeBoard()
	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 4, Row: 0}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)
	whiteRook := board.Figure{FigureType: board.Rook, FigureSide: board.White, Moved: false}
	whiteRookCords := board.Cords{Col: 0, Row: 0}
	whiteRookField := board.Field{Figure: whiteRook, Cords: whiteRookCords, Filled: true}
	chessBoard.SetField(whiteRookField)
	castleCords := board.Cords{Col: 2, Row: 0}
	futureRookCords := board.Cords{Col: 3, Row: 0}

	gameSession := session.MakeSession(&chessBoard)
	moveRequest := session.MoveRequest{
		DepartureCords:   whiteKingCords,
		DestinationCords: castleCords,
		PromoteToType:    board.EmptyType,
	}
	isMoved := gameSession.Move(moveRequest)

	actualBoard := gameSession.ActualBoard

	assert.True(t, isMoved)
	actualCastleField := actualBoard.GetField(castleCords)
	assert.True(t, actualCastleField.Filled)
	assert.Equal(
		t,
		board.Figure{FigureType: board.King, FigureSide: board.White, Moved: true},
		actualCastleField.Figure,
	)
	actualWhiteRookDepartureField := actualBoard.GetField(whiteRookCords)
	assert.False(t, actualWhiteRookDepartureField.Filled)
	assert.Equal(t, board.Figure{}, actualWhiteRookDepartureField.Figure)
	actualWhiteRookField := actualBoard.GetField(futureRookCords)
	assert.True(t, actualWhiteRookField.Filled)
	assert.Equal(
		t,
		board.Figure{FigureType: board.Rook, FigureSide: board.White, Moved: true},
		actualWhiteRookField.Figure,
	)
}

func TestShortCastleMove(t *testing.T) {
	chessBoard := board.MakeBoard()
	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 4, Row: 0}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)
	whiteRook := board.Figure{FigureType: board.Rook, FigureSide: board.White, Moved: false}
	whiteRookCords := board.Cords{Col: 7, Row: 0}
	whiteRookField := board.Field{Figure: whiteRook, Cords: whiteRookCords, Filled: true}
	chessBoard.SetField(whiteRookField)
	castleCords := board.Cords{Col: 6, Row: 0}
	futureRookCords := board.Cords{Col: 5, Row: 0}

	gameSession := session.MakeSession(&chessBoard)
	moveRequest := session.MoveRequest{
		DepartureCords:   whiteKingCords,
		DestinationCords: castleCords,
		PromoteToType:    board.EmptyType,
	}
	isMoved := gameSession.Move(moveRequest)

	actualBoard := gameSession.ActualBoard

	assert.True(t, isMoved)
	actualCastleField := actualBoard.GetField(castleCords)
	assert.True(t, actualCastleField.Filled)
	assert.Equal(
		t,
		board.Figure{FigureType: board.King, FigureSide: board.White, Moved: true},
		actualCastleField.Figure,
	)
	actualWhiteRookDepartureField := actualBoard.GetField(whiteRookCords)
	assert.False(t, actualWhiteRookDepartureField.Filled)
	assert.Equal(t, board.Figure{}, actualWhiteRookDepartureField.Figure)
	actualWhiteRookField := actualBoard.GetField(futureRookCords)
	assert.True(t, actualWhiteRookField.Filled)
	assert.Equal(
		t,
		board.Figure{FigureType: board.Rook, FigureSide: board.White, Moved: true},
		actualWhiteRookField.Figure,
	)
}

func TestWhiteLongCastleValidation(t *testing.T) {
	chessBoard := board.MakeBoard()
	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 4, Row: 0}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)
	whiteRook := board.Figure{FigureType: board.Rook, FigureSide: board.White, Moved: false}
	whiteRookCords := board.Cords{Col: 0, Row: 0}
	whiteRookField := board.Field{Figure: whiteRook, Cords: whiteRookCords, Filled: true}
	chessBoard.SetField(whiteRookField)
	castlingMoveValidator := board.CastlingMoveValidator{ActualBoard: &chessBoard}
	destinationCastleField := chessBoard.GetField(board.Cords{Col: 2, Row: 0})
	castlingMove := board.MakeMove(whiteKingField, destinationCastleField, board.EmptyType)

	isCastled := castlingMoveValidator.Validate(castlingMove)

	assert.True(t, isCastled)
}

func TestBlackLongCastleValidation(t *testing.T) {
	chessBoard := board.MakeBoard()
	blackKing := board.Figure{FigureType: board.King, FigureSide: board.Black, Moved: false}
	blackKingCords := board.Cords{Col: 4, Row: 7}
	blackKingField := board.Field{Figure: blackKing, Cords: blackKingCords, Filled: true}
	chessBoard.SetField(blackKingField)
	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black, Moved: false}
	blackRookCords := board.Cords{Col: 0, Row: 7}
	blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
	chessBoard.SetField(blackRookField)
	castlingMoveValidator := board.CastlingMoveValidator{ActualBoard: &chessBoard}
	destinationCastleField := chessBoard.GetField(board.Cords{Col: 2, Row: 7})
	castlingMove := board.MakeMove(blackKingField, destinationCastleField, board.EmptyType)

	isCastled := castlingMoveValidator.Validate(castlingMove)

	assert.True(t, isCastled)
}

func TestWhiteLongCastleValidation_FailKnightFilledBetween(t *testing.T) {
	chessBoard := board.MakeBoard()

	whiteKnight := board.Figure{FigureType: board.Knight, FigureSide: board.White, Moved: false}
	whiteKnightCords := board.Cords{Col: 1, Row: 0}
	whiteKnightField := board.Field{Figure: whiteKnight, Cords: whiteKnightCords, Filled: true}
	chessBoard.SetField(whiteKnightField)
	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 4, Row: 0}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)
	whiteRook := board.Figure{FigureType: board.Rook, FigureSide: board.White, Moved: false}
	whiteRookCords := board.Cords{Col: 0, Row: 0}
	whiteRookField := board.Field{Figure: whiteRook, Cords: whiteRookCords, Filled: true}
	chessBoard.SetField(whiteRookField)
	castlingMoveValidator := board.CastlingMoveValidator{ActualBoard: &chessBoard}
	destinationCastleField := chessBoard.GetField(board.Cords{Col: 2, Row: 0})
	castlingMove := board.MakeMove(whiteKingField, destinationCastleField, board.EmptyType)

	isCastled := castlingMoveValidator.Validate(castlingMove)

	assert.False(t, isCastled)
}

func TestBlackLongCastleValidation_FailKingMovedBefore(t *testing.T) {
	chessBoard := board.MakeBoard()
	blackKing := board.Figure{FigureType: board.King, FigureSide: board.Black, Moved: true}
	blackKingCords := board.Cords{Col: 4, Row: 7}
	blackKingField := board.Field{Figure: blackKing, Cords: blackKingCords, Filled: true}
	chessBoard.SetField(blackKingField)
	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black, Moved: false}
	blackRookCords := board.Cords{Col: 0, Row: 7}
	blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
	chessBoard.SetField(blackRookField)
	castlingMoveValidator := board.CastlingMoveValidator{ActualBoard: &chessBoard}
	destinationCastleField := chessBoard.GetField(board.Cords{Col: 2, Row: 7})
	castlingMove := board.MakeMove(blackKingField, destinationCastleField, board.EmptyType)

	isCastled := castlingMoveValidator.Validate(castlingMove)

	assert.False(t, isCastled)
}

func TestBlackLongCastleValidation_FailRookMovedBefore(t *testing.T) {
	chessBoard := board.MakeBoard()
	blackKing := board.Figure{FigureType: board.King, FigureSide: board.Black, Moved: false}
	blackKingCords := board.Cords{Col: 4, Row: 7}
	blackKingField := board.Field{Figure: blackKing, Cords: blackKingCords, Filled: true}
	chessBoard.SetField(blackKingField)
	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black, Moved: true}
	blackRookCords := board.Cords{Col: 0, Row: 7}
	blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
	chessBoard.SetField(blackRookField)
	castlingMoveValidator := board.CastlingMoveValidator{ActualBoard: &chessBoard}
	destinationCastleField := chessBoard.GetField(board.Cords{Col: 2, Row: 7})
	castlingMove := board.MakeMove(blackKingField, destinationCastleField, board.EmptyType)

	isCastled := castlingMoveValidator.Validate(castlingMove)

	assert.False(t, isCastled)
}
