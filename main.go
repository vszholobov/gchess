package main

import (
	"chess/board"
	"fmt"
)

func main() {
	chessboard := initDefaultBoard()
	validators := initValidators()

	fmt.Println(chessboard)
	fmt.Println(validators)
}

func initValidators() map[board.FigureType][]board.MoveValidator {
	validators := make(map[board.FigureType][]board.MoveValidator, 6)
	validators[board.King] = []board.MoveValidator{
		board.BordersBreachValidator{},
		board.DepartureEqualsDestinationValidator{},
		board.NotAllyChessmanValidator{},
		board.KingMoveValidator{},
	}
	validators[board.Pawn] = []board.MoveValidator{
		board.BordersBreachValidator{},
		board.DepartureEqualsDestinationValidator{},
		board.NotAllyChessmanValidator{},
		board.LinePathValidator{},
		board.PawnMoveValidator{},
	}
	validators[board.Rook] = []board.MoveValidator{
		board.BordersBreachValidator{},
		board.DepartureEqualsDestinationValidator{},
		board.NotAllyChessmanValidator{},
		board.LinePathValidator{},
		board.RookMoveValidator{},
	}
	validators[board.Knight] = []board.MoveValidator{
		board.BordersBreachValidator{},
		board.DepartureEqualsDestinationValidator{},
		board.NotAllyChessmanValidator{},
		board.KingMoveValidator{},
	}
	validators[board.Bishop] = []board.MoveValidator{
		board.BordersBreachValidator{},
		board.DepartureEqualsDestinationValidator{},
		board.NotAllyChessmanValidator{},
		board.DiagonalPathValidator{},
		board.BishopMoveValidator{},
	}
	validators[board.Queen] = []board.MoveValidator{
		board.BordersBreachValidator{},
		board.DepartureEqualsDestinationValidator{},
		board.NotAllyChessmanValidator{},
		board.LinePathValidator{},
		board.DiagonalPathValidator{},
		board.QueenMoveValidator{},
	}
	return validators
}

func move(currentBoard board.Board, departure board.Field, destination board.Field, moveValidators []board.MoveValidator) {
	move := board.MakeMove(departure, destination)

	// TODO: получать цепочку валидаторов по типу фигуры
	for _, validator := range moveValidators {
		validMove := validator.Validate(currentBoard, move)
		if !validMove {
			panic("Wrong Move")
		}
	}

	// TODO: заполнить историю бордов
	actualBoard := currentBoard.Copy()
	//append(boardHist, actualBoard)

	// Перемещение. TODO: учесть рокировку и повышение фигуры
	newDestination := board.Field{Figure: departure.Figure, Cords: destination.Cords, Filled: true}
	newDeparture := board.Field{Cords: departure.Cords, Filled: false}
	actualBoard.SetField(newDestination)
	actualBoard.SetField(newDeparture)
}

func initDefaultBoard() board.Board {
	chessboard := board.MakeBoard()
	whitePawn := board.Figure{FigureType: board.Pawn, FigureSide: board.White}
	blackPawn := board.Figure{FigureType: board.Pawn, FigureSide: board.Black}
	for col := 0; col < board.SIZE; col++ {
		chessboard.SetField(board.Field{Figure: whitePawn, Cords: board.Cords{Col: col, Row: 1}, Filled: true})
		chessboard.SetField(board.Field{Figure: blackPawn, Cords: board.Cords{Col: col, Row: 6}, Filled: true})
	}
	whiteRook := board.Figure{FigureType: board.Rook, FigureSide: board.Black}
	chessboard.SetField(board.Field{Figure: whiteRook, Cords: board.Cords{Col: 0, Row: 0}, Filled: true})
	chessboard.SetField(board.Field{Figure: whiteRook, Cords: board.Cords{Col: 7, Row: 0}, Filled: true})
	blackRook := board.Figure{FigureType: board.Rook, FigureSide: board.White}
	chessboard.SetField(board.Field{Figure: blackRook, Cords: board.Cords{Col: 0, Row: 7}, Filled: true})
	chessboard.SetField(board.Field{Figure: blackRook, Cords: board.Cords{Col: 7, Row: 7}, Filled: true})
	return chessboard
}
