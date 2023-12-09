package test

import (
	"chess/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasAvailableMoves_PawnNoAvailableMoves_CoversKing(t *testing.T) {
	chessBoard := board.MakeBoard()
	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 2, Row: 1}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)

	whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
	whitePawnCords := board.Cords{Col: 1, Row: 1}
	whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
	chessBoard.SetField(whitePawnField)

	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black, Moved: false}
	blackRookCords := board.Cords{Col: 0, Row: 1}
	blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
	chessBoard.SetField(blackRookField)

	validators := board.InitValidators(&chessBoard)
	generator := board.MakeMoveGenerator(validators)

	hasAvailableMoves := generator.HasAvailableMoves(chessBoard, whitePawnField)
	assert.False(t, hasAvailableMoves)
}

func TestHasAvailableMoves_PawnHasAvailableMoves(t *testing.T) {
	chessBoard := board.MakeBoard()
	whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
	whitePawnCords := board.Cords{Col: 0, Row: 5}
	whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
	chessBoard.SetField(whitePawnField)

	validators := board.InitValidators(&chessBoard)
	generator := board.MakeMoveGenerator(validators)

	hasAvailableMoves := generator.HasAvailableMoves(chessBoard, whitePawnField)
	assert.True(t, hasAvailableMoves)
}

func TestHasAvailableMoves_PawnNoAvailableMoves(t *testing.T) {
	chessBoard := board.MakeBoard()
	whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
	whitePawnCords := board.Cords{Col: 0, Row: 5}
	whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
	chessBoard.SetField(whitePawnField)

	blackPawn := board.Figure{FigureType: board.Pawn, FigureSide: board.Black, Moved: false}
	blackPawnCords := board.Cords{Col: 0, Row: 6}
	blackPawnField := board.Field{Figure: blackPawn, Cords: blackPawnCords, Filled: true}
	chessBoard.SetField(blackPawnField)

	validators := board.InitValidators(&chessBoard)
	generator := board.MakeMoveGenerator(validators)

	hasAvailableMoves := generator.HasAvailableMoves(chessBoard, whitePawnField)
	assert.False(t, hasAvailableMoves)
}
