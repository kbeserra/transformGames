package representation

import "strings"

type OutcomeState interface {
	String() string
	Copy() OutcomeState
}

type SequenceOutcomeState struct {
	States []OutcomeState
}

func (sgs SequenceOutcomeState) String() string {
	var sb strings.Builder
	for _, s := range sgs.States {
		sb.WriteString(s.String())
	}
	return sb.String()
}

func (sgs SequenceOutcomeState) Copy() OutcomeState {
	states := make([]OutcomeState, len(sgs.States))
	for i, s := range sgs.States {
		states[i] = s.Copy()
	}
	return &SequenceOutcomeState{States: states}
}
