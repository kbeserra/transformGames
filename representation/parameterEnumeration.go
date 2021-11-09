package representation

type ParameterEnumeration interface {
	Get(uint64) Parameter
	Length() uint64
	AssociatedMorphism() Morphism
}

/*

 */
type ParameterEnumerateionBase struct {
	M Morphism
}

/*

 */
type ConstantParameterEnumerateion struct {
	ParameterEnumerateionBase
	C interface{}
}

func (cpe *ConstantParameterEnumerateion) Get(n uint64) Parameter {
	return &ConstantParameter{C: cpe.C}
}

func (cpe *ConstantParameterEnumerateion) Length() uint64 {
	return 1
}

func (cpe *ConstantParameterEnumerateion) AssociatedMorphism() Morphism {
	return cpe.M
}

/*

 */
type ConcatenationParameterEnumerateion struct {
	ParameterEnumerateionBase
	Enumerations []ParameterEnumeration

	cached  bool
	length  uint64
	lengths []uint64
}

func (cpe *ConcatenationParameterEnumerateion) cache() {
	cpe.lengths = make([]uint64, len(cpe.Enumerations))
	cpe.length = 1
	for i, e := range cpe.Enumerations {
		cpe.lengths[i] = e.Length()
		cpe.length *= cpe.lengths[i]
	}
}

func (cpe *ConcatenationParameterEnumerateion) Get(n uint64) Parameter {
	if !cpe.cached {
		cpe.cache()
	}
	rtn := make([]Parameter, len(cpe.Enumerations))
	for i, e := range cpe.Enumerations {
		rtn[i] = e.Get(n % cpe.lengths[i])
		n /= cpe.lengths[i]
	}

	return &SequenceParameter{
		Sigmas: rtn,
	}
}

func (cpe *ConcatenationParameterEnumerateion) Length() uint64 {
	if !cpe.cached {
		cpe.cache()
	}
	return cpe.length
}

func (cpe *ConcatenationParameterEnumerateion) AssociatedMorphism() Morphism {
	return cpe.M
}
