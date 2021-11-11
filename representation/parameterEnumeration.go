package representation

type ParameterEnumeration interface {
	Get(uint64) (p Parameter, weight uint64)
	GetWeighted(uint64) Parameter
	Length() uint64
	Weight() uint64
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

func (cpe *ConstantParameterEnumerateion) Get(n uint64) (p Parameter, weight uint64) {
	return &ConstantParameter{C: cpe.C}, 1
}

func (cpe *ConstantParameterEnumerateion) GetWeighted(uint64) Parameter {
	return &ConstantParameter{C: cpe.C}
}

func (cpe *ConstantParameterEnumerateion) Length() uint64 {
	return 1
}

func (cpe *ConstantParameterEnumerateion) Weight() uint64 {
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

	cached bool
	length uint64
	weight uint64
}

func (cpe *ConcatenationParameterEnumerateion) cache() {
	cpe.length = 1
	cpe.weight = 1
	for _, e := range cpe.Enumerations {
		cpe.length *= e.Length()
		cpe.weight *= e.Weight()
	}
	cpe.cached = true
}

func (cpe *ConcatenationParameterEnumerateion) Get(n uint64) (Parameter, uint64) {
	if !cpe.cached {
		cpe.cache()
	}
	rtn := make([]Parameter, len(cpe.Enumerations))
	var w, weight uint64 = 0, 0
	for i, e := range cpe.Enumerations {
		rtn[i], w = e.Get(n % e.Length())
		weight *= w
		n /= e.Length()
	}

	return &SequenceParameter{
		Sigmas: rtn,
	}, weight
}

func (cpe *ConcatenationParameterEnumerateion) GetWeighted(n uint64) Parameter {
	if !cpe.cached {
		cpe.cache()
	}
	rtn := make([]Parameter, len(cpe.Enumerations))
	for i, e := range cpe.Enumerations {
		rtn[i] = e.GetWeighted(n % e.Weight())
		n /= e.Weight()
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

func (cpe *ConcatenationParameterEnumerateion) Weight() uint64 {
	if !cpe.cached {
		cpe.cache()
	}
	return cpe.weight
}

func (cpe *ConcatenationParameterEnumerateion) AssociatedMorphism() Morphism {
	return cpe.M
}

/*

 */
type IntegerIntervalParameterEnumerateion struct {
	ParameterEnumerateionBase
	// first coordinate is the lower bound, second is the upper bound.
	// We include the lower bound, exclude the upper.
	LowerBound int
	UpperBound int
	Weights    []int

	cached bool
	weight uint64
}

func (iipe *IntegerIntervalParameterEnumerateion) cache() {

	if iipe.Weights != nil {
		iipe.weight = 0
		for i := 0; i < (iipe.UpperBound - iipe.LowerBound); i++ {
			iipe.weight += uint64(iipe.Weights[i])
		}
	} else {
		iipe.weight = uint64(iipe.UpperBound - iipe.LowerBound)
	}

}

func (iipe *IntegerIntervalParameterEnumerateion) Get(n uint64) (Parameter, uint64) {
	if !iipe.cached {
		iipe.cache()
	}
	n %= uint64(iipe.UpperBound - iipe.LowerBound)
	var weight uint64 = 1
	if iipe.Weights != nil {
		weight = uint64(iipe.Weights[n])
	}

	return &ConstantParameter{
		C: iipe.LowerBound + int(n),
	}, weight
}

func (iipe *IntegerIntervalParameterEnumerateion) GetWeighted(n uint64) Parameter {
	if !iipe.cached {
		iipe.cache()
	}

	var i uint64 = 0
	if iipe.Weights != nil {
		for n > 0 {
			n -= uint64(iipe.Weights[i])
			i++
		}
	} else {
		i = n
	}
	return &ConstantParameter{
		C: iipe.LowerBound + int(i),
	}
}

func (iipe *IntegerIntervalParameterEnumerateion) Length() uint64 {
	if !iipe.cached {
		iipe.cache()
	}
	return uint64(iipe.UpperBound - iipe.LowerBound)
}

func (iipe *IntegerIntervalParameterEnumerateion) Weight() uint64 {
	if !iipe.cached {
		iipe.cache()
	}
	return iipe.weight
}

func (iipe *IntegerIntervalParameterEnumerateion) AssociatedMorphism() Morphism {
	return iipe.M
}
