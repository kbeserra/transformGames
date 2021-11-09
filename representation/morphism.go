package representation

type Morphism interface {
	Init() error
	String() string
	Apply(*Outcome, *Parameter) (*Outcome, error)
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

func (i *IdentityMorphism) Apply(o *Outcome, sigma *Parameter) (*Outcome, error) {
	return o, nil
}

func (i *IdentityMorphism) EnumerateParameters() (ParameterEnumeration, error) {
	return nil, nil
}
