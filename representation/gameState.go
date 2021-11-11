package representation

import "strings"

type GameState interface {
	String() string
	Copy() GameState
}

type SequenceGameState struct {
	States []GameState
}

func (sgs SequenceGameState) String() string {
	var sb strings.Builder
	for _, s := range sgs.States {
		sb.WriteString(s.String())
	}
	return sb.String()
}

func (sgs SequenceGameState) Copy() GameState {
	states := make([]GameState, len(sgs.States))
	for i, s := range sgs.States {
		states[i] = s.Copy()
	}
	return &SequenceGameState{States: states}
}
