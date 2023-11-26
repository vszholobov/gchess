package board

type Move struct {
	Departure      Field
	Destination    Field
	representation string
}

// MakeMove TODO: make representation
func MakeMove(departure Field, destination Field) Move {
	return Move{Departure: departure, Destination: destination, representation: ""}
}

func (move Move) String() string {
	return move.representation
}

//type KillMove struct{}
//
//func (KillMove) stringRepresentation() string {
//	//TODO implement me
//	panic("implement me")
//}
//
//type PromoteMove struct{}
//
//func (PromoteMove) stringRepresentation() string {
//	//TODO implement me
//	panic("implement me")
//}
//
//type CastleMove struct{}
//
//func (CastleMove) stringRepresentation() string {
//	//TODO implement me
//	panic("implement me")
//}
