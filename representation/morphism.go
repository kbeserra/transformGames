package representation

import "errors"

var (
	ErrFailedToCastParameter                  = errors.New("failed to cast parameter to necessary type")
	errLengthOfParametersNotLengthOfMorphisms = errors.New("length of parameters is not equal to the length of morphisms")
)

type Morphism interface {
	Init() error
	String() string
	Apply(*Outcome, Parameter) (*Outcome, error)
	EnumerateParameters() (Parameterization, error)
	Awards() []Award
}

/*

 */
type IdentityMorphism struct {
	Name string
}

func (M *IdentityMorphism) Init() error {
	return nil
}

func (M *IdentityMorphism) String() string {
	return M.Name
}

func (M *IdentityMorphism) Apply(o *Outcome, sigma Parameter) (*Outcome, error) {
	return o, nil
}

func (M *IdentityMorphism) EnumerateParameters() (Parameterization, error) {
	return &ConstantParameterization{
			ParameterizationBase: ParameterizationBase{
				M: M,
			},
			C: nil},
		nil
}

func (M *IdentityMorphism) Awards() []Award {
	return nil
}

/*

 */
type ConcatenationMorphism struct {
	Name      string
	Morphisms []Morphism

	Expand       bool
	IgnoreStates bool
	IgnoreAwards bool
}

func (M *ConcatenationMorphism) Init() error {
	for _, m := range M.Morphisms {
		if err := m.Init(); err != nil {
			return err
		}
	}
	return nil
}

func (M *ConcatenationMorphism) String() string {
	return M.Name
}

func (M *ConcatenationMorphism) tidyOutcome(root, end *Outcome) (*Outcome, error) {
	if M.Expand {
		return &Outcome{
			Previous: end,
			M:        M,
			State:    nil,
			Awards:   nil,
		}, nil
	} else {
		var states GameState
		var awards []Award

		if M.IgnoreStates {
			states = end.State
		} else {
			states = &SequenceGameState{States: end.AccumulateStates(root)}
		}
		if M.IgnoreAwards {
			awards = end.Awards
		} else {
			awards = end.AccumulateAwards(root)
		}

		return &Outcome{
			Previous: root,
			M:        M,
			State:    states,
			Awards:   awards,
		}, nil
	}
}

func (M *ConcatenationMorphism) Apply(root *Outcome, sigma Parameter) (*Outcome, error) {
	sigmas, ok := sigma.Value().([]Parameter)
	if !ok {
		return nil, ErrFailedToCastParameter
	}
	if len(sigmas) != len(M.Morphisms) {
		return nil, errLengthOfParametersNotLengthOfMorphisms
	}

	p := root
	var err error
	for i, m := range M.Morphisms {
		p, err = m.Apply(p, sigmas[i])
		if err != nil {
			return nil, err
		}
	}

	return M.tidyOutcome(root, p)
}

func (M *ConcatenationMorphism) EnumerateParameters() (Parameterization, error) {
	enumerations := make([]Parameterization, len(M.Morphisms))
	for i, m := range M.Morphisms {
		e, err := m.EnumerateParameters()
		if err != nil {
			return nil, err
		}
		enumerations[i] = e
	}
	return nil, nil
}

func (M *ConcatenationMorphism) Awards() []Award {
	rtn := make([]Award, 0)
	for _, a := range M.Morphisms {
		rtn = append(rtn, a.Awards()...)
	}
	return rtn
}
