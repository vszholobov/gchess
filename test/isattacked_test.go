package test

import (
	"chess/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFieldIsAttackedByBishop_MainDiagonal(t *testing.T) {
	chessBoard := board.MakeBoard()

	whiteBishop := board.Figure{FigureType: board.Bishop, FigureSide: board.White, Moved: true}
	whiteBishopCords := board.Cords{Col: 0, Row: 0}
	whiteBishopField := board.Field{Figure: whiteBishop, Cords: whiteBishopCords, Filled: true}
	chessBoard.SetField(whiteBishopField)

	isAttacked := chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 7, Row: 7}, board.Black)
	assert.True(t, isAttacked)
}

func TestFieldIsAttackedByBishop_SideDiagonal(t *testing.T) {
	chessBoard := board.MakeBoard()

	whiteBishop := board.Figure{FigureType: board.Bishop, FigureSide: board.White, Moved: true}
	whiteBishopCords := board.Cords{Col: 0, Row: 7}
	whiteBishopField := board.Field{Figure: whiteBishop, Cords: whiteBishopCords, Filled: true}
	chessBoard.SetField(whiteBishopField)

	isAttacked := chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 7, Row: 0}, board.Black)
	assert.True(t, isAttacked)
}

func TestFieldIsNotAttackedByBishop_AnotherDiagonal(t *testing.T) {
	chessBoard := board.MakeBoard()

	whiteBishop := board.Figure{FigureType: board.Bishop, FigureSide: board.White, Moved: true}
	whiteBishopCords := board.Cords{Col: 1, Row: 0}
	whiteBishopField := board.Field{Figure: whiteBishop, Cords: whiteBishopCords, Filled: true}
	chessBoard.SetField(whiteBishopField)

	isAttacked := chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 7, Row: 7}, board.Black)
	assert.False(t, isAttacked)
}

func TestFieldIsAttackedByRook_SameCol(t *testing.T) {
	chessBoard := board.MakeBoard()
	whiteRook := board.Figure{FigureType: board.Rook, FigureSide: board.White, Moved: false}
	whiteRookCords := board.Cords{Col: 0, Row: 0}
	whiteRookField := board.Field{Figure: whiteRook, Cords: whiteRookCords, Filled: true}
	chessBoard.SetField(whiteRookField)

	assert.True(t, chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 7, Row: 0}, board.Black))
	assert.True(t, chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 0, Row: 7}, board.Black))
}

func TestFieldIsAttackedByPawn(t *testing.T) {
	chessBoard := board.MakeBoard()
	whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White, Moved: false}
	whitePawnCords := board.Cords{Col: 1, Row: 1}
	whitePawnField := board.Field{Figure: whitePawn, Cords: whitePawnCords, Filled: true}
	chessBoard.SetField(whitePawnField)

	assert.True(t, chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 0, Row: 2}, board.Black))
	assert.True(t, chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 2, Row: 2}, board.Black))
}

func TestFieldIsAttackedByKing(t *testing.T) {
	chessBoard := board.MakeBoard()
	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 1, Row: 1}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)

	assert.True(t, chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 0, Row: 0}, board.Black))
	assert.True(t, chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 0, Row: 1}, board.Black))
	assert.True(t, chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 0, Row: 2}, board.Black))
	assert.True(t, chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 1, Row: 2}, board.Black))
	assert.True(t, chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 2, Row: 2}, board.Black))
	assert.True(t, chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 2, Row: 1}, board.Black))
	assert.True(t, chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 2, Row: 0}, board.Black))
	assert.True(t, chessBoard.IsFieldAttackedByOpposedSide(board.Cords{Col: 1, Row: 0}, board.Black))
}

func TestKingIsNotAttackedAfterKingMove(t *testing.T) {
	chessBoard := board.MakeBoard()
	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 4, Row: 0}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)
	destinationCords := board.Cords{Col: 3, Row: 0}

	destinationField := chessBoard.GetField(destinationCords)
	kingMove := board.MakeMove(whiteKingField, destinationField, board.EmptyType)

	validator := board.KingIsNotAttackedAfterMoveValidator{ActualBoard: &chessBoard}
	kingIsAttacked := !validator.Validate(kingMove)

	assert.False(t, kingIsAttacked)
}

func TestKingIsAttackedAfterKingMove_KingAlreadyAttacked(t *testing.T) {
	chessBoard := board.MakeBoard()

	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black, Moved: true}
	blackRookCords := board.Cords{Col: 0, Row: 0}
	blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
	chessBoard.SetField(blackRookField)

	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 4, Row: 0}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)
	destinationCords := board.Cords{Col: 3, Row: 0}

	destinationField := chessBoard.GetField(destinationCords)
	kingMove := board.MakeMove(whiteKingField, destinationField, board.EmptyType)

	validator := board.KingIsNotAttackedAfterMoveValidator{ActualBoard: &chessBoard}
	kingIsAttacked := !validator.Validate(kingMove)

	assert.True(t, kingIsAttacked)
}

func TestKingIsAttackedAfterKingMove_KingMovedToAttackedField(t *testing.T) {
	chessBoard := board.MakeBoard()

	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black, Moved: true}
	blackRookCords := board.Cords{Col: 0, Row: 1}
	blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
	chessBoard.SetField(blackRookField)

	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 4, Row: 0}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)
	destinationCords := board.Cords{Col: 4, Row: 1}

	destinationField := chessBoard.GetField(destinationCords)
	kingMove := board.MakeMove(whiteKingField, destinationField, board.EmptyType)

	validator := board.KingIsNotAttackedAfterMoveValidator{ActualBoard: &chessBoard}
	kingIsAttacked := !validator.Validate(kingMove)

	assert.True(t, kingIsAttacked)
}

func TestKingIsAttacked_CoveringFigureMoved(t *testing.T) {
	chessBoard := board.MakeBoard()

	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black, Moved: true}
	blackRookCords := board.Cords{Col: 0, Row: 0}
	blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
	chessBoard.SetField(blackRookField)

	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 4, Row: 0}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)

	whiteBishop := board.Figure{FigureType: board.Bishop, FigureSide: board.White, Moved: false}
	whiteBishopCords := board.Cords{Col: 3, Row: 0}
	whiteBishopField := board.Field{Figure: whiteBishop, Cords: whiteBishopCords, Filled: true}
	chessBoard.SetField(whiteBishopField)

	destinationCords := board.Cords{Col: 4, Row: 1}

	destinationField := chessBoard.GetField(destinationCords)
	bishopMove := board.MakeMove(whiteBishopField, destinationField, board.EmptyType)

	validator := board.KingIsNotAttackedAfterMoveValidator{ActualBoard: &chessBoard}
	kingIsAttacked := !validator.Validate(bishopMove)

	assert.True(t, kingIsAttacked)
}

func TestKingIsNotAttacked_KingAlreadyAttacked_AttackingFigureKilled(t *testing.T) {
	chessBoard := board.MakeBoard()

	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black, Moved: true}
	blackRookCords := board.Cords{Col: 0, Row: 0}
	blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
	chessBoard.SetField(blackRookField)

	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 4, Row: 0}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)

	whiteBishop := board.Figure{FigureType: board.Bishop, FigureSide: board.White, Moved: false}
	whiteBishopCords := board.Cords{Col: 1, Row: 1}
	whiteBishopField := board.Field{Figure: whiteBishop, Cords: whiteBishopCords, Filled: true}
	chessBoard.SetField(whiteBishopField)

	bishopMove := board.MakeMove(whiteBishopField, blackRookField, board.EmptyType)

	validator := board.KingIsNotAttackedAfterMoveValidator{ActualBoard: &chessBoard}
	kingIsAttacked := !validator.Validate(bishopMove)

	assert.False(t, kingIsAttacked)
}

func TestKingIsAttacked_KingAlreadyAttacked_AttackingFigureKilledByCoveringPiece_AnotherAttackerOpened(t *testing.T) {
	chessBoard := board.MakeBoard()

	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black, Moved: true}
	blackRookCords := board.Cords{Col: 0, Row: 0}
	blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
	chessBoard.SetField(blackRookField)

	blackQueen := board.Figure{FigureType: board.Queen, FigureSide: board.Black, Moved: true}
	blackQueenCords := board.Cords{Col: 4, Row: 5}
	blackQueenField := board.Field{Figure: blackQueen, Cords: blackQueenCords, Filled: true}
	chessBoard.SetField(blackQueenField)

	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 4, Row: 0}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)

	whiteBishop := board.Figure{FigureType: board.Bishop, FigureSide: board.White, Moved: false}
	whiteBishopCords := board.Cords{Col: 4, Row: 4}
	whiteBishopField := board.Field{Figure: whiteBishop, Cords: whiteBishopCords, Filled: true}
	chessBoard.SetField(whiteBishopField)

	bishopMove := board.MakeMove(whiteBishopField, blackRookField, board.EmptyType)

	validator := board.KingIsNotAttackedAfterMoveValidator{ActualBoard: &chessBoard}
	kingIsAttacked := !validator.Validate(bishopMove)

	assert.True(t, kingIsAttacked)
}

func TestKingIsNotAttacked_KingAlreadyAttacked_KingCovered(t *testing.T) {
	chessBoard := board.MakeBoard()

	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black, Moved: true}
	blackRookCords := board.Cords{Col: 0, Row: 0}
	blackRookField := board.Field{Figure: blackRook, Cords: blackRookCords, Filled: true}
	chessBoard.SetField(blackRookField)

	whiteKing := board.Figure{FigureType: board.King, FigureSide: board.White, Moved: false}
	whiteKingCords := board.Cords{Col: 4, Row: 0}
	whiteKingField := board.Field{Figure: whiteKing, Cords: whiteKingCords, Filled: true}
	chessBoard.SetField(whiteKingField)

	whiteBishop := board.Figure{FigureType: board.Bishop, FigureSide: board.White, Moved: false}
	whiteBishopCords := board.Cords{Col: 4, Row: 1}
	whiteBishopField := board.Field{Figure: whiteBishop, Cords: whiteBishopCords, Filled: true}
	chessBoard.SetField(whiteBishopField)

	destinationCords := board.Cords{Col: 3, Row: 0}
	destinationField := chessBoard.GetField(destinationCords)
	bishopMove := board.MakeMove(whiteBishopField, destinationField, board.EmptyType)

	validator := board.KingIsNotAttackedAfterMoveValidator{ActualBoard: &chessBoard}
	kingIsAttacked := !validator.Validate(bishopMove)

	assert.False(t, kingIsAttacked)
}
