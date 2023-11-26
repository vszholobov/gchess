package test

import (
	"chess/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeBoard(t *testing.T) {
	b1 := board.MakeBoard()
	assert.NotNil(t, b1)
}

func TestCopyBoard(t *testing.T) {
	b1 := board.MakeBoard()
	b2 := b1.Copy()
	assert.Equal(t, b1, b2)
}

func TestChangeCopyBoard(t *testing.T) {
	b1 := board.MakeBoard()
	b2 := b1.Copy()
	b2.SetField(board.Field{Filled: true})
	assert.NotEqual(t, b1, b2)
}
