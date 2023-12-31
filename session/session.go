package session

import "chess/board"

type Session struct {
	ActualBoard   *board.Board
	BoardHistory  []board.Board
	moveSide      board.FigureSide
	moveGenerator board.MoveGenerator
}

type MoveRequest struct {
	DepartureCords   board.Cords
	DestinationCords board.Cords
	PromoteToType    board.FigureType
}

func MakeDefaultSession() Session {
	chessboard := board.InitDefaultBoard()
	return Session{
		ActualBoard:   chessboard,
		BoardHistory:  make([]board.Board, 0, 50),
		moveSide:      board.White,
		moveGenerator: board.MakeMoveGenerator(board.InitValidators(chessboard)),
	}
}

func MakeSession(chessBoard *board.Board) Session {
	return Session{
		ActualBoard:   chessBoard,
		BoardHistory:  make([]board.Board, 0, 50),
		moveSide:      board.White,
		moveGenerator: board.MakeMoveGenerator(board.InitValidators(chessBoard)),
	}
}

func (session *Session) Move(moveRequest MoveRequest) bool {
	departure := session.ActualBoard.GetField(moveRequest.DepartureCords)
	destination := session.ActualBoard.GetField(moveRequest.DestinationCords)
	if departure.Figure.FigureSide != session.moveSide {
		return false
	}

	move := board.MakeMove(departure, destination, moveRequest.PromoteToType)

	if !session.moveGenerator.IsValidMove(move) {
		return false
	}

	newActualBoard := session.ActualBoard.Move(move)

	if session.moveSide == board.White {
		session.moveSide = board.Black
	} else {
		session.moveSide = board.White
	}
	session.BoardHistory = append(session.BoardHistory, *session.ActualBoard)
	session.ActualBoard = &newActualBoard
	return true
}
