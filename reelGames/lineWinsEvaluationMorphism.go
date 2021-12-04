package reelGames

import (
	"github.com/kbeserra/transformGames/representation"
)

type LineWinsEvaluationMorphism struct {
	Name  string
	Lines [][]uint
}

func (M *LineWinsEvaluationMorphism) Init() error {
	return nil
}

func (M *LineWinsEvaluationMorphism) String() string {
	return M.Name
}

func (M *LineWinsEvaluationMorphism) Apply(o *representation.Outcome, sigma representation.Parameter) (*representation.Outcome, error) {
	return nil, nil
}

func (M *LineWinsEvaluationMorphism) EnumerateParameters() (representation.Parameterization, error) {
	return &representation.ConstantParameterization{
			ParameterizationBase: representation.ParameterizationBase{
				M: M,
			},
			C: nil},
		nil
}

func (M *LineWinsEvaluationMorphism) Awards() []representation.Award {
	return nil
}
