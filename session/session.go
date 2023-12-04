package session

import "chess/board"

type Session struct {
	ActualBoard  *board.Board
	BoardHistory []board.Board
	moveSide     board.FigureSide
}

func MakeSession() Session {
	return Session{ActualBoard: board.InitDefaultBoard(), BoardHistory: make([]board.Board, 0, 50), moveSide: board.White}
}

func (session *Session) Move(departureCords board.Cords, destinationCords board.Cords) bool {
	if session.ActualBoard.GetField(departureCords).Figure.FigureSide != session.moveSide {
		return false
	}

	isMoved, newActualBoard := session.ActualBoard.Move(departureCords, destinationCords)
	if !isMoved {
		return false
	}
	if session.moveSide == board.White {
		session.moveSide = board.Black
	} else {
		session.moveSide = board.White
	}
	session.BoardHistory = append(session.BoardHistory, *session.ActualBoard)
	session.ActualBoard = newActualBoard
	return true
}
