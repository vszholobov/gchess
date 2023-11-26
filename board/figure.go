package board

type Figure struct {
	FigureType FigureType
	FigureSide FigureSide
	Moved      bool
}

type FigureType int

const (
	EmptyType FigureType = iota
	King      FigureType = iota
	Pawn      FigureType = iota
	Rook      FigureType = iota
	Knight    FigureType = iota
	Bishop    FigureType = iota
	Queen     FigureType = iota
)

type FigureSide int

const (
	EmptySide FigureSide = iota
	White     FigureSide = iota
	Black     FigureSide = iota
)
