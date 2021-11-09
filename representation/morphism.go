package representation

import "errors"

var (
	errFailedToCastParameter                  = errors.New("failed to cast parameter to necessary type")
	errLengthOfParametersNotLengthOfMorphisms = errors.New("length of parameters is not equal to the length of morphisms")
)

type Morphism interface {
	Init() error
	String() string
	Apply(*Outcome, Parameter) (*Outcome, error)
	EnumerateParameters() (ParameterEnumeration, error)
}

/*

 */
type IdentityMorphism struct {
	Name string
}

func (i *IdentityMorphism) Init() error {
	return nil
}

func (i *IdentityMorphism) String() string {
	return i.Name
}

func (i *IdentityMorphism) Apply(o *Outcome, sigma Parameter) (*Outcome, error) {
	return o, nil
}

func (i *IdentityMorphism) EnumerateParameters() (ParameterEnumeration, error) {
	return &ConstantParameterEnumerateion{
			ParameterEnumerateionBase: ParameterEnumerateionBase{
				M: i,
			},
			C: nil},
		nil
}

/*

 */
type ConcatenationMorphism struct {
	Name      string
	Morphisms []Morphism

	Expand       bool
	GatherStates bool
	GatherAwards bool
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
	return nil, nil
}

func (M *ConcatenationMorphism) Apply(root *Outcome, sigma Parameter) (*Outcome, error) {
	sigmas, ok := sigma.Value().([]Parameter)
	if !ok {
		return nil, errFailedToCastParameter
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

func (M *ConcatenationMorphism) EnumerateParameters() (ParameterEnumeration, error) {
	enumerations := make([]ParameterEnumeration, len(M.Morphisms))
	for i, m := range M.Morphisms {
		e, err := m.EnumerateParameters()
		if err != nil {
			return nil, err
		}
		enumerations[i] = e
	}
	return nil, nil
}
