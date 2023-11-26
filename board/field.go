package board

type Field struct {
	Figure Figure
	Cords  Cords
	Filled bool
	Moved  bool
}

type Cords struct {
	Col int
	Row int
}

func (cords Cords) Equal(cords2 Cords) bool {
	return cords.Col == cords2.Col && cords.Row == cords2.Row
}
