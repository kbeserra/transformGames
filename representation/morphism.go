package representation

import (
	"errors"
	"github.com/kbeserra/tjson"
)

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

type IdentityMorphism struct {
	Name string
}

func (M *IdentityMorphism) Init() error {
	return nil
}

func (M *IdentityMorphism) String() string {
	return M.Name
}

func (M *IdentityMorphism) Apply(o *Outcome, _ Parameter) (*Outcome, error) {
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

type ConcatenationMorphism struct {
	Name      string
	Morphisms []Morphism

	Expand       bool
	IgnoreStates bool
	IgnoreAwards bool
}

func (M *ConcatenationMorphism) UnmarshalJSON(b []byte) error {
	if err := tjson.ParseJsonTags(b, M); err != nil {
		return err
	}
	if err := tjson.ParseTJsonTags(b, M); err != nil {
		return err
	}
	return nil
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
		var states OutcomeState
		var awards []Award

		if M.IgnoreStates {
			states = end.State
		} else {
			states = &SequenceOutcomeState{States: end.AccumulateStates(root)}
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

	s, ok := sigma.(*SequenceParameter)

	if !ok {
		return nil, ErrFailedToCastParameter
	}

	sigmas := s.Sigmas

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

	return &SequenceParameterization{
		Parameterizations: enumerations,
	}, nil
}

func (M *ConcatenationMorphism) Awards() []Award {
	rtn := make([]Award, 0)
	for _, a := range M.Morphisms {
		rtn = append(rtn, a.Awards()...)
	}
	return rtn
}

/*
	Applies a transformation conditional on the state of the passed outcome, otherwise acts as the identity transform.
*/
type ConditionalFunction func(*Outcome) bool
type ConditionalMorphism struct {
	Name string              `json:"name"`
	M    Morphism            `tjson:"transformation"`
	P    ConditionalFunction `tjson:"conditionalFunction"`
}

func (M *ConditionalMorphism) UnmarshalJSON(b []byte) error {
	if err := tjson.ParseJsonTags(b, M); err != nil {
		return err
	}
	if err := tjson.ParseTJsonTags(b, M); err != nil {
		return err
	}
	return nil
}

func (M ConditionalMorphism) String() string {
	return M.Name
}

func (M ConditionalMorphism) Init() error {
	return M.M.Init()
}

func (M ConditionalMorphism) Apply(o *Outcome, parameter Parameter) (*Outcome, error) {
	if (M.P)(o) {
		return M.M.Apply(o, parameter)
	} else {
		return o, nil
	}
}

func (M ConditionalMorphism) Awards() []Award {
	return M.M.Awards()
}

func (M ConditionalMorphism) EnumerateParameters() (Parameterization, error) {
	return M.M.EnumerateParameters()
}

/*
	Similar to the ConditionalMorphism, but as a loop!

	This is really a meta transformation. As such, there is an argument that this transformation should not contribute to the outcome sequence.
*/
// var (
// 	WhileMorphismImageMaximumRecursionDepth = 1 << 6
// )
type WhileMorphism struct {
	Name                string              `json:"name"`
	M                   Morphism            `tjson:"transformation"`
	ConditionalFunction func(*Outcome) bool `tjson:"conditionalFunction"`
	AppendOutcome       bool                `json:"appendOutcome"`
}

func (M *WhileMorphism) UnmarshalJSON(b []byte) error {
	if err := tjson.ParseJsonTags(b, M); err != nil {
		return err
	}
	if err := tjson.ParseTJsonTags(b, M); err != nil {
		return err
	}
	return nil
}

func (M WhileMorphism) String() string {
	return M.Name
}

func (M WhileMorphism) Init() error {
	return M.M.Init()
}

func (M WhileMorphism) Apply(o *Outcome, parameter Parameter) (*Outcome, error) {
	p := o
	var err error
	for M.ConditionalFunction(p) {
		p, err = M.M.Apply(p, parameter)
		if err != nil {
			return nil, err
		}
	}
	return p, nil
}

func (M WhileMorphism) Awards() []Award {
	return M.M.Awards()
}

func (M WhileMorphism) EnumerateParameters() (Parameterization, error) {
	return M.M.EnumerateParameters()
}
