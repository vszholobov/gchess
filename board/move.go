package board

import "math"

type Move interface {
	Departure() Field
	Destination() Field
	String() string
}

// MakeMove TODO: make representation
func MakeMove(departure Field, destination Field) Move {
	colDistance := math.Abs(float64(departure.Cords.Col - destination.Cords.Col))
	if departure.Figure.FigureType == King && colDistance > 1 {
		row := GetDefaultRowBySide(departure.Figure.FigureSide)
		var rookDepartureCords Cords
		var rookDestinationCords Cords
		if destination.Cords.Col == 2 {
			// long side castle
			rookDepartureCords = Cords{Col: 0, Row: row}
			rookDestinationCords = Cords{Col: 3, Row: row}
		} else {
			// short side castle
			rookDepartureCords = Cords{Col: 7, Row: row}
			rookDestinationCords = Cords{Col: 5, Row: row}
		}
		return CastleMove{
			departure:            departure,
			destination:          destination,
			stringRepresentation: "",
			rookDepartureCords:   rookDepartureCords,
			rookDestinationCords: rookDestinationCords,
		}

		//} else if departure.Figure.FigureType == Pawn && (destination.Cords.Row == 0 || destination.Cords.Row == 7) {
		//
	} else {
		return DefaultMove{departure: departure, destination: destination, stringRepresentation: ""}
	}
}

type DefaultMove struct {
	departure            Field
	destination          Field
	stringRepresentation string
}

func (move DefaultMove) Departure() Field {
	return move.departure
}

func (move DefaultMove) Destination() Field {
	return move.destination
}

func (move DefaultMove) String() string {
	return move.stringRepresentation
}

type CastleMove struct {
	departure            Field
	destination          Field
	stringRepresentation string
	rookDepartureCords   Cords
	rookDestinationCords Cords
}

func (move CastleMove) String() string {
	return move.stringRepresentation
}

func (move CastleMove) Departure() Field {
	return move.departure
}

func (move CastleMove) Destination() Field {
	return move.destination
}

func (move CastleMove) RookDepartureCords() Cords {
	return move.rookDepartureCords
}

func (move CastleMove) RookDestinationCords() Cords {
	return move.rookDestinationCords
}

//type KillMove struct {
//	departure            Field
//	destination          Field
//	stringRepresentation string
//}
//
//func (move KillMove) Departure() Field {
//	return move.departure
//}
//
//func (move KillMove) Destination() Field {
//	return move.destination
//}
//
//func (move KillMove) String() string {
//	return move.stringRepresentation
//}
