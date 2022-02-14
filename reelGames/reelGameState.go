package reelGames

import (
	"github.com/kbeserra/transformGames/representation"
)

type ReelGameState struct {
	B Board
}

func (S ReelGameState) String() string {
	return S.B.String()
}

func (S ReelGameState) Copy() representation.OutcomeState {
	return &ReelGameState{
		B: *(S.B.Copy().(*Board)),
	}
}
