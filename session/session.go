package session

import "chess/board"

type Session struct {
	ActualBoard  *board.Board
	BoardHistory []board.Board
}

func MakeSession() Session {
	return Session{ActualBoard: board.InitDefaultBoard(), BoardHistory: make([]board.Board, 0, 50)}
}

func (session *Session) Move(departureCords board.Cords, destinationCords board.Cords) {
	isMoved, newActualBoard := session.ActualBoard.Move(departureCords, destinationCords)
	if !isMoved {
		return
	}
	session.BoardHistory = append(session.BoardHistory, *session.ActualBoard)
	session.ActualBoard = newActualBoard
}
