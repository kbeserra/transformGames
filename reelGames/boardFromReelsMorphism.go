package reelGames

import (
	"github.com/kbeserra/transformGames/representation"
)

type BoardFromReelsMorphism struct {
	Name    string
	Heights []uint
	Reels   [][]string
	Weights [][]uint64
}

func (M *BoardFromReelsMorphism) Init() error {
	return nil
}

func (M *BoardFromReelsMorphism) String() string {
	return M.Name
}

func (M *BoardFromReelsMorphism) Apply(o *representation.Outcome, sigma representation.Parameter) (*representation.Outcome, error) {
	unpacked, ok := sigma.Value().([]interface{})
	if !ok {
		return nil, representation.ErrFailedToCastParameter
	}
	stops := make([]int, len(unpacked))
	for i, v := range unpacked {
		s, ok := v.(int)
		if !ok {
			return nil, representation.ErrFailedToCastParameter
		}
		stops[i] = s
	}

	board := Board{}
	err := board.FromShape(M.Heights)
	if err != nil {
		return nil, err
	}
	board.FillFromReels(M.Reels, stops)
	if err != nil {
		return nil, err
	}

	return &representation.Outcome{
		Previous: o,
		M:        M,
		State: &ReelGameState{
			B: board,
		},
		Awards: nil,
	}, nil
}

func (M *BoardFromReelsMorphism) EnumerateParameters() (representation.Parameterization, error) {

	reelParameterizations := make([]representation.Parameterization, len(M.Reels))
	for i, reel := range M.Reels {
		var weights []uint64
		if M.Weights != nil {
			weights = M.Weights[i]
		}
		reelParameterizations[i] = &representation.IntegerIntervalParameterization{
			LowerBound: 0,
			UpperBound: len(reel),
			Weights:    weights,
		}
	}

	return &representation.SequenceParameterization{
		ParameterizationBase: representation.ParameterizationBase{
			M: M,
		},
		Parameterizations: reelParameterizations,
	}, nil

}

func (M *BoardFromReelsMorphism) Awards() []representation.Award {
	return nil
}
