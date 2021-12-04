package reelGames

import (
	"github.com/kbeserra/transformGames/representation"
)

type RemoveRowMorphism struct {
	Name string
	Row  int
}

func (M *RemoveRowMorphism) Init() error {
	return nil
}

func (M *RemoveRowMorphism) String() string {
	return M.Name
}

func (M *RemoveRowMorphism) Apply(o *representation.Outcome, sigma representation.Parameter) (*representation.Outcome, error) {
	S, ok := o.State.Copy().(*ReelGameState)
	if !ok {
		return nil, representation.ErrFailedToCastFromTo(o.State, &ReelGameState{})
	}
	if err := S.B.RemoveRow(M.Row); err != nil {
		return nil, err
	}

	return &representation.Outcome{
		Previous: o,
		M:        M,
		State:    S,
		Awards:   nil,
	}, nil
}

func (M *RemoveRowMorphism) EnumerateParameters() (representation.Parameterization, error) {
	return &representation.ConstantParameterization{
			ParameterizationBase: representation.ParameterizationBase{
				M: M,
			},
			C: nil},
		nil
}

func (M *RemoveRowMorphism) Awards() []representation.Award {
	return nil
}
