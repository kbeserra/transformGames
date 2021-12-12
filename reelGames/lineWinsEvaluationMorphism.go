package reelGames

import (
	"sort"

	"github.com/kbeserra/transformGames/representation"
)

type LineWinsEvaluationMorphism struct {
	Name  string
	Lines [][]uint
	PaySequences []SymbolSequence
	PayTable     map[string][]representation.Award
}

func (M *LineWinsEvaluationMorphism) Init() error {

	paySeqOrder := func(sIndex, tIndex int) bool {
		s := M.PaySequences[sIndex]
		t := M.PaySequences[tIndex]
		sAwards, sIn := M.PayTable[s.Name]
		if !sIn {
			sAwards = nil
		}
		tAwards, tIn := M.PayTable[t.Name]
		if !tIn {
			tAwards = nil
		}
		sValue := representation.AwardValueSum(sAwards)
		tValue := representation.AwardValueSum(tAwards)

		if sValue == tValue {
			return s.subSetEq(t)
		}
		return sValue > tValue
	}

	sort.SliceStable(M.PaySequences, paySeqOrder)

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
