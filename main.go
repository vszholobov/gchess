package main

import (
	"chess/board"
	"fmt"
)

func main() {
	chessboard := board.InitDefaultBoard()
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
