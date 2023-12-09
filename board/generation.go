package board

type MoveGenerator struct {
	validatorsMap map[FigureType][]MoveValidator
}

func MakeMoveGenerator(validators map[FigureType][]MoveValidator) MoveGenerator {
	return MoveGenerator{validatorsMap: validators}
}

func (moveGenerator MoveGenerator) HasAvailableMoves(chessBoard Board, field Field) bool {
	for col := 0; col < ChessboardSize; col++ {
		for row := 0; row < ChessboardSize; row++ {
			move := MakeMove(field, chessBoard.GetField(Cords{Col: col, Row: row}), EmptyType)
			if moveGenerator.IsValidMove(move) {
				return true
			}
		}
	}
	return false
}

func (moveGenerator MoveGenerator) IsValidMove(move Move) bool {
	validators := moveGenerator.validatorsMap[move.Departure().Figure.FigureType]
	for _, validator := range validators {
		isValid := validator.Validate(move)
		if !isValid {
			return false
		}
	}
	return true
}
