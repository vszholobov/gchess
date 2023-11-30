package test

import (
	"chess/board"
	"chess/session"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeBoard(t *testing.T) {
	b1 := board.MakeBoard()
	assert.NotNil(t, b1)
}

func TestCopyBoard(t *testing.T) {
	b1 := board.MakeBoard()
	b2 := b1.Copy()
	assert.Equal(t, b1, b2)
}

func TestChangeCopyBoard(t *testing.T) {
	b1 := board.MakeBoard()
	b2 := b1.Copy()
	b2.SetField(board.Field{Filled: true})
	assert.NotEqual(t, b1, b2)
}

func TestInitDefaultBoard(t *testing.T) {
	defaultBoard := board.InitDefaultBoard()
	// pawn
	whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
	blackPawn := board.Figure{FigureType: board.Pawn, FigureSide: board.Black, Moved: false}
	for col := 0; col < board.SIZE; col++ {
		assert.Equal(t, whitePawn, defaultBoard.GetField(board.Cords{Col: col, Row: 1}).Figure)
		assert.Equal(t, blackPawn, defaultBoard.GetField(board.Cords{Col: col, Row: 6}).Figure)
	}

	// rook
	whiteRook := board.Figure{FigureType: board.Rook, FigureSide: board.White, Moved: false}
	assert.Equal(t, whiteRook, defaultBoard.GetField(board.Cords{Col: 0, Row: 0}).Figure)
	assert.Equal(t, whiteRook, defaultBoard.GetField(board.Cords{Col: 7, Row: 0}).Figure)
	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black, Moved: false}
	assert.Equal(t, blackRook, defaultBoard.GetField(board.Cords{Col: 0, Row: 7}).Figure)
	assert.Equal(t, blackRook, defaultBoard.GetField(board.Cords{Col: 7, Row: 7}).Figure)

	// knight
	whiteKnight := board.Figure{FigureType: board.Knight, FigureSide: board.White, Moved: false}
	assert.Equal(t, whiteKnight, defaultBoard.GetField(board.Cords{Col: 1, Row: 0}).Figure)
	assert.Equal(t, whiteKnight, defaultBoard.GetField(board.Cords{Col: 6, Row: 0}).Figure)
	blackKnight := board.Figure{FigureType: board.Knight, FigureSide: board.Black, Moved: false}
	assert.Equal(t, blackKnight, defaultBoard.GetField(board.Cords{Col: 1, Row: 7}).Figure)
	assert.Equal(t, blackKnight, defaultBoard.GetField(board.Cords{Col: 6, Row: 7}).Figure)

	// bishop
	whiteBishop := board.Figure{FigureType: board.Bishop, FigureSide: board.White, Moved: false}
	assert.Equal(t, whiteBishop, defaultBoard.GetField(board.Cords{Col: 2, Row: 0}).Figure)
	assert.Equal(t, whiteBishop, defaultBoard.GetField(board.Cords{Col: 5, Row: 0}).Figure)
	blackBishop := board.Figure{FigureType: board.Bishop, FigureSide: board.Black, Moved: false}
	assert.Equal(t, blackBishop, defaultBoard.GetField(board.Cords{Col: 2, Row: 7}).Figure)
	assert.Equal(t, blackBishop, defaultBoard.GetField(board.Cords{Col: 5, Row: 7}).Figure)

	// queen
	whiteQueen := board.Figure{FigureType: board.Queen, FigureSide: board.White, Moved: false}
	assert.Equal(t, whiteQueen, defaultBoard.GetField(board.Cords{Col: 3, Row: 0}).Figure)
	blackQueen := board.Figure{FigureType: board.Queen, FigureSide: board.Black, Moved: false}
	assert.Equal(t, blackQueen, defaultBoard.GetField(board.Cords{Col: 3, Row: 7}).Figure)

	// king
	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	assert.Equal(t, whiteKing, defaultBoard.GetField(board.Cords{Col: 4, Row: 0}).Figure)
	blackKing := board.Figure{FigureType: board.King, FigureSide: board.Black, Moved: false}
	assert.Equal(t, blackKing, defaultBoard.GetField(board.Cords{Col: 4, Row: 7}).Figure)

	// empty
	for row := 2; row < 6; row++ {
		for col := 0; col < board.SIZE; col++ {
			assert.False(t, defaultBoard.GetField(board.Cords{Col: col, Row: row}).Filled)
		}
	}
}

func TestSessionMove(t *testing.T) {
	chessSession := session.MakeSession()
	departureCords := board.Cords{Col: 0, Row: 1}
	destinationCords := board.Cords{Col: 0, Row: 3}

	chessSession.Move(departureCords, destinationCords)

	departure := chessSession.ActualBoard.GetField(departureCords)
	destination := chessSession.ActualBoard.GetField(destinationCords)
	assert.False(t, departure.Filled)
	assert.Equal(t, departure.Figure, board.Figure{})
	assert.True(t, destination.Filled)
	assert.Equal(t, destination.Figure, board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: true})
	assert.Len(t, chessSession.BoardHistory, 1)
}

func TestSessionMove_FailSecondWhiteMove(t *testing.T) {
	chessSession := session.MakeSession()
	firstWhiteMoveDepartureCords := board.Cords{Col: 0, Row: 1}
	firstWhiteMoveDestinationCords := board.Cords{Col: 0, Row: 3}
	secondWhiteMoveDepartureCords := board.Cords{Col: 1, Row: 1}
	secondWhiteMoveDestinationCords := board.Cords{Col: 1, Row: 3}

	firstWhiteMoveIsMoved := chessSession.Move(firstWhiteMoveDepartureCords, firstWhiteMoveDestinationCords)
	secondWhiteMoveIsMoved := chessSession.Move(secondWhiteMoveDepartureCords, secondWhiteMoveDestinationCords)

	assert.True(t, firstWhiteMoveIsMoved)
	assert.False(t, secondWhiteMoveIsMoved)
	assert.Len(t, chessSession.BoardHistory, 1)
}

func TestSessionMove_FailFirstBlackMove(t *testing.T) {
	chessSession := session.MakeSession()
	firstBlackMoveDepartureCords := board.Cords{Col: 0, Row: 6}
	firstBlackMoveDestinationCords := board.Cords{Col: 0, Row: 4}

	firstBlackMoveIsMoved := chessSession.Move(firstBlackMoveDepartureCords, firstBlackMoveDestinationCords)

	assert.False(t, firstBlackMoveIsMoved)
	assert.Len(t, chessSession.BoardHistory, 0)
}

func TestSessionMove_SecondBlackMove(t *testing.T) {
	chessSession := session.MakeSession()
	whiteMoveDepartureCords := board.Cords{Col: 0, Row: 1}
	whiteMoveDestinationCords := board.Cords{Col: 0, Row: 3}
	blackMoveDepartureCords := board.Cords{Col: 0, Row: 6}
	blackMoveDestinationCords := board.Cords{Col: 0, Row: 4}

	whiteMoveIsMoved := chessSession.Move(whiteMoveDepartureCords, whiteMoveDestinationCords)
	blackMoveIsMoved := chessSession.Move(blackMoveDepartureCords, blackMoveDestinationCords)

	assert.True(t, whiteMoveIsMoved)
	assert.True(t, blackMoveIsMoved)
	assert.Len(t, chessSession.BoardHistory, 2)

	whiteDeparture := chessSession.ActualBoard.GetField(whiteMoveDepartureCords)
	whiteDestination := chessSession.ActualBoard.GetField(whiteMoveDestinationCords)
	assert.False(t, whiteDeparture.Filled)
	assert.Equal(t, whiteDeparture.Figure, board.Figure{})
	assert.True(t, whiteDestination.Filled)
	assert.Equal(t, whiteDestination.Figure, board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: true})

	blackDeparture := chessSession.ActualBoard.GetField(blackMoveDepartureCords)
	blackDestination := chessSession.ActualBoard.GetField(blackMoveDestinationCords)
	assert.False(t, blackDeparture.Filled)
	assert.Equal(t, blackDeparture.Figure, board.Figure{})
	assert.True(t, blackDestination.Filled)
	assert.Equal(t, blackDestination.Figure, board.Figure{FigureType: board.Pawn, FigureSide: board.Black, Moved: true})
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
	castlingMoveValidator := board.CastlingMoveValidator{}
	destinationCastleField := chessBoard.GetField(board.Cords{Col: 2, Row: 0})
	castlingMove := board.Move{Departure: whiteKingField, Destination: destinationCastleField}

	isCastled := castlingMoveValidator.Validate(&chessBoard, castlingMove)

	assert.True(t, isCastled)
}

func TestBlackLongCastleValidation(t *testing.T) {
	chessBoard := board.MakeBoard()
	blackKing := board.Figure{FigureType: board.King, FigureSide: board.Black, Moved: false}
	blackKingCords := board.Cords{Col: 4, Row: 7}
	blackKingField := board.Field{Figure: blackKing, Cords: blackKingCords, Filled: true}
	chessBoard.SetField(blackKingField)
	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.White, Moved: false}
	blackRookCords := board.Cords{Col: 0, Row: 7}
	blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
	chessBoard.SetField(blackRookField)
	castlingMoveValidator := board.CastlingMoveValidator{}
	destinationCastleField := chessBoard.GetField(board.Cords{Col: 2, Row: 7})
	castlingMove := board.Move{Departure: blackKingField, Destination: destinationCastleField}

	isCastled := castlingMoveValidator.Validate(&chessBoard, castlingMove)

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
	castlingMoveValidator := board.CastlingMoveValidator{}
	destinationCastleField := chessBoard.GetField(board.Cords{Col: 2, Row: 0})
	castlingMove := board.Move{Departure: whiteKingField, Destination: destinationCastleField}

	isCastled := castlingMoveValidator.Validate(&chessBoard, castlingMove)

	assert.False(t, isCastled)
}

func TestBlackLongCastleValidation_FailKingMovedBefore(t *testing.T) {
	chessBoard := board.MakeBoard()
	blackKing := board.Figure{FigureType: board.King, FigureSide: board.Black, Moved: true}
	blackKingCords := board.Cords{Col: 4, Row: 7}
	blackKingField := board.Field{Figure: blackKing, Cords: blackKingCords, Filled: true}
	chessBoard.SetField(blackKingField)
	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.White, Moved: false}
	blackRookCords := board.Cords{Col: 0, Row: 7}
	blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
	chessBoard.SetField(blackRookField)
	castlingMoveValidator := board.CastlingMoveValidator{}
	destinationCastleField := chessBoard.GetField(board.Cords{Col: 2, Row: 7})
	castlingMove := board.Move{Departure: blackKingField, Destination: destinationCastleField}

	isCastled := castlingMoveValidator.Validate(&chessBoard, castlingMove)

	assert.False(t, isCastled)
}

func TestBlackLongCastleValidation_FailRookMovedBefore(t *testing.T) {
	chessBoard := board.MakeBoard()
	blackKing := board.Figure{FigureType: board.King, FigureSide: board.Black, Moved: false}
	blackKingCords := board.Cords{Col: 4, Row: 7}
	blackKingField := board.Field{Figure: blackKing, Cords: blackKingCords, Filled: true}
	chessBoard.SetField(blackKingField)
	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.White, Moved: true}
	blackRookCords := board.Cords{Col: 0, Row: 7}
	blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
	chessBoard.SetField(blackRookField)
	castlingMoveValidator := board.CastlingMoveValidator{}
	destinationCastleField := chessBoard.GetField(board.Cords{Col: 2, Row: 7})
	castlingMove := board.Move{Departure: blackKingField, Destination: destinationCastleField}

	isCastled := castlingMoveValidator.Validate(&chessBoard, castlingMove)

	assert.False(t, isCastled)
}

// TODO: attacked between validator
