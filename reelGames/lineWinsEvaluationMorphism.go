package reelGames

import (
	"sort"

	"github.com/kbeserra/transformGames/representation"
)

type LineWinsEvaluationMorphism struct {
	Name         string
	Lines        [][]int
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

func (M *LineWinsEvaluationMorphism) Apply(o *representation.Outcome, _ representation.Parameter) (*representation.Outcome, error) {

	S, ok := o.State.Copy().(*ReelGameState)
	if !ok {
		return nil, representation.ErrFailedToCastFromTo(o.State, &ReelGameState{})
	}

	awards := make([]representation.Award, 0)

	for _, l := range M.Lines {
		line, err := S.B.ProjectToLine(l)
		if err != nil {
			return nil, err
		}
		for _, seq := range M.PaySequences {
			if seq.matchesLine(line) {
				if A, in := M.PayTable[seq.Name]; in {
					cells := make([][]int, len(line))
					for i, r := range l {
						cells[i] = []int{r}
					}
					award := BoardAward{
						name:   seq.Name,
						cells:  cells,
						awards: A,
					}

					awards = append(awards, award)
				}
				break
			}
		}
	}

	return &representation.Outcome{
		Previous: o,
		M:        M,
		State:    S,
		Awards:   awards,
	}, nil
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
	rtn := make([]representation.Award, 0)
	for _, awards := range M.PayTable {
		rtn = append(rtn, awards...)
	}
	return rtn
}
